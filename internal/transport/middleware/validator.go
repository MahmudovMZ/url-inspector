package middleware

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func ValidateURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("url")
		log.Printf("Validating: %s", target)
		if target == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		host, err := parseTarget(target)
		if err != nil || target == "" {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}
		host = strings.ToLower(host)
		host = strings.TrimSpace(host)

		ips, err := net.LookupIP(host)
		if err != nil {
			http.Error(w, "Invalid domain or host not found", http.StatusBadRequest)
			return
		}
		for _, ip := range ips {
			if isBlockedIp(ip) {
				http.Error(w, "Blocked host", http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)

	})
}

func parseTarget(target string) (string, error) {
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		target = "http://" + target
	}

	u, err := url.Parse(target)
	if err != nil {
		return "", err
	}

	host := u.Hostname()
	if host == "" {
		host = u.Path
	}
	return host, nil
}

func isBlockedIp(ip net.IP) bool {
	return ip.IsLoopback() || ip.IsPrivate() || ip.IsUnspecified() //checking for localhost, loopback or unspecified url/host
}

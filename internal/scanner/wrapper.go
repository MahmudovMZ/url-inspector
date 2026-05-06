package wrapper

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MahmudovMZ/go-scanner/pkg/models"
	scanner "github.com/MahmudovMZ/go-scanner/pkg/port"
)

func ScanTarget(target string) ([]models.Result, string, string, int, error) {
	start := time.Now()
	statusChan := make(chan int)
	portsChan := make(chan []models.Result)

	host := target
	u, err := url.Parse(target)
	if err != nil {
		host = u.Host
	} else {
		host = strings.TrimPrefix(target, "http://")
		host = strings.TrimPrefix(target, "https://")
		host = strings.Split(target, "/")[0]
	}

	iPs, err := net.LookupIP(host)
	if err != nil || len(iPs) == 0 {
		return nil, "", "", 0, fmt.Errorf("could not resolve host: %v", err)

	}
	targetIp := iPs[0].String()

	whiteList := []int{80, 443, 8080, 8443, 3000, 5000, 8000, 21, 22, 23, 3389, 3306, 5432, 6379, 27017, 25, 53}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	go func() {
		resp, err := client.Get("http://" + host)
		statusCode := 0
		if err == nil {
			statusCode = resp.StatusCode
			defer resp.Body.Close()
		}
		statusChan <- statusCode
	}()

	go func() {
		s := scanner.NewScanner(targetIp, time.Second*2, 10)

		var openPorts []models.Result
		for _, p := range whiteList {
			res := s.ScanPort(p)

			if res.Open {
				openPorts = append(openPorts, res)
			}
		}
		portsChan <- openPorts
	}()
	duration := time.Since(start)
	code := <-statusChan
	result := <-portsChan

	return result, targetIp, duration.String(), code, nil //Method .string will automaticaly convert time into string if less the second wil be some like 450m
}

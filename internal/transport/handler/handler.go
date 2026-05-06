package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MahmudovMZ/url-inspector/internal/models"
	wrapper "github.com/MahmudovMZ/url-inspector/internal/scanner"
)

func Inspect(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("url")
	if target == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
	}

	results, ip, rTime, statusCode, err := wrapper.ScanTarget(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalResponse := models.InspectResult{
		TargetURL:    target,
		IpAddress:    ip,
		ScanDetails:  results,
		ResponseTime: rTime,
		IsReacheble:  true,
		StatusCode:   statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResponse)
}

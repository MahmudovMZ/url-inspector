package models

import (
	"github.com/MahmudovMZ/go-scanner/pkg/models"
)

type InspectResult struct {
	TargetURL    string          `json:"target_url"`
	IsReacheble  bool            `json:"is_reacheble"`
	StatusCode   int             `json:"status_code"`
	IpAddress    string          `json:"ip_address"`
	ScanDetails  []models.Result `json:"scan_details"`
	ResponseTime string          `json:"response_time"`
}

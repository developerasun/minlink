package api

import (
// "encoding/json"
)

type HealthResponse struct {
	Message string `json:"message"`
}

type CrawlResponse struct {
	Dxy string `json:"dxy"`
}

type SseStatsResponse struct {
	Data string `json:"data"`
}

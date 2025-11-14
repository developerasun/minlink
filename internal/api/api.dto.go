package api

import (
// "encoding/json"
)

type HealthResponse struct {
	Message string `json:"message"`
}

type CrawlResponse struct {
	Data string `json:"data"`
}

type SseStatsResponse struct {
	Data string `json:"data"`
}

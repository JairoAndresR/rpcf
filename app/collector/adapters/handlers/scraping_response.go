package handlers

import "time"

type ScrapingResponse struct {
	Time int64
}

func NewScrapingResponse() *ScrapingResponse {
	return &ScrapingResponse{
		Time: time.Now().Unix(),
	}
}

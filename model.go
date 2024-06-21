package uranai

import (
	"time"
)

type Result struct {
	Rank         int32     `json:"rank"`
	Name         string    `json:"name"`
	LuckyItem    string    `json:"lucky_item"`
	LuckyColor   string    `json:"lucky_color"`
	LuckyService string    `json:"lucky_service"`
	CareerLuck   int32     `json:"career_luck"`
	LoveLuck     int32     `json:"love_luck"`
	HealthLuck   int32     `json:"health_luck"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}

type ResultSet struct {
	Results []Result `json:"results"`
}

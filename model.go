package uranai

import "google.golang.org/genproto/googleapis/type/date"

type Result struct {
	Rank         int32     `json:"rank"`
	Name         string    `json:"name"`
	LuckyItem    string    `json:"lucky_item"`
	LuckyColor   string    `json:"lucky_color"`
	LuckyService string    `json:"lucky_service"`
	Description  string    `json:"description"`
	CreatedAt    date.Date `json:"created_at"`
}

type ResultSet struct {
	Results []Result `json:"results"`
}

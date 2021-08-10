package entity

type Post struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	User int    `json:"user"`
}

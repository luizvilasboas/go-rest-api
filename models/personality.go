package models

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

func (p *Personality) TableName() string {
	return "personality"
}

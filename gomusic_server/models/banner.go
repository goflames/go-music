package models

type Banner struct {
	ID  int    `json:"id"`
	Pic string `json:"pic"`
}

func (Banner) TableName() string {
	return "banner"
}

package models

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}

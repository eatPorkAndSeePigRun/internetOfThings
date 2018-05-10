package model

type User struct {
	Id string `json:"id"`
	TrueName string `json:"trueName"`
	Password string `json:"password"`
	Power string `json:"power"`
}

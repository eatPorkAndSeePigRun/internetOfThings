package model

import "time"

type User struct {
	Id string `json:"id"`
	TrueName string `json:"trueName"`
	Password string `json:"password"`
	Power string `json:"power"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

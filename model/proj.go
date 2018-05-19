package model

import "time"

type Proj struct {
	ProjId int `json:"proj_id"`
	ProjName string `json:"proj_name"`
	ProjRemark string `json:"proj_remark"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

package model

import "time"

type Area struct {
	AreaId     int    `json:"area_id"`
	AreaName   string `json:"area_name"`
	AreaRemark string `json:"area_remark"`
	//TODO 这里pic的数据类型可能需要修改
	Pic        []byte    `json:"pic"`
	ProjId     int       `json:"proj_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

package model

import "time"

type Node struct {
	NodeId     int       `json:"node_id"`
	NodeName   string    `json:"node_name"`
	NodeRemark string    `json:"node_remark"`
	NodeX      int       `json:"node_x"`
	NodeY      int       `json:"node_y"`
	AreaId     int       `json:"area_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

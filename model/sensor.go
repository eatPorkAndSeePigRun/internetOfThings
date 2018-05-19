package model

import "time"

type Sensor struct {
	SensorId int `json:"sensor_id"`
	SensorName string `json:"sensor_name"`
	Max float32 `json:"max"`
	Min float32 `json:"min"`
	NodeId int `json:"node_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

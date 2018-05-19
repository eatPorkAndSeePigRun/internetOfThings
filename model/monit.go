package model

import "time"

type Monit struct {
	Monit int `json:"monit"`
	Value float32 `json:"value"`
	SensorId int `json:"sensor_id"`
	CreateTime time.Time `json:"create_time"`
}

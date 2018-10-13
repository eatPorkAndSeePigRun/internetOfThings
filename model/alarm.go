package model

import "time"

type Alarm struct {
	AlarmId    int       `json:"alarm_id"`
	HasSolved  int       `json:"has_solved"`
	SensorId2  int       `json:"sensor_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

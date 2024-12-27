// models/run.go
package models

import "gorm.io/gorm"

type Run struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Distance  float64 `json:"distance"`
	Duration  int64   `json:"duration"`
	GPSPoints string  `json:"gps_points"`
}

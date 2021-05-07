package po

import "time"

type Users struct {
	Id         int
	Name       string
	Password   string
	Status     int
	Fee        float64
	FeeStatus  int
	FeeTotal   int64
	CreateDate *time.Time
	CreateTime *time.Time
}

package po

import (
	"time"
)

//type Users struct {
//	Id         int
//	Name       *sql.NullString
//	Password   *sql.NullString
//	Status     *sql.NullInt32
//	Fee        *sql.NullFloat64
//	FeeStatus  *sql.NullInt32
//	FeeTotal   *sql.NullInt64
//	CreateDate *sql.NullTime
//	CreateTime *sql.NullTime
//}

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

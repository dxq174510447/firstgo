package po

import (
	"database/sql"
)

type Users struct {
	Id         int
	Name       *sql.NullString
	Password   *sql.NullString
	Status     *sql.NullInt32
	Fee        *sql.NullFloat64
	FeeStatus  *sql.NullInt32
	FeeTotal   *sql.NullInt64
	CreateDate *sql.NullTime
	CreateTime *sql.NullTime
}

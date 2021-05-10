package po

import "time"

AUTO
//Indicates that the persistence provider should pick an appropriate strategy for the particular database.
//IDENTITY
//Indicates that the persistence provider must assign primary keys for the entity using a database identity column.
//SEQUENCE
//Indicates that the persistence provider must assign primary keys for the entity using a database sequence.
//TABLE
//Indicates that the persistence provider must assign primary keys for the entity using an underlying database table to ensure uniqueness.
//`column:"Id" id:"true" transient:"false" updatable:"false" table:"" columnDefinition:"" GenerationType:"AUTO IDENTITY SEQUENCE TABLE"`
import (
	"time"
)

//`species:"gopher" color:"blue"`
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
	Id         int      `column:"id" id:"true" transient:"false" updatable:"false" table:"users" columnDefinition:"" GenerationType:"IDENTITY"`
	Name       string	`column:"name" updatable:"false" columnDefinition:""`
	Password   string
	Status     int
	Fee        float64
	FeeStatus  int
	FeeTotal   int64
	CreateDate *time.Time
	CreateTime *time.Time
	NameIn     []string
}

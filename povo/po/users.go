package po

import "time"

//Indicates that the persistence provider should pick an appropriate strategy for the particular database.
//IDENTITY
//Indicates that the persistence provider must assign primary keys for the entity using a database identity column.
//SEQUENCE
//Indicates that the persistence provider must assign primary keys for the entity using a database sequence.
//TABLE
//Indicates that the persistence provider must assign primary keys for the entity using an underlying database table to ensure uniqueness.
//`column:"Id" id:"true" transient:"false" updatable:"false" table:"" columnDefinition:"" GenerationType:"AUTO IDENTITY SEQUENCE TABLE"`

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
	Id         int        `column:"id" id:"true" transient:"false" updatable:"false" table:"users" columnDefinition:"" GenerationType:"IDENTITY"`
	Name       string     `column:"name" columnDefinition:""`
	Password   string     `column:"password" columnDefinition:""`
	Status     int        `column:"status" columnDefinition:""`
	Fee        float64    `column:"fee" columnDefinition:""`
	FeeStatus  int        `column:"fee_status" columnDefinition:""`
	FeeTotal   int64      `column:"fee_total" columnDefinition:""`
	CreateDate *time.Time `column:"create_date" columnDefinition:""`
	CreateTime *time.Time `column:"create_time" updatable:"false" columnDefinition:""`
	NameIn     []string   `transient:"true"`
}

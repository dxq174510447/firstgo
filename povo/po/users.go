package po

import "time"

// AUTO
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
	Id         int        `column:"id" id:"true" transient:"false" updatable:"true" table:"users" columnDefinition:"" GenerationType:"IDENTITY" json:"id,omitempty"`
	Name       string     `column:"name" columnDefinition:"" json:"name,omitempty"`
	Password   string     `column:"password" columnDefinition:"" json:"password,omitempty"`
	Status     int        `column:"status" columnDefinition:"" json:"status,omitempty"`
	Fee        float64    `column:"fee" columnDefinition:"" json:"fee,omitempty"`
	FeeStatus  int        `column:"fee_status" columnDefinition:"" json:"fee_status,omitempty"`
	FeeTotal   int64      `column:"fee_total" columnDefinition:"" json:"fee_total,omitempty"`
	CreateDate *time.Time `column:"create_date" updatable:"false" columnDefinition:"" json:"create_date,omitempty"`
	CreateTime *time.Time `column:"create_time"  columnDefinition:"" json:"create_time,omitempty"`
	NameIn     []string   `transient:"true" json:"name_in,omitempty"`
}

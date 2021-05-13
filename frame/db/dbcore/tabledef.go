package dbcore

import "reflect"

type TableDef struct {
	Name           string
	DbName         string
	GenerationType string
	IdColumn       *TableColumnDef
	Columns        []*TableColumnDef
}

type TableColumnDef struct {
	FieldName  string
	Field      *reflect.StructField
	ColumnName string
	Transient  bool
	Updatable  bool
}

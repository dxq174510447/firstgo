package po

type Users struct {
	Id       int `dbprimary:"1" dbtable:"Users" dbfield:"id"`
	Name     string
	Password string
	Status   int
}

package impl

type AuthServicer interface {
	AuthValid(name string, pwd string) (bool, error)
}

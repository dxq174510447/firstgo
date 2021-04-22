package db

import "firstgo/frame/context"

type databaseUtil struct {
}

func (d *databaseUtil) setDbConIfNotExist(local context.LocalStack) {
	con := local.Get(DataBaseConnectKey)
	if con == nil {

	}
}

var DatabaseUtil databaseUtil = databaseUtil{}

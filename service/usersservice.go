package service

import "firstgo/dao"

type UsersService struct {
	usersDao *dao.UsersDao
}

func (c *UsersService) Save() {

}

func (c *UsersService) Update() {

}

func (c *UsersService) Delete() {

}

func (c *UsersService) ChangeStatus() {

}

var UsersServiceImpl UsersService = UsersService{}

func init() {
	UsersServiceImpl.usersDao = &dao.UsersDaoImpl
}

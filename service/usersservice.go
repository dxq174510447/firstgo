package service

import (
	"firstgo/dao"
	"firstgo/frame"
	"firstgo/povo/po"
	"firstgo/povo/vo"
)

type UsersService struct {
	usersDao *dao.UsersDao
}

func (c *UsersService) Save(local *frame.FrameStack, data *po.Users) *vo.UsersVo {

	var total int = c.usersDao.FindByName(local, data.Name)
	if total > 0 {
		panic(frame.NewException(502, "名称重复"))
	}

	c.usersDao.Save(local, data)
	return c.Get(local, data.Id)
}

func (c *UsersService) Update(local *frame.FrameStack, data *po.Users) *vo.UsersVo {

	var total int = c.usersDao.FindByNameExcludeId(local, data.Name, data.Id)
	if total > 0 {
		panic(frame.NewException(502, "名称重复"))
	}

	c.usersDao.Update(local, data)
	return c.Get(local, data.Id)
}

func (c *UsersService) Delete(local *frame.FrameStack, id int) {
	c.usersDao.Delete(local, id)
}

func (c *UsersService) Get(local *frame.FrameStack, id int) *vo.UsersVo {
	return c.usersDao.Get(local, id)
}

func (c *UsersService) ChangeStatus(local *frame.FrameStack, id int, status int) {
	c.usersDao.ChangeStatus(local, id, status)
}
func (c *UsersService) List(local *frame.FrameStack, param *vo.UsersParam) ([]*vo.UsersVo, int) {
	return c.usersDao.List(local, param)
}

var UsersServiceImpl UsersService = UsersService{}

func init() {
	UsersServiceImpl.usersDao = &dao.UsersDaoImpl
}

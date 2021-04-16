package controller

import (
	"firstgo/povo/vo"
	"firstgo/service"
	"firstgo/servlet"
	"fmt"
	"net/http"
)

type UsersController struct {
	usersService *service.UsersService
}

func (c *UsersController) Save(response http.ResponseWriter,
	request *http.Request,
	name string,
	status int,
	param *vo.UsersAdd) {
	c.usersService.Delete()
	fmt.Println("get", request.URL.Path, "  ", name, "  ", status, param)
	response.Write([]byte("123"))
}

func (c *UsersController) Update() {
	c.usersService.Update()
	fmt.Println("Update")
}

func (c *UsersController) Delete() {
	c.usersService.Delete()
	fmt.Println("delete")
}

func (c *UsersController) Get(response http.ResponseWriter,
	request *http.Request,
	name string,
	status int) {
	c.usersService.Delete()
	fmt.Println("get", request.URL.Path, "  ", name, "  ", status)
	response.Write([]byte("123"))
}

func (c *UsersController) List() {
	c.usersService.Delete()
}

func (c *UsersController) ChangeStatus() {
	c.usersService.ChangeStatus()
}

var UsersControllerImpl UsersController = UsersController{}

var UsersRequestController servlet.RequestController = servlet.RequestController{
	HttpPath: "/users",
	Target:   &UsersControllerImpl,
	Methods: []servlet.RequestMethod{
		{
			HttpMethod: "post",
			HttpPath:   "/",
			MethodName: "Save",
			Paramter:   "_,_,name,status,_",
		},
		{
			HttpMethod: "put",
			HttpPath:   "/",
			MethodName: "Update",
		},
		{
			HttpMethod: "delete",
			HttpPath:   "/",
			MethodName: "Delete",
		},
		{
			HttpMethod: "post",
			HttpPath:   "/change",
			MethodName: "ChangeStatus",
		},
		{
			HttpMethod: "get",
			HttpPath:   "/",
			MethodName: "Get",
			Paramter:   "_,_,name,status,_",
		},
		{
			HttpMethod: "post",
			HttpPath:   "/list",
			MethodName: "List",
		},
	},
}

func init() {
	UsersControllerImpl.usersService = &service.UsersServiceImpl
	servlet.DispatchServlet.AddRequestMapping(&UsersRequestController)
}

package v1

import (
	"context"
	"firstgo/src/main/golang/com/firstgo/povo/vo"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/application"
)

/*
UsersController
@RestController
@RequestMapping()
 */
type UsersController struct {

}

// Save 新增
func (c *UsersController) Save(ctx context.Context, param *vo.UsersAdd) (*vo.UsersVo,error) {
	fmt.Println("Save")
	return nil,nil
}

// Update 修改
func (c *UsersController) Update(ctx context.Context, param *vo.UsersEdit) (*vo.UsersVo,error) {
	fmt.Println("Update")
	return nil,nil
}

// Delete 删除
func (c *UsersController) Delete(ctx context.Context, id int) error {
	fmt.Printf("Delete %d \n",id)
	return nil
}

// Get 查看
func (c *UsersController) Get(ctx context.Context, id int) (*vo.UsersVo,error) {
	fmt.Printf("Get %d \n",id)
	return nil,nil
}

// List 列表
func (c *UsersController) List(ctx context.Context, param *vo.UsersParam) ([]*vo.UsersVo,error) {
	fmt.Printf("List\n")
	return nil,nil
}


func init() {
	usersController := &UsersController{}
	application.GetResourcePool().RegisterInstance("",usersController)
}

package dao

import (
	"firstgo/povo/po"
	"firstgo/povo/vo"
	"github.com/dxq174510447/goframe/lib/frame/context"
	"github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	"github.com/dxq174510447/goframe/lib/frame/proxy"
)

type UsersDao struct {
	dbcore.BaseDao
	Proxy_ *proxy.ProxyClass
	//test
	GetById_             func(local *context.LocalStack, id int) (*po.Users, error)
	FindByNameAndStatus_ func(local *context.LocalStack, name string, status int, statusList []string) ([]*po.Users, error)
	FindIds_             func(local *context.LocalStack, users *po.Users) ([]int, error)
	FindNames_           func(local *context.LocalStack, users *po.Users) ([]string, error)
	FindFees_            func(local *context.LocalStack, users *po.Users) ([]float64, error)
	GetMaxFees_          func(local *context.LocalStack, users *po.Users) (float64, error)
	GetMaxId_            func(local *context.LocalStack, users *po.Users) (int, error)
	UpdateNameByField_   func(local *context.LocalStack, name string, id int) (int, error)
	UpdateNameByEntity_  func(local *context.LocalStack, users *po.Users) (int, error)
	DeleteNameByField_   func(local *context.LocalStack, name string, id int) (int, error)
	DeleteNameByEntity_  func(local *context.LocalStack, users *po.Users) (int, error)
	InsertSingle_        func(local *context.LocalStack, users *po.Users) (int, error)
	InsertBatch_         func(local *context.LocalStack, users []*po.Users) (int, error)

	//users
	Save1_                func(local *context.LocalStack, data *po.Users) (int, error)
	Update1_              func(local *context.LocalStack, data *po.Users) (int, error)
	Delete1_              func(local *context.LocalStack, id int) (int, error)
	Get1_                 func(local *context.LocalStack, id int) (*vo.UsersVo, error)
	ChangeStatus1_        func(local *context.LocalStack, id int, status int) (int, error)
	List1_                func(local *context.LocalStack, param *vo.UsersParam) (*vo.UsersPage, error)
	FindByNameExcludeId1_ func(local *context.LocalStack, name string, id int) (int, error)
	FindByName1_          func(local *context.LocalStack, name string) (int, error)
	QueryAddon_           func(local *context.LocalStack, users *vo.QueryAddon) ([]*vo.QueryResult, error)
}

// 测试案例
func (c *UsersDao) GetById(local *context.LocalStack, id int) (*po.Users, error) {
	return c.GetById_(local, id)
}

func (c *UsersDao) FindByNameAndStatus(local *context.LocalStack, name string, status int, statusList []string) ([]*po.Users, error) {
	return c.FindByNameAndStatus_(local, name, status, statusList)
}

func (c *UsersDao) FindIds(local *context.LocalStack, users *po.Users) ([]int, error) {
	return c.FindIds_(local, users)
}
func (c *UsersDao) FindNames(local *context.LocalStack, users *po.Users) ([]string, error) {
	return c.FindNames_(local, users)
}
func (c *UsersDao) FindFees(local *context.LocalStack, users *po.Users) ([]float64, error) {
	return c.FindFees_(local, users)
}
func (c *UsersDao) GetMaxFees(local *context.LocalStack, users *po.Users) (float64, error) {
	return c.GetMaxFees_(local, users)
}
func (c *UsersDao) GetMaxId(local *context.LocalStack, users *po.Users) (int, error) {
	return c.GetMaxId_(local, users)
}
func (c *UsersDao) UpdateNameByField(local *context.LocalStack, name string, id int) (int, error) {
	return c.UpdateNameByField_(local, name, id)
}
func (c *UsersDao) UpdateNameByEntity(local *context.LocalStack, users *po.Users) (int, error) {
	return c.UpdateNameByEntity_(local, users)
}
func (c *UsersDao) DeleteNameByField(local *context.LocalStack, name string, id int) (int, error) {
	return c.DeleteNameByField_(local, name, id)
}
func (c *UsersDao) DeleteNameByEntity(local *context.LocalStack, users *po.Users) (int, error) {
	return c.DeleteNameByEntity_(local, users)
}
func (c *UsersDao) InsertSingle(local *context.LocalStack, users *po.Users) (int, error) {
	return c.InsertSingle_(local, users)
}
func (c *UsersDao) InsertBatch(local *context.LocalStack, users []*po.Users) (int, error) {
	return c.InsertBatch_(local, users)
}

func (c *UsersDao) QueryAddon(local *context.LocalStack, users *vo.QueryAddon) ([]*vo.QueryResult, error) {
	return c.QueryAddon_(local, users)
}

// Get override
func (c *UsersDao) Get(local *context.LocalStack, id int) (*po.Users, error) {
	m, err := c.BaseDao.Get(local, id)
	if err != nil {
		return nil, err
	} else {
		if m == nil {
			return nil, err
		}
		return m.(*po.Users), err
	}
}

// Get override
func (c *UsersDao) Find(local *context.LocalStack, users *po.Users) ([]*po.Users, error) {
	var result []*po.Users
	m, err := c.BaseDao.Find(local, users)
	if err != nil {
		return result, err
	} else {
		if len(m) == 0 {
			return result, err
		}
		result = make([]*po.Users, 0, len(m))
		for _, e := range m {
			result = append(result, e.(*po.Users))
		}
		return result, nil
	}
}

// user case
func (c *UsersDao) Save1(local *context.LocalStack, data *po.Users) (int, error) {
	return c.Save1_(local, data)
}
func (c *UsersDao) Update1(local *context.LocalStack, data *po.Users) (int, error) {
	return c.Update1_(local, data)
}

func (c *UsersDao) Delete1(local *context.LocalStack, id int) (int, error) {
	return c.Delete1_(local, id)
}

func (c *UsersDao) Get1(local *context.LocalStack, id int) (*vo.UsersVo, error) {
	return c.Get1_(local, id)
}

func (c *UsersDao) ChangeStatus1(local *context.LocalStack, id int, status int) (int, error) {
	return c.ChangeStatus1_(local, id, status)
}

func (c *UsersDao) List1(local *context.LocalStack, param *vo.UsersParam) (*vo.UsersPage, error) {
	return c.List1_(local, param)
}

func (c *UsersDao) FindByNameExcludeId1(local *context.LocalStack, name string, id int) (int, error) {
	return c.FindByNameExcludeId1_(local, name, id)
}

func (c *UsersDao) FindByName1(local *context.LocalStack, name string) (int, error) {
	return c.FindByName1_(local, name)
}

func (c *UsersDao) ProxyTarget() *proxy.ProxyClass {
	return c.Proxy_
}

var usersDao UsersDao = UsersDao{
	Proxy_: &proxy.ProxyClass{
		Annotations: []*proxy.AnnotationClass{
			proxy.NewSingleAnnotation(proxy.AnnotationDao, nil),
		},
		Methods: []*proxy.ProxyMethod{
			//测试
			{
				Name: "GetById",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,id"),
				},
			},
			{
				Name: "FindByNameAndStatus",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,name,status,statusList"),
				},
			},
			{
				Name: "UpdateNameByField",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,name,id"),
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			{
				Name: "UpdateNameByEntity",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			{
				Name: "DeleteNameByField",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,name,id"),
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			{
				Name: "DeleteNameByEntity",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			{
				Name: "InsertSingle",
				Annotations: []*proxy.AnnotationClass{
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			{
				Name: "InsertBatch",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,values"),
					proxy.NewSingleAnnotation(dbcore.TransactionRequire, nil),
				},
			},
			//users
			{
				Name: "Delete1",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,id"),
				},
			},
			{
				Name: "Get1",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,id"),
				},
			},
			{
				Name: "ChangeStatus1",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,id,status"),
				},
			},
			{
				Name: "FindByNameExcludeId1",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,name,id"),
				},
			},
			{
				Name: "FindByName1",
				Annotations: []*proxy.AnnotationClass{
					dbcore.NewSqlProvierConfigAnnotation("_,name"),
				},
			},
		},
	},
}

func GetUsersDao() *UsersDao {
	return &usersDao
}

func init() {
	dbcore.AddMapperProxyTarget(proxy.ProxyTarger(&usersDao), UsersXml, &po.Users{})
}

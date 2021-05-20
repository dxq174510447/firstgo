package dao

import (
	"firstgo/povo/po"
	"firstgo/util"
	"fmt"
	"github.com/dxq174510447/goframe/lib/frame/context"
	"github.com/dxq174510447/goframe/lib/frame/db/dbcore"
	_ "github.com/dxq174510447/goframe/lib/frame/db/filter"
	"testing"
	"time"
)

func printRow(v *po.Users) {
	fmt.Println("-->",
		v.Id,
		dbcore.GetSqlNullTypeValue(v.Name),
		dbcore.GetSqlNullTypeValue(v.Password),
		dbcore.GetSqlNullTypeValue(v.Status),
		dbcore.GetSqlNullTypeValue(v.Fee),
		dbcore.GetSqlNullTypeValue(v.FeeStatus),
		dbcore.GetSqlNullTypeValue(v.CreateDate),
		dbcore.GetSqlNullTypeValue(v.CreateTime))
}

func ATestUsersDao_FindIds(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().FindIds(local, &po.Users{})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) > 0 {
			for _, v := range m5 {
				fmt.Println(v)
			}

		} else {
			fmt.Println("result length 0")
		}
	}
}

func ATestUsersDao_FindNames(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().FindNames(local, &po.Users{NameIn: []string{"w1322", "w13a232"}})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) > 0 {
			for _, v := range m5 {
				fmt.Println(v)
			}
		} else {
			fmt.Println("result length 0")
		}
	}
}

func ATestUsersDao_FindFees(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().FindFees(local, &po.Users{})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) > 0 {
			for _, v := range m5 {
				fmt.Println(v)
			}
		} else {
			fmt.Println("result length 0")
		}
	}
}

func ATestUsersDao_GetMaxFees(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().GetMaxFees(local, &po.Users{})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println(m5)
	}
}

func ATestUsersDao_GetMaxId(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().GetMaxId(local, &po.Users{})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println(m5)
	}
}

func ATestUsersDao_GetById(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().GetById(local, 94)
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if m5 != nil {
			printRow(m5)
		} else {
			fmt.Println("no row")
		}

	}
}

func ATestUsersDao_FindByNameAndStatus(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().FindByNameAndStatus(local, "w1322", 1, []string{"102", "100", "1"})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) > 0 {
			for _, v := range m5 {
				printRow(v)
			}
		} else {
			fmt.Println("result length 0")
		}
	}

}

func ATestUsersDao_UpdateNameByEntity(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().UpdateNameByEntity(local, &po.Users{Id: 93, Name: "xxxx1"})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_UpdateNameByField(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().UpdateNameByField(local, "xxxx2", 94)
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_DeleteNameByEntity(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().DeleteNameByEntity(local, &po.Users{Id: 93, Name: "xxxx1", NameIn: []string{"aaa", "bbb"}})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_DeleteNameByField(t *testing.T) {
	local := context.NewLocalStack()

	m5, err5 := GetUsersDao().DeleteNameByField(local, "xxxx2", 94)
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_InsertSingle(t *testing.T) {
	local := context.NewLocalStack()

	n := time.Now()
	m5, err5 := GetUsersDao().InsertSingle(local, &po.Users{Name: "nnn21", Password: "ppp2", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_InsertBatch(t *testing.T) {
	local := context.NewLocalStack()

	n := time.Now()
	m5, err5 := GetUsersDao().InsertBatch(local, []*po.Users{
		{Name: "nnn2", Password: "ppp2", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n},
		{Name: "nnn3", Password: "ppp3"},
		{Name: "nnn4", Password: "ppp4"},
	})
	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func TestUsersDao_Save(t *testing.T) {
	local := context.NewLocalStack()

	n := time.Now()
	//m5, err5 := GetUsersDao().Save(local, &po.Users{Name: "new22", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	p := &po.Users{Name: "new22", Fee: 333.333, CreateDate: &n}
	m5, err5 := GetUsersDao().Save(local, p)

	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println(m5)
		printRow(p)
	}
}

func ATestUsersDao_Update(t *testing.T) {
	local := context.NewLocalStack()

	n := time.Now()
	//m5, err5 := GetUsersDao().Save(local, &po.Users{Name: "new22", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	m5, err5 := GetUsersDao().Update(local, &po.Users{Id: 155, Name: "new212", Fee: 333.333, CreateDate: &n, CreateTime: &n})

	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_Delete(t *testing.T) {
	local := context.NewLocalStack()

	//n := time.Now()
	//m5, err5 := GetUsersDao().Save(local, &po.Users{Name: "new22", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	m5, err5 := GetUsersDao().Delete(local, 155)

	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		fmt.Println("result", m5)
	}
}

func ATestUsersDao_Get(t *testing.T) {
	local := context.NewLocalStack()

	//n := time.Now()
	//m5, err5 := GetUsersDao().Save(local, &po.Users{Name: "new22", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	m5, err5 := GetUsersDao().Get(local, 154)

	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if m5 == nil {
			fmt.Println("result nil")
		} else {
			printRow(m5)
			fmt.Println("result", m5)
		}
	}
}

func ATestUsersDao_Find(t *testing.T) {
	local := context.NewLocalStack()

	//n := time.Now()
	//m5, err5 := GetUsersDao().Save(local, &po.Users{Name: "new22", Status: 1, Fee: 333.333, CreateTime: &n, CreateDate: &n})
	m5, err5 := GetUsersDao().Find(local, &po.Users{Name: "new22"})

	if err5 != nil {
		fmt.Println(err5)
		panic(err5)
	} else {
		if len(m5) == 0 {
			fmt.Println("result length 0")
		} else {
			for _, r := range m5 {
				printRow(r)
			}
		}
	}
}

func init() {
	var defaultFactory dbcore.DatabaseFactory = dbcore.DatabaseFactory{
		DbUser: util.ConfigUtil.Get("DB_USER", "platform"),
		DbPwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxcxcx"),
		DbName: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
		DbPort: util.ConfigUtil.Get("DB_PORT", "3306"),
		DbHost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com"),
	}
	db := defaultFactory.NewDatabase()
	dbcore.AddDatabaseRouter(dbcore.DataBaseDefaultKey, db)
}

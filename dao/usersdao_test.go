package dao

import (
	"firstgo/frame/context"
	"firstgo/frame/db/dbcore"
	_ "firstgo/frame/db/filter"
	"firstgo/povo/po"
	"fmt"
	"testing"
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

func TestUsersDao_FindNames(t *testing.T) {
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

	m5, err5 := GetUsersDao().FindFees_(local, &po.Users{})
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

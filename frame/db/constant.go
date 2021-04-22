package db

const (
	// TransactionReadOnly 只读 用于注解
	TransactionReadOnly = "TransactionReadOnly_"

	// TransactionRequire 如果有事物的话无需创建 没有的话就创建一个
	TransactionRequire = "TransactionRequire_"

	// TransactionRequireNew 如果有事物的话 在创建一个
	TransactionRequireNew = "TransactionRequireNew_"

	//DataBaseDefaultKey db 路由 默认key
	DataBaseDefaultKey = "default"

	//上下文中保存的数据库连接
	DataBaseConnectKey = "DataBaseConnectKey_"
)

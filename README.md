# firstgo
设置db环境变量key
dbuser: util.ConfigUtil.Get("DB_USER", "platform"),
dbpwd:  util.ConfigUtil.Get("DB_PASSWORD", "xxxx"),
dbname: util.ConfigUtil.Get("DB_NAME", "plat_base1"),
dbport: util.ConfigUtil.Get("DB_PORT", "3306"),
dbhost: util.ConfigUtil.Get("DB_HOST", "rm-bp1thh63s5tx33q0kio.mysql.rds.aliyuncs.com")




## 不使用翻墙的时候，需要设置go env
GOPROXY="direct"  
GOSUMDB="off"   
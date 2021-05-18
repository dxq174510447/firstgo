# firstgo

### sql

```sql

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL COMMENT '用户名',
  `password` varchar(45) DEFAULT NULL COMMENT '密码',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '0:有效帐户 1:无效帐户',
  `fee` decimal(11,2) DEFAULT NULL,
  `fee_status` int(5) DEFAULT NULL,
  `fee_total` int(11) DEFAULT NULL,
  `create_date` date DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8 COMMENT='用户表';


```

### 启动
```
$ export DB_PASSWORD=xxxxx
$ go run ./
$ curl -XGET "http://localhost:8080/api/v1/users?id=152"
```
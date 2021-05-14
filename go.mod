module firstgo

//replace github.com/jinzhu/copier v0.0.0 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/copier v0.2.9
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/urfave/negroni v1.0.0
	klook.libs v0.0.0-20210302022848-dc28d5fe108b
)

replace (
	github.com/boombuler/barcode => bitbucket.org/klook/klook-libs/src/github.com/boombuler/barcode v0.0.0-20210401085317-25d36a72c19d
	klook.libs => bitbucket.org/klook/klook-libs/src/klook.libs v0.0.0-20210401085317-25d36a72c19d
	klook.proto => bitbucket.org/klook/klook-libs/src/klook.proto v0.0.0-20210401085317-25d36a72c19d
)

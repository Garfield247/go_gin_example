module github.com/Garfield247/go_gin_example

go 1.13

require (
	github.com/astaxie/beego v1.12.2
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/mailru/easyjson v0.7.2 // indirect
	github.com/swaggo/swag v1.6.7 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.4 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200728102440-3e129f6d46b1 // indirect
	golang.org/x/tools v0.0.0-20200729194436-6467de6f59a7 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/Garfield247/go_gin_example/conf => ./conf
	github.com/Garfield247/go_gin_example/middleware => ./middleware
	github.com/Garfield247/go_gin_example/models => ./models
	github.com/Garfield247/go_gin_example/pkg => ./pkg
	github.com/Garfield247/go_gin_example/routers => ./routers
	github.com/Garfield247/go_gin_example/util => ./util
)

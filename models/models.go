package models

import (
	"fmt"
	"github.com/Garfield247/go_gin_example/pkg/logging"
	"github.com/Garfield247/go_gin_example/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	log.Println("初始化数据库连接中")
	var (
		err error
		databaseType = setting.DatabaseSetting.Type
		user = setting.DatabaseSetting.User
		pass = setting.DatabaseSetting.Password
		host = setting.DatabaseSetting.Host
		name = setting.DatabaseSetting.Name
	)
	/*
	如果使用语法糖 := 赋值，全局变量 db 的作用域只在 Init{} 函数内，其他函数内调用会报错空指针。
	db,err := gorm.Open()
	因此若避免全局变量变成局部变量，应采用 “=” 写法：
	*/
	db,err = gorm.Open(databaseType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",user,pass,host,name))
	if err != nil {
		logging.Fatal("数据库连接失败",err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string)string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)

	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Auth{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB()  {
	defer db.Close()
}
package model

import (
	"fmt"
	"gin_restfull/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func init()  {
	var conn = config.Config.GetString("sqlConn")
	db, err := gorm.Open("mysql", conn)
	if err != nil{
		fmt.Println(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	fmt.Println("db init succeed")
	DB.AutoMigrate(&User{})

}

package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 127.0.0.1:3306 是地址和端口，go_login_demo 是库名
	dsn := "root:root@tcp(127.0.0.1:3306)/go_login_demo?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}

	fmt.Println("数据库连接成功！")
}

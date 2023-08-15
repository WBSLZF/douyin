package model

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	var cfg *ini.File
	cfg, err = ini.Load("./config/dbConfig.ini")
	if err != nil {
		fmt.Printf("Faild to read config file: %v", err)
		os.Exit(1)
	}
	// Reading mysql Config
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		QueryFields: true, //打印sql
	})
	if err != nil {
		fmt.Println(err)
	}
	err = DB.AutoMigrate(&Video{}, &UserInfo{}, &UserLogin{}, &Comment{})
	if err != nil {
		panic(err)
	}
}

package gorm_utilities

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Utility struct {
	dbMain  *gorm.DB
	dbSlave *gorm.DB
	action  string
}

func NewUtility(dbMain, dbSlave *gorm.DB) *Utility {
	return &Utility{dbMain: dbMain, dbSlave: dbSlave}
}

func NewMySQL(addr,user,password,dbName string) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=500s&readTimeout=500s&writeTimeout=500s&parseTime=true&loc=Local&charset=utf8,utf8mb4",
		user, password, addr, dbName))
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(1800 * time.Second)
	db.SingularTable(true) // 禁用默认的让表名为模型名的复数形式
	db.LogMode(true)
	return db
}
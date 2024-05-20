package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"

	"xlsqac/iachina-crm/pkg/setting"
)

var gormDB *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "加载 section 失败：'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		DisableAutomaticPing:   true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println(err)
	}

	db, _ := gormDB.DB()

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

}

//func CloseDB() {
//	defer sqlDB.Close()
//}

package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func InitDatabase(config *MySqlConfig) {
	var err error

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
		config.TimeZone,
	)

	conn := mysql.Open(dsn)
	db, err = gorm.Open(conn, &gorm.Config{})

	if err != nil {
		log.Fatal("数据库链接失败，原因：" + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("数据库链接失败，原因：" + err.Error())
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Second * 5)

	if config.Debug {
		db.Debug()
	}
}


func GetDB() *gorm.DB {
	return db
}
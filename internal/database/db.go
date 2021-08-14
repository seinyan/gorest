package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func NewDBConn(dsn string) (*gorm.DB, error) {

	//dsn := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
	//	conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SslMode, conf.TimeZone,
	//)

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
	}

	conn, err := DB.DB()
	//conn.Ping()

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	return DB, nil
}

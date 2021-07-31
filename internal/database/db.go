package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func NewDBConn() (*gorm.DB, error) {

	var err error
	//var DB *gorm.DB


	//dsn := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	c.Host, c.Port, c.User, c.Password, c.DbName)
	dsn := "host=localhost user=postgres password=8888 dbname=gorest port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
	}


	conn, err := DB.DB()
	conn.Ping()

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	return DB, nil
}

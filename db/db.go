package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接配置
const (
	user     = "root"
	password = "root"
	host     = "localhost"
	port     = 3306
	dbname   = "public"
)

// InitDB 初始化 MySQL 数据库连接
func InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// fmt.Println("db connected successfully!")
	return db, nil
}

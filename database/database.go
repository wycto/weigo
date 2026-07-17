package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库封装
type DB struct {
	*gorm.DB
	Prefix string
}

// Connect 连接数据库
func Connect(cfg DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("数据库连接成功")
	return &DB{DB: db, Prefix: cfg.Prefix}, nil
}

// TableName 自动加前缀
func (db *DB) TableName(tableName string) *gorm.DB {
	return db.Table(db.Prefix + tableName)
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
	Prefix   string
}

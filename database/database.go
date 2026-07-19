package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB 数据库封装
type DB struct {
	*gorm.DB
	Prefix string
}

// Connect 连接数据库
func Connect(cfg DBConfig) (*DB, error) {
	var dialector gorm.Dialector

	switch cfg.Driver {
	case "sqlite":
		dialector = sqlite.Open(cfg.Database)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
		dialector = mysql.Open(dsn)
	case "postgres":
		var dsn string
		if cfg.Password != "" {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
				cfg.Username, cfg.Host, cfg.Port, cfg.Database)
		}
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 配置连接池 (SQLite不需要连接池配置)
	if cfg.Driver != "sqlite" {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

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

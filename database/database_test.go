package database

import (
	"testing"
)

func TestConnectUnsupportedDriver(t *testing.T) {
	_, err := Connect(DBConfig{
		Driver:   "oracle",
		Host:     "localhost",
		Port:     "1521",
		Username: "test",
		Password: "test",
		Database: "testdb",
	})
	if err == nil {
		t.Error("Connect() with unsupported driver should return error")
	}
	if err.Error() != "unsupported database driver: oracle" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConnectSQLite(t *testing.T) {
	db, err := Connect(DBConfig{
		Driver:   "sqlite",
		Database: ":memory:",
	})
	if err != nil {
		t.Fatalf("Connect() with sqlite error: %v", err)
	}
	if db == nil {
		t.Fatal("db should not be nil")
	}
	if db.DB == nil {
		t.Fatal("db.DB should not be nil")
	}
}

func TestDBPrefix(t *testing.T) {
	// 验证 DB 结构体 Prefix 字段
	db := &DB{Prefix: "wp_"}
	if db.Prefix != "wp_" {
		t.Errorf("Prefix = %q, want 'wp_'", db.Prefix)
	}
}

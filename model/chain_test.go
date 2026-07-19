package model

import (
	"testing"

	"github.com/wycto/weigo/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *database.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	return &database.DB{DB: db, Prefix: ""}
}

type TestUser struct {
	BaseModel
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status int    `json:"status"`
}

func (TestUser) TableName() string {
	return "users"
}

func TestNewChain(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	if chain == nil {
		t.Fatal("NewChain() returned nil")
	}
	if chain.table != "users" {
		t.Errorf("table = %q, want 'users'", chain.table)
	}
}

func TestChainFields(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	result := chain.Fields("id,name")
	if result != chain {
		t.Error("Fields() should return self for chaining")
	}
	if chain.fields != "id,name" {
		t.Errorf("fields = %q, want 'id,name'", chain.fields)
	}
}

func TestChainWhere(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Where("status = ?", 1).Where("age > ?", 18)
	if len(chain.where) != 2 {
		t.Errorf("len(where) = %d, want 2", len(chain.where))
	}
	if chain.where[0] != "status = ?" {
		t.Errorf("where[0] = %q", chain.where[0])
	}
}

func TestChainOrder(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Order("id desc")
	if chain.order != "id desc" {
		t.Errorf("order = %q, want 'id desc'", chain.order)
	}
}

func TestChainGroup(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Group("status")
	if chain.group != "status" {
		t.Errorf("group = %q, want 'status'", chain.group)
	}
}

func TestChainHaving(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Having("count(*) > 1")
	if chain.having != "count(*) > 1" {
		t.Errorf("having = %q", chain.having)
	}
}

func TestChainLimit(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Limit(10)
	if chain.limit != 10 {
		t.Errorf("limit = %d, want 10", chain.limit)
	}
}

func TestChainPage(t *testing.T) {
	db := setupTestDB(t)
	chain := NewChain(db, "users")
	chain.Page(2, 20)
	if chain.page != 2 {
		t.Errorf("page = %d, want 2", chain.page)
	}
	if chain.pageSize != 20 {
		t.Errorf("pageSize = %d, want 20", chain.pageSize)
	}
}

func TestChainInsertAndFind(t *testing.T) {
	db := setupTestDB(t)

	// 自动迁移
	db.AutoMigrate(&TestUser{})

	// 插入数据
	user := &TestUser{Name: "Alice", Email: "alice@test.com", Status: 1}
	err := NewChain(db, "users").Insert(user)
	if err != nil {
		t.Fatalf("Insert() error: %v", err)
	}

	// 查询
	var users []TestUser
	err = NewChain(db, "users").Find(&users)
	if err != nil {
		t.Fatalf("Find() error: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Find() returned %d users, want 1", len(users))
	}
	if users[0].Name != "Alice" {
		t.Errorf("users[0].Name = %q, want 'Alice'", users[0].Name)
	}
}

func TestChainFindOne(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	NewChain(db, "users").Insert(&TestUser{Name: "Bob", Email: "bob@test.com", Status: 1})
	NewChain(db, "users").Insert(&TestUser{Name: "Carol", Email: "carol@test.com", Status: 0})

	var user TestUser
	err := NewChain(db, "users").Where("name = ?", "Bob").FindOne(&user)
	if err != nil {
		t.Fatalf("FindOne() error: %v", err)
	}
	if user.Name != "Bob" {
		t.Errorf("FindOne().Name = %q, want 'Bob'", user.Name)
	}
}

func TestChainCount(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	NewChain(db, "users").Insert(&TestUser{Name: "A", Status: 1})
	NewChain(db, "users").Insert(&TestUser{Name: "B", Status: 1})
	NewChain(db, "users").Insert(&TestUser{Name: "C", Status: 0})

	var count int64
	err := NewChain(db, "users").Where("status = ?", 1).Count(&count)
	if err != nil {
		t.Fatalf("Count() error: %v", err)
	}
	if count != 2 {
		t.Errorf("Count() = %d, want 2", count)
	}
}

func TestChainUpdate(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	NewChain(db, "users").Insert(&TestUser{Name: "Dave", Email: "dave@test.com", Status: 0})

	err := NewChain(db, "users").Where("name = ?", "Dave").Update(map[string]interface{}{"status": 1})
	if err != nil {
		t.Fatalf("Update() error: %v", err)
	}

	var user TestUser
	NewChain(db, "users").Where("name = ?", "Dave").FindOne(&user)
	if user.Status != 1 {
		t.Errorf("Status = %d, want 1 after update", user.Status)
	}
}

func TestChainDelete(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	NewChain(db, "users").Insert(&TestUser{Name: "Eve", Status: 1})

	err := NewChain(db, "users").Where("name = ?", "Eve").Delete()
	if err != nil {
		t.Fatalf("Delete() error: %v", err)
	}

	var count int64
	NewChain(db, "users").Count(&count)
	if count != 0 {
		t.Errorf("Count() after delete = %d, want 0", count)
	}
}

func TestChainPagePagination(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	for i := 0; i < 10; i++ {
		NewChain(db, "users").Insert(&TestUser{Name: "User" + string(rune('A'+i)), Status: 1})
	}

	// 第1页，每页3条
	var users []TestUser
	err := NewChain(db, "users").Page(1, 3).Find(&users)
	if err != nil {
		t.Fatalf("Find() error: %v", err)
	}
	if len(users) != 3 {
		t.Errorf("Page 1 returned %d users, want 3", len(users))
	}

	// 第2页
	var users2 []TestUser
	err = NewChain(db, "users").Page(2, 3).Find(&users2)
	if err != nil {
		t.Fatalf("Find() error: %v", err)
	}
	if len(users2) != 3 {
		t.Errorf("Page 2 returned %d users, want 3", len(users2))
	}

	// 第4页（只有1条）
	var users3 []TestUser
	err = NewChain(db, "users").Page(4, 3).Find(&users3)
	if err != nil {
		t.Fatalf("Find() error: %v", err)
	}
	if len(users3) != 1 {
		t.Errorf("Page 4 returned %d users, want 1", len(users3))
	}
}

func TestChainChaining(t *testing.T) {
	db := setupTestDB(t)
	db.AutoMigrate(&TestUser{})

	for i := 0; i < 5; i++ {
		NewChain(db, "users").Insert(&TestUser{
			Name:   "User" + string(rune('A'+i)),
			Status: i % 2,
		})
	}

	// 链式查询：status=1, 排序, 分页
	var users []TestUser
	err := NewChain(db, "users").
		Where("status = ?", 1).
		Order("id desc").
		Page(1, 2).
		Find(&users)
	if err != nil {
		t.Fatalf("Chained Find() error: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Chained query returned %d users, want 2", len(users))
	}
}

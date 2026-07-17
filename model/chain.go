package model

import (
	"github.com/wycto/weigo/database"
	"gorm.io/gorm"
)

// ChainQuery 链式查询（保留 PHP 风格）
type ChainQuery struct {
	db       *database.DB
	table    string
	fields   string
	where    []string
	args     []interface{}
	order    string
	group    string
	having   string
	limit    int
	offset   int
	page     int
	pageSize int
}

// NewChain 创建链式查询
func NewChain(db *database.DB, table string) *ChainQuery {
	return &ChainQuery{
		db:    db,
		table: table,
	}
}

// Fields 设置查询字段（替代 SetFields）
func (c *ChainQuery) Fields(fields string) *ChainQuery {
	c.fields = fields
	return c
}

// Where 添加条件
func (c *ChainQuery) Where(condition string, args ...interface{}) *ChainQuery {
	c.where = append(c.where, condition)
	c.args = append(c.args, args...)
	return c
}

// Order 排序
func (c *ChainQuery) Order(order string) *ChainQuery {
	c.order = order
	return c
}

// Group 分组
func (c *ChainQuery) Group(group string) *ChainQuery {
	c.group = group
	return c
}

// Having 聚合条件
func (c *ChainQuery) Having(having string) *ChainQuery {
	c.having = having
	return c
}

// Limit 限制条数
func (c *ChainQuery) Limit(limit int) *ChainQuery {
	c.limit = limit
	return c
}

// Page 分页
func (c *ChainQuery) Page(page, pageSize int) *ChainQuery {
	c.page = page
	c.pageSize = pageSize
	return c
}

// Find 查询多条
func (c *ChainQuery) Find(result interface{}) error {
	query := c.buildQuery()
	return query.Find(result).Error
}

// FindOne 查询单条
func (c *ChainQuery) FindOne(result interface{}) error {
	query := c.buildQuery()
	return query.First(result).Error
}

// Insert 新增
func (c *ChainQuery) Insert(data interface{}) error {
	return c.db.Table(c.table).Create(data).Error
}

// Update 更新
func (c *ChainQuery) Update(data map[string]interface{}) error {
	query := c.buildQuery()
	return query.Updates(data).Error
}

// Delete 删除
func (c *ChainQuery) Delete() error {
	query := c.buildQuery()
	return query.Delete(nil).Error
}

// Count 计数
func (c *ChainQuery) Count(result *int64) error {
	query := c.buildQuery()
	return query.Count(result).Error
}

// buildQuery 构建查询
func (c *ChainQuery) buildQuery() *gorm.DB {
	query := c.db.Table(c.table)

	if c.fields != "" {
		query = query.Select(c.fields)
	}

	for i, w := range c.where {
		if i == 0 {
			query = query.Where(w, c.args...)
		} else {
			query = query.Where(w, c.args[i])
		}
	}

	if c.order != "" {
		query = query.Order(c.order)
	}
	if c.group != "" {
		query = query.Group(c.group)
	}
	if c.having != "" {
		query = query.Having(c.having)
	}
	if c.page > 0 && c.pageSize > 0 {
		offset := (c.page - 1) * c.pageSize
		query = query.Offset(offset).Limit(c.pageSize)
	} else if c.limit > 0 {
		query = query.Limit(c.limit)
	}

	return query
}

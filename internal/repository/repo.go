package repository

import (
	"context"
	"fmt"

	"github.com/lllllan02/iam/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(conf *config.Config) *Repo {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/iam?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Repo{db: db}
}

type contextTxKey struct{}

func (d *Repo) InTx(c context.Context, fn func(c context.Context) error) error {
	return d.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		c = context.WithValue(c, contextTxKey{}, tx)
		return fn(c)
	})
}

func (d *Repo) DB(c context.Context) *gorm.DB {
	if tx, ok := c.Value(contextTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return d.db.WithContext(c)
}

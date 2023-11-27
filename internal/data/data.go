package data

import (
	"context"
	"fmt"
	"time"

	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type contextKey string

const contextTxKey contextKey = "ctxTxKey"

type Data struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewData(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Data {
	return &Data{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

type Transaction interface {
	Transaction(c context.Context, fn func(c context.Context) error) error
}

func NewTransaction(d *Data) Transaction {
	return d
}

// DB return tx
// If you need to create a Transaction, you must call DB(c) and Transaction(c, fn)
func (d *Data) DB(c context.Context) *gorm.DB {
	v := c.Value(contextTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return d.db.WithContext(c)
}

func (d *Data) Transaction(c context.Context, fn func(c context.Context) error) error {
	return d.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		c = context.WithValue(c, contextTxKey, tx)
		return fn(c)
	})
}

func NewDB(conf *config.Config, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		conf.Data.Mysql.User,
		conf.Data.Mysql.Password,
		conf.Data.Mysql.Addr,
		conf.Data.Mysql.DBName,
	)

	// 上海时区
	dsn = dsn + "&loc=Asia%2FShanghai"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}

	// debug mode
	if conf.Data.Mysql.Debug {
		db = db.Debug()
	}

	// auto migrate if test env
	if conf.IsTestEnv() {
		var tables = []interface{}{
			&model.User{},
		}

		if err := db.Migrator().DropTable(tables...); err != nil {
			panic(err)
		}

		if err := db.AutoMigrate(tables...); err != nil {
			panic(err)
		}
	}

	return db
}

func NewRedis(conf *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Data.Redis.Addr,
		Password: conf.Data.Redis.Password,
		DB:       conf.Data.Redis.DB,
	})

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(c).Result(); err != nil {
		panic(err)
	}
	return rdb
}

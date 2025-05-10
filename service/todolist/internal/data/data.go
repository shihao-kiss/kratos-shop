package data

import (
	"todolist/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewGreeterRepo, NewTodoListRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: c,
	}, cleanup, nil
}

// NewDB 创建数据库连接
func NewDB(c *conf.Data) *gorm.DB {
	switch c.Database.Driver {
	case "mysql":
		dsn := c.Database.Source
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db
	default:
		panic("unsupported database driver: " + c.Database.Driver)
	}
}

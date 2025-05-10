package data

import (
	"todolist/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var MysqlProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type MysqlData struct {
	// TODO wrapped database client
}

// NewData .
func NewMysqlData(c *conf.Data, logger log.Logger) (*MysqlData, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &MysqlData{}, cleanup, nil
}

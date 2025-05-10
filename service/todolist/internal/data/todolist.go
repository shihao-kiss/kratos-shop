package data

import (
	"context"

	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todoListRepo struct {
	data *MysqlData
	log  *log.Helper
}

// NewGreeterRepo .
func NewTodoListRepo(data *MysqlData, logger log.Logger) biz.TodoListRepo {
	return &todoListRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todoListRepo) Save(ctx context.Context, g *biz.TodoList) (*biz.TodoList, error) {
	r.log.Infof("Save: %+v", g)
	return g, nil
}
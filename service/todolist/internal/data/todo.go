package data

import (
	"context"

	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todoListRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewTodoListRepo(data *Data, logger log.Logger) biz.TodoListRepo {
	return &todoListRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todoListRepo) Save(ctx context.Context, g *biz.TodoList) (*biz.TodoList, error) {
	r.log.Infof("Save: %+v", g)
	err := r.data.db.Create(g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *todoListRepo) Update(ctx context.Context, g *biz.TodoList) (*biz.TodoList, error) {
	r.log.Infof("Update: %+v", g)
	err := r.data.db.Model(&biz.TodoList{}).Where("id = ?", g.ID).Updates(g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *todoListRepo) Delete(ctx context.Context, id int64) error {
	r.log.Infof("Delete: %+v", id)
	err := r.data.db.Where("id = ?", id).Delete(&biz.TodoList{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *todoListRepo) FindByID(ctx context.Context, id int64) (*biz.TodoList, error) {
	r.log.Infof("FindByID: %+v", id)
	var todo biz.TodoList
	err := r.data.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoListRepo) ListAll(ctx context.Context) ([]*biz.TodoList, error) {
	r.log.Infof("ListAll")
	var todos []*biz.TodoList
	err := r.data.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
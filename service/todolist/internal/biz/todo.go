package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type TodoList struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string `gorm:"column:title;type:varchar(256);not null"`
	Completed bool   `gorm:"column:completed;not null;default:false"`
}

// 指定表名（可选）
func (TodoList) TableName() string {
	return "todo_list"
}

// GreeterRepo is a Greater repo.
type TodoListRepo interface {
	Save(context.Context, *TodoList) (*TodoList, error)
	Update(context.Context, *TodoList) (*TodoList, error)
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*TodoList, error)
	ListAll(context.Context) ([]*TodoList, error)
}

// GreeterUsecase is a Greeter usecase.
type TodoListUsecase struct {
	repo TodoListRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewTodoListUsecase(repo TodoListRepo, logger log.Logger) *TodoListUsecase {
	return &TodoListUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *TodoListUsecase) CreateTodoList(ctx context.Context, g *TodoList) (*TodoList, error) {
	uc.log.WithContext(ctx).Infof("CreateTodoList: %+v", g)
	return uc.repo.Save(ctx, g)
}

// UpdateTodoList updates a TodoList, and returns the updated TodoList.
func (uc *TodoListUsecase) UpdateTodoList(ctx context.Context, g *TodoList) (*TodoList, error) {
	uc.log.WithContext(ctx).Infof("UpdateTodoList: %+v", g)
	return uc.repo.Update(ctx, g)
}

// DeleteTodoList deletes a TodoList, and returns the deleted TodoList.
func (uc *TodoListUsecase) DeleteTodoList(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteTodoList: %+v", id)
	return uc.repo.Delete(ctx, id)
}

// FindByID finds a TodoList by ID, and returns the TodoList.
func (uc *TodoListUsecase) FindByID(ctx context.Context, id int64) (*TodoList, error) {
	uc.log.WithContext(ctx).Infof("FindByID: %+v", id)
	return uc.repo.FindByID(ctx, id)
}

// ListAll lists all TodoLists, and returns the TodoLists.
func (uc *TodoListUsecase) ListAll(ctx context.Context) ([]*TodoList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	return uc.repo.ListAll(ctx)
}

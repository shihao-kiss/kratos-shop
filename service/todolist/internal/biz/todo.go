package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type TodoList struct {
	ID     int64
	Title  string
	Completed bool
}

// GreeterRepo is a Greater repo.
type TodoListRepo interface {
	Save(context.Context, *TodoList) (*TodoList, error)
	// Update(context.Context, *Greeter) (*Greeter, error)
	// FindByID(context.Context, int64) (*Greeter, error)
	// ListByHello(context.Context, string) ([]*Greeter, error)
	// ListAll(context.Context) ([]*Greeter, error)
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

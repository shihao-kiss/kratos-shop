package service

import (
	"context"
	"errors"

	pb "todolist/api/todolist/v1"
	"todolist/internal/biz"
)

type TodoService struct {
	pb.UnimplementedTodoServer

	uc *biz.TodoListUsecase
}

func NewTodoService(uc *biz.TodoListUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	// 参数校验
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
    // 业务逻辑 
	resp, err := s.uc.CreateTodoList(ctx, &biz.TodoList{
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}
    // 返回值 
	return &pb.CreateTodoReply{
		Todo: &pb.TodoBody{
			Id:        resp.ID,
			Title:     resp.Title,
			Completed: resp.Completed,
		},
	}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoReply, error) {
	return &pb.UpdateTodoReply{}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}

func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}

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
	// 参数校验
	if req.Id == 0 {
		return nil, errors.New("id and title and completed are required")
	}
	// 业务逻辑
	_, err := s.uc.UpdateTodoList(ctx, &biz.TodoList{
		ID:        req.Id,
		Title:     req.Title,
		Completed: req.Completed,
	})
	if err != nil {
		return nil, err
	}
	// 返回值
	return &pb.UpdateTodoReply{}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	// 参数校验
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}
	// 业务逻辑
	err := s.uc.DeleteTodoList(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 返回值
	return &pb.DeleteTodoReply{}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	// 参数校验
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}
	// 业务逻辑
	resp, err := s.uc.FindByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 返回值
	return &pb.GetTodoReply{
		Todo: &pb.TodoBody{
			Id:        resp.ID,
			Title:     resp.Title,
			Completed: resp.Completed,
		},
	}, nil
}

func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	// 业务逻辑
	resp, err := s.uc.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	// 返回值
	var todos []*pb.TodoBody
	for _, todo := range resp {
		todos = append(todos, &pb.TodoBody{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
		})
	}
	return &pb.ListTodoReply{
		Todos: todos,
	}, nil
}

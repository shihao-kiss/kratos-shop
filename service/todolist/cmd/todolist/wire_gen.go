// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"todolist/internal/biz"
	"todolist/internal/conf"
	"todolist/internal/data"
	"todolist/internal/server"
	"todolist/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	todoListRepo := data.NewTodoListRepo(dataData, logger)
	todoListUsecase := biz.NewTodoListUsecase(todoListRepo, logger)
	todoService := service.NewTodoService(todoListUsecase)
	grpcServer := server.NewGRPCServer(confServer, todoService, logger)
	httpServer := server.NewHTTPServer(confServer, todoService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

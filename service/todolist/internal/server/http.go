package server

import (
	"context"
	"fmt"
	"time"

	todolistV1 "todolist/api/todolist/v1"
	"todolist/internal/conf"
	"todolist/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func Middleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 统计耗时
				fmt.Printf("Middleware1 Operation %+v\n", tr.Operation())
				start := time.Now()
				defer func() {
					elapsed := time.Since(start)
					fmt.Printf("Middleware1 Operation %+v elapsed: %s\n", tr.Operation(), elapsed)
				}()
			}
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, todoService *service.TodoService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			Middleware(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts = append(opts, http.ResponseEncoder(responseEncoder))
	opts = append(opts, http.ErrorEncoder(responseErrorEncoder))

	srv := http.NewServer(opts...)
	todolistV1.RegisterTodoHTTPServer(srv, todoService)
	return srv
}

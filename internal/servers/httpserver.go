package servers

import (
	"errors"
	"fmt"
	"graphql-api/internal/graphql"
	routers "graphql-api/internal/router"
	"net/http"
)

var (
	instance *HTTPServer
)

type (
	HTTPServer struct {
		server *http.Server
	}
)

func GetInstance() *HTTPServer {
	if instance == nil {
		instance = &HTTPServer{}

		//gin 路由初始化
		r := routers.DefaultRouter(true)

		//r.Use(middleware.Cors())
		//注册graphql路由

		graphql.RegisterGraphQlRoutes(r, true)

		//开启监听服务端口
		maxHeaderBytes := 1 << 20

		instance.server = &http.Server{
			Addr:    "127.0.0.1:8080",
			Handler: r,
			//ReadTimeout:    time.Duration(60),
			//WriteTimeout:   time.Duration(120),
			MaxHeaderBytes: maxHeaderBytes,
		}

		fmt.Println("[info] start http server listening： ", "127.0.0.1:8080")
	}
	return instance
}

func (hs *HTTPServer) Start() error {
	go func() {
		err := hs.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server.http : start fail", "addr", "127.0.0.1:8080", "error", err)
		}
	}()
	return nil
}

package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"graphql-api/internal/graphql/schema"
	"graphql-api/pkg/base/api_basic"
	"graphql-api/pkg/facility/graphql/handle"
)

type (
	RestHook struct {
	}
)

// Hook hook接口，可自定义需要心跳检测的功能
func (s *RestHook) Hook() bool {
	//fmt.Println("check hook is fly")
	return true
}

// RegisterGraphQlRoutes graphql api gin 路由注册入口
func RegisterGraphQlRoutes(r *gin.Engine, debug bool) *gin.Engine {
	restHook := &RestHook{}
	api_basic.RegisterBaseRoute(r, restHook, debug)

	graphqlG := r.Group("/graphql")

	//注册各模块resolve
	graphqlResolver := RegisterResolver()
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), graphqlResolver)
	handle.GraphqlServeResource(graphqlG, graphqlSchema, true)
	return r
}

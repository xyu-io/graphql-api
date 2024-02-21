package api_basic

import (
	"github.com/gin-gonic/gin"
	"graphql-api/pkg/facility/graphql/playground"
)

type (
	HA struct {
	}
	HookFun interface {
		Hook() bool
		//GetSSOConfig() sso.SConfig
		//GetSystem() base_struct.SystemConfig
		//GetServer() base_struct.ServerConfig
	}
)

func RegisterBaseRoute(router *gin.Engine, HookFun HookFun, debug bool) {
	//注册 check
	apiG := router.Group("/check")
	registerRestfulCheck(apiG, HookFun)

	//sso路由注册
	//sSOConfig := HookFun.GetSSOConfig()
	//system := HookFun.GetSystem()
	//if sSOConfig.Switch == true { // debug 开启时才允许访问iql
	//	ssoG := router.Group("/account/")
	//	sso_restful.RegisterRoute(ssoG, sSOConfig, system.Debug)
	//}

	//swagger路由注册
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if debug { // debug 开启时才允许访问iql
		router.GET("/graphiql", gin.WrapF(playground.Handler("GraphQL", "/graphql/query")))
		router.GET("/playground", gin.WrapF(playground.AltairHandler("Altair", "/graphql/query")))
	}
}

func graphqlHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func registerRestfulCheck(router *gin.RouterGroup, HookFun HookFun) {
	checkR := MustNew(HookFun)
	router.GET("", checkR.Check)
}

package api_basic

import (
	"github.com/gin-gonic/gin"
	"graphql-api/pkg/base/e"
	"graphql-api/pkg/facility/httptool"
	"net/http"
)

var instance *Check

type (
	Check struct {
		HookFun HookFun
	}
)

// MustNew connects clients,
func MustNew(HookFun HookFun) *Check {
	if instance == nil {
		instance = &Check{HookFun: HookFun}
	}
	return instance
}

func (s Check) Check(c *gin.Context) {
	appG := httptool.Gin{C: c}
	if s.HookFun != nil {
		s.HookFun.Hook()
	}
	result := make(map[string]interface{})

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

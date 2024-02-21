package httptool

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"graphql-api/pkg/base/e"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.InvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		return http.StatusBadRequest, e.InvalidParams
	}

	return http.StatusOK, e.SUCCESS
}

package errors

import (
	"github.com/gin-gonic/gin"
	"graphql-api/pkg/facility/logger"
	"testing"
)

func TestErrorHandle(t *testing.T) {
	config := map[string]interface{}{}
	efunc := ErrorHandle(config, logger.NewLogger())

	efunc(&gin.Context{})
}

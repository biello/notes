package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignCheck(ctx *gin.Context) {
	userName := ctx.Param("user")

	if len(userName) == 0 {
		ctx.String(http.StatusUnauthorized, "please login first")
		return
	}
	ctx.Next()
}

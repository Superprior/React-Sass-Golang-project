package headers

import "github.com/gin-gonic/gin"

func DefaultHeaders(ctx *gin.Context, method string) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Access-Control-Allow-Methods", method)
	ctx.Header("Content-Type", "application/json; charset=UTF-8")
}

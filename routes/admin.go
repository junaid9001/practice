package routes

import (
	"github.com/gin-gonic/gin"
)

func Adminroute(r *gin.Engine) {
	admin := r.Group("/admin")

	admin.GET("/users", func(ctx *gin.Context) {
		ctx.String(200, "these are the users")
	})
}

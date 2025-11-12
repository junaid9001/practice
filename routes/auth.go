package routes

import "github.com/gin-gonic/gin"

func Authrouter(r *gin.Engine) {

	auth := r.Group("/auth")

	auth.GET("/signup", func(c *gin.Context) {
		c.String(200, "please signup")
	})

	auth.GET("/login", func(c *gin.Context) {
		c.String(200, "please signup")
	})

}

package apiv1

import (
	"gin-rest-api-sample/api/v1.0/auth"
	"gin-rest-api-sample/api/v1.0/posts"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		posts.ApplyRoutes(v1)
	}
}

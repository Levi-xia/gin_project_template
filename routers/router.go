package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/app/controller"
	"project/middleware"
)

func SetRouter(r *gin.Engine) {

	commentRouter := r.Group("/comment")
	{
		commentRouter.GET("/getCommentById/:id", middleware.NeedLogin, controller.GetCommentById)
		commentRouter.GET("/addComment", middleware.CheckLogin, controller.AddComment)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "no router config",
			"data":    "",
		})
	})
}

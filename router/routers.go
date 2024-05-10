package router

import (
	"net/http"

	"github.com/CyberMidori/gin-ranking/controllers"
	"github.com/CyberMidori/gin-ranking/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 日志中间件
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)

		user.POST("/list", controllers.UserController{}.GetList)

		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Add User")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Delete User")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}

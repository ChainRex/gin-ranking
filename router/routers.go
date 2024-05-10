package router

import (
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
		user.GET("/info/:id", controllers.UserController{}.GetUserById)

		user.GET("/list", controllers.UserController{}.GetUserList)

		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)

		user.POST("/delete", controllers.UserController{}.DeleteUser)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}

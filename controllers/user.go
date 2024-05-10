package controllers

import (
	"github.com/CyberMidori/gin-ranking/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	logger.Write("日志信息", "user")
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(c, 4004, num3)
}

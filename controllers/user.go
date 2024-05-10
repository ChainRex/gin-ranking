package controllers

import (
	"strconv"

	"github.com/CyberMidori/gin-ranking/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserById(c *gin.Context) {
	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserById(id)
	ReturnSuccess(c, 0, "success", user, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.PostForm("username")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "添加用户失败")
		return
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	idStr := c.PostForm("id")
	username := c.PostForm("username")

	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4003, "删除用户失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", true, 1)
}

func (u UserController) GetUserList(c *gin.Context) {
	users, err := models.GetUserList()
	if err != nil {
		ReturnError(c, 4004, "获取用户列表失败")
		return
	}
	ReturnSuccess(c, 0, "获取成功", users, int64(len(users)))
}

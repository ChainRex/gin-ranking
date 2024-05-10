package controllers

import (
	"strconv"

	"github.com/CyberMidori/gin-ranking/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	// 接受用户名 密码 确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "用户名或密码不能为空")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码不一致")
		return
	}

	// 判断用户是否存在
	user, err := models.GetUserInfoByUsername(username)
	if user.Id > 0 {
		ReturnError(c, 4001, "用户已存在")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "保存失败，请联系管理员")
		return
	}

	ReturnSuccess(c, 200, "注册成功", nil, 0)
}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4001, "用户名或密码不能为空")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}

	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}
	if err != nil {
		ReturnError(c, 4004, "登录失败，请联系管理员")
		return
	}

	// 保存到session中
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Username}

	ReturnSuccess(c, 0, "登录成功", data, 1)
}

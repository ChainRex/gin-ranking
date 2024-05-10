package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	search := Search{}
	err := c.BindJSON(&search)
	if err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	ReturnSuccess(c, 0, search.Name, search.Cid, 1)

}

package controllers

import (
	"strconv"

	"github.com/CyberMidori/gin-ranking/models"
	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoById(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfoById(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "投票选手不存在")
		return
	}

	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "已经投过票了")
		return
	}

	rs, err := models.AddVote(userId, playerId)
	if err != nil {
		ReturnError(c, 4004, "投票失败，请联系管理员")
		return
	}
	models.UpdatePlayerScore(playerId)
	ReturnSuccess(c, 0, "投票成功", rs, 1)
}

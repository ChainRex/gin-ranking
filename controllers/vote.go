package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/CyberMidori/gin-ranking/cache"
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

	// 更新redis
	// 新的数据
	player, _ = models.GetPlayerInfoById(playerId)
	playerJson, _ := json.Marshal(player)
	score := player.Score

	// 递增前
	pastPlayer := player
	pastPlayer.Score = player.Score - 1
	pastPlayerJson, _ := json.Marshal(pastPlayer)

	var redisKey string
	redisKey = "ranking:" + strconv.Itoa(player.Aid)
	// 先删掉原来的
	cache.Rdb.ZRem(cache.Rctx, redisKey, string(pastPlayerJson))
	// 再添加新的
	cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(playerJson, score))

	ReturnSuccess(c, 0, "投票成功", rs, 1)
}

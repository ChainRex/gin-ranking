package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/CyberMidori/gin-ranking/cache"
	"github.com/CyberMidori/gin-ranking/models"
	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {

	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	// redis有序集合速度更快
	var redisKey string
	redisKey = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		// redis中有数据
		var players []models.Player
		for _, value := range rs {
			var player models.Player
			json.Unmarshal([]byte(value), &player)
			players = append(players, player)
		}

		ReturnSuccess(c, 0, "success", players, 1)
		return
	}

	// redis中没有数据，从数据库中获取

	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		// 先删除redis中的数据
		cache.Rdb.Del(cache.Rctx, redisKey)

		for _, value := range rsDb {
			// value 转json
			valueJson, _ := json.Marshal(value)
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(valueJson, value.Score)).Err()
		}
		// 设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 0, "success", rsDb, 1)
		return
	}
	ReturnError(c, 4004, "没有相关信息")

}

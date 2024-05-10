package models

import (
	"github.com/CyberMidori/gin-ranking/dao"
	"github.com/jinzhu/gorm"
)

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func GetPlayers(aid int) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Find(&players).Error
	return players, err
}

func GetPlayerInfoById(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

func UpdatePlayerScore(id int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	return err
}

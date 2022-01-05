package models

import (
	"gorm.io/gorm"
	"github.com/cyanvestige/game-management-in-golang/config"
	"time"
)

var db *gorm.DB

type Game struct {
	gorm.Model
	Name string `json:"name"`
	Genre string `json:"genre"`
	Publisher string `json:"publisher"`
	ReleaseDate time.Time `json:"releasedate"`
}

func init()  {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Game{})
}

func (g *Game) CreateGame() *Game{
	db.Create(&g)
	return g
}

func GetAllGames() []Game{
	var Games []Game
	db.Find(&Games)
	return Games
}

func GetGameById(ID int64) (*Game, *gorm.DB){
	var getGame Game
	db := db.Where("ID=?", ID).Find(&getGame)
	return &getGame, db
}

func DeleteGame(ID int64) Game{
	var game Game
	db.Where("ID=?", ID).Delete(&game)
	return game
}
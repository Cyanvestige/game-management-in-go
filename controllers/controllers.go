package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/cyanvestige/game-management-in-golang/utils"
	"github.com/cyanvestige/game-management-in-golang/models"
)

var NewGame models.Game

func GetGame(w http.ResponseWriter, r *http.Request){
	newGames := models.GetAllGames()
	res, _ := json.Marshal(newGames)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGameById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	ID, err := strconv.ParseInt(gameId,0,0)
	if err != nil{
		fmt.Println("Parsing Error")
	}
	gameInfo, _ := models.GetGameById(ID)
	res, _ := json.Marshal(gameInfo)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateGame(w http.ResponseWriter, r *http.Request){
	createGame := &models.Game{}
	utils.ParseBody(r, createGame)
	b := createGame.CreateGame()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	ID, err := strconv.ParseInt(gameId,0,0)
	if err != nil{
		fmt.Println("Parsing Error")
	}
	deletedGame := models.DeleteGame(ID)
	res, _ := json.Marshal(deletedGame)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateGame(w http.ResponseWriter, r *http.Request){
	updateGame := &models.Game{}
	utils.ParseBody(r, updateGame)
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	ID, err := strconv.ParseInt(gameId,0,0)
	if err != nil{
		fmt.Println("Parsing Error!")
	}
	gameInfo, db := models.GetGameById(ID)
	if updateGame.Name != ""{
		gameInfo.Name = updateGame.Name
	}
	if updateGame.Genre != ""{
		gameInfo.Genre = updateGame.Genre
	}
	if updateGame.Publisher != ""{
		gameInfo.Publisher = updateGame.Publisher
	}
	if !updateGame.ReleaseDate.IsZero(){
		gameInfo.ReleaseDate = updateGame.ReleaseDate
	}
	db.Save(&gameInfo)
	res, _ := json.Marshal(gameInfo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
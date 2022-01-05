package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/cyanvestige/game-management-in-golang/router"
)

func main(){
	r := mux.NewRouter()
	router.RegisterRoutes(r)
	http.Handle("/", r)
}
package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"github.com/cyanvestige/game-management-in-golang/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterGameManagementRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5000"), r)
}
package main

import (
	_"os"
	"net/http"
	"flag"
	"fmt"

	"./handler"
	"./config"
	"./db"
	_"./vkservice"
	_"./model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	option = flag.String("o", "", "Choose option: start")
)

func main() {
	flag.Parse()

	config.LoadConfiguration("./config.json")
	

	switch op := *option; op {
	case "start":
		startServer()
	case "migrate":
		migrateDB()
	default:
		fmt.Println("You need to choose option")
	}
}

func startServer() {
	db.ConnectToDB()
	r := mux.NewRouter()

	r.Handle("/", handler.StatusHandler).Methods("GET")
	r.Handle("/upload_image", handler.UploadImage).Methods("POST")
	r.Handle("/auth", handler.Authorize).Methods("GET")
	r.Handle("/logout", handler.Logout).Methods("GET")
	r.Handle("/verify", handler.VerifySession).Methods("GET")
	r.Handle("/get_images", handler.GetImages).Methods("GET")
	
	http.ListenAndServe(":8002", handlers.CORS()(r))
}


func migrateDB() {
	db.Migrate()
}

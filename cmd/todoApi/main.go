package main

import (
	"fmt"
	"log"
	"os"
	"ovhTest/app"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	var db *gorm.DB
	var dbType string

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := &app.OVHTestServer{
		DbType:      dbType,
		Router:      InitialiseRouter(),
		TodosStore:  app.NewTodoStore(),
		AllowOrigin: "*",
	}
	if dbType := os.Getenv("DB_TYPE"); dbType == "postgrsql" {
		if db, err = app.OpenDbPostgres(os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBNAME"), os.Getenv("USER"), os.Getenv("PASSWORD")); err != nil {
			log.Fatal("Error opening DB")
		}
		server.Database = db
	}

	app.SetServer(server)
	app.InitRouter(server.Router)
	fmt.Println("Serveur running:" + os.Getenv("LISTEN_PORT"))
	server.ListenAndServe(os.Getenv("LISTEN_PORT"))
}

package main

import (
	"database/sql"
	"fmt"
	"os"
	"financial-journey/database"
	"financial-journey/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)


func main(){

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Loaded .env file")
	}


	psqlInfo := "host=" + os.Getenv("PGHOST") + " port=" + os.Getenv("PGPORT") + " user=" + os.Getenv("PGUSER") + " password=" + os.Getenv("PGPASSWORD") + " dbname=" + os.Getenv("PGDATABASE") + " sslmode=disable"
	DB, err = sql.Open("postgres", psqlInfo)

	err = DB.Ping()
	if err != nil{
		fmt.Println("Error connecting to database")
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}
	database.DbMigrate(DB)
	defer DB.Close()	

	router := gin.Default()
	routes.SetupRouter(router)
	
}
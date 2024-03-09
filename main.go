package main

import (
	"database/sql"
	"fmt"
	"quiz3/routers"

	// "miniproject/controllers"
	"quiz3/database"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	} else {
		fmt.Println("success read file environtment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connected Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.ConnectionDatabase(DB)

	defer DB.Close()

	// var PORT = ":8080"
	routers.StartServer().Run("PORT")
}

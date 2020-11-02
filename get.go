package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Car struct {
	ID         int    `json:"id"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	HorsePower string `json:"horsepower"`
}

var db *sql.DB
var err error

func connect() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load env file! %v", err)
	}
	fmt.Println(os.Getenv("DB_HOST"))
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = sql.Open("mysql", DBURL)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully Connected to MySQL database!")
}

func getCar(w http.ResponseWriter, r *http.Request) {
	connect()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, brand, model, horse_power FROM cars WHERE id = ?", params["ident"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	defer db.Close()
	var car Car

	for result.Next() {
		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.HorsePower)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(car)
	fmt.Println("Successfully Selected from myMariadb database")
}

func handleRequests() {
	fmt.Println("You are using my-get-go-app service!")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter := mux.NewRouter()
	myRouter.HandleFunc("/service/v1/cars/{ident}", getCar).Methods("GET")
	//pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequests()
}

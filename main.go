package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// --------------------------------------------------------------------
// structs
// --------------------------------------------------------------------
type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Msg struct {
	Status string
	Msg    string
}

// --------------------------------------------------------------------
// vars and configuration
// --------------------------------------------------------------------

// server port
var serverport string = ":8080"

// logging
var logging bool = true

// database
var dbdriver string = "mysql"
var dbuser string = "test"
var dbpass string = "test"
var dbname string = "test"

//var dbhost string = "127.0.0.1"

// --------------------------------------------------------------------
// logger
// --------------------------------------------------------------------
func logger(msg any) {
	if logging {
		log.Println(msg)
	}
}

// --------------------------------------------------------------------
// database
// --------------------------------------------------------------------
func connDB() (connDB *sql.DB) {

	//var connstring string = dbuser + ":" + dbpass + "@" + dbhost + "/" + dbname
	var connstring string = dbuser + ":" + dbpass + "@/" + dbname
	connDB, err := sql.Open(dbdriver, connstring)

	if err != nil {
		panic(err.Error())
	}
	return connDB
}

// --------------------------------------------------------------------
// Main
// --------------------------------------------------------------------
func main() {
	//fmt.Println("Hello World!")

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/items", ListItemsHandler).Methods("GET")
	router.HandleFunc("/api/item/{id}", GetItemHandler).Methods("GET")
	router.HandleFunc("/api/item", CreateItemHandler).Methods("POST")
	router.HandleFunc("/api/item/{id}", UpdateItemHandler).Methods("POST")
	router.HandleFunc("/api/item/{id}", DeleteItemHandler).Methods("DELETE")

	// start the server
	logger("Server running on port " + serverport + ".")
	http.ListenAndServe(serverport, router)
}

// --------------------------------------------------------------------
// handlers
// --------------------------------------------------------------------

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	/*
		var response Msg
		response.Status = "Ok"
		response.Msg = "It works!"
		json.NewEncoder(w).Encode(response)
	*/

	json.NewEncoder(w).Encode(&Msg{"Ok", "It works!"})
}

func ListItemsHandler(w http.ResponseWriter, r *http.Request) {

	// response
	itemArray := GetList()
	json.NewEncoder(w).Encode(itemArray)

}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	connDB := connDB()
	recordset, err := connDB.Query("SELECT * FROM items where id=?", id)

	if err != nil {
		panic(err.Error())
	}

	item := Item{}
	itemArray := []Item{}

	for recordset.Next() {
		var id int
		var name, description string
		err = recordset.Scan(&id, &name, &description)

		if err != nil {
			panic(err.Error())
		}

		item.Id = id
		item.Name = name
		item.Description = description

		itemArray = append(itemArray, item)
	}

	logger("Item list:")
	logger(itemArray)

	// response
	json.NewEncoder(w).Encode(itemArray)

}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	logger("Insert item:")
	logger(item)

	connDB := connDB()
	insertRecords, err := connDB.Prepare("INSERT INTO items (name, description) VALUES (?,?)")

	if err != nil {
		panic(err.Error())
	}

	// insert
	insertRecords.Exec(item.Name, item.Description)

	// response
	itemArray := GetList()
	json.NewEncoder(w).Encode(itemArray)

}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	logger("Updating id " + id + ": ")
	logger(item)

	connDB := connDB()
	updateRecords, err := connDB.Prepare("UPDATE items SET name=?, description=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	// update
	updateRecords.Exec(item.Name, item.Description, id)

	// response
	itemArray := GetList()
	json.NewEncoder(w).Encode(itemArray)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	logger("Deleting id: " + id)

	connDB := connDB()
	deleteRecords, err := connDB.Prepare("DELETE FROM items WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	// delete
	deleteRecords.Exec(id)

	// response
	itemArray := GetList()
	json.NewEncoder(w).Encode(itemArray)
}

// --------------------------------------------------------------------
// Get item list
// --------------------------------------------------------------------
func GetList() []Item {
	connDB := connDB()
	recordset, err := connDB.Query("SELECT * FROM items")

	if err != nil {
		panic(err.Error())
	}

	item := Item{}
	itemArray := []Item{}

	for recordset.Next() {
		var id int
		var name, description string
		err = recordset.Scan(&id, &name, &description)

		if err != nil {
			panic(err.Error())
		}

		item.Id = id
		item.Name = name
		item.Description = description

		itemArray = append(itemArray, item)
	}

	logger("Item list:")
	logger(itemArray)

	return itemArray
}

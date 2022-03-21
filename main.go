package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilhamabdlh/simple_rest/helper"
	"github.com/ilhamabdlh/simple_rest/transfer"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/labstack/echo"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var books []models.Book

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book models.Book
		// & character returns the memory address of the following variable.
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books) // encode similar to serialize process.
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var book models.Account
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	e := echo.Context
	id:= e.Param("account_number")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"account_number": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Account

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

}

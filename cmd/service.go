package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"fmt"
	"errors"

	"github.com/ilhamabdlh/simple_rest/helper"
	"github.com/ilhamabdlh/simple_rest/transfer"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)
func CheckAccount(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("ini w: ", w)
	// fmt.Println("ini r: ", r)
	w.Header().Set("Content-Type", "application/json")
	var collection = helper.ConnectDB("check")
	var book transfer.Account
	var id = strings.Split(r.URL.Path, "/")
	filter := bson.M{"account_number": id[len(id)-1]}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	// fmt.Println(err)

	if err != nil {
		helper.GetError(errors.New("account not found"), w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collection = helper.ConnectDB("check")
	var acc transfer.Account
	_ = json.NewDecoder(r.Body).Decode(&acc)
	result, err := collection.InsertOne(context.TODO(), acc)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func Transfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var connect = helper.ConnectDB("check")
	var acc transfer.Account
	var id = strings.Split(r.URL.Path, "/")

	var trf transfer.Transfer
	_= json.NewDecoder(r.Body).Decode(&trf)
	
	filter := bson.M{"account_number": id[len(id)-2]}
	plus := bson.M{"account_number": trf.ToAccountNumber}
	var dcc transfer.Account

	err := connect.FindOne(context.TODO(), filter).Decode(&acc)
	rr := connect.FindOne(context.TODO(), plus).Decode(&dcc)
	
	// fmt.Println("ini err", err)
	
	if err != nil || rr != nil  {
		helper.GetError(errors.New("account not found"), w)	
		defer r.Body.Close()
	} else if acc.Balance < trf.Amount{
		helper.GetError(errors.New("no balance"), w)
	} else{
		acc.Balance -= trf.Amount
		dcc.Balance += trf.Amount
		json.NewEncoder(w).Encode(acc)
	}
	// fmt.Println("ballance1 : ", acc.Balance)
	
    update := bson.D{{"$set",
        bson.D{
            {"balance", acc.Balance},
        },
    }}
	_ = connect.FindOneAndUpdate(context.TODO(), filter, update).Decode(&acc)
	// fmt.Println("ballance2 : ", dcc.Balance)
	updated := bson.D{{"$set",
        bson.D{
            {"balance", dcc.Balance},
        },
    }}
	_ = connect.FindOneAndUpdate(context.TODO(), plus, updated).Decode(&dcc)
}
func main() {
	
	r := mux.NewRouter()

	r.HandleFunc("/account/{account_number}", CheckAccount).Methods("GET")
	r.HandleFunc("/account", createAccount).Methods("POST")
	r.HandleFunc("/account/{from_account_number}/transfer", Transfers).Methods("POST")

	fmt.Println("start on: 3001")
	log.Fatal(http.ListenAndServe(":3001", r))

}

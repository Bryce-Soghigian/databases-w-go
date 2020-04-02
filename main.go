package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "net/http"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)


func get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/",get).Methods(http.MethodGet)
	PORT := ":8080"
	fmt.Println("//======Server is running on port",PORT)
	log.Fatal(http.ListenAndServe(PORT,router))
}
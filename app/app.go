package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(){
	router := mux.NewRouter()
	router.HandleFunc("/api/time",getLocalTime).Methods(http.MethodGet)
	router.HandleFunc("/api/time/{tz}",getTime).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:8001",router))
}
package app

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

type times struct {
	CurrTime time.Time `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	varstr := fmt.Sprint(vars)
	loc, _ := time.LoadLocation(string(varstr))
	time1 := []times{
		{time.Now()},
		{time.Now().In(loc)},
	}
	json.NewEncoder(w).Encode(time1)
}

func getLocalTime(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Local current time: ",time.Now().UTC())
}

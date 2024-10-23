package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)


func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string,0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz,",")

	for _, tzdb := range timezones{
		loc,err := time.LoadLocation(tzdb)
		if(err != nil){
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone")))
			continue
		}
		now := time.Now().In(loc)
		response[tzdb] = now.String()
	}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}


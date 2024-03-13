package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("FAILED TO MARSHALL %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application.json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Res with 5xx err", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	// ung may ganto `` para yan dun sa marshal or unmarshall para malaman kung ano ibabalik
	//so sa scenario na to {"error":"something went wrong"}

	respWithJSON(w, code, errResponse{
		Error: msg,
	})

}

package views

import (
        "net/http"
	"log"
	"encoding/json"
	"github.com/go_simurgh/models"
        )

type report struct {
  Text string
}

func reqlog(r *http.Request){
	log.Printf("[%s %s] %s %s \n", r.Proto, r.Method, r.RemoteAddr, r.URL.String())
}

// url /
func Index(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	report := report{"Success"}

	js, err := json.Marshal(report)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// url /SignUp
func Signup(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	U,err := models.AddUser(decoder);if err != nil {
	  http.Error(w, err.Error(), http.StatusBadRequest)
	  return
	}

	report := report{"Waiting for Email Varification"}

	js, err := json.Marshal(report)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}



















//

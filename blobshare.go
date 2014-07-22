package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// indexHandler servers the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app.html")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

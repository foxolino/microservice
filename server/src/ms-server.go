package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler())
	http.ListenAndServe(":8000", nil)
}

func myHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rv := hellohttp{namerotz{"christof", "Fox"}, 33, "christof.fox@googlemail.com"}
		jsonrv, _ := json.Marshal(&rv)
		fmt.Fprint(w, string(jsonrv))
	}
}

type hellohttp struct {
	Name  namerotz
	Age   int32
	Email string
}

type namerotz struct {
	Vorname  string `json:"Penis"`
	Nachname string
}

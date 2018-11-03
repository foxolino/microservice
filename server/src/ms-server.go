package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"./catalog"
)

func main() {
	http.HandleFunc("/", myHandler())
	port := 8000
	fmt.Println("Server started on port", port, "...")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func myHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Init catalog
		s := catalog.DefaultScedule{}
		tecinf := catalog.Training{11, "Tech Inf", "Technische Informatik", catalog.Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, 250}
		teoinf := catalog.Training{11, "Theo Inf", "Theoretische Informatik", catalog.Teacher{101, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, 400}
		s.AddTraining(tecinf)
		s.AddTraining(teoinf)

		for _, t := range s.AllTrainings() {
			jsonrv, _ := json.Marshal(&t)
			fmt.Fprint(w, string(jsonrv))
		}

		// fmt.Println(string(jsonrv))
	}
}

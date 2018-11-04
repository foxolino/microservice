package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"./catalog"
	"./handler"
)

func main() {
	// Fill scedule catalag
	s := catalog.DefaultScedule{}
	fillScedules(&s)

	// Create http-Listener
	port := 8000
	fmt.Println("Server started on port", port, "...")
	http.ListenAndServe(":"+strconv.Itoa(port), handler.TrainingsRouter(s))
}

// fillScedules with some trainings to demo
func fillScedules(s *catalog.DefaultScedule) {
	var startTime time.Time
	startTime, _ = time.Parse("10:30", "10:30")
	s.AddTraining(catalog.Training{11, "Tech Inf", "Technische Informatik", catalog.Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250})
	s.AddTraining(catalog.Training{12, "Theo Inf", "Theoretische Informatik", catalog.Teacher{102, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, startTime, 400})
	s.AddTraining(catalog.Training{13, "Kom Tech", "Kommunikationstechnik", catalog.Teacher{103, "Kerstin Uhde", 49, "uhde@fbrs.de"}, startTime, 150})
	test := catalog.Training{11, "Tech Inf", "Technische Informatik", catalog.Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250}
	jsonrv, _ := json.Marshal(&test)
	fmt.Println(jsonrv)
}

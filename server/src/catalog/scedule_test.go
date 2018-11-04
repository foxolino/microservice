package catalog

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// func TestAddTraining(t *testing.T) {

// 	// Setup
// 	s := DefaultScedule{}

// 	// Method under test
// 	s.AddTraining(Training{11, "Theo Inf", "Theoretische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, 250})

// 	// Assertions
// 	if len(s.AllTrainings()) != 1 {
// 		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
// 	}
// }

// func TestDelete(t *testing.T) {

// 	// Init catalog
// 	s := DefaultScedule{}
// 	s.AddTraining(Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, 250})
// 	s.AddTraining(Training{12, "Theo Inf", "Theoretische Informatik", Teacher{101, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, 400})

// 	// Assertion for 2 elements
// 	if len(s.AllTrainings()) != 2 {
// 		t.Fatalf("wanted %d, got %d", 2, len(s.AllTrainings()))
// 	}

// 	s.Delete(12)

// 	// Assertions for 1 elements
// 	if len(s.AllTrainings()) != 1 {
// 		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
// 	}
// }

func TestFindTrainings(t *testing.T) {

	// Init catalog
	s := DefaultScedule{}

	layout := "15:04"
	var startTime time.Time
	startTime, _ = time.Parse(layout, "10:30")

	s.AddTraining(Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250})
	startTime, _ = time.Parse(layout, "11:30")
	s.AddTraining(Training{12, "Theo Inf", "Theoretische Informatik", Teacher{102, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, startTime, 400})
	startTime, _ = time.Parse(layout, "12:30")
	s.AddTraining(Training{13, "Kom Tech", "Kommunikationstechnik", Teacher{103, "Kerstin Uhde", 49, "uhde@fbrs.de"}, startTime, 150})
	startTime, _ = time.Parse(layout, "14:30")
	test := Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250}
	jsonrv, _ := json.Marshal(&test)
	fmt.Println(string(jsonrv))

	// Assertions for 1 elements
	if len(s.AllTrainings()) != 3 {
		t.Fatalf("wanted %d, got %d", 3, len(s.AllTrainings()))
	}
}

package catalog

import (
	"testing"
	"time"
)

func TestAddTraining(t *testing.T) {
	layout := "15:04"
	var startTime time.Time
	startTime, _ = time.Parse(layout, "10:30")

	// Setup
	s := DefaultScedule{}

	// Method under test
	s.AddTraining(Training{11, "Theo Inf", "Theoretische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, startTime, 250})

	// Assertions
	if len(s.AllTrainings()) != 1 {
		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
	}
}

func TestDelete(t *testing.T) {
	layout := "15:04"
	var startTime time.Time
	startTime, _ = time.Parse(layout, "10:30")

	// Init catalog
	s := DefaultScedule{}
	s.AddTraining(Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, startTime, 250})
	s.AddTraining(Training{12, "Theo Inf", "Theoretische Informatik", Teacher{101, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, startTime, 400})

	// Assertion for 2 elements
	if len(s.AllTrainings()) != 2 {
		t.Fatalf("wanted %d, got %d", 2, len(s.AllTrainings()))
	}

	s.Delete(12)

	// Assertions for 1 elements
	if len(s.AllTrainings()) != 1 {
		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
	}
}

func TestFindTrainings(t *testing.T) {

	//Init catalog
	s := DefaultScedule{}

	layout := "15:04"
	var startTime time.Time
	startTime, _ = time.Parse(layout, "10:30")
	s.AddTraining(Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250})
	startTime, _ = time.Parse(layout, "11:30")
	s.AddTraining(Training{12, "Theo Inf", "Theoretische Informatik", Teacher{102, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, startTime, 400})
	startTime, _ = time.Parse(layout, "12:30")
	s.AddTraining(Training{13, "Kom Tech", "Kommunikationstechnik", Teacher{103, "Kerstin Uhde", 49, "uhde@fbrs.de"}, startTime, 150})

	// Find with start equal to training
	testcase := "start equal"
	ts := s.FindTrainings("11:30", "12:29")
	wanted := int32(12)
	if len(ts) != 1 {
		t.Fatalf(testcase+": wanted %d, got %d", 1, len(ts))
	}
	if ts[0].ID != wanted {
		t.Fatalf(testcase+": wanted ID %d, got %d", wanted, ts[0].ID)
	}

	// Find with start smaller than training
	testcase, ts = "start smaller", s.FindTrainings("11:29", "12:29")
	wanted = int32(12)
	if len(ts) != 1 {
		t.Fatalf(testcase+": wanted %d, got %d", 1, len(ts))
	}
	if ts[0].ID != wanted {
		t.Fatalf(testcase+": wanted ID %d, got %d", wanted, ts[0].ID)
	}

	// Find with end equal to training
	testcase, ts = "end equal", s.FindTrainings("12:30", "12:31")
	wanted = int32(13)
	if len(ts) != 1 {
		t.Fatalf(testcase+": wanted %d, got %d", 1, len(ts))
	}
	if ts[0].ID != wanted {
		t.Fatalf(testcase+": wanted ID %d, got %d", wanted, ts[0].ID)
	}

	// Find with end greater than training
	testcase, ts = "end greater", s.FindTrainings("11:29", "11:30")
	wanted = int32(12)
	if len(ts) != 1 {
		t.Fatalf(testcase+": wanted %d, got %d", 1, len(ts))
	}
	if ts[0].ID != wanted {
		t.Fatalf(testcase+": wanted ID %d, got %d", wanted, ts[0].ID)
	}
	// Find with start equal to training and end equal to training
	testcase, ts = "start equal, end equal", s.FindTrainings("11:30", "12:30")
	if len(ts) != 2 {
		t.Fatalf(testcase+": wanted %d, got %d", 2, len(ts))
	}
	if ts[0].ID != 12 {
		t.Fatalf(testcase+": wanted ID %d, got %d", 12, ts[0].ID)
	}
	if ts[1].ID != 13 {
		t.Fatalf(testcase+": wanted ID %d, got %d", 13, ts[1].ID)
	}

}

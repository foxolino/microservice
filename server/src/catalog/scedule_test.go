package catalog

import "testing"

func TestAddTraining(t *testing.T) {

	// Setup
	s := DefaultScedule{}

	// Method under test
	teoinf := Training{11, "Theo Inf", "Theoretische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, 250}
	s.AddTraining(teoinf)

	// Assertions
	if len(s.AllTrainings()) != 1 {
		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
	}
}

func TestDelete(t *testing.T) {

	// Init catalog
	s := DefaultScedule{}
	tecinf := Training{11, "Tech Inf", "Technische Informatik", Teacher{101, "Norbert Jung", 55, "njung@fbrs.de"}, 250}
	teoinf := Training{12, "Theo Inf", "Theoretische Informatik", Teacher{101, "Kurt-Ulrich Witt", 59, "witt@fbrs.de"}, 400}
	s.AddTraining(tecinf)
	s.AddTraining(teoinf)

	// Assertions
	if len(s.AllTrainings()) != 2 {
		t.Fatalf("wanted %d, got %d", 2, len(s.AllTrainings()))
	}

	s.Delete(12)

	// Assertions
	if len(s.AllTrainings()) != 1 {
		t.Fatalf("wanted %d, got %d", 1, len(s.AllTrainings()))
	}
}

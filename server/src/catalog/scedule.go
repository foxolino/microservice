package catalog

import "time"

// Scedule all Trainings
type Scedule interface {
	AddTraining(t Training) int
	AllTrainings() []Training
	FindTrainings(startTime string, endTime string) []Training
	Delete(id int32) bool
	Update(id int32, t Training) bool
	TrainingByID(id int32) (Training, bool)
}

// DefaultScedule for all Trainings
type DefaultScedule struct {
	trainings []Training
}

// NewMyDefaultScedule for getting an Pointer back
func NewMyDefaultScedule() *DefaultScedule {
	r := DefaultScedule{}
	return &r
}

// AddTraining to scedule
func (s *DefaultScedule) AddTraining(t Training) int {
	s.trainings = append(s.trainings, t)
	return 0
}

// AllTrainings in default scedule
func (s *DefaultScedule) AllTrainings() []Training {
	return s.trainings
}

// FindTrainings by time
func (s *DefaultScedule) FindTrainings(startTime string, endTime string) []Training {
	var start, end time.Time
	layout := "15:04"
	start, _ = time.Parse(layout, startTime)
	end, _ = time.Parse(layout, endTime)
	var rv []Training
	for _, t := range s.trainings {
		if (t.Time.After(start) || t.Time.Equal(start)) && (t.Time.Before(end) || t.Time.Equal(end)) {
			rv = append(rv, t)
		}
	}

	return rv
}

// Delete training
func (s *DefaultScedule) Delete(id int32) bool {
	if i := index(id, s.trainings); i != -1 {
		s.trainings = append(s.trainings[:i], s.trainings[i+1:]...)
		return true
	}

	return false
}

// Returns Index of training
func index(id int32, t []Training) int {
	for i, p := range t {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// Update Training (by Id) in Scedule
func (s *DefaultScedule) Update(id int32, t Training) bool {
	if i := index(id, s.trainings); i != -1 {
		s.trainings[i] = t
		return true
	}

	return false
}

// TrainingByID in default scedule
func (s *DefaultScedule) TrainingByID(id int32) (Training, bool) {
	for _, t := range s.trainings {
		if t.ID == id {
			return t, true
		}
	}

	return Training{}, false
}

package catalog

// Scedule all Trainings
type Scedule interface {
	AddTraining(t Training) int
	AllTrainings() []Training
	FindTraining(time string) []Training
	TrainingById(id int) (Training, bool)
	Delete(id int) bool
	Contains(name string) bool
	Update(id int, t Training) bool
}

// DefaultScedule for all Trainings
type DefaultScedule struct {
	trainings []Training
}

// AddTraining to scedule
func (s *DefaultScedule) AddTraining(p Training) int {
	s.trainings = append(s.trainings, p)
	return 0
}

// AllTrainings in default scedule
func (s *DefaultScedule) AllTrainings() []Training {
	return s.trainings
}

// FindTrainings by time
func (s *DefaultScedule) FindTrainings(time string) []Training {
	return s.trainings
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

// Contains name
func (s *DefaultScedule) Contains(name string) bool {
	for _, p := range s.AllTrainings() {
		if p.Name == name {
			return true
		}
	}
	return false
}

// Update Training (by Id) in Scedule
func (s *DefaultScedule) Update(id int32, p Training) bool {
	if i := index(id, s.trainings); i != -1 {
		s.trainings[i] = p
		return true
	}

	return false
}

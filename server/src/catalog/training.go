package catalog

import "time"

// Training structure, depend on Teacher
type Training struct {
	ID          int32     `json:"Id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Teacher     Teacher   `json:"Lecturer"`
	Time        time.Time `json:"Time"`
	Price       float32   `json:"Price"`
}

// Teacher structure
type Teacher struct {
	ID    int32  `json:"Id"`
	Name  string `json:"Name"`
	Age   int32  `json:"Age"`
	Email string `json:"EMail"`
}

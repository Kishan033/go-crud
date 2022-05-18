package models

// User schema of the user table
type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Age      int64  `json:"age,omitempty"`
}

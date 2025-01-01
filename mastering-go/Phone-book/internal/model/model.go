package model

type PhoeBookError struct {
	Message    string
	StatusCode int32
}

type Entry struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
}

type ListResponse struct {
	Entries []Entry `json:"entries"`
}

type InsertResponse struct {
	ID int64 `json:"id"`
}

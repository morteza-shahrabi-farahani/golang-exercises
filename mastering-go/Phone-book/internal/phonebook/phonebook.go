package phonebook

type PhoeBookError struct {
	Message    string
	StatusCode int32
}

type Entry struct {
	ID          int64
	Name        string
	Surname     string
	PhoneNumber string
}

type ListResponse struct {
	entries []Entry
}

type InsertResponse struct {
	id int64
}

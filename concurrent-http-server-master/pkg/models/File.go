package models

type Filee struct {
	Id       		string `json:"file_id"`
	Location		string  `json:"-"`
	//location string `json:"file_location"`
}

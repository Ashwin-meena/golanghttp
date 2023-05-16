package model

type User_detail struct {
	ID          uint    `json:"id"`
    User_name     string `json:"name"`
    Location string `json:"location"`
    Age      int64  `json:"age"`
	Password string `json:"password"`
}
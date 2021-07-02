package model

type User struct {
	ID      int     `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int     `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"Married"`
}

type Error struct{
	Message string
}
package domain

import (
	"time"
)

type User struct {
	ID        	int64     `json:"id"`
	Username		string    `json:"username"`
	Password		string    `json:"password"`
	Email				string    `json:"email"`
	PhoneNumber	string    `json:"phone_number"`
	CreatedOn 	time.Time `json:"created_on"`
	CreatedBy 	int    		`json:"created_by"`
	UpdatedOn 	*time.Time `json:"updated_on"`
	UpdatedBy 	*int    		`json:"updated_by"`
}
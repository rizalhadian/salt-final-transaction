package infrastructure_customer_http_response

import "time"

type Customer struct {
	Id         int64     `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Dob        time.Time `json:"dob"`
	Address    string    `json:"address"`
	Created_at time.Time `json:"created_at"`
}

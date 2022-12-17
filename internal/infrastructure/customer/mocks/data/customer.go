package mocks_data

import (
	infrastructure_customer_http_response "salt-final-transaction/internal/infrastructure/customer/http_response"
	"time"
)

func GetById(customer_id int64) infrastructure_customer_http_response.Customer {
	customer := infrastructure_customer_http_response.Customer{
		Id:         customer_id,
		Email:      "mock_customer@gmail.com",
		Name:       "mock customer name",
		Phone:      "111222333",
		Dob:        time.Now(),
		Address:    "Jimbaran",
		Created_at: time.Now(),
	}

	return customer
}

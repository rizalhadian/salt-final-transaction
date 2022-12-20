package http_response

import (
	"salt-final-transaction/domain/entity"
	"time"
)

type CustomersTransactionCountResponse struct {
	Id                         int64     `json:"id"`
	Customer_id                int64     `json:"customer_id"`
	Total_transaction_spend    float64   `json:"total_transaction_spend"`
	Total_transaction_count    int32     `json:"total_transaction_count"`
	First_transaction_datetime time.Time `json:"first_transaction_datetime"`
	Last_transaction_datetime  time.Time `json:"last_transaction_datetime"`
}

func MapCustomersTransactionCountResponse(entity_customers_transaction_count *entity.CustomersTransactionCount) *CustomersTransactionCountResponse {

	return &CustomersTransactionCountResponse{
		Id:                         entity_customers_transaction_count.GetId(),
		Customer_id:                entity_customers_transaction_count.GetCustomerId(),
		Total_transaction_spend:    entity_customers_transaction_count.GetTotalTransactionSpend(),
		Total_transaction_count:    entity_customers_transaction_count.GetTransactionCount(),
		First_transaction_datetime: entity_customers_transaction_count.GetFirstTransactionDatetime(),
		Last_transaction_datetime:  entity_customers_transaction_count.GetLastTransactionDatetime(),
	}
}

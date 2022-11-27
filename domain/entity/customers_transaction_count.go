package entity

import "time"

type CustomersTransactionCount struct {
	id                         int64
	customer_id                int64
	total_transaction_spend    float64
	transaction_count          int32
	first_transaction_datetime time.Time
	last_transaction_datetime  time.Time
}

type DTOCustomersTransactionCount struct {
	Id                         int64
	Customer_id                int64
	Total_transaction_spend    float64
	Total_transaction_count    int32
	First_transaction_datetime time.Time
	Last_transaction_datetime  time.Time
}

// Customer's Transaction Count Getter

func (ctc *CustomersTransactionCount) GetId() int64 {
	return ctc.id
}

func (ctc *CustomersTransactionCount) GetCustomerId() int64 {
	return ctc.customer_id
}

func (ctc *CustomersTransactionCount) GetTotalTransactionSpend() float64 {
	return ctc.total_transaction_spend
}

func (ctc *CustomersTransactionCount) GetTransactionCount() int32 {
	return ctc.transaction_count
}

func (ctc *CustomersTransactionCount) GetFirstTransactionDatetime() time.Time {
	return ctc.first_transaction_datetime
}

func (ctc *CustomersTransactionCount) GetLastTransactionDatetime() time.Time {
	return ctc.last_transaction_datetime
}

// Customer's Transaction Count Setter

func (ctc *CustomersTransactionCount) SetId(value int64) {
	ctc.id = value
}

func (ctc *CustomersTransactionCount) SetCustomerId(value int64) {
	ctc.customer_id = value
}

func (ctc *CustomersTransactionCount) SetTotalTransactionSpend(value float64) {
	ctc.total_transaction_spend = value
}

func (ctc *CustomersTransactionCount) SetTransactionCount(value int32) {
	ctc.transaction_count = value
}

func (ctc *CustomersTransactionCount) SetFirstTransactionDatetime(value time.Time) {
	ctc.first_transaction_datetime = value
}

func (ctc *CustomersTransactionCount) SetLastTransactionDatetime(value time.Time) {
	ctc.last_transaction_datetime = value
}

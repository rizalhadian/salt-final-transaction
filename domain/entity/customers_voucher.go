package entity

import (
	"database/sql"
	"time"
)

type CustomersVoucher struct {
	id             int64
	customer_id    int64
	voucher_id     int64
	code           string
	expired_at     time.Time
	transaction_id sql.NullInt64
	created_at     time.Time
	updated_at     sql.NullTime
	deleted_at     sql.NullTime
}

type DTOCustomersVoucher struct {
	Id             int64
	Customer_id    int64
	Voucher_id     int64
	Code           string
	Expired_at     time.Time
	Transaction_id sql.NullInt64
	Created_at     time.Time
	Updated_at     sql.NullTime
	Deleted_at     sql.NullTime
}

// Getter

func (cv *CustomersVoucher) GetId() int64 {
	return cv.id
}

func (cv *CustomersVoucher) GetCustomerId() int64 {
	return cv.customer_id
}

func (cv *CustomersVoucher) GetVoucherId() int64 {
	return cv.voucher_id
}

func (cv *CustomersVoucher) GetCode() string {
	return cv.code
}

func (cv *CustomersVoucher) GetExpiredAt() time.Time {
	return cv.expired_at
}

func (cv *CustomersVoucher) GetTransactionId() sql.NullInt64 {
	return cv.transaction_id
}

func (cv *CustomersVoucher) GetCreatedAt() time.Time {
	return cv.created_at
}

func (cv *CustomersVoucher) GetUpdatedAt() sql.NullTime {
	return cv.updated_at
}

func (cv *CustomersVoucher) GetDeletedAt() sql.NullTime {
	return cv.deleted_at
}

// Setter

func (cv *CustomersVoucher) SetId(value int64) {
	cv.id = value
}

func (cv *CustomersVoucher) SetCustomerId(value int64) {
	cv.customer_id = value
}

func (cv *CustomersVoucher) SetVoucherId(value int64) {
	cv.voucher_id = value
}

func (cv *CustomersVoucher) SetCode(value string) {
	cv.code = value
}

func (cv *CustomersVoucher) SetExpiredAt(value time.Time) {
	cv.expired_at = value
}

func (cv *CustomersVoucher) SetTransactionId(value sql.NullInt64) {
	cv.transaction_id = value
}

func (cv *CustomersVoucher) SetCreatedAt(value time.Time) {
	cv.created_at = value
}

func (cv *CustomersVoucher) SetUpdatedAt(value sql.NullTime) {
	cv.updated_at = value
}

func (cv *CustomersVoucher) SetDeletedAt(value sql.NullTime) {
	cv.deleted_at = value
}

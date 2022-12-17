package infrastructure_voucher_http_response

import "time"

type Voucher struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CustomersVoucher struct {
	Id             int64     `json:"id"`
	Customer_id    int64     `json:"customer_id"`
	Voucher_id     int64     `json:"voucher_id"`
	Code           string    `json:"code"`
	Expired_at     time.Time `json:"expired_at"`
	Transaction_id int64     `json:"transaction_id"`
	Created_at     time.Time `json:"created_at"`
	// Updated_at     time.Time `json:"updated_at"`
	// Deleted_at     time.Time `json:"deleted_at"`
}

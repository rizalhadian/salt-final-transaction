package infrastructure_voucher_http_response

import "time"

type Voucher struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CustomersVoucher struct {
	Id                    int64     `json:"id"`
	Customer_id           int64     `json:"customer_id"`
	Voucher_id            int64     `json:"voucher_id"`
	Voucher_name          string    `json:"voucher_name"`
	Code                  string    `json:"code"`
	Expired_at            time.Time `json:"expired_at"`
	Transaction_id        int64     `json:"transaction_id"`
	Total_amount          float64   `json:"total_amount"`
	Total_discount_amount float64   `json:"total_discount_amount"`
	Final_total_amount    float64   `json:"final_total_amount"`
	Status                int16     `json:"status"`
	Created_at            time.Time `json:"created_at"`
}

type SuccedRedeemResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    []*CustomersVoucher `json:"data"`
}

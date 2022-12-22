package http_request

type Transaction struct {
	Id                           int64                         `json:"id,omitempty"`
	Customer_id                  int64                         `json:"customer_id,omitempty"`
	Total_amount                 float64                       `json:"total_amount,omitempty"`
	TransactionsItems            []TransactionsItem            `json:"items,omitempty"`
	TransactionsVouchersRedeemed []TransactionsVoucherRedeemed `json:"vouchers_redeemed,omitempty"`
}

type TransactionsItem struct {
	Item_id       int64   `json:"item_id,omitempty"`
	Items_type_id int64   `json:"items_type_id,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Qty           int32   `json:"qty,omitempty"`
	Note          string  `json:"note,omitempty"`
}

type TransactionsVoucherRedeemed struct {
	Code string `json:"code,omitempty"`
}

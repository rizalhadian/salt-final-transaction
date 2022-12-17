package http_request

import (
	"time"
)

type Transaction struct {
	Customer_id       int                `json:"customer_id,omitempty"`
	Note              string             `json:"note,omitempty"`
	Status            string             `json:"status,omitempty"`
	TesDate           time.Time          `json:"date,omitempty"`
	TransactionsItems []TransactionsItem `json:"items,omitempty"`
}

type TransactionsItem struct {
	Item_id int64   `json:"item_id,omitempty"`
	Price   float64 `json:"price,omitempty"`
	Qty     int32   `json:"qty,omitempty"`
	Note    string  `json:"note,omitempty"`
}

// type Transaction struct {
// 	Customer_id       int                `validate:"required"`
// 	Note              string             `validate:"omitempty"`
// 	Status            string             `validate:"omitempty"`
// 	TesDate           time.Time          `validate:"omitempty"`
// 	TransactionsItems []TransactionsItem `validate:"omitempty"`
// }

// type TransactionsItem struct {
// 	Item_id int64   `validate:"required"`
// 	Price   float64 `validate:"omitempty,numeric"`
// 	Qty     int32   `validate:"omitempty,min=1,numeric"`
// 	Note    string  `validate:"omitempty"`
// }

// func (t *Transaction) Validate() error {
// 	validate := validator.New()
// 	err := validate.Struct(t)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (ti *TransactionsItem) Validate() error {
// 	validate := validator.New()
// 	err := validate.Struct(ti)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

package entity

import (
	"database/sql"
	"errors"
	"time"
)

type TransactionsItem struct {
	id                   int64
	transaction_id       int64
	item_id              int64
	items_type_id        int64
	price                float64
	qty                  int32
	total_price          float64
	note                 string
	customers_voucher_id sql.NullInt64
	voucher_id           sql.NullInt64
	voucher_code         string
	discount_percentage  int32
	discount_amount      float64
	final_price          float64
	created_at           time.Time
	updated_at           sql.NullTime
	deleted_at           sql.NullTime
}

type DTOTransactionsItem struct {
	Id                   int64
	Transaction_id       int64
	Item_id              int64
	Items_type_id        int64
	Price                float64
	Qty                  int32
	Total_price          float64
	Note                 string
	Customers_voucher_id sql.NullInt64
	Voucher_id           sql.NullInt64
	Voucher_code         string
	Discount_percentage  int32
	Discount_amount      float64
	Final_price          float64
	Created_at           time.Time
	Updated_at           sql.NullTime
	Deleted_at           sql.NullTime
}

func NewTransactionsItem(dto DTOTransactionsItem) (*TransactionsItem, error) {

	if dto.Transaction_id == 0 {
		return nil, errors.New("Transaction_id is required")
	}

	if dto.Items_type_id == 0 {
		return nil, errors.New("Items_type_id is required")
	}

	// This validation should be on usecase
	// if dto.Total_price != (dto.Price * float64(dto.Qty)) {
	// 	return nil, errors.New("Total_price is incorrect")
	// }

	if dto.Customers_voucher_id.Valid == true {
		if dto.Customers_voucher_id.Int64 == 0 {
			return nil, errors.New("If Customers_voucher_id not null, Customers_voucher_id cannot be 0")
		}
		if dto.Voucher_id.Valid == false {
			return nil, errors.New("If Customers_voucher_id not null, Voucher_id cannot be null")
		}
		if dto.Voucher_id.Int64 == 0 {
			return nil, errors.New("If Customers_voucher_id not null, Voucher_id cannot be 0")
		}
		if dto.Voucher_code == "" {
			return nil, errors.New("If Customers_voucher_id not null, Voucher_code cannot be blank")
		}
	}

	if dto.Voucher_id.Valid == true {
		if dto.Voucher_id.Int64 == 0 {
			return nil, errors.New("If Voucher_id not null, Voucher_id cannot be 0")
		}
		if dto.Customers_voucher_id.Valid == false {
			return nil, errors.New("If Voucher_id not null, Customers_voucher_id cannot be null")
		}
		if dto.Customers_voucher_id.Int64 == 0 {
			return nil, errors.New("If Voucher_id not null, Customers_voucher_id cannot be 0")
		}
		if dto.Voucher_code == "" {
			return nil, errors.New("If Voucher_id not null, Voucher_code cannot be blank")
		}
	}

	if dto.Voucher_code == "" {
		if dto.Voucher_id.Valid == false {
			return nil, errors.New("If Voucher_code not blank, Voucher_id cannot be null")
		}
		if dto.Voucher_id.Int64 == 0 {
			return nil, errors.New("If Voucher_code not blank, Voucher_id cannot be 0")
		}
		if dto.Customers_voucher_id.Valid == false {
			return nil, errors.New("If Voucher_code not blank, Customers_voucher_id cannot be null")
		}
		if dto.Customers_voucher_id.Int64 == 0 {
			return nil, errors.New("If Voucher_code not blank, Customers_voucher_id cannot be 0")
		}
	}

	transactions_item := &TransactionsItem{
		id:                   dto.Id,
		transaction_id:       dto.Transaction_id,
		item_id:              dto.Item_id,
		items_type_id:        dto.Items_type_id,
		price:                dto.Price,
		qty:                  dto.Qty,
		total_price:          dto.Total_price,
		note:                 dto.Note,
		customers_voucher_id: dto.Customers_voucher_id,
		voucher_id:           dto.Voucher_id,
		voucher_code:         dto.Voucher_code,
		discount_percentage:  dto.Discount_percentage,
		discount_amount:      dto.Discount_amount,
		final_price:          dto.Final_price,
		created_at:           dto.Created_at,
		updated_at:           dto.Updated_at,
		deleted_at:           dto.Deleted_at,
	}

	return transactions_item, nil

}

// Transaction's Item Getter
func (t *TransactionsItem) GetId() int64 {
	return t.id
}

func (t *TransactionsItem) GetTransactionId() int64 {
	return t.transaction_id
}

func (t *TransactionsItem) GetItemId() int64 {
	return t.item_id
}

func (t *TransactionsItem) GetItemsTypeId() int64 {
	return t.items_type_id
}

func (t *TransactionsItem) GetPrice() float64 {
	return t.price
}

func (t *TransactionsItem) GetQty() int32 {
	return t.qty
}

func (t *TransactionsItem) GetTotalPrice() float64 {
	return t.total_price
}

func (t *TransactionsItem) GetNote() string {
	return t.note
}

func (t *TransactionsItem) GetCustomerVoucherId() sql.NullInt64 {
	return t.customers_voucher_id
}

func (t *TransactionsItem) GetVoucherId() sql.NullInt64 {
	return t.voucher_id
}

func (t *TransactionsItem) GetVoucherCode() string {
	return t.voucher_code
}

func (t *TransactionsItem) GetDiscountPercentage() int32 {
	return t.discount_percentage
}

func (t *TransactionsItem) GetDiscountAmount() float64 {
	return t.discount_amount
}

func (t *TransactionsItem) GetFinalPrice() float64 {
	return t.final_price
}

func (t *TransactionsItem) GetCreatedAt() time.Time {
	return t.created_at
}

func (t *TransactionsItem) GetUpdatedAt() sql.NullTime {
	return t.updated_at
}

func (t *TransactionsItem) GetDeletedAt() sql.NullTime {
	return t.deleted_at
}

// Transaction's Item Setter
func (t *TransactionsItem) SetId(value int64) {
	t.id = value
}

func (t *TransactionsItem) SetTransactionId(value int64) {
	t.transaction_id = value
}

func (t *TransactionsItem) SetItemId(value int64) {
	t.item_id = value
}

func (t *TransactionsItem) SetItemsTypeId(value int64) {
	t.items_type_id = value
}

func (t *TransactionsItem) SetPrice(value float64) {
	t.price = value
}

func (t *TransactionsItem) SetQty(value int32) {
	t.qty = value
}

func (t *TransactionsItem) SetTotalPrice(value float64) {
	t.total_price = value
}

func (t *TransactionsItem) SetNote(value string) {
	t.note = value
}

func (t *TransactionsItem) SetCustomerVoucherId(value sql.NullInt64) {
	t.customers_voucher_id = value
}

func (t *TransactionsItem) SetVoucherId(value sql.NullInt64) {
	t.voucher_id = value
}

func (t *TransactionsItem) SetVoucherCode(value string) {
	t.voucher_code = value
}

func (t *TransactionsItem) SetDiscountPercentage(value int32) {
	t.discount_percentage = value
}

func (t *TransactionsItem) SetDiscountAmount(value float64) {
	t.discount_amount = value
}

func (t *TransactionsItem) SetFinalPrice(value float64) {
	t.final_price = value
}

func (t *TransactionsItem) SetCreatedAt(value time.Time) {
	t.created_at = value
}

func (t *TransactionsItem) SetUpdatedAt(value sql.NullTime) {
	t.updated_at = value
}

func (t *TransactionsItem) SetDeletedAt(value sql.NullTime) {
	t.deleted_at = value
}

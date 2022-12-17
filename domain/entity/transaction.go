package entity

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type TransactionStatus struct {
	code        int16
	section     string
	name        string
	description string
}

var transactionStatuses = []TransactionStatus{
	{
		code:        110,
		section:     "transaction",
		name:        "failed",
		description: "",
	},
	{
		code:        111,
		section:     "transaction",
		name:        "submitted",
		description: "",
	},
	{
		code:        112,
		section:     "transaction",
		name:        "pending",
		description: "",
	},
	{
		code:        113,
		section:     "transaction",
		name:        "revised",
		description: "",
	},
	{
		code:        114,
		section:     "transaction",
		name:        "revise",
		description: "",
	},
	{
		code:        115,
		section:     "transaction",
		name:        "rollback",
		description: "",
	},
	{
		code:        116,
		section:     "transaction",
		name:        "confirmed",
		description: "",
	},
}

type Transaction struct {
	id                           int64
	customer_id                  int64
	total_amount                 float64
	total_discount_amount        float64
	final_total_amount           float64
	note                         string
	status                       int16
	rollback_transaction_id      int64
	update_transaction_id        int64
	items                        []*TransactionsItem
	created_at                   time.Time
	updated_at                   sql.NullTime
	deleted_at                   sql.NullTime
	is_generated_voucher_succeed bool
}

type DTOTransaction struct {
	Id                           int64
	Customer_id                  int64
	Total_amount                 float64
	Total_discount_amount        float64
	Final_total_amount           float64
	Note                         string
	Status                       int16
	Rollback_transaction_id      int64
	Update_transaction_id        int64
	Items                        []*DTOTransactionsItem
	Created_at                   time.Time
	Updated_at                   sql.NullTime
	Deleted_at                   sql.NullTime
	Is_generated_voucher_succeed bool
}

func ConvertTransactionStatusStringToInt(value string) (int16, error) {
	value_lowercase := strings.ToLower(value)
	for _, transaction_status := range transactionStatuses {
		if transaction_status.name == value_lowercase {
			return transaction_status.code, nil
		}
	}
	return 0, errors.New("Transaction Status Not Found")
}

func ConvertTransactionStatusIntToString(value int) (string, error) {
	for _, transaction_status := range transactionStatuses {
		if int(transaction_status.code) == value {
			return transaction_status.name, nil
		}
	}
	return "", errors.New("Transaction Status Not Found")
}

type TotalAmountPerCategory struct {
	items_type_id int64
	total_amount  float64
}

func NewTransaction(dto *DTOTransaction) (*Transaction, error) {
	// if dto.Status == 0 {
	// 	return nil, errors.New("Status is required")
	// }

	// This validation should be on usecase
	// if dto.Final_total_amount != (dto.Total_amount - dto.Total_discount_amount) {
	// 	return nil, errors.New("Final_total_amount is incorrect")
	// }

	// var items_entity []*TransactionsItem

	// var total_amount float64

	// for _, item := range dto.Items {
	// 	item_entity, error_item_entity := NewTransactionsItem(*item)
	// 	if error_item_entity != nil {
	// 		return nil, error_item_entity
	// 	}
	// 	total_amount += item_entity.GetTotalPrice()
	// 	items_entity = append(items_entity, item_entity)
	// }

	// var transaction_items []*TransactionsItem
	// var total_amount float64

	// for _, dto_transactions_item := range dto.Items {
	// 	transaction_item_entity, transaction_item_entity_errs := NewTransactionsItem(dto_transactions_item)
	// 	if transaction_item_entity_errs != nil {
	// 		return nil, transaction_item_entity_errs
	// 	}
	// 	total_amount += dto_transactions_item.Total_price
	// 	transaction_items = append(transaction_items, transaction_item_entity)
	// }

	transaction := &Transaction{
		id:                           dto.Id,
		customer_id:                  dto.Customer_id,
		note:                         dto.Note,
		status:                       dto.Status,
		rollback_transaction_id:      dto.Rollback_transaction_id,
		update_transaction_id:        dto.Update_transaction_id,
		is_generated_voucher_succeed: false,
		// final_total_amount:           total_amount - dto.Total_discount_amount,
		// total_amount:                 total_amount,
		// items:                        transaction_items,
		// created_at:                   dto.Created_at,
		// updated_at:                   dto.Updated_at,
		// deleted_at:                   dto.Deleted_at,
		// total_discount_amount:        dto.Total_discount_amount,

	}
	return transaction, nil
}

// Transaction's Getter
func (t *Transaction) GetId() int64 {
	return t.id
}

func (t *Transaction) GetCustomerId() int64 {
	return t.customer_id
}

func (t *Transaction) GetTotalAmount() float64 {
	return t.total_amount
}

func (t *Transaction) GetTotalDiscountAmount() float64 {
	return t.total_discount_amount
}

func (t *Transaction) GetFinalTotalAmount() float64 {
	return t.final_total_amount
}

func (t *Transaction) GetNote() string {
	return t.note
}

func (t *Transaction) GetStatus() int16 {
	return t.status
}

func (t *Transaction) GetStatusString() string {
	status_string, err := ConvertTransactionStatusIntToString(int(t.status))
	if err != nil {
		panic(err)
	}
	return status_string
}

func (t *Transaction) GetRollbackTransactionId() int64 {
	return t.rollback_transaction_id
}

func (t *Transaction) GetUpdateTransactionId() int64 {
	return t.update_transaction_id
}

func (t *Transaction) GetItems() []*TransactionsItem {
	return t.items
}

func (t *Transaction) GetCreatedAt() time.Time {
	return t.created_at
}

func (t *Transaction) GetUpdatedAt() sql.NullTime {
	return t.updated_at
}

func (t *Transaction) GetDeletedAt() sql.NullTime {
	return t.deleted_at
}

func (t *Transaction) GetIsGeneratedVoucherSucceed() bool {
	return t.is_generated_voucher_succeed
}

// Transaction's Setter
func (t *Transaction) SetId(value int64) {
	t.id = value
}

func (t *Transaction) SetCustomerId(value int64) {
	t.customer_id = value
}

func (t *Transaction) SetTotalAmount(value float64) {
	t.total_amount = value
	t.final_total_amount = t.total_amount - t.total_discount_amount
}

func (t *Transaction) SetTotalDiscountAmount(value float64) {
	t.total_discount_amount = value
	t.final_total_amount = t.total_amount - t.total_discount_amount
}

func (t *Transaction) SetFinalTotalAmount(value float64) {
	t.final_total_amount = value
}

func (t *Transaction) SetNote(value string) {
	t.note = value
}

func (t *Transaction) SetStatus(value int16) {
	t.status = value
}

func (t *Transaction) SetRollbackTransactionId(value int64) {
	t.rollback_transaction_id = value
}

func (t *Transaction) SetUpdateTransactionId(value int64) {
	t.update_transaction_id = value
}

func (t *Transaction) SetItems(values []*TransactionsItem) {
	// var total_amount float64
	// for _, item := range values {
	// 	total_amount += item.GetTotalPrice()
	// }
	// t.total_amount = total_amount
	t.items = values
}

func (t *Transaction) SetCreatedAt(value time.Time) {
	t.created_at = value
}

func (t *Transaction) SetUpdatedAt(value sql.NullTime) {
	t.updated_at = value
}

func (t *Transaction) SetDeletedAt(value sql.NullTime) {
	t.deleted_at = value
}

func (t *Transaction) SetIsGeneratedVoucherSucceed(value bool) {
	t.is_generated_voucher_succeed = value
}

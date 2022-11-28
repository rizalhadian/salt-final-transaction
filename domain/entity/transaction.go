package entity

import (
	"database/sql"
	"errors"
	"time"
)

type Transaction struct {
	id                           int64
	customer_id                  int64
	total_amount                 float64
	total_discount_amount        float64
	final_total_amount           float64
	note                         string
	status                       int16
	rollback_transaction_id      sql.NullInt64
	update_transaction_id        sql.NullInt64
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
	Rollback_transaction_id      sql.NullInt64
	Update_transaction_id        sql.NullInt64
	Created_at                   time.Time
	Updated_at                   sql.NullTime
	Deleted_at                   sql.NullTime
	Is_generated_voucher_succeed bool
}

func NewTransaction(dto DTOTransaction) (*Transaction, error) {

	if dto.Status == 0 {
		return nil, errors.New("Status is required")
	}

	// This validation should be on usecase
	// if dto.Final_total_amount != (dto.Total_amount - dto.Total_discount_amount) {
	// 	return nil, errors.New("Final_total_amount is incorrect")
	// }

	transaction := &Transaction{
		id:                           dto.Id,
		customer_id:                  dto.Customer_id,
		total_amount:                 dto.Total_amount,
		total_discount_amount:        dto.Total_discount_amount,
		final_total_amount:           dto.Final_total_amount,
		note:                         dto.Note,
		status:                       dto.Status,
		rollback_transaction_id:      dto.Rollback_transaction_id,
		update_transaction_id:        dto.Update_transaction_id,
		created_at:                   dto.Created_at,
		updated_at:                   dto.Updated_at,
		deleted_at:                   dto.Deleted_at,
		is_generated_voucher_succeed: dto.Is_generated_voucher_succeed,
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

func (t *Transaction) GetRollbackTransactionId() sql.NullInt64 {
	return t.rollback_transaction_id
}

func (t *Transaction) GetUpdateTransactionId() sql.NullInt64 {
	return t.update_transaction_id
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
}

func (t *Transaction) SetTotalDiscountAmount(value float64) {
	t.total_discount_amount = value
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

func (t *Transaction) SetRollbackTransactionId(value sql.NullInt64) {
	t.rollback_transaction_id = value
}

func (t *Transaction) SetUpdateTransactionId(value sql.NullInt64) {
	t.update_transaction_id = value
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

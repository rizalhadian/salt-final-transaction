package model

import (
	"errors"
	"time"
)

type ModelTransaction struct {
	Id                           int64     `dbq:"id,omitempty"`
	Customer_id                  int64     `dbq:"customer_id,omitempty"`
	Total_amount                 float64   `dbq:"total_amount,omitempty"`
	Total_discount_amount        float64   `dbq:"total_discount_amount,omitempty"`
	Final_total_amount           float64   `dbq:"final_total_amount,omitempty"`
	Note                         string    `dbq:"note,omitempty"`
	Status                       int16     `dbq:"status,omitempty"`
	Rollback_transaction_id      int64     `dbq:"rollback_transaction_id,omitempty"`
	Update_transaction_id        int64     `dbq:"update_transaction_id,omitempty"`
	Created_at                   time.Time `dbq:"created_at,omitempty"`
	Updated_at                   string    `dbq:"updated_at,omitempty"`
	Deleted_at                   string    `dbq:"deleted_at,omitempty"`
	Is_generated_voucher_succeed bool      `dbq:"is_generated_voucher_succeed,omitempty"`
}

func (ModelTransaction) GetTableName() string {
	return "transaction"
}

func (mt *ModelTransaction) ToArrInterface(value string) []interface{} {

	data := make([]interface{}, 0)

	if value == "save" {
		for _, tag := range mt.GetFieldsNeededToStoreProcess() {
			val, err := mt.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "update" {
		for _, tag := range mt.GetFieldsNeededToUpdateProcess() {
			val, err := mt.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "delete" {
		for _, tag := range mt.GetFieldsNeededToSoftDeleteProcess() {
			val, err := mt.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	return data
}

func (mt *ModelTransaction) GetValueFromDbqTag(tag string) (any, error) {
	switch {
	case tag == "id":
		return mt.Id, nil
	case tag == "customer_id":
		return mt.Customer_id, nil
	case tag == "total_amount":
		return mt.Total_amount, nil
	case tag == "total_discount_amount":
		return mt.Total_discount_amount, nil
	case tag == "final_total_amount":
		return mt.Final_total_amount, nil
	case tag == "note":
		return mt.Note, nil
	case tag == "status":
		return mt.Status, nil
	case tag == "rollback_transaction_id":
		return mt.Rollback_transaction_id, nil
	case tag == "update_transaction_id":
		return mt.Update_transaction_id, nil
	case tag == "created_at":
		return mt.Created_at, nil
	case tag == "updated_at":
		return mt.Updated_at, nil
	case tag == "deleted_at":
		return mt.Deleted_at, nil
	case tag == "is_generated_voucher_succeed":
		return mt.Is_generated_voucher_succeed, nil
	default:
		return nil, errors.New("Value Not Found")
	}
}

func (ModelTransaction) GetFieldsNeededToStoreProcess() []string {
	return []string{
		"customer_id",
		"total_amount",
		"total_discount_amount",
		"final_total_amount",
		"note",
		"status",
		"rollback_transaction_id",
		"update_transaction_id",
		"created_at",
		"is_generated_voucher_succeed",
	}
}

func (ModelTransaction) GetFieldsNeededToGetProcess() []string {
	return []string{
		"id",
		"customer_id",
		"total_amount",
		"total_discount_amount",
		"final_total_amount",
		"note",
		"status",
		"rollback_transaction_id",
		"update_transaction_id",
		"created_at",
		"updated_at",
		"is_generated_voucher_succeed",
	}
}

func (ModelTransaction) GetFieldsNeededToUpdateProcess() []string {
	return []string{
		"customer_id",
		"total_amount",
		"total_discount_amount",
		"final_total_amount",
		"note",
		"status",
		"rollback_transaction_id",
		"update_transaction_id",
		"updated_at",
		"is_generated_voucher_succeed",
	}
}

func (ModelTransaction) GetFieldsNeededToSoftDeleteProcess() []string {
	return []string{
		"deleted_at",
	}
}

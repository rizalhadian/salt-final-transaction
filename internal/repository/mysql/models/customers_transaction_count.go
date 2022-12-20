package model

import (
	"errors"
	"time"
)

type ModelCustomersTransactionCount struct {
	Id                         int64     `dbq:"id"`
	Customer_id                int64     `dbq:"customer_id"`
	Total_transaction_spend    float64   `dbq:"total_transaction_spend"`
	Transaction_count          int32     `dbq:"transaction_count"`
	First_transaction_datetime time.Time `dbq:"first_transaction_datetime"`
	Last_transaction_datetime  time.Time `dbq:"last_transaction_datetime"`
}

func (ModelCustomersTransactionCount) GetTableName() string {
	return "customers_transaction_count"
}

func (mctc *ModelCustomersTransactionCount) ToArrInterface(value string) []interface{} {

	data := make([]interface{}, 0)

	if value == "save" {
		for _, tag := range mctc.GetFieldsNeededToStoreProcess() {
			val, err := mctc.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "update" {
		for _, tag := range mctc.GetFieldsNeededToUpdateProcess() {
			val, err := mctc.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "delete" {
		for _, tag := range mctc.GetFieldsNeededToSoftDeleteProcess() {
			val, err := mctc.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	return data
}

func (mctc *ModelCustomersTransactionCount) GetValueFromDbqTag(tag string) (any, error) {
	switch {
	case tag == "id":
		return mctc.Id, nil
	case tag == "customer_id":
		return mctc.Customer_id, nil
	case tag == "total_transaction_spend":
		return mctc.Total_transaction_spend, nil
	case tag == "transaction_count":
		return mctc.Transaction_count, nil
	case tag == "first_transaction_datetime":
		return mctc.First_transaction_datetime, nil
	case tag == "last_transaction_datetime":
		return mctc.Last_transaction_datetime, nil
	default:
		return nil, errors.New("Value Not Found")
	}
}

func (ModelCustomersTransactionCount) GetFieldsNeededToStoreProcess() []string {
	return []string{
		"customer_id",
		"total_transaction_spend",
		"transaction_count",
		"first_transaction_datetime",
		"last_transaction_datetime",
	}
}

func (ModelCustomersTransactionCount) GetFieldsNeededToGetProcess() []string {
	return []string{
		"id",
		"customer_id",
		"total_transaction_spend",
		"transaction_count",
		"first_transaction_datetime",
		"last_transaction_datetime",
	}
}

func (ModelCustomersTransactionCount) GetFieldsNeededToUpdateProcess() []string {
	return []string{
		"customer_id",
		"total_transaction_spend",
		"transaction_count",
		"first_transaction_datetime",
		"last_transaction_datetime",
	}
}

func (ModelCustomersTransactionCount) GetFieldsNeededToSoftDeleteProcess() []string {
	//Belum Kepake Karena Ga Buat Fitur Ini
	return []string{}
}

package model

import (
	"errors"
	"time"
)

type ModelTransactionsItem struct {
	Id             int64     `dbq:"id"`
	Transaction_id int64     `dbq:"transaction_id"`
	Item_id        int64     `dbq:"item_id"`
	Items_type_id  int64     `dbq:"items_type_id"`
	Price          float64   `dbq:"price"`
	Qty            int32     `dbq:"qty"`
	Total_price    float64   `dbq:"total_price"`
	Note           string    `dbq:"note"`
	Created_at     time.Time `dbq:"created_at"`
	Updated_at     string    `dbq:"updated_at"`
	Deleted_at     string    `dbq:"deleted_at"`
}

func (ModelTransactionsItem) GetTableName() string {
	return "transactions_item"
}

func (mti *ModelTransactionsItem) ToArrInterface(value string) []interface{} {

	data := make([]interface{}, 0)

	if value == "save" {
		for _, tag := range mti.GetFieldsNeededToStoreProcess() {
			val, err := mti.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "update" {
		for _, tag := range mti.GetFieldsNeededToUpdateProcess() {
			val, err := mti.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	if value == "delete" {
		for _, tag := range mti.GetFieldsNeededToSoftDeleteProcess() {
			val, err := mti.GetValueFromDbqTag(tag)
			if err != nil {
				panic(err)
			}
			data = append(data, val)
		}
	}

	return data
}

func (mti *ModelTransactionsItem) GetValueFromDbqTag(tag string) (any, error) {
	switch {
	case tag == "id":
		return mti.Id, nil
	case tag == "transaction_id":
		return mti.Transaction_id, nil
	case tag == "item_id":
		return mti.Item_id, nil
	case tag == "items_type_id":
		return mti.Items_type_id, nil
	case tag == "price":
		return mti.Price, nil
	case tag == "qty":
		return mti.Qty, nil
	case tag == "total_price":
		return mti.Total_price, nil
	case tag == "note":
		return mti.Note, nil
	case tag == "created_at":
		return mti.Created_at, nil
	case tag == "updated_at":
		return mti.Updated_at, nil
	case tag == "deleted_at":
		return mti.Deleted_at, nil
	default:
		return nil, errors.New("Value Not Found")
	}
}

func (ModelTransactionsItem) GetFieldsNeededToStoreProcess() []string {
	return []string{
		"transaction_id",
		"item_id",
		"items_type_id",
		"price",
		"qty",
		"total_price",
		"note",
		"created_at",
	}
}

func (ModelTransactionsItem) GetFieldsNeededToGetProcess() []string {
	return []string{
		"id",
		"transaction_id",
		"item_id",
		"items_type_id",
		"price",
		"qty",
		"total_price",
		"note",
		"created_at",
		"updated_at",
	}
}

func (ModelTransactionsItem) GetFieldsNeededToUpdateProcess() []string {
	return []string{
		"id",
		"transaction_id",
		"item_id",
		"items_type_id",
		"price",
		"qty",
		"total_price",
		"note",
		"updated_at",
	}
}

func (ModelTransactionsItem) GetFieldsNeededToSoftDeleteProcess() []string {
	return []string{
		"deleted_at",
	}
}

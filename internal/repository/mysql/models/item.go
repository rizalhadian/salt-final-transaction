package model

import (
	"database/sql"
	"time"
)

type ModelItem struct {
	Id            int64        `dbq:"id"`
	Items_type_id int64        `dbq:"items_type_id"`
	Name          string       `dbq:"name"`
	Description   string       `dbq:"description"`
	Price         float64      `dbq:"price"`
	Stock         int32        `dbq:"stock"`
	Status        int32        `dbq:"status"`
	Created_at    time.Time    `dbq:"created_at"`
	Updated_at    sql.NullTime `dbq:"updated_at"`
	Deleted_at    sql.NullTime `dbq:"delete_at"`
}

func (ModelItem) GetTableName() string {
	return "item"
}

func (ModelItem) GetFieldsNeededToStoreProcess() []string {
	//Belum Kepake Karena Ga Buat Fitur Ini
	return []string{}
}

func (ModelItem) GetFieldsNeededToGetProcess() []string {
	return []string{
		"id",
		"items_type_id",
		"name",
		"description",
		"price",
		"stock",
		"status",
		"created_at",
		"updated_at",
	}
}

func (ModelItem) GetFieldsNeededToUpdateProcess() []string {
	//Belum Kepake Karena Ga Buat Fitur Ini
	return []string{}
}

func (ModelItem) GetFieldsNeededToSoftDeleteProcess() []string {
	//Belum Kepake Karena Ga Buat Fitur Ini
	return []string{}
}

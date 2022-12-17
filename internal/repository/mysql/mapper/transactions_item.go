package mapper_mysql

import (
	"database/sql"
	"salt-final-transaction/domain/entity"
	model "salt-final-transaction/internal/repository/mysql/models"
	"time"
)

func TransactionsItemModelToEntity(model_transactions_item *model.ModelTransactionsItem) *entity.TransactionsItem {
	entity_transactions_item := &entity.TransactionsItem{}
	entity_transactions_item.SetId(model_transactions_item.Id)
	entity_transactions_item.SetTransactionId(model_transactions_item.Transaction_id)
	entity_transactions_item.SetItemId(model_transactions_item.Item_id)
	entity_transactions_item.SetItemsTypeId(model_transactions_item.Items_type_id)
	entity_transactions_item.SetPrice(model_transactions_item.Price)
	entity_transactions_item.SetQty(model_transactions_item.Qty)
	entity_transactions_item.SetTotalPrice(model_transactions_item.Total_price)
	entity_transactions_item.SetNote(model_transactions_item.Note)
	entity_transactions_item.SetCreatedAt(model_transactions_item.Created_at)
	if model_transactions_item.Updated_at != "" {
		update_at_time, err_parse := time.Parse(time.RFC3339, model_transactions_item.Updated_at)
		if err_parse != nil {
			panic(err_parse)
		}

		update_at_sql_null := sql.NullTime{Time: update_at_time, Valid: true}
		entity_transactions_item.SetUpdatedAt(update_at_sql_null)
	}

	if model_transactions_item.Deleted_at != "" {
		deleted_at_time, err_parse := time.Parse(time.RFC3339, model_transactions_item.Deleted_at)
		if err_parse != nil {
			panic(err_parse)
		}

		deleted_at_sql_null := sql.NullTime{Time: deleted_at_time, Valid: true}
		entity_transactions_item.SetUpdatedAt(deleted_at_sql_null)
	}

	return entity_transactions_item
}

func TransactionsItemEntityToModel(entity_transactions_item *entity.TransactionsItem) *model.ModelTransactionsItem {
	var updated_at_string string
	var deleted_at_string string

	if entity_transactions_item.GetUpdatedAt().Valid == true {
		updated_at_string = entity_transactions_item.GetUpdatedAt().Time.String()
	}

	if entity_transactions_item.GetDeletedAt().Valid == true {
		deleted_at_string = entity_transactions_item.GetDeletedAt().Time.String()
	}

	return &model.ModelTransactionsItem{
		Id:             entity_transactions_item.GetId(),
		Transaction_id: entity_transactions_item.GetTransactionId(),
		Item_id:        entity_transactions_item.GetItemId(),
		Items_type_id:  entity_transactions_item.GetItemsTypeId(),
		Price:          entity_transactions_item.GetPrice(),
		Qty:            entity_transactions_item.GetQty(),
		Total_price:    entity_transactions_item.GetTotalPrice(),
		Note:           entity_transactions_item.GetNote(),
		Created_at:     entity_transactions_item.GetCreatedAt(),
		Updated_at:     updated_at_string,
		Deleted_at:     deleted_at_string,
	}
}

func TransactionsItemModelListToEntityList(model_transactions_items []*model.ModelTransactionsItem) []*entity.TransactionsItem {
	entities := make([]*entity.TransactionsItem, 0)
	for _, model_transactions_item := range model_transactions_items {
		d := TransactionsItemModelToEntity(model_transactions_item)
		entities = append(entities, d)
	}
	return entities
}

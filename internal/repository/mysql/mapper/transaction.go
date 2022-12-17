package mapper_mysql

import (
	"database/sql"
	"salt-final-transaction/domain/entity"
	model "salt-final-transaction/internal/repository/mysql/models"
	"time"
)

func TransactionModelToEntity(model_transaction *model.ModelTransaction) *entity.Transaction {
	entity_transaction := &entity.Transaction{}
	entity_transaction.SetId(model_transaction.Id)
	entity_transaction.SetCustomerId(model_transaction.Customer_id)
	entity_transaction.SetTotalAmount(model_transaction.Total_amount)
	entity_transaction.SetTotalDiscountAmount(model_transaction.Total_discount_amount)
	entity_transaction.SetFinalTotalAmount(model_transaction.Final_total_amount)
	entity_transaction.SetNote(model_transaction.Note)
	entity_transaction.SetStatus(model_transaction.Status)
	entity_transaction.SetRollbackTransactionId(model_transaction.Rollback_transaction_id)
	entity_transaction.SetUpdateTransactionId(model_transaction.Update_transaction_id)
	entity_transaction.SetCreatedAt(model_transaction.Created_at)
	if model_transaction.Updated_at != "" {
		update_at_time, err_parse := time.Parse(time.RFC3339, model_transaction.Updated_at)
		if err_parse != nil {
			panic(err_parse)
		}

		update_at_sql_null := sql.NullTime{Time: update_at_time, Valid: true}

		entity_transaction.SetUpdatedAt(update_at_sql_null)
	}

	if model_transaction.Deleted_at != "" {
		deleted_at_time, err_parse := time.Parse(time.RFC3339, model_transaction.Deleted_at)
		if err_parse != nil {
			panic(err_parse)
		}

		deleted_at_sql_null := sql.NullTime{Time: deleted_at_time, Valid: true}

		entity_transaction.SetDeletedAt(deleted_at_sql_null)
	}
	entity_transaction.SetIsGeneratedVoucherSucceed(model_transaction.Is_generated_voucher_succeed)

	return entity_transaction
}

func TransactionEntityToModel(entity_transaction *entity.Transaction) *model.ModelTransaction {
	var updated_at_string string
	var deleted_at_string string

	if entity_transaction.GetUpdatedAt().Valid == true {
		updated_at_string = entity_transaction.GetUpdatedAt().Time.String()
	}

	if entity_transaction.GetDeletedAt().Valid == true {
		deleted_at_string = entity_transaction.GetDeletedAt().Time.String()
	}

	return &model.ModelTransaction{
		Id:                           entity_transaction.GetId(),
		Customer_id:                  entity_transaction.GetCustomerId(),
		Total_amount:                 entity_transaction.GetTotalAmount(),
		Total_discount_amount:        entity_transaction.GetTotalDiscountAmount(),
		Final_total_amount:           entity_transaction.GetFinalTotalAmount(),
		Note:                         entity_transaction.GetNote(),
		Status:                       entity_transaction.GetStatus(),
		Rollback_transaction_id:      entity_transaction.GetRollbackTransactionId(),
		Update_transaction_id:        entity_transaction.GetUpdateTransactionId(),
		Created_at:                   entity_transaction.GetCreatedAt(),
		Updated_at:                   updated_at_string,
		Deleted_at:                   deleted_at_string,
		Is_generated_voucher_succeed: entity_transaction.GetIsGeneratedVoucherSucceed(),
	}
}

func TransactionModelListToEntityList(model_transactions []*model.ModelTransaction) []*entity.Transaction {
	entities := make([]*entity.Transaction, 0)
	for _, model_transaction := range model_transactions {
		d := TransactionModelToEntity(model_transaction)
		entities = append(entities, d)
	}
	return entities
}

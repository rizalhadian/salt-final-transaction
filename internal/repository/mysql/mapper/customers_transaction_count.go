package mapper_mysql

import (
	"salt-final-transaction/domain/entity"
	model "salt-final-transaction/internal/repository/mysql/models"
)

func CustomersTransactionCountModelToEntity(model_customers_transaction_count *model.ModelCustomersTransactionCount) *entity.CustomersTransactionCount {
	entity_customer_transaction_count := &entity.CustomersTransactionCount{}
	entity_customer_transaction_count.SetId(model_customers_transaction_count.Id)
	entity_customer_transaction_count.SetCustomerId(model_customers_transaction_count.Customer_id)
	entity_customer_transaction_count.SetTotalTransactionSpend(model_customers_transaction_count.Total_transaction_spend)
	entity_customer_transaction_count.SetTransactionCount(model_customers_transaction_count.Transaction_count)
	entity_customer_transaction_count.SetFirstTransactionDatetime(model_customers_transaction_count.First_transaction_datetime)
	entity_customer_transaction_count.SetLastTransactionDatetime(model_customers_transaction_count.Last_transaction_datetime)

	return entity_customer_transaction_count
}

func CustomersTransactionCountEntityToModel(entity_customer_transaction_count *entity.CustomersTransactionCount) *model.ModelCustomersTransactionCount {

	return &model.ModelCustomersTransactionCount{
		Id:                         entity_customer_transaction_count.GetId(),
		Customer_id:                entity_customer_transaction_count.GetCustomerId(),
		Total_transaction_spend:    entity_customer_transaction_count.GetTotalTransactionSpend(),
		Transaction_count:          entity_customer_transaction_count.GetTransactionCount(),
		First_transaction_datetime: entity_customer_transaction_count.GetFirstTransactionDatetime(),
		Last_transaction_datetime:  entity_customer_transaction_count.GetLastTransactionDatetime(),
	}
}

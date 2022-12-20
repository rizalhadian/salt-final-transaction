package interface_usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceUsecaseCustomersTransactionCount interface {
	GetByCustomerId(ctx context.Context, customer_id int64) (*entity.CustomersTransactionCount, error)
}

package interface_repo

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceRepoCustomersTransactionCount interface {
	Store(ctx context.Context, t *entity.CustomersTransactionCount) error
	Update(ctx context.Context, t *entity.CustomersTransactionCount) error
	GetByCustomerId(ctx context.Context, customer_id int64) (*entity.CustomersTransactionCount, error)
}

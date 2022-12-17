package interface_repo

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceRepoTransactionsItem interface {
	Store(ctx context.Context, transaction_id int64, trx_items []*entity.TransactionsItem) error
	// Update(ctx context.Context, transaction_id int64, trx_items []*entity.TransactionsItem) error
	// Delete(ctx context.Context, ids []int64) error
	// HardDelete(ctx context.Context, ids []int64) error
	// GetById(ctx context.Context, id int64) (*entity.TransactionsItem, error)
	GetByTransactionId(ctx context.Context, transaction_id int64) (res []*entity.TransactionsItem, err error)
}

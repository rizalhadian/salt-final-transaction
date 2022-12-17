package interface_usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceUsecaseTransactionsItem interface {
	Store(ctx context.Context, transaction_id int64, dto_transaction *entity.DTOTransactionsItem) (id int64, errs error)
	Update(ctx context.Context, transaction_id int64, dto_transaction *entity.DTOTransactionsItem) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*entity.TransactionsItem, error)
	GetList(ctx context.Context, transaction_id int64) (res []*entity.TransactionsItem, err error)
}

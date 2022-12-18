package interface_usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceUsecaseTransaction interface {
	Store(ctx context.Context, dto_transaction *entity.DTOTransaction) (res *entity.Transaction, errs error)
	Update(ctx context.Context, dto_transaction *entity.DTOTransaction) error
	Delete(ctx context.Context, customer_id int64, id int64) error
	GetById(ctx context.Context, customer_id int64, id int64) (*entity.Transaction, error)
	GetByCustomerIdList(ctx context.Context, customer_id int64, page int32) (res []*entity.Transaction, err error)
	GetList(ctx context.Context, page int32) (res []*entity.Transaction, err error)
}

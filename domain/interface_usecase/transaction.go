package interface_usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceUsecaseTransaction interface {
	Store(ctx context.Context, a *entity.Transaction) (id int64, errs error)
	Update(ctx context.Context, a *entity.Transaction) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	GetList(ctx context.Context, page int32) (res []*entity.Transaction, err error)
}

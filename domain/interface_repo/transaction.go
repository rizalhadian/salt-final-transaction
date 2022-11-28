package interface_repo

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceRepoTransaction interface {
	Store(ctx context.Context, a *entity.Transaction) (id int64, errs error)
	Update(ctx context.Context, a *entity.Transaction) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	GetPaginate(ctx context.Context, limit int32, offset int32) (res []*entity.Transaction, err error)
}

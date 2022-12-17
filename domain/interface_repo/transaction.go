package interface_repo

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceRepoTransaction interface {
	Store(ctx context.Context, t *entity.Transaction) error
	Update(ctx context.Context, t *entity.Transaction) error
	Delete(ctx context.Context, id int64) error
	HardDelete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	GetByCustomerIdList(ctx context.Context, customer_id int64, limit int32, offset int32) (res []*entity.Transaction, err error)
	GetList(ctx context.Context, limit int32, offset int32) (res []*entity.Transaction, err error)
}

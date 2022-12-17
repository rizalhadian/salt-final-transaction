package interface_usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceUsecaseItem interface {
	GetById(ctx context.Context, id int64) (*entity.Item, error)
}

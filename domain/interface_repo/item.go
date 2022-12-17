package interface_repo

import (
	"context"
	"salt-final-transaction/domain/entity"
)

type InterfaceRepoItem interface {
	GetById(ctx context.Context, id int64) (*entity.Item, error)
}

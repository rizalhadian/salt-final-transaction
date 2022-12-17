package usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	"salt-final-transaction/domain/interface_usecase"
)

type UsecaseItem struct {
	repoItem interface_repo.InterfaceRepoItem
}

func NewUsecaseItem(interfaceRepoItem interface_repo.InterfaceRepoItem) interface_usecase.InterfaceUsecaseItem {
	return &UsecaseItem{
		repoItem: interfaceRepoItem,
	}
}

func (ui *UsecaseItem) GetById(ctx context.Context, id int64) (*entity.Item, error) {
	item, err := ui.repoItem.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

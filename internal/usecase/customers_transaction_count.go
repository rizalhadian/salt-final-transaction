package usecase

import (
	"context"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	"salt-final-transaction/domain/interface_usecase"
)

type UsecaseCustomersTransactionCount struct {
	repoCustomersTransactionCount interface_repo.InterfaceRepoCustomersTransactionCount
}

func NewUsecaseCustomersTransactionCount(interfaceRepoCustomersTransactionCount interface_repo.InterfaceRepoCustomersTransactionCount) interface_usecase.InterfaceUsecaseCustomersTransactionCount {
	return &UsecaseCustomersTransactionCount{
		repoCustomersTransactionCount: interfaceRepoCustomersTransactionCount,
	}
}

func (ui *UsecaseCustomersTransactionCount) GetByCustomerId(ctx context.Context, customer_id int64) (*entity.CustomersTransactionCount, error) {
	item, err := ui.repoCustomersTransactionCount.GetByCustomerId(ctx, customer_id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

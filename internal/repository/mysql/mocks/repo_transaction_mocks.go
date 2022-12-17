package mocks

import (
	"context"
	"salt-final-transaction/domain/entity"

	"github.com/stretchr/testify/mock"
)

type RepoCustomerMocks struct {
	mock.Mock
}

func (rcm RepoCustomerMocks) Store(ctx context.Context, a *entity.Transaction) (id int64, errs error) {
	args := rcm.Called(ctx, a)
	return int64(args.Int(0)), args.Error(1)
}

func (rcm RepoCustomerMocks) Update(ctx context.Context, a *entity.Transaction) error {
	args := rcm.Called(ctx, a)
	return args.Error(0)
}

func (rcm RepoCustomerMocks) Delete(ctx context.Context, id int64) error {
	args := rcm.Called(ctx, id)
	return args.Error(0)
}

func (rcm RepoCustomerMocks) GetById(ctx context.Context, id int64) (*entity.Transaction, error) {
	args := rcm.Called(ctx, id)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (rcm RepoCustomerMocks) GetPaginate(ctx context.Context, limit int32, offset int32) (res []*entity.Transaction, err error) {
	args := rcm.Called(ctx, limit, offset)
	return args.Get(0).([]*entity.Transaction), args.Error(1)
}

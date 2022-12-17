package mocks

import (
	"context"
	infrastructure_customer_http_response "salt-final-transaction/internal/infrastructure/customer/http_response"

	"github.com/stretchr/testify/mock"
)

type InfrastructureCustomerMocks struct {
	mock.Mock
}

func (icm *InfrastructureCustomerMocks) GetById(ctx context.Context, customer_id int64) (customer *infrastructure_customer_http_response.Customer, http_response_code int, err error) {
	args := icm.Called(ctx, customer_id)
	return args.Get(0).(*infrastructure_customer_http_response.Customer), args.Int(1), args.Error(2)
}

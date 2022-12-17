package infrastructure_customer_interface

import (
	"context"
	infrastructure_customer_http_response "salt-final-transaction/internal/infrastructure/customer/http_response"
)

type InterfaceInfrastructureCustomer interface {
	GetById(ctx context.Context, id int64) (customer *infrastructure_customer_http_response.Customer, err error)
}

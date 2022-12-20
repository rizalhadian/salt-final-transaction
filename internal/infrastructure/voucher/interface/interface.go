package infrastructure_customer_interface

import (
	"context"
)

// type InterfaceInfrastructureVoucher interface {
// 	GetById(ctx context.Context, id int64) (customer *infrastructure_customer_http_response.Customer, http_response_code int, err error)
// }

type InterfaceInfrastructureVoucher interface {
	// GetByCode(ctx context.Context, code string, transaction *entity.Transaction) (customer *infrastructure_voucher_http_response.CustomersVoucher, http_response_code int, err error)
	// SetVoucherIsUsedByCode(ctx context.Context, code string, transaction_id int64) (http_response_code int, err error)
	GenerateVoucher(ctx context.Context, customer_id int64) (err error)
}

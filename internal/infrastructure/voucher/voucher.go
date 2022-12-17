package infrastructure_voucher

import (
	"context"
	"net/http"
	"salt-final-transaction/domain/entity"
	infrastructure_voucher_http_response "salt-final-transaction/internal/infrastructure/voucher/http_response"
	infrastructure_voucher_interface "salt-final-transaction/internal/infrastructure/voucher/interface"
)

type InfrastructureVoucher struct {
	http_client   http.Client
	base_endpoint string
}

func NewInfrastructureVoucher(http_client_value http.Client, base_endpoint_value string) infrastructure_voucher_interface.InterfaceInfrastructureVoucher {
	return &InfrastructureVoucher{
		// base_endpoint: "http://localhost:8080/customer",
		http_client:   http_client_value,
		base_endpoint: base_endpoint_value,
	}
}

func (icv *InfrastructureVoucher) GetByCode(ctx context.Context, code string, transaction *entity.Transaction) (customer *infrastructure_voucher_http_response.CustomersVoucher, http_response_code int, err error)
func (icv *InfrastructureVoucher) SetVoucherIsUsedByCode(ctx context.Context, code string, transaction_id int64) (http_response_code int, err error)
func (icv *InfrastructureVoucher) GenerateVoucher(ctx context.Context, transaction *entity.Transaction) (http_response_code int, err error)

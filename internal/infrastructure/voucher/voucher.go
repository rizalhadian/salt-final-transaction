package infrastructure_voucher

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	infrastructure_voucher_interface "salt-final-transaction/internal/infrastructure/voucher/interface"
	"strconv"
	"time"
)

type InfrastructureVoucher struct {
	http_client   http.Client
	base_endpoint string
}

func NewInfrastructureVoucher(http_client_value http.Client, base_endpoint_value string) infrastructure_voucher_interface.InterfaceInfrastructureVoucher {
	return &InfrastructureVoucher{
		// base_endpoint: "http://localhost:8080/api/voucher/",
		http_client:   http_client_value,
		base_endpoint: base_endpoint_value,
	}
}

func (icv *InfrastructureVoucher) GenerateVoucher(ctx context.Context, customer_id int64) (err error) {
	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	customer_id_string := strconv.Itoa(int(customer_id))

	endpoint := icv.base_endpoint + "/generate/" + customer_id_string
	fmt.Println("!! Generate Voucher !!")
	fmt.Println(endpoint)

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return errors.New("500")
	}

	response, err := icv.http_client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return errors.New("404")
	}

	return nil
}

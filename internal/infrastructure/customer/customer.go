package infrastructure_customer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	infrastructure_customer_http_response "salt-final-transaction/internal/infrastructure/customer/http_response"
	infrastructure_customer_interface "salt-final-transaction/internal/infrastructure/customer/interface"
	"strconv"
	"time"
)

type InfrastructureCustomer struct {
	http_client   *http.Client
	base_endpoint string
}

func NewInfrastructureCustomer(http_client_value http.Client, base_endpoint_value string) infrastructure_customer_interface.InterfaceInfrastructureCustomer {
	return &InfrastructureCustomer{
		// base_endpoint: "http://localhost:8080/customer",
		http_client:   &http_client_value,
		base_endpoint: base_endpoint_value,
	}
}

func (ic InfrastructureCustomer) GetById(ctx context.Context, customer_id int64) (customer *infrastructure_customer_http_response.Customer, err error) {
	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	endpoint := fmt.Sprint(ic.base_endpoint+"/%s", strconv.Itoa(int(customer_id)))
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("500")
	}

	response, err := ic.http_client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, errors.New("404")
	}

	customer_resp := infrastructure_customer_http_response.Customer{}
	err = json.NewDecoder(response.Body).Decode(&customer_resp)
	if err != nil {
		return nil, err
	}

	return &customer_resp, nil
}

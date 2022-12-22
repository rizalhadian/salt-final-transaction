package infrastructure_voucher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"salt-final-transaction/domain/entity"
	infrastructure_voucher_http_request "salt-final-transaction/internal/infrastructure/voucher/http_request"
	infrastructure_voucher_http_response "salt-final-transaction/internal/infrastructure/voucher/http_response"
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

func (icv *InfrastructureVoucher) Redeem(ctx context.Context, transaction *entity.Transaction, vouchers_codes []string) (vouchers []*infrastructure_voucher_http_response.CustomersVoucher, err error) {
	req_transaction := infrastructure_voucher_http_request.Transaction{
		Id:           transaction.GetId(),
		Customer_id:  transaction.GetCustomerId(),
		Total_amount: transaction.GetTotalAmount(),
		// TransactionsItems            []TransactionsItem            `json:"items,omitempty"`
		// TransactionsVouchersRedeemed []TransactionsVoucherRedeemed `json:"vouchers_redeemed,omitempty"`
	}

	var req_transactions_items []infrastructure_voucher_http_request.TransactionsItem
	for _, transaction_item := range transaction.GetItems() {
		req_transactions_item := infrastructure_voucher_http_request.TransactionsItem{
			Item_id:       transaction_item.GetId(),
			Items_type_id: transaction_item.GetItemsTypeId(),
			Price:         transaction_item.GetPrice(),
			Qty:           transaction_item.GetQty(),
		}

		req_transactions_items = append(req_transactions_items, req_transactions_item)
	}
	req_transaction.TransactionsItems = req_transactions_items

	var req_transactions_vouchers_redeemed []infrastructure_voucher_http_request.TransactionsVoucherRedeemed
	for _, vouchers_code := range vouchers_codes {
		req_transactions_voucher_redeemed := infrastructure_voucher_http_request.TransactionsVoucherRedeemed{
			Code: vouchers_code,
		}

		req_transactions_vouchers_redeemed = append(req_transactions_vouchers_redeemed, req_transactions_voucher_redeemed)
	}
	req_transaction.TransactionsVouchersRedeemed = req_transactions_vouchers_redeemed

	var req bytes.Buffer
	err_encode := json.NewEncoder(&req).Encode(req_transaction)
	if err_encode != nil {
		log.Fatal(err_encode)
	}

	endpoint := icv.base_endpoint + "/redeem"
	fmt.Println("!! Redeem Voucher !!")
	fmt.Println(endpoint)

	request, err := http.NewRequest(http.MethodPost, endpoint, &req)
	if err != nil {
		fmt.Println("AAAA")
		return nil, errors.New("500")
	}

	response, response_err := icv.http_client.Do(request)
	if response_err != nil {
		log.Fatal(response_err.Error())
		fmt.Println("BBBBB")

		return nil, errors.New("500")
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, errors.New("404")
	}

	if response.StatusCode == 500 {
		fmt.Println("CCCCCC")

		return nil, errors.New("500")
	}

	if response.StatusCode == 200 {
		var resp infrastructure_voucher_http_response.SuccedRedeemResponse
		resp_decoder := json.NewDecoder(response.Body)
		err_resp_decode := resp_decoder.Decode(&resp)
		if err_resp_decode != nil {
			log.Fatal("Err Decode Resp")
			fmt.Println("DDDDDD")

			return nil, errors.New("500")
		} else {
			fmt.Println("Total Discount ====")
			fmt.Println(resp)
		}

		return resp.Data, nil
	}
	fmt.Println("ZZZZZ")
	return nil, errors.New("500")
}

package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"salt-final-transaction/domain/interface_usecase"
	http_response "salt-final-transaction/internal/delivery/http/response"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerCustomersTransactionCount struct {
	usecase_customers_transaction_count interface_usecase.InterfaceUsecaseCustomersTransactionCount
}

func NewHandlerCustomersTransactionCount(router *mux.Router, usecase_customers_transaction_count_value interface_usecase.InterfaceUsecaseCustomersTransactionCount) {
	HandlerCustomersTransnCount := &HandlerCustomersTransactionCount{
		usecase_customers_transaction_count: usecase_customers_transaction_count_value,
	}

	router.HandleFunc("/api/customer/{customer_id}/transaction-count", HandlerCustomersTransnCount.GetCount).Methods(http.MethodGet)
}

func (ht *HandlerCustomersTransactionCount) GetCount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======| Request | URI : " + r.RequestURI + " | Method : " + r.Method + " |======")

	customer_id_string := mux.Vars(r)["customer_id"]
	customer_id, customer_id_conv_err := strconv.Atoi(customer_id_string)

	if customer_id_conv_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		entities, err := ht.usecase_customers_transaction_count.GetByCustomerId(r.Context(), int64(customer_id))
		if err != nil {
			err_int, err_convert := strconv.Atoi(err.Error())
			if err_convert != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(err_int)
		}

		if entities == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusOK)

			response_transaction := http_response.MapCustomersTransactionCountResponse(entities)

			resp_skeleton := http_response.SkeletonSingleResponse{
				Success: true,
				Message: "",
				Data:    response_transaction,
			}
			resp, resp_json_err := json.Marshal(resp_skeleton)
			if resp_json_err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write(resp)
		}
	}
}

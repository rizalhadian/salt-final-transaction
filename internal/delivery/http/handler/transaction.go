package http_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_usecase"
	http_request "salt-final-transaction/internal/delivery/http/request"
	http_response "salt-final-transaction/internal/delivery/http/response"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HandlerTransaction struct {
	usecase_transaction interface_usecase.InterfaceUsecaseTransaction
}

func NewHandlerTransaction(router *mux.Router, usecase_transaction_value interface_usecase.InterfaceUsecaseTransaction) {
	HandlerTrans := &HandlerTransaction{
		usecase_transaction: usecase_transaction_value,
	}

	router.HandleFunc("/api/customer/{customer_id}/transaction", HandlerTrans.GetList).Methods(http.MethodGet)
	router.HandleFunc("/api/customer/{customer_id}/transaction", HandlerTrans.Store).Methods(http.MethodPost)
	router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.Delete).Methods(http.MethodDelete)
}

func (ht *HandlerTransaction) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======| Request | URI : " + r.RequestURI + " | Method : " + r.Method + " |======")

}

func (ht *HandlerTransaction) Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======| Request | URI : " + r.RequestURI + " | Method : " + r.Method + " |======")

	var (
		req     http_request.Transaction
		decoder = json.NewDecoder(r.Body)
		ctx_req = r.Context()
	)

	ctx_handler, cancel := context.WithTimeout(ctx_req, 60*time.Second)
	defer cancel()

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	// status_int, err := entity.ConvertTransactionStatusStringToInt(req.Status)
	// if err != nil {

	// }

	var dto_transactions_items []*entity.DTOTransactionsItem

	for _, transaction_item := range req.TransactionsItems {
		dto_transactions_item := entity.DTOTransactionsItem{
			Item_id: transaction_item.Item_id,
			Price:   transaction_item.Price,
			Qty:     transaction_item.Qty,
			Note:    transaction_item.Note,
		}

		dto_transactions_items = append(dto_transactions_items, &dto_transactions_item)
	}

	dto_transaction := &entity.DTOTransaction{
		Customer_id: int64(req.Customer_id),
		// Note:        req.Note,
		// Status:      status_int,
		Items: dto_transactions_items,
	}

	transaction_id, usecase_store_err := ht.usecase_transaction.Store(ctx_handler, dto_transaction)
	if usecase_store_err != nil {
		response_skeleton_error := http_response.SkeletonError{
			Success: false,
		}

		// response_error, err_marshal := json.Marshal(http_response.SkeletonError)
		if usecase_store_err.Error() == "Customer Not Found" {
			w.WriteHeader(http.StatusNotFound)
			response_skeleton_error.Message = "Customer Not Found"
		} else if usecase_store_err.Error() == "Internal Server Error" {
			w.WriteHeader(http.StatusInternalServerError)
			response_skeleton_error.Message = "Internal Server Error"
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response_skeleton_error.Message = "Error"

			errs_json, _ := json.Marshal(usecase_store_err.Error())

			response_article_json_map := map[string]interface{}{}
			json.Unmarshal([]byte(errs_json), &response_article_json_map)

			fmt.Println(response_article_json_map)
		}
		// resp := http_response.SkeletonError{
		// 	Success: false,
		// 	Message: usecase_store_err.Error(),
		// }
		resp_json, _ := json.Marshal(response_skeleton_error)
		w.Write(resp_json)
	}
	fmt.Println(transaction_id)
}

func (ht *HandlerTransaction) FindById(w http.ResponseWriter, r *http.Request) {
	transaction_id_string := mux.Vars(r)["id"]
	customer_id_string := mux.Vars(r)["customer_id"]

	transaction_id, transaction_id_conv_err := strconv.Atoi(transaction_id_string)
	customer_id, customer_id_conv_err := strconv.Atoi(customer_id_string)

	if transaction_id_conv_err != nil || customer_id_conv_err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		entity, err := ht.usecase_transaction.GetById(r.Context(), int64(customer_id), int64(transaction_id))
		if err != nil {
			err_int, err_convert := strconv.Atoi(err.Error())
			if err_convert != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(err_int)
		}

		if entity == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)

			response_transaction := http_response.MapTransactionResponse(entity)

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

func (ht *HandlerTransaction) Update(w http.ResponseWriter, r *http.Request) {

}

func (ht *HandlerTransaction) Delete(w http.ResponseWriter, r *http.Request) {

}

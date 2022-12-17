package http_response

import (
	"salt-final-transaction/domain/entity"
	"time"
)

type TransactionResponse struct {
	Id                      int64                       `json:"id"`
	Customer_id             int64                       `json:"customer_id"`
	Total_amount            float64                     `json:"total_amount"`
	Total_discount_amount   float64                     `json:"total_discount_amount"`
	Final_total_amount      float64                     `json:"final_total_amount"`
	Note                    string                      `json:"note"`
	Status                  int16                       `json:"status"`
	Rollback_transaction_id int64                       `json:"rollback_transaction_id"`
	Update_transaction_id   int64                       `json:"update_transaction_id"`
	Items                   []*TransactionsItemResponse `json:"items"`
	Created_at              time.Time                   `json:"created_at"`
}

type TransactionsItemResponse struct {
	Id            int64     `json:"id"`
	Item_id       int64     `json:"item_id"`
	Items_type_id int64     `json:"item_type_id"`
	Price         float64   `json:"price"`
	Qty           int32     `json:"qty"`
	Total_price   float64   `json:"total_price"`
	Note          string    `json:"note"`
	Created_at    time.Time `json:"created_at"`
}

func MapTransactionResponse(entity_transaction *entity.Transaction) *TransactionResponse {

	var transaction_items []*TransactionsItemResponse

	for _, entity_transactions_item := range entity_transaction.GetItems() {
		transaction_item := TransactionsItemResponse{
			Id:            entity_transactions_item.GetId(),
			Item_id:       entity_transactions_item.GetItemId(),
			Items_type_id: entity_transactions_item.GetItemsTypeId(),
			Price:         entity_transactions_item.GetTotalPrice(),
			Qty:           entity_transactions_item.GetQty(),
			Total_price:   entity_transactions_item.GetTotalPrice(),
			Note:          entity_transactions_item.GetNote(),
			Created_at:    entity_transactions_item.GetCreatedAt(),
		}

		transaction_items = append(transaction_items, &transaction_item)
	}

	return &TransactionResponse{
		Id:                      entity_transaction.GetId(),
		Customer_id:             entity_transaction.GetCustomerId(),
		Total_amount:            entity_transaction.GetTotalAmount(),
		Total_discount_amount:   entity_transaction.GetTotalDiscountAmount(),
		Final_total_amount:      entity_transaction.GetFinalTotalAmount(),
		Note:                    entity_transaction.GetNote(),
		Status:                  entity_transaction.GetStatus(),
		Rollback_transaction_id: entity_transaction.GetRollbackTransactionId(),
		Items:                   transaction_items,
		Update_transaction_id:   entity_transaction.GetUpdateTransactionId(),
		Created_at:              entity_transaction.GetCreatedAt(),
	}

	// response_transaction_json, err_marshal := json.Marshal(transaction)
	// if err_marshal != nil {
	// 	return nil, errors.New("500")
	// }
	// response_transaction_json_map := map[string]interface{}{}
	// json.Unmarshal([]byte(response_transaction_json), &response_transaction_json_map)

	// return response_transaction_json_map, nil
}

// func MapTransactionsItemResponse(entity_transaction *entity.Transaction) ([]map[string]interface{}, error) {

// 	transaction := TransactionResponse{
// 		Id:                      entity_transaction.GetId(),
// 		Customer_id:             entity_transaction.GetCustomerId(),
// 		Total_amount:            entity_transaction.GetTotalAmount(),
// 		Total_discount_amount:   entity_transaction.GetTotalDiscountAmount(),
// 		Final_total_amount:      entity_transaction.GetFinalTotalAmount(),
// 		Note:                    entity_transaction.GetNote(),
// 		Status:                  entity_transaction.GetStatus(),
// 		Rollback_transaction_id: entity_transaction.GetRollbackTransactionId(),
// 		Update_transaction_id:   entity_transaction.GetUpdateTransactionId(),
// 		Created_at:              entity_transaction.GetCreatedAt(),
// 	}

// 	response_transaction_json, err_marshal := json.Marshal(transaction)
// 	if err_marshal != nil {
// 		return nil, errors.New("500")
// 	}
// 	response_transaction_json_map := map[string]interface{}{}
// 	json.Unmarshal([]byte(response_transaction_json), &response_transaction_json_map)

// 	return response_transaction_json_map, nil
// }

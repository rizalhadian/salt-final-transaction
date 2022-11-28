package entity_test

import (
	"database/sql"
	"salt-final-transaction/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewTransactionsItem_Positive(t *testing.T) {

	type test_data_transactions_item_struct struct {
		request        entity.DTOTransactionsItem
		expected_error error
	}

	test_dto_transactions_items := []test_data_transactions_item_struct{
		test_data_transactions_item_struct{
			request: entity.DTOTransactionsItem{
				Id:                   1,
				Transaction_id:       1,
				Item_id:              1,
				Items_type_id:        1,
				Price:                100000.00,
				Qty:                  1,
				Total_price:          100000.00,
				Note:                 "",
				Customers_voucher_id: sql.NullInt64{Valid: false},
				Voucher_id:           sql.NullInt64{Valid: false},
				Voucher_code:         "",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: nil,
		},
	}

	for _, test_dto_transactions_item := range test_dto_transactions_items {
		data, err := entity.NewTransactionsItem(test_dto_transactions_item.request)
		assert.NotNil(t, data)
		assert.Nil(t, err)
	}
}

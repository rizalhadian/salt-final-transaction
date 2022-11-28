package entity_test

import (
	"database/sql"
	"errors"
	"salt-final-transaction/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewTransactionsItem_Negative(t *testing.T) {

	type test_data_transactions_item_struct struct {
		request        entity.DTOTransactionsItem
		expected_error error
	}

	test_dto_transactions_items := []test_data_transactions_item_struct{
		test_data_transactions_item_struct{
			request: entity.DTOTransactionsItem{
				Id:                   1,
				Transaction_id:       0,
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
			expected_error: errors.New("Transaction_id is required"),
		},
		test_data_transactions_item_struct{
			request: entity.DTOTransactionsItem{
				Id:                   1,
				Transaction_id:       1,
				Item_id:              0,
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
			expected_error: errors.New("Item_id is required"),
		},
		test_data_transactions_item_struct{
			request: entity.DTOTransactionsItem{
				Id:                   1,
				Transaction_id:       1,
				Item_id:              1,
				Items_type_id:        0,
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
			expected_error: errors.New("Items_type_id is required"),
		},
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
				Customers_voucher_id: sql.NullInt64{Valid: true},
				Voucher_id:           sql.NullInt64{Valid: true, Int64: 1},
				Voucher_code:         "",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Customers_voucher_id not null, Customers_voucher_id cannot be 0"),
		},
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
				Customers_voucher_id: sql.NullInt64{Valid: true, Int64: 1},
				Voucher_id:           sql.NullInt64{Valid: true},
				Voucher_code:         "",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher_id not null, Voucher_id cannot be 0"),
		},
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
				Customers_voucher_id: sql.NullInt64{Valid: true, Int64: 1},
				Voucher_id:           sql.NullInt64{Valid: true, Int64: 1},
				Voucher_code:         "",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher is used, Customers_voucher_id, Voucher_id, and Voucher_code is required"),
		},
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
				Customers_voucher_id: sql.NullInt64{Valid: true, Int64: 1},
				Voucher_id:           sql.NullInt64{Valid: false},
				Voucher_code:         "asd",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher is used, Customers_voucher_id, Voucher_id, and Voucher_code is required"),
		},
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
				Voucher_id:           sql.NullInt64{Valid: true, Int64: 1},
				Voucher_code:         "asd",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher is used, Customers_voucher_id, Voucher_id, and Voucher_code is required"),
		},
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
				Voucher_code:         "asd",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher is used, Customers_voucher_id, Voucher_id, and Voucher_code is required"),
		},
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
				Voucher_code:         "asd",
				Discount_percentage:  0,
				Discount_amount:      0.00,
				Final_price:          100000.00,
				Created_at:           time.Now(),
				Updated_at:           sql.NullTime{Valid: false},
				Deleted_at:           sql.NullTime{Valid: false},
			},
			expected_error: errors.New("If Voucher is used, Customers_voucher_id, Voucher_id, and Voucher_code is required"),
		},
	}

	for _, test_dto_transactions_item := range test_dto_transactions_items {
		data, err := entity.NewTransactionsItem(test_dto_transactions_item.request)
		assert.Nil(t, data)
		assert.NotNil(t, err)
		assert.Equal(t, test_dto_transactions_item.expected_error, err)
	}
}

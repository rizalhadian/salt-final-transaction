package entity_test

import (
	"database/sql"
	"errors"
	"salt-final-transaction/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewTransaction_Negative(t *testing.T) {

	type test_data_transaction_struct struct {
		request        entity.DTOTransaction
		expected_error error
	}

	test_dto_transactions := []test_data_transaction_struct{
		test_data_transaction_struct{
			request: entity.DTOTransaction{
				Id:                           1,
				Customer_id:                  0,
				Total_amount:                 100000.0,
				Total_discount_amount:        0.0,
				Final_total_amount:           100000.0,
				Note:                         "",
				Status:                       0,
				Rollback_transaction_id:      sql.NullInt64{Valid: false},
				Update_transaction_id:        sql.NullInt64{Valid: false},
				Created_at:                   time.Now(),
				Updated_at:                   sql.NullTime{Valid: false},
				Deleted_at:                   sql.NullTime{Valid: false},
				Is_generated_voucher_succeed: true,
			},
			expected_error: errors.New("Status is required"),
		},
		// This test should be on usecase
		// test_data_transaction_struct{
		// 	request: entity.DTOTransaction{
		// 		Id:                    2,
		// 		Customer_id:           0,
		// 		Total_amount:          -100000.0,
		// 		Total_discount_amount: 0.0,
		// 		Final_total_amount:    -200000.0,
		// 		Note:                  "",
		// 		Status:                1,
		// 		Rollback_transaction_id: sql.NullInt64{
		// 			Int64: int64(1),
		// 			Valid: true,
		// 		},
		// 		Update_transaction_id: sql.NullInt64{
		// 			Valid: false,
		// 		},
		// 		Created_at: time.Now(),
		// 		Updated_at: sql.NullTime{
		// 			Valid: false,
		// 		},
		// 		Deleted_at: sql.NullTime{
		// 			Valid: false,
		// 		},
		// 		Is_generated_voucher_succeed: true,
		// 	},
		// 	expected_error: errors.New("Final_total_amount is incorrect"),
		// },
	}

	for _, test_dto_transaction := range test_dto_transactions {
		data, err := entity.NewTransaction(test_dto_transaction.request)
		assert.NotNil(t, err)
		assert.Equal(t, test_dto_transaction.expected_error, err)
		assert.Nil(t, data)
	}
}

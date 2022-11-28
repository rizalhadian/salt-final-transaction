package entity_test

import (
	"database/sql"
	"salt-final-transaction/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewTransaction_Positive(t *testing.T) {

	type test_data_transaction_struct struct {
		request        entity.DTOTransaction
		expected_error error
	}

	test_dto_transactions_positive := []test_data_transaction_struct{
		test_data_transaction_struct{
			request: entity.DTOTransaction{
				Id:                           1,
				Customer_id:                  0,
				Total_amount:                 100000.0,
				Total_discount_amount:        0.0,
				Final_total_amount:           100000.0,
				Note:                         "",
				Status:                       1,
				Rollback_transaction_id:      sql.NullInt64{Valid: false},
				Update_transaction_id:        sql.NullInt64{Valid: false},
				Created_at:                   time.Now(),
				Updated_at:                   sql.NullTime{Valid: false},
				Deleted_at:                   sql.NullTime{Valid: false},
				Is_generated_voucher_succeed: true,
			},
			expected_error: nil,
		},
		test_data_transaction_struct{
			request: entity.DTOTransaction{
				Id:                           2,
				Customer_id:                  0,
				Total_amount:                 -100000.0,
				Total_discount_amount:        0.0,
				Final_total_amount:           -100000.0,
				Note:                         "",
				Status:                       1,
				Rollback_transaction_id:      sql.NullInt64{Valid: true, Int64: int64(1)},
				Update_transaction_id:        sql.NullInt64{Valid: false},
				Created_at:                   time.Now(),
				Updated_at:                   sql.NullTime{Valid: false},
				Deleted_at:                   sql.NullTime{Valid: false},
				Is_generated_voucher_succeed: true,
			},
			expected_error: nil,
		},
		test_data_transaction_struct{
			request: entity.DTOTransaction{
				Id:                           3,
				Customer_id:                  0,
				Total_amount:                 1000000.0,
				Total_discount_amount:        0.0,
				Final_total_amount:           1000000.0,
				Note:                         "",
				Status:                       1,
				Rollback_transaction_id:      sql.NullInt64{Valid: false},
				Update_transaction_id:        sql.NullInt64{Valid: true, Int64: int64(1)},
				Created_at:                   time.Now(),
				Updated_at:                   sql.NullTime{Valid: false},
				Deleted_at:                   sql.NullTime{Valid: false},
				Is_generated_voucher_succeed: true,
			},
			expected_error: nil,
		},
	}

	for _, test_dto_transaction_positive := range test_dto_transactions_positive {
		data, err := entity.NewTransaction(test_dto_transaction_positive.request)
		assert.NotNil(t, data)
		assert.Nil(t, err)
	}
}

package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"salt-final-transaction/domain/entity"
	infrastructure_customer "salt-final-transaction/internal/infrastructure/customer"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	usecase "salt-final-transaction/internal/usecase"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type UsecaseTransactionNegativeTest struct {
	dto_transaction         entity.DTOTransaction
	expected_transaction_id int64
	expected_err            error
}

func Test_Transaction_Store_Negative(t *testing.T) {

	usecase_transaction_negative_test_datas := []UsecaseTransactionNegativeTest{

		UsecaseTransactionNegativeTest{
			dto_transaction: entity.DTOTransaction{
				Customer_id:           0,
				Total_amount:          0,
				Total_discount_amount: 0,
				Final_total_amount:    0,
				Note:                  "",
				Items: []*entity.DTOTransactionsItem{
					&entity.DTOTransactionsItem{
						Item_id:       999,
						Items_type_id: 4,
						Price:         100000.00,
						Qty:           1,
						Total_price:   100000.00,
						Note:          "Note",
					},
					&entity.DTOTransactionsItem{
						Item_id:       11,
						Items_type_id: 4,
						Price:         100000.00,
						Qty:           1,
						Total_price:   100000.00,
						Note:          "Note",
					},
				},
				Status:                       112,
				Created_at:                   time.Now(),
				Is_generated_voucher_succeed: false,
			},
			expected_transaction_id: 0,
			expected_err:            errors.New("400"),
		},
		UsecaseTransactionNegativeTest{
			dto_transaction: entity.DTOTransaction{
				Customer_id:           0,
				Total_amount:          0,
				Total_discount_amount: 0,
				Final_total_amount:    0,
				Note:                  "",
				Items: []*entity.DTOTransactionsItem{
					&entity.DTOTransactionsItem{
						Item_id:       10,
						Items_type_id: 4,
						Price:         100000.00,
						Qty:           10000,
						Total_price:   1000000000.00,
						Note:          "Note",
					},
					&entity.DTOTransactionsItem{
						Item_id:       11,
						Items_type_id: 4,
						Price:         100000.00,
						Qty:           1,
						Total_price:   100000.00,
						Note:          "Note",
					},
				},
				Status:                       112,
				Created_at:                   time.Now(),
				Is_generated_voucher_succeed: false,
			},
			expected_transaction_id: 0,
			expected_err:            errors.New("Insufficient Stock"),
		},
		UsecaseTransactionNegativeTest{
			dto_transaction: entity.DTOTransaction{
				Customer_id:           0,
				Total_amount:          0,
				Total_discount_amount: 0,
				Final_total_amount:    0,
				Note:                  "",
				Items: []*entity.DTOTransactionsItem{
					&entity.DTOTransactionsItem{
						Item_id:       10,
						Items_type_id: 4,
						Price:         50000.00,
						Qty:           1,
						Total_price:   50000.00,
						Note:          "Note",
					},
					&entity.DTOTransactionsItem{
						Item_id:       11,
						Items_type_id: 4,
						Price:         100000.00,
						Qty:           1,
						Total_price:   100000.00,
						Note:          "Note",
					},
				},
				Status:                       112,
				Created_at:                   time.Now(),
				Is_generated_voucher_succeed: false,
			},
			expected_transaction_id: 0,
			expected_err:            errors.New("Price Changed"),
		},
		UsecaseTransactionNegativeTest{
			dto_transaction: entity.DTOTransaction{
				Customer_id:                  0,
				Total_amount:                 0,
				Total_discount_amount:        0,
				Final_total_amount:           0,
				Note:                         "",
				Items:                        []*entity.DTOTransactionsItem{},
				Status:                       112,
				Created_at:                   time.Now(),
				Is_generated_voucher_succeed: false,
			},
			expected_transaction_id: 0,
			expected_err:            errors.New("Transaction's Items is required"),
		},
		// UsecaseTransactionNegativeTest{
		// 	dto_transaction: entity.DTOTransaction{
		// 		Customer_id:           9999999,
		// 		Total_amount:          0,
		// 		Total_discount_amount: 0,
		// 		Final_total_amount:    0,
		// 		Note:                  "",
		// 		Items: []*entity.DTOTransactionsItem{
		// 			&entity.DTOTransactionsItem{
		// 				Item_id:       10,
		// 				Items_type_id: 4,
		// 				Price:         100000.00,
		// 				Qty:           1,
		// 				Total_price:   100000.00,
		// 				Note:          "Note",
		// 			},
		// 			&entity.DTOTransactionsItem{
		// 				Item_id:       11,
		// 				Items_type_id: 4,
		// 				Price:         100000.00,
		// 				Qty:           1,
		// 				Total_price:   100000.00,
		// 				Note:          "Note",
		// 			},
		// 		},
		// 		Status:                       112,
		// 		Created_at:                   time.Now(),
		// 		Is_generated_voucher_succeed: false,
		// 	},
		// 	expected_transaction_id: 0,
		// 	expected_err:            errors.New("404"),
		// },
	}

	var (
		ctx             = context.Background()
		connectionMysql = pkg_database_mysql.InitDBMysql()
		http_client     = http.Client{}
		// ============ Infrastructure
		infrastrctureCustomer = infrastructure_customer.NewInfrastructureCustomer(http_client, "http://localhost:8080/customer")
		repoTransaction       = repository_mysql.NewRepoTransaction(connectionMysql)
		repoTransactionsItem  = repository_mysql.NewRepoTransactionsItem(connectionMysql)
		repoItem              = repository_mysql.NewRepoItem(connectionMysql)

		usecaseTransaction = usecase.NewUsecaseTransaction(infrastrctureCustomer, repoTransaction, repoTransactionsItem, repoItem)
	)

	for _, usecase_transaction_negative_test_data := range usecase_transaction_negative_test_datas {
		transaction_id, usecase_transaction_store_err := usecaseTransaction.Store(ctx, &usecase_transaction_negative_test_data.dto_transaction)
		assert.Equal(t, int64(usecase_transaction_negative_test_data.expected_transaction_id), transaction_id)
		assert.NotNil(t, usecase_transaction_store_err)
		assert.Equal(t, usecase_transaction_negative_test_data.expected_err, usecase_transaction_store_err)
	}

}

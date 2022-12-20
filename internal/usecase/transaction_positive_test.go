package usecase_test

import (
	"context"
	"net/http"
	"salt-final-transaction/domain/entity"
	infrastructure_customer "salt-final-transaction/internal/infrastructure/customer"
	infrastructure_voucher "salt-final-transaction/internal/infrastructure/voucher"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	usecase "salt-final-transaction/internal/usecase"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type UsecaseTransactionPositiveTest struct {
	dto_transaction entity.DTOTransaction
	expected_err    error
}

var (
	ctx             = context.Background()
	connectionMysql = pkg_database_mysql.InitDBMysql()
	http_client     = http.Client{}
	// ============ Infrastructure
	infrastrctureCustomer = infrastructure_customer.NewInfrastructureCustomer(http_client, "http://localhost:8080/api/customer")
	infrastrctureVoucher  = infrastructure_voucher.NewInfrastructureVoucher(http_client, "http://localhost:8080/api/voucher")
	repoTransaction       = repository_mysql.NewRepoTransaction(connectionMysql)
	repoTransactionsItem  = repository_mysql.NewRepoTransactionsItem(connectionMysql)
	repoItem              = repository_mysql.NewRepoItem(connectionMysql)

	repoCustomersTransactionCount = repository_mysql.NewRepoCustomersTransactionCount(connectionMysql)
	usecaseTransaction            = usecase.NewUsecaseTransaction(infrastrctureCustomer, infrastrctureVoucher, repoTransaction, repoTransactionsItem, repoItem, repoCustomersTransactionCount)
)

func Test_Transaction_Store_Positive(t *testing.T) {

	dto_transaction := &entity.DTOTransaction{
		Customer_id: 0,
		Note:        "",
		Items: []*entity.DTOTransactionsItem{
			&entity.DTOTransactionsItem{
				Item_id:       10,
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
	}

	transaction_id, usecase_transaction_store_err := usecaseTransaction.Store(ctx, dto_transaction)
	assert.NotEqual(t, 0, transaction_id)
	assert.Nil(t, usecase_transaction_store_err)
}

func Test_Transaction_Update_Positive(t *testing.T) {

	dto_transaction := &entity.DTOTransaction{
		Id:          1,
		Customer_id: 0,
		Note:        "",
		Items: []*entity.DTOTransactionsItem{
			&entity.DTOTransactionsItem{
				Item_id:       10,
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
			&entity.DTOTransactionsItem{
				Item_id:       12,
				Items_type_id: 4,
				Price:         250000.00,
				Qty:           1,
				Total_price:   250000.00,
				Note:          "Note",
			},
		},
		Status:                       112,
		Created_at:                   time.Now(),
		Is_generated_voucher_succeed: false,
	}
	usecase_transaction_update_err := usecaseTransaction.Update(ctx, dto_transaction)
	assert.Nil(t, usecase_transaction_update_err)
}

func Test_Transaction_Delete_Positive(t *testing.T) {

	usecase_transaction_delete_err := usecaseTransaction.Delete(ctx, 0, 4)
	assert.Nil(t, usecase_transaction_delete_err)
}

func Test_Transaction_GetById_Positive(t *testing.T) {
	usecase_transaction_get, usecase_transaction_get_err := usecaseTransaction.GetById(ctx, 0, 4)
	assert.NotNil(t, usecase_transaction_get)
	assert.Nil(t, usecase_transaction_get_err)
}

func Test_Transaction_GetByCustomerIdList_Positive(t *testing.T) {

	usecase_transactions_get, usecase_transactions_get_err := usecaseTransaction.GetByCustomerIdList(ctx, 0, 1)
	assert.NotNil(t, usecase_transactions_get)
	assert.Nil(t, usecase_transactions_get_err)
}

func Test_Transaction_GetList_Positive(t *testing.T) {

	usecase_transactions_get, usecase_transactions_get_err := usecaseTransaction.GetList(ctx, 2)
	assert.NotNil(t, usecase_transactions_get)
	assert.Nil(t, usecase_transactions_get_err)
}

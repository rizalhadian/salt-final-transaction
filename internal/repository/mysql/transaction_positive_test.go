package repository_mysql_test

import (
	"context"
	"fmt"
	"salt-final-transaction/domain/entity"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func Test_Transaction_Store_Positive(t *testing.T) {

	ctx := context.Background()

	dto_transaction := &entity.DTOTransaction{
		Customer_id:                  1,
		Total_amount:                 0,
		Total_discount_amount:        0,
		Final_total_amount:           0,
		Note:                         "",
		Status:                       112,
		Created_at:                   time.Now(),
		Is_generated_voucher_succeed: false,
	}

	entity_transaction, err_entity_transaction := entity.NewTransaction(dto_transaction)
	if err_entity_transaction != nil {
		panic(err_entity_transaction)
	}

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	repo_err := repoTransaction.Store(ctx, entity_transaction)
	assert.NotZero(t, entity_transaction.GetId())
	assert.Nil(t, repo_err)
}

func Test_Transaction_GetById_Positive(t *testing.T) {

	ctx := context.Background()

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	data, repo_err := repoTransaction.GetById(ctx, 3)
	fmt.Println(data.GetUpdatedAt())
	assert.NotNil(t, data)
	assert.Nil(t, repo_err)
}

func Test_Transaction_GetList_Positive(t *testing.T) {
	ctx := context.Background()

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	data, repo_err := repoTransaction.GetList(ctx, 10, 0)
	fmt.Println(data[0].GetStatus())
	assert.NotNil(t, data)
	assert.Nil(t, repo_err)
}

func Test_Transaction_Update_Positive(t *testing.T) {
	ctx := context.Background()
	dto_transaction := &entity.DTOTransaction{
		Id:                           1,
		Customer_id:                  1,
		Total_amount:                 0,
		Total_discount_amount:        0,
		Final_total_amount:           0,
		Note:                         "",
		Status:                       113,
		Created_at:                   time.Now(),
		Is_generated_voucher_succeed: false,
	}

	entity_transaction, err_entity_transaction := entity.NewTransaction(dto_transaction)
	if err_entity_transaction != nil {
		panic(err_entity_transaction)
	}

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	repo_err := repoTransaction.Update(ctx, entity_transaction)
	assert.Nil(t, repo_err)
}

func Test_Transaction_Soft_Delete_Positive(t *testing.T) {
	ctx := context.Background()

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	repo_err := repoTransaction.Delete(ctx, 1)
	assert.Nil(t, repo_err)
}

func Test_Transaction_Hard_Delete_Positive(t *testing.T) {
	ctx := context.Background()

	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoTransaction = repository_mysql.NewRepoTransaction(connectionMysql)
	)

	repo_err := repoTransaction.HardDelete(ctx, 1)
	assert.Nil(t, repo_err)
}

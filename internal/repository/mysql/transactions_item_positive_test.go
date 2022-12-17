package repository_mysql_test

import (
	"context"
	"salt-final-transaction/domain/entity"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Transactions_Item_Store_Positive(t *testing.T) {

	ctx := context.Background()

	var entity_transactions_items []*entity.TransactionsItem
	for i := 0; i < 10; i++ {
		dto_transactions_item := &entity.DTOTransactionsItem{
			Item_id:       4,
			Items_type_id: 2,
			Price:         1800000.00,
			Qty:           1,
			Total_price:   1800000.00,
			Note:          "Note index " + strconv.Itoa(i),
		}
		entity_transactions_item, err_entity_transactions_item := entity.NewTransactionsItem(dto_transactions_item)
		if err_entity_transactions_item != nil {
			panic(err_entity_transactions_item)
		}
		entity_transactions_items = append(entity_transactions_items, entity_transactions_item)
	}

	var (
		connectionMysql      = pkg_database_mysql.InitDBMysql()
		RepoTransactionsItem = repository_mysql.NewRepoTransactionsItem(connectionMysql)
	)

	repo_err := RepoTransactionsItem.Store(ctx, 1, entity_transactions_items)
	assert.Nil(t, repo_err)
}

func Test_Transactions_Item_Get_By_Transactions_Id_Positive(t *testing.T) {

	ctx := context.Background()

	var (
		connectionMysql      = pkg_database_mysql.InitDBMysql()
		RepoTransactionsItem = repository_mysql.NewRepoTransactionsItem(connectionMysql)
	)

	transactions_items, repo_err := RepoTransactionsItem.GetByTransactionId(ctx, 1)

	assert.NotNil(t, transactions_items)
	assert.Nil(t, repo_err)
}

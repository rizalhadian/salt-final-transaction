package repository_mysql_test

import (
	"context"
	"fmt"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Item_GetById_Positive(t *testing.T) {
	ctx := context.Background()
	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoItem        = repository_mysql.NewRepoItem(connectionMysql)
	)
	data, repo_err := repoItem.GetById(ctx, 10)
	fmt.Println(data)
	assert.NotNil(t, data)
	assert.Nil(t, repo_err)
}

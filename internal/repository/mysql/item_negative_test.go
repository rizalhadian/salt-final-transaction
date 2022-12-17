package repository_mysql_test

import (
	"context"
	"errors"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Item_GetById_Negative(t *testing.T) {
	ctx := context.Background()
	var (
		connectionMysql = pkg_database_mysql.InitDBMysql()
		repoItem        = repository_mysql.NewRepoItem(connectionMysql)
	)
	data, repo_err := repoItem.GetById(ctx, 999999)

	assert.Nil(t, data)
	assert.NotNil(t, repo_err)
	assert.Equal(t, errors.New("404"), repo_err)
}

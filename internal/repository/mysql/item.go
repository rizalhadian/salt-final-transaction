package repository_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	mapper_mysql "salt-final-transaction/internal/repository/mysql/mapper"
	model "salt-final-transaction/internal/repository/mysql/models"
	"strconv"
	"strings"

	"github.com/rocketlaunchr/dbq/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RepoItem struct {
	db *sql.DB
}

func NewRepoItem(db *sql.DB) interface_repo.InterfaceRepoItem {
	return &RepoItem{
		db: db,
	}
}

func (ri *RepoItem) GetById(ctx context.Context, id int64) (*entity.Item, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	id_string := strconv.Itoa(int(id))

	items_fields_selected_arr := model.ModelItem{}.GetFieldsNeededToGetProcess()
	items_fields_selected := strings.Join(items_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = %s AND deleted_at IS NULL LIMIT 1 ", items_fields_selected, model.ModelItem{}.GetTableName(), id_string)
	fmt.Println(id_string)
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.ModelItem{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, ri.db, query, opts)

	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		item := mapper_mysql.ItemModelToEntity(result.(*model.ModelItem))
		return item, nil
	}
	return nil, errors.New("404")
}

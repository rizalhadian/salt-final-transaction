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
	dbqx "github.com/rocketlaunchr/dbq/v2/x"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RepoCustomersTransactionCount struct {
	db *sql.DB
}

func NewRepoCustomersTransactionCount(db *sql.DB) interface_repo.InterfaceRepoCustomersTransactionCount {
	return &RepoCustomersTransactionCount{
		db: db,
	}
}

func (rctc *RepoCustomersTransactionCount) Store(ctx context.Context, t *entity.CustomersTransactionCount) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	model_cust_trx_count := mapper_mysql.CustomersTransactionCountEntityToModel(t)

	stmt := dbq.INSERT(model_cust_trx_count.GetTableName(), model_cust_trx_count.GetFieldsNeededToStoreProcess(), 1, dbq.MySQL)
	res := dbq.MustE(ctx, rctc.db, stmt, nil, model_cust_trx_count.ToArrInterface("save"))

	last_id, err_last_id := res.LastInsertId()
	if err_last_id != nil {
		log.Error().Msg(err_last_id.Error())
		return err_last_id
	}
	t.SetId(last_id)

	return nil
}

func (rctc *RepoCustomersTransactionCount) Update(ctx context.Context, t *entity.CustomersTransactionCount) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	model_cust_trx_count := mapper_mysql.CustomersTransactionCountEntityToModel(t)

	opts := dbqx.BulkUpdateOptions{
		Table:      model_cust_trx_count.GetTableName(),
		Columns:    model_cust_trx_count.GetFieldsNeededToUpdateProcess(),
		PrimaryKey: "customer_id",
	}

	data := map[interface{}]interface{}{
		t.GetCustomerId(): model_cust_trx_count.ToArrInterface("update"),
	}

	_, err_update := dbqx.BulkUpdate(ctx, rctc.db, data, opts)
	if err_update != nil {
		log.Error().Msg(err_update.Error())
		return err_update
	}

	return nil
}

func (rctc *RepoCustomersTransactionCount) GetByCustomerId(ctx context.Context, customer_id int64) (*entity.CustomersTransactionCount, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	customer_id_string := strconv.Itoa(int(customer_id))

	items_fields_selected_arr := model.ModelCustomersTransactionCount{}.GetFieldsNeededToGetProcess()
	items_fields_selected := strings.Join(items_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE customer_id = %s LIMIT 1 ", items_fields_selected, model.ModelCustomersTransactionCount{}.GetTableName(), customer_id_string)

	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.ModelCustomersTransactionCount{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, rctc.db, query, opts)

	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		item := mapper_mysql.CustomersTransactionCountModelToEntity(result.(*model.ModelCustomersTransactionCount))
		return item, nil
	}
	return nil, errors.New("404")
}

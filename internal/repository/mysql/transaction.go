package repository_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	mapper_mysql "salt-final-transaction/internal/repository/mysql/mapper"
	model "salt-final-transaction/internal/repository/mysql/models"
	"strconv"
	"strings"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
	dbqx "github.com/rocketlaunchr/dbq/v2/x"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RepoTransaction struct {
	db *sql.DB
}

func NewRepoTransaction(db *sql.DB) interface_repo.InterfaceRepoTransaction {
	return &RepoTransaction{
		db: db,
	}
}

func (rt *RepoTransaction) Store(ctx context.Context, t *entity.Transaction) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	current_time := time.Now()
	t.SetCreatedAt(current_time)

	model_transaction := mapper_mysql.TransactionEntityToModel(t)

	stmt := dbq.INSERT(model_transaction.GetTableName(), model_transaction.GetFieldsNeededToStoreProcess(), 1, dbq.MySQL)
	res := dbq.MustE(ctx, rt.db, stmt, nil, model_transaction.ToArrInterface("save"))

	last_id, err_last_id := res.LastInsertId()
	if err_last_id != nil {
		log.Error().Msg(err_last_id.Error())
		return err_last_id
	}
	t.SetId(last_id)

	return nil
}
func (rt *RepoTransaction) Update(ctx context.Context, t *entity.Transaction) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	current_time := time.Now()
	updated_at := sql.NullTime{
		Time:  current_time,
		Valid: true,
	}

	t.SetUpdatedAt(updated_at)

	model_transaction := mapper_mysql.TransactionEntityToModel(t)

	opts := dbqx.BulkUpdateOptions{
		Table:      model_transaction.GetTableName(),
		Columns:    model_transaction.GetFieldsNeededToUpdateProcess(),
		PrimaryKey: "id",
	}

	data := map[interface{}]interface{}{
		t.GetId(): model_transaction.ToArrInterface("update"),
	}

	_, err_update := dbqx.BulkUpdate(ctx, rt.db, data, opts)
	if err_update != nil {
		log.Error().Msg(err_update.Error())
		return err_update
	}

	return nil
}

// Softdelete
func (rt *RepoTransaction) Delete(ctx context.Context, id int64) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	current_time := time.Now()
	// deleted_at := sql.NullTime{
	// 	Time:  current_time,
	// 	Valid: true,
	// }

	// t := entity.Transaction{}
	// t.SetDeletedAt(deleted_at)
	// model_transaction := mapper_mysql.TransactionEntityToModel(&t)

	model_transaction := model.ModelTransaction{Deleted_at: current_time.String()}

	opts := dbqx.BulkUpdateOptions{
		Table:      model_transaction.GetTableName(),
		Columns:    model_transaction.GetFieldsNeededToSoftDeleteProcess(),
		PrimaryKey: "id",
	}

	data := map[interface{}]interface{}{
		id: model_transaction.ToArrInterface("delete"),
	}

	_, err_update := dbqx.BulkUpdate(ctx, rt.db, data, opts)
	if err_update != nil {
		log.Error().Msg(err_update.Error())
		return err_update
	}

	return nil
}

func (rt *RepoTransaction) HardDelete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", "transaction", strconv.Itoa(int(id)))

	_, err_delete := dbq.E(ctx, rt.db, query, nil)
	if err_delete != nil {
		log.Error().Msg(err_delete.Error())
		return err_delete
	}
	return nil
}

func (rt *RepoTransaction) GetById(ctx context.Context, id int64) (*entity.Transaction, error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	id_string := strconv.Itoa(int(id))

	articles_fields_selected_arr := model.ModelTransaction{}.GetFieldsNeededToGetProcess()
	articles_fields_selected := strings.Join(articles_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = %s AND deleted_at IS NULL LIMIT 1 ", articles_fields_selected, model.ModelTransaction{}.GetTableName(), id_string)
	fmt.Println(query)
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.ModelTransaction{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, rt.db, query, opts)
	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		transaction := mapper_mysql.TransactionModelToEntity(result.(*model.ModelTransaction))
		return transaction, nil
	}
	return nil, nil
}

func (rt *RepoTransaction) GetByCustomerIdList(ctx context.Context, customer_id int64, limit int32, offset int32) (res []*entity.Transaction, err error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	customer_id_string := strconv.Itoa(int(customer_id))
	limit_string := strconv.Itoa(int(limit))
	offset_string := strconv.Itoa(int(offset))

	transcations_fields_selected_arr := model.ModelTransaction{}.GetFieldsNeededToGetProcess()
	transcations_fields_selected := strings.Join(transcations_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE customer_id = %s AND deleted_at IS NULL LIMIT %s OFFSET %s", transcations_fields_selected, model.ModelTransaction{}.GetTableName(), customer_id_string, limit_string, offset_string)

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.ModelTransaction{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, rt.db, query, opts)
	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		transactions := mapper_mysql.TransactionModelListToEntityList(result.([]*model.ModelTransaction))
		return transactions, nil
	}
	return nil, nil
}

func (rt *RepoTransaction) GetList(ctx context.Context, limit int32, offset int32) (res []*entity.Transaction, err error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	limit_string := strconv.Itoa(int(limit))
	offset_string := strconv.Itoa(int(offset))

	transcations_fields_selected_arr := model.ModelTransaction{}.GetFieldsNeededToGetProcess()
	transcations_fields_selected := strings.Join(transcations_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at IS NULL LIMIT %s OFFSET %s", transcations_fields_selected, model.ModelTransaction{}.GetTableName(), limit_string, offset_string)

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.ModelTransaction{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, rt.db, query, opts)
	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		transactions := mapper_mysql.TransactionModelListToEntityList(result.([]*model.ModelTransaction))
		return transactions, nil
	}
	return nil, nil
}

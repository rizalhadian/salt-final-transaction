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
	"time"

	"github.com/rocketlaunchr/dbq/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RepoTransactionsItem struct {
	db *sql.DB
}

func NewRepoTransactionsItem(db *sql.DB) interface_repo.InterfaceRepoTransactionsItem {
	return &RepoTransactionsItem{
		db: db,
	}
}

func (rti *RepoTransactionsItem) Store(ctx context.Context, transaction_id int64, trx_items []*entity.TransactionsItem) (errs error) {
	current_time := time.Now()

	err := dbq.Tx(ctx, rti.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		for _, trx_item := range trx_items {
			trx_item.SetCreatedAt(current_time)
			trx_item.SetTransactionId(transaction_id)

			model_transactions_item := mapper_mysql.TransactionsItemEntityToModel(trx_item)
			stmt := dbq.INSERTStmt(model_transactions_item.GetTableName(), model_transactions_item.GetFieldsNeededToStoreProcess(), 1)

			res, err := E(ctx, stmt, nil, model_transactions_item.ToArrInterface("save"))
			if err != nil {
				log.Error().Msg(err.Error())
				return
			}
			last_id, err_last_id := res.LastInsertId()
			if err_last_id != nil {
				log.Error().Msg(err_last_id.Error())
				return
			}
			trx_item.SetId(last_id)

		}
		txCommit() // Commit
	})
	if err != nil {
		return err
	}
	return nil
}

//	func (rti *RepoTransactionsItem) Update(ctx context.Context, transaction_id int64, trx_items []*entity.TransactionsItem) error {
//		return nil
//	}
//
//	func (rti *RepoTransactionsItem) Delete(ctx context.Context, ids []int64) error {
//		return nil
//	}
//
//	func (rti *RepoTransactionsItem) HardDelete(ctx context.Context, ids []int64) error {
//		return nil
//	}
//
//	func (rti *RepoTransactionsItem) GetById(ctx context.Context, id int64) (*entity.TransactionsItem, error) {
//		return nil, nil
//	}

func (rti *RepoTransactionsItem) GetByTransactionId(ctx context.Context, transaction_id int64) (res []*entity.TransactionsItem, err error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	transactions_id_string := strconv.Itoa(int(transaction_id))

	transactions_items_fields_selected_arr := model.ModelTransactionsItem{}.GetFieldsNeededToGetProcess()
	transactions_items_fields_selected := strings.Join(transactions_items_fields_selected_arr, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE transaction_id = %s AND deleted_at IS NULL", transactions_items_fields_selected, model.ModelTransactionsItem{}.GetTableName(), transactions_id_string)

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.ModelTransactionsItem{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err_get := dbq.Q(ctx, rti.db, query, opts)
	if err_get != nil {
		log.Error().Msg(err_get.Error())
		return nil, err_get
	}

	if result != nil {
		transactions_items := mapper_mysql.TransactionsItemModelListToEntityList(result.([]*model.ModelTransactionsItem))
		return transactions_items, nil
	}
	return nil, errors.New("404")
}

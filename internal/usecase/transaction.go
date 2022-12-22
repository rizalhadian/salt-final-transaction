package usecase

import (
	"context"
	"errors"
	"fmt"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	"salt-final-transaction/domain/interface_usecase"
	interface_infrastructure_customer "salt-final-transaction/internal/infrastructure/customer/interface"
	interface_infrastructure_voucher "salt-final-transaction/internal/infrastructure/voucher/interface"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UsecaseTransaction struct {
	infraCustomer                 interface_infrastructure_customer.InterfaceInfrastructureCustomer
	infraVoucher                  interface_infrastructure_voucher.InterfaceInfrastructureVoucher
	repoTransaction               interface_repo.InterfaceRepoTransaction
	repoTransactionsItem          interface_repo.InterfaceRepoTransactionsItem
	repoItem                      interface_repo.InterfaceRepoItem
	repoCustomersTransactionCount interface_repo.InterfaceRepoCustomersTransactionCount
}

type TotalAmountPerCategory struct {
	items_type_id int64
	total_amount  float64
}

type Transactions_Items_Err struct {
	item_index int
	item_id    int64
	field      string
	err        string
}

func NewUsecaseTransaction(interfaceInfraCustomer interface_infrastructure_customer.InterfaceInfrastructureCustomer, interfaceInfraVoucher interface_infrastructure_voucher.InterfaceInfrastructureVoucher, interfaceRepoTransaction interface_repo.InterfaceRepoTransaction, interfaceRepoTransactionsItem interface_repo.InterfaceRepoTransactionsItem, interfaceRepoItem interface_repo.InterfaceRepoItem, interfaceCustomersTransactionCount interface_repo.InterfaceRepoCustomersTransactionCount) interface_usecase.InterfaceUsecaseTransaction {
	return &UsecaseTransaction{
		infraCustomer:                 interfaceInfraCustomer,
		infraVoucher:                  interfaceInfraVoucher,
		repoTransaction:               interfaceRepoTransaction,
		repoTransactionsItem:          interfaceRepoTransactionsItem,
		repoItem:                      interfaceRepoItem,
		repoCustomersTransactionCount: interfaceCustomersTransactionCount,
	}
}

func (ut *UsecaseTransaction) Store(ctx context.Context, dto_transaction *entity.DTOTransaction) (transaction *entity.Transaction, errs error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// >>> Infrastructure Check Customer By Id || Lebih dari 0 karena yang perlu di cek adalah id lebih dari 0.. Customer_id 0 dianggap tanpa member
	if dto_transaction.Customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, dto_transaction.Customer_id)
		if infra_cust_err.Error() == "404" {
			return nil, errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return nil, errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	if len(dto_transaction.Items) == 0 {
		return nil, errors.New("Transaction's Items is required")
	}

	entity_transaction, err_entity_transaction := entity.NewTransaction(dto_transaction)
	if err_entity_transaction != nil {
		return nil, err_entity_transaction
	}

	// >>> Check Transaction's Item

	var entity_transactions_items []*entity.TransactionsItem
	var total_items_amount float64
	for _, dto_transactions_item := range dto_transaction.Items {
		// >>> Create Entity And Validate
		entity_transactions_item, entity_transactions_item_err := entity.NewTransactionsItem(dto_transactions_item)
		if entity_transactions_item_err != nil {
			return nil, entity_transactions_item_err
		}
		// <<< Create Entity And Validate

		// >>> Check Repo Item
		item, repo_item_err := ut.repoItem.GetById(ctx, dto_transactions_item.Item_id)

		if repo_item_err != nil {
			if repo_item_err.Error() == "404" {
				log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
				return nil, errors.New("Item Not Found")
			}

			log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
			return nil, errors.New("500")
		}

		if item.GetIsService() == false && item.GetStock() < dto_transactions_item.Qty {
			return nil, errors.New("Insufficient Stock")
		}
		if item.GetPrice() != dto_transactions_item.Price {
			return nil, errors.New("Price Changed")
		}
		if item.GetItemsTypeId() != dto_transactions_item.Items_type_id {
			return nil, errors.New("Items Type Id Is Not Match")
		}
		// <<< Check Repo Item
		total_items_amount += entity_transactions_item.GetTotalPrice()
		entity_transactions_items = append(entity_transactions_items, entity_transactions_item)

	}
	// <<< Check Transaction's Item

	dto_transaction.Is_generated_voucher_succeed = false

	// >>> Count Final Transaction's Amount
	dto_transaction.Final_total_amount = dto_transaction.Total_amount - dto_transaction.Total_discount_amount
	// <<< Count Final Transaction's Amount

	// >>> Repository Store
	repo_transaction_err := ut.repoTransaction.Store(ctx, entity_transaction)
	if repo_transaction_err != nil {
		log.Error().Msg(repo_transaction_err.Error())
		return nil, repo_transaction_err
	}

	repo_transactions_items_err := ut.repoTransactionsItem.Store(ctx, entity_transaction.GetId(), entity_transactions_items)
	if repo_transactions_items_err != nil {
		log.Error().Msg(repo_transactions_items_err.Error())
		return nil, repo_transactions_items_err
	}

	// >>> Infrastructure Check & Count Voucher Discount
	if len(dto_transaction.Vouchers_redeemed) > 0 {
		entity_transaction.SetTotalAmount(total_items_amount)
		entity_transaction.SetItems(entity_transactions_items)

		fmt.Println("Ada Voucher")
		var vouchers_codes []string
		for _, entity_transactions_voucher := range dto_transaction.Vouchers_redeemed {
			vouchers_code := entity_transactions_voucher.Code
			vouchers_codes = append(vouchers_codes, vouchers_code)
		}
		redeem_vouchers, redeem_voucher_err := ut.infraVoucher.Redeem(ctx, entity_transaction, vouchers_codes)
		if redeem_voucher_err != nil {
			if redeem_voucher_err.Error() == "404" {
				return nil, errors.New("Voucher Not Found")
			} else {
				return nil, redeem_voucher_err
			}
		} else {
			fmt.Println("Redeem Voucher Succeed")

			total_dicount_amount_all_vouchers := 0.00
			for _, redeem_voucher := range redeem_vouchers {
				total_dicount_amount_all_vouchers += redeem_voucher.Total_discount_amount
			}
			entity_transaction.SetStatus(111)
			entity_transaction.SetTotalDiscountAmount(total_dicount_amount_all_vouchers)
			repo_transaction_update_err := ut.repoTransaction.Update(ctx, entity_transaction)
			if repo_transaction_update_err != nil {
				return nil, repo_transaction_update_err
			}
		}
	}

	if len(dto_transaction.Vouchers_redeemed) == 0 {
		// Kalau tidak ada voucher langsung ubah status jadi 111 (submitted) dan set total discount amount 0
		entity_transaction.SetStatus(111)
		entity_transaction.SetTotalAmount(total_items_amount)
		entity_transaction.SetItems(entity_transactions_items)
		entity_transaction.SetTotalDiscountAmount(0)
		repo_transaction_update_err := ut.repoTransaction.Update(ctx, entity_transaction)
		if repo_transaction_update_err != nil {
			return nil, repo_transaction_update_err
		}
	}
	// >>> Infrastructure Check & Count Voucher Discount

	// <<< Repository Store

	// >>> Customers Transaction Count
	cust_trx_count_existed, cust_trx_count_existed_err := ut.repoCustomersTransactionCount.GetByCustomerId(ctx, entity_transaction.GetCustomerId())
	if cust_trx_count_existed_err != nil {
		if cust_trx_count_existed_err.Error() == "404" {
			entity_cust_trx_count := entity.CustomersTransactionCount{}
			entity_cust_trx_count.SetFirstTransactionDatetime(entity_transaction.GetCreatedAt())
			entity_cust_trx_count.SetLastTransactionDatetime(entity_transaction.GetCreatedAt())
			entity_cust_trx_count.SetTotalTransactionSpend(entity_transaction.GetFinalTotalAmount())
			entity_cust_trx_count.SetTransactionCount(1)
			entity_cust_trx_count_store_err := ut.repoCustomersTransactionCount.Store(ctx, &entity_cust_trx_count)
			if entity_cust_trx_count_store_err != nil {
				log.Error().Msg(entity_cust_trx_count_store_err.Error())
				return nil, errors.New("500")
			}
		}
	} else {
		last_total_trx_spend := cust_trx_count_existed.GetTotalTransactionSpend()
		last_trx_count := cust_trx_count_existed.GetTransactionCount()

		cust_trx_count_existed.SetLastTransactionDatetime(entity_transaction.GetCreatedAt())
		cust_trx_count_existed.SetTotalTransactionSpend(last_total_trx_spend + entity_transaction.GetFinalTotalAmount())
		cust_trx_count_existed.SetTransactionCount(last_trx_count + 1)

		entity_cust_trx_count_update_err := ut.repoCustomersTransactionCount.Update(ctx, cust_trx_count_existed)
		if entity_cust_trx_count_update_err != nil {
			log.Error().Msg(entity_cust_trx_count_update_err.Error())
			return nil, errors.New("500")
		}

	}
	// <<< Customers Transaction Count

	generate_voucher_err := ut.infraVoucher.GenerateVoucher(ctx, entity_transaction.GetCustomerId())
	if generate_voucher_err != nil {
		log.Error().Msg("!!! Generate Voucher Failed !!!")
		log.Error().Msg(generate_voucher_err.Error())
	} else {
		entity_transaction.SetIsGeneratedVoucherSucceed(true)
		repo_transaction_update_generate_voucher_err := ut.repoTransaction.Update(ctx, entity_transaction)
		if repo_transaction_update_generate_voucher_err != nil {
			return nil, repo_transaction_update_generate_voucher_err
		}
	}
	return entity_transaction, nil
}

func (ut *UsecaseTransaction) Update(ctx context.Context, dto_transaction *entity.DTOTransaction) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// >>> Infrastructure Check Customer By Id || Lebih dari 0 karena yang perlu di cek adalah id lebih dari 0.. Customer_id 0 dianggap tanpa member
	if dto_transaction.Customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, dto_transaction.Customer_id)
		if infra_cust_err.Error() == "404" {
			return errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	if len(dto_transaction.Items) == 0 {
		return errors.New("Transaction's Items is required")
	}

	entity_transaction, err_entity_transaction := entity.NewTransaction(dto_transaction)
	if err_entity_transaction != nil {
		return err_entity_transaction
	}

	// >>> Check Transaction's Item

	var entity_transactions_items []*entity.TransactionsItem
	var total_items_amount float64
	for _, dto_transactions_item := range dto_transaction.Items {
		// >>> Create Entity And Validate
		entity_transactions_item, entity_transactions_item_err := entity.NewTransactionsItem(dto_transactions_item)
		if entity_transactions_item_err != nil {
			return entity_transactions_item_err
		}
		// <<< Create Entity And Validate

		// >>> Check Repo Item
		item, repo_item_err := ut.repoItem.GetById(ctx, dto_transactions_item.Item_id)

		if repo_item_err != nil {
			if repo_item_err.Error() == "404" {
				log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
				return errors.New("Item Not Found")
			}

			log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
			return errors.New("500")
		}

		if item.GetIsService() == false && item.GetStock() < dto_transactions_item.Qty {
			return errors.New("Insufficient Stock")
		}
		if item.GetPrice() != dto_transactions_item.Price {
			return errors.New("Price Changed")
		}
		// <<< Check Repo Item
		total_items_amount += entity_transactions_item.GetTotalPrice()
		entity_transactions_items = append(entity_transactions_items, entity_transactions_item)
	}
	// <<< Check Transaction's Item

	// Code is here
	// <<< Infrastructure Check & Count Voucher

	dto_transaction.Is_generated_voucher_succeed = false

	// >>> Count Final Transaction's Amount
	dto_transaction.Final_total_amount = dto_transaction.Total_amount - dto_transaction.Total_discount_amount
	// <<< Count Final Transaction's Amount

	// >>> Repository Store
	old_transaction_id := entity_transaction.GetId()

	entity_rollback_transaction, entity_rollback_transaction_err := ut.repoTransaction.GetById(ctx, old_transaction_id)
	if entity_rollback_transaction_err != nil {
		return entity_rollback_transaction_err
	}

	if entity_rollback_transaction.GetStatus() != 114 && entity_rollback_transaction.GetStatus() != 111 {
		return errors.New("Update Only On Submitted Transaction or Revise Transaction")
	}

	// Add Updated Transaction
	entity_transaction.SetStatus(114) //Status Revise
	repo_transaction_err := ut.repoTransaction.Store(ctx, entity_transaction)
	if repo_transaction_err != nil {
		log.Error().Msg(repo_transaction_err.Error())
		return repo_transaction_err
	}
	repo_transactions_items_err := ut.repoTransactionsItem.Store(ctx, entity_transaction.GetId(), entity_transactions_items)
	if repo_transactions_items_err != nil {
		log.Error().Msg(repo_transactions_items_err.Error())
		return repo_transactions_items_err
	}
	entity_transaction.SetTotalAmount(total_items_amount)
	entity_transaction.SetTotalDiscountAmount(0)
	repo_transaction_update_err := ut.repoTransaction.Update(ctx, entity_transaction)
	if repo_transaction_update_err != nil {
		return repo_transaction_update_err
	}
	updated_transaction_id := entity_transaction.GetId()

	// Add Rollback Transaction
	entity_rollback_transaction.SetUpdateTransactionId(old_transaction_id)
	entity_rollback_transaction.SetTotalAmount(float64(entity_rollback_transaction.GetTotalAmount() * -1.0))
	entity_rollback_transaction.SetTotalDiscountAmount(float64(entity_rollback_transaction.GetTotalDiscountAmount() * -1.0))
	entity_rollback_transaction.SetStatus(115) //Status Rollback
	repo_transaction__rollback_err := ut.repoTransaction.Store(ctx, entity_rollback_transaction)
	if repo_transaction__rollback_err != nil {
		log.Error().Msg(repo_transaction__rollback_err.Error())
		return repo_transaction__rollback_err
	}
	rollbacked_transaction_id := entity_rollback_transaction.GetId()

	// Update Old Transaction
	entity_old_transaction, old_transaction_err := ut.repoTransaction.GetById(ctx, old_transaction_id)
	if old_transaction_err != nil {
		return old_transaction_err
	}
	entity_old_transaction.SetId(old_transaction_id)
	entity_old_transaction.SetStatus(113)
	entity_old_transaction.SetRollbackTransactionId(rollbacked_transaction_id)
	entity_old_transaction.SetUpdateTransactionId(updated_transaction_id)
	repo_old_transaction_update_err := ut.repoTransaction.Update(ctx, entity_old_transaction)
	if repo_old_transaction_update_err != nil {
		return repo_old_transaction_update_err
	}
	// <<<< Repository Store

	// >>> Customers Transaction Count
	cust_trx_count_existed, cust_trx_count_existed_err := ut.repoCustomersTransactionCount.GetByCustomerId(ctx, entity_rollback_transaction.GetCustomerId())
	if cust_trx_count_existed_err != nil {
		log.Error().Msg("Update Transaction But Customers Transaction Count Is Null | Need To count all transaction")
	} else {
		last_total_trx_spend := cust_trx_count_existed.GetTotalTransactionSpend()
		cust_trx_count_existed.SetLastTransactionDatetime(entity_rollback_transaction.GetCreatedAt())
		cust_trx_count_existed.SetTotalTransactionSpend(last_total_trx_spend + (entity_old_transaction.GetFinalTotalAmount() * -1.0) + entity_transaction.GetFinalTotalAmount())

		entity_cust_trx_count_update_err := ut.repoCustomersTransactionCount.Update(ctx, cust_trx_count_existed)
		if entity_cust_trx_count_update_err != nil {
			log.Error().Msg(entity_cust_trx_count_update_err.Error())
			return errors.New("500")
		}

	}
	// <<< Customers Transaction Count

	return nil
}
func (ut *UsecaseTransaction) Delete(ctx context.Context, customer_id int64, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// >>> Infrastructure Check Customer By Id || Lebih dari 0 karena yang perlu di cek adalah id lebih dari 0.. Customer_id 0 dianggap tanpa member
	if customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, customer_id)
		if infra_cust_err.Error() == "404" {
			return errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	// >>> Repository Store
	old_transaction_id := id

	// Add Rollback Transaction
	entity_rollback_transaction, entity_rollback_transaction_err := ut.repoTransaction.GetById(ctx, old_transaction_id)
	if entity_rollback_transaction_err != nil {
		return entity_rollback_transaction_err
	}

	entity_rollback_transaction.SetUpdateTransactionId(0)
	entity_rollback_transaction.SetTotalAmount(float64(entity_rollback_transaction.GetTotalAmount() * -1.0))
	entity_rollback_transaction.SetTotalDiscountAmount(float64(entity_rollback_transaction.GetTotalDiscountAmount() * -1.0))
	entity_rollback_transaction.SetStatus(115) //Status Rollback
	repo_transaction__rollback_err := ut.repoTransaction.Store(ctx, entity_rollback_transaction)
	if repo_transaction__rollback_err != nil {
		log.Error().Msg(repo_transaction__rollback_err.Error())
		return repo_transaction__rollback_err
	}
	rollbacked_transaction_id := entity_rollback_transaction.GetId()

	// Update Old Transaction
	entity_old_transaction, old_transaction_err := ut.repoTransaction.GetById(ctx, old_transaction_id)
	if old_transaction_err != nil {
		return old_transaction_err
	}
	entity_old_transaction.SetId(old_transaction_id)
	entity_old_transaction.SetStatus(113)
	entity_old_transaction.SetRollbackTransactionId(rollbacked_transaction_id)
	entity_old_transaction.SetUpdateTransactionId(0)
	repo_old_transaction_update_err := ut.repoTransaction.Update(ctx, entity_old_transaction)
	if repo_old_transaction_update_err != nil {
		return repo_old_transaction_update_err
	}
	// <<<< Repository Store

	// >>> Customers Transaction Count
	cust_trx_count_existed, cust_trx_count_existed_err := ut.repoCustomersTransactionCount.GetByCustomerId(ctx, entity_rollback_transaction.GetCustomerId())
	if cust_trx_count_existed_err != nil {
		log.Error().Msg("Delete Transaction But Customers Transaction Count Is Null | Need To count all transaction")
	} else {
		last_total_trx_spend := cust_trx_count_existed.GetTotalTransactionSpend()
		last_trx_count := cust_trx_count_existed.GetTransactionCount()

		cust_trx_count_existed.SetLastTransactionDatetime(entity_rollback_transaction.GetCreatedAt())
		cust_trx_count_existed.SetTotalTransactionSpend(last_total_trx_spend + entity_rollback_transaction.GetFinalTotalAmount())
		cust_trx_count_existed.SetTransactionCount(last_trx_count - 1)

		entity_cust_trx_count_update_err := ut.repoCustomersTransactionCount.Update(ctx, cust_trx_count_existed)
		if entity_cust_trx_count_update_err != nil {
			log.Error().Msg(entity_cust_trx_count_update_err.Error())
			return errors.New("500")
		}

	}
	// <<< Customers Transaction Count

	return nil
}
func (ut *UsecaseTransaction) GetById(ctx context.Context, customer_id int64, id int64) (*entity.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	// var errs error
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// >>> Infrastructure Check Customer By Id
	if customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, customer_id)
		if infra_cust_err.Error() == "404" {
			return nil, errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return nil, errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	entity_transaction, entity_transaction_err := ut.repoTransaction.GetById(ctx, id)
	if entity_transaction_err != nil {
		log.Error().Msg(entity_transaction_err.Error())
		return nil, errors.New("500")
	}

	if entity_transaction != nil {
		entity_transactions_item, entity_transactions_item_err := ut.repoTransactionsItem.GetByTransactionId(ctx, id)
		if entity_transactions_item_err != nil {
			log.Error().Msg(entity_transactions_item_err.Error())
			return nil, errors.New("500")
		}
		entity_transaction.SetItems(entity_transactions_item)
	} else {
		return nil, errors.New("404")
	}

	return entity_transaction, nil
}
func (ut *UsecaseTransaction) GetByCustomerIdList(ctx context.Context, customer_id int64, page int32) (res []*entity.Transaction, err error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	// var errs error
	item_per_page := int32(10)
	limit := item_per_page
	offset := (page * item_per_page) - item_per_page

	// >>> Infrastructure Check Customer By Id
	if customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, customer_id)
		if infra_cust_err.Error() == "404" {
			return nil, errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return nil, errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	entity_transactions, entity_transactions_err := ut.repoTransaction.GetByCustomerIdList(ctx, customer_id, int32(limit), offset)
	if entity_transactions_err != nil {
		log.Error().Msg(entity_transactions_err.Error())
		return nil, errors.New("500")
	} else {
		for _, entity_transaction := range entity_transactions {
			entity_transactions_item, entity_transactions_item_err := ut.repoTransactionsItem.GetByTransactionId(ctx, entity_transaction.GetId())
			if entity_transactions_item_err != nil {
				log.Error().Msg(entity_transactions_item_err.Error())
				return nil, errors.New("500")
			}
			entity_transaction.SetItems(entity_transactions_item)
		}
	}

	return entity_transactions, nil
}

func (ut *UsecaseTransaction) GetList(ctx context.Context, page int32) (res []*entity.Transaction, err error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	// var errs error
	item_per_page := int32(10)
	limit := item_per_page
	offset := (page * item_per_page) - item_per_page

	entity_transactions, entity_transactions_err := ut.repoTransaction.GetList(ctx, int32(limit), offset)
	if entity_transactions_err != nil {
		log.Error().Msg(entity_transactions_err.Error())
		return nil, errors.New("500")
	}
	for _, entity_transaction := range entity_transactions {
		entity_transactions_item, entity_transactions_item_err := ut.repoTransactionsItem.GetByTransactionId(ctx, entity_transaction.GetId())
		if entity_transactions_item_err != nil {
			log.Error().Msg(entity_transactions_item_err.Error())
			return nil, errors.New("500")
		}
		entity_transaction.SetItems(entity_transactions_item)
	}

	return entity_transactions, nil

}

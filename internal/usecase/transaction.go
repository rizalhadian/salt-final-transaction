package usecase

import (
	"context"
	"errors"
	"salt-final-transaction/domain/entity"
	"salt-final-transaction/domain/interface_repo"
	"salt-final-transaction/domain/interface_usecase"
	interface_infrastructure_customer "salt-final-transaction/internal/infrastructure/customer/interface"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UsecaseTransaction struct {
	infraCustomer        interface_infrastructure_customer.InterfaceInfrastructureCustomer
	repoTransaction      interface_repo.InterfaceRepoTransaction
	repoTransactionsItem interface_repo.InterfaceRepoTransactionsItem
	repoItem             interface_repo.InterfaceRepoItem
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

func NewUsecaseTransaction(interfaceInfraCustomer interface_infrastructure_customer.InterfaceInfrastructureCustomer, interfaceRepoTransaction interface_repo.InterfaceRepoTransaction, interfaceRepoTransactionsItem interface_repo.InterfaceRepoTransactionsItem, interfaceRepoItem interface_repo.InterfaceRepoItem) interface_usecase.InterfaceUsecaseTransaction {
	return &UsecaseTransaction{
		infraCustomer:        interfaceInfraCustomer,
		repoTransaction:      interfaceRepoTransaction,
		repoTransactionsItem: interfaceRepoTransactionsItem,
		repoItem:             interfaceRepoItem,
	}
}

func (ut *UsecaseTransaction) Store(ctx context.Context, dto_transaction *entity.DTOTransaction) (id int64, errs error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// >>> Infrastructure Check Customer By Id || Lebih dari 0 karena yang perlu di cek adalah id lebih dari 0.. Customer_id 0 dianggap tanpa member
	if dto_transaction.Customer_id > 0 {
		_, infra_cust_err := ut.infraCustomer.GetById(ctx, dto_transaction.Customer_id)
		if infra_cust_err.Error() == "404" {
			return 0, errors.New("404")
		}
		if infra_cust_err != nil {
			log.Error().Msg("Transaction=>Infrastructure=>Customer : " + infra_cust_err.Error())
			return 0, errors.New("500")
		}
	}
	// <<< Infrastructure Check Customer By Id

	if len(dto_transaction.Items) == 0 {
		return 0, errors.New("Transaction's Items is required")
	}

	entity_transaction, err_entity_transaction := entity.NewTransaction(dto_transaction)
	if err_entity_transaction != nil {
		return 0, err_entity_transaction
	}

	// >>> Check Transaction's Item

	var entity_transactions_items []*entity.TransactionsItem
	var total_items_amount float64
	for _, dto_transactions_item := range dto_transaction.Items {
		// >>> Create Entity And Validate
		entity_transactions_item, entity_transactions_item_err := entity.NewTransactionsItem(dto_transactions_item)
		if entity_transactions_item_err != nil {
			return 0, entity_transactions_item_err
		}
		// <<< Create Entity And Validate

		// >>> Check Repo Item
		item, repo_item_err := ut.repoItem.GetById(ctx, dto_transactions_item.Item_id)

		if repo_item_err != nil {
			if repo_item_err.Error() == "404" {
				log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
				return 0, errors.New("400")
			}

			log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
			return 0, errors.New("500")
		}

		if item.GetStock() < dto_transactions_item.Qty {
			return 0, errors.New("Insufficient Stock")
		}
		if item.GetPrice() != dto_transactions_item.Price {
			return 0, errors.New("Price Changed")
		}
		if item.GetItemsTypeId() != dto_transactions_item.Items_type_id {
			return 0, errors.New("Items Type Id Is Not Match")
		}
		// <<< Check Repo Item
		total_items_amount += entity_transactions_item.GetTotalPrice()
		entity_transactions_items = append(entity_transactions_items, entity_transactions_item)

		// >>> Count Total Amount per Category | Move this to infrastructure voucher
		// if index == 0 {
		// 	total_amount_per_category = append(total_amount_per_category, TotalAmountPerCategory{
		// 		items_type_id: dto_transactions_item.Item_id,
		// 		total_amount:  dto_transactions_item.Total_price,
		// 	})
		// } else {
		// 	for _, total_amount_per_category_item := range total_amount_per_category {
		// 		if total_amount_per_category_item.items_type_id == dto_transactions_item.Item_id {
		// 			total_amount_per_category_item.total_amount += dto_transactions_item.Total_price
		// 		} else {
		// 			total_amount_per_category = append(total_amount_per_category, TotalAmountPerCategory{
		// 				items_type_id: dto_transactions_item.Item_id,
		// 				total_amount:  dto_transactions_item.Total_price,
		// 			})
		// 		}
		// 	}
		// }
		// <<< Count Total Amount per Category
	}
	// <<< Check Transaction's Item

	// >>> Infrastructure Check & Count Voucher Discount
	// Code is here
	// <<< Infrastructure Check & Count Voucher

	dto_transaction.Is_generated_voucher_succeed = false

	// >>> Count Final Transaction's Amount
	dto_transaction.Final_total_amount = dto_transaction.Total_amount - dto_transaction.Total_discount_amount
	// <<< Count Final Transaction's Amount

	// >>> Repository Store
	repo_transaction_err := ut.repoTransaction.Store(ctx, entity_transaction)
	if repo_transaction_err != nil {
		log.Error().Msg(repo_transaction_err.Error())
		return 0, repo_transaction_err
	}

	repo_transactions_items_err := ut.repoTransactionsItem.Store(ctx, entity_transaction.GetId(), entity_transactions_items)
	if repo_transactions_items_err != nil {
		log.Error().Msg(repo_transactions_items_err.Error())
		return 0, repo_transactions_items_err
	}

	entity_transaction.SetStatus(111)
	entity_transaction.SetTotalAmount(total_items_amount)
	entity_transaction.SetTotalDiscountAmount(0)
	repo_transaction_update_err := ut.repoTransaction.Update(ctx, entity_transaction)
	if repo_transaction_update_err != nil {
		return 0, repo_transaction_update_err
	}
	// <<<< Repository Store

	return 0, nil
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
				return errors.New("400")
			}

			log.Error().Msg("Transaction=>Repository=>Item : " + repo_item_err.Error())
			return errors.New("500")
		}

		if item.GetStock() < dto_transactions_item.Qty {
			return errors.New("Insufficient Stock")
		}
		if item.GetPrice() != dto_transactions_item.Price {
			return errors.New("Price Changed")
		}
		// <<< Check Repo Item
		total_items_amount += entity_transactions_item.GetTotalPrice()
		entity_transactions_items = append(entity_transactions_items, entity_transactions_item)

		// >>> Count Total Amount per Category | Move this to infrastructure voucher
		// if index == 0 {
		// 	total_amount_per_category = append(total_amount_per_category, TotalAmountPerCategory{
		// 		items_type_id: dto_transactions_item.Item_id,
		// 		total_amount:  dto_transactions_item.Total_price,
		// 	})
		// } else {
		// 	for _, total_amount_per_category_item := range total_amount_per_category {
		// 		if total_amount_per_category_item.items_type_id == dto_transactions_item.Item_id {
		// 			total_amount_per_category_item.total_amount += dto_transactions_item.Total_price
		// 		} else {
		// 			total_amount_per_category = append(total_amount_per_category, TotalAmountPerCategory{
		// 				items_type_id: dto_transactions_item.Item_id,
		// 				total_amount:  dto_transactions_item.Total_price,
		// 			})
		// 		}
		// 	}
		// }
		// <<< Count Total Amount per Category
	}
	// <<< Check Transaction's Item

	// >>> Infrastructure Check & Count Voucher Discount
	// Code is here
	// <<< Infrastructure Check & Count Voucher

	dto_transaction.Is_generated_voucher_succeed = false

	// >>> Count Final Transaction's Amount
	dto_transaction.Final_total_amount = dto_transaction.Total_amount - dto_transaction.Total_discount_amount
	// <<< Count Final Transaction's Amount

	// >>> Repository Store
	old_transaction_id := entity_transaction.GetId()

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
	entity_old_transaction.SetUpdateTransactionId(updated_transaction_id)
	repo_old_transaction_update_err := ut.repoTransaction.Update(ctx, entity_old_transaction)
	if repo_old_transaction_update_err != nil {
		return repo_old_transaction_update_err
	}
	// <<<< Repository Store

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

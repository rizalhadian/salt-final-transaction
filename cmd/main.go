package main

import (
	"fmt"
	"log"
	"net/http"
	http_handler "salt-final-transaction/internal/delivery/http/handler"
	infrastructure_customer "salt-final-transaction/internal/infrastructure/customer"
	infrastructure_voucher "salt-final-transaction/internal/infrastructure/voucher"
	repository_mysql "salt-final-transaction/internal/repository/mysql"
	"salt-final-transaction/internal/usecase"
	pkg_database_mysql "salt-final-transaction/pkg/database/mysql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (

	// ============ Connection to Storage / Cache
	http_client     = http.Client{}
	connectionMysql = pkg_database_mysql.InitDBMysql()
	// connectionRedis =
	// ============ Infrastructure
	infrastrctureCustomer = infrastructure_customer.NewInfrastructureCustomer(http_client, "http://27.0.0.1:8000/api/customer")
	infrastrctureVoucher  = infrastructure_voucher.NewInfrastructureVoucher(http_client, "http://127.0.0.3:8000/api/voucher")
	// ============ Repos
	repoTransaction               = repository_mysql.NewRepoTransaction(connectionMysql)
	repoTransactionsItem          = repository_mysql.NewRepoTransactionsItem(connectionMysql)
	repoItem                      = repository_mysql.NewRepoItem(connectionMysql)
	repoCustomersTransactionCount = repository_mysql.NewRepoCustomersTransactionCount(connectionMysql)
	// ============ Usecasese
	UsecaseTransaction               = usecase.NewUsecaseTransaction(infrastrctureCustomer, infrastrctureVoucher, repoTransaction, repoTransactionsItem, repoItem, repoCustomersTransactionCount)
	UsecaseCustomersTransactionCount = usecase.NewUsecaseCustomersTransactionCount(repoCustomersTransactionCount)
)

func main() {
	router := mux.NewRouter()

	http_handler.NewHandlerTransaction(router, UsecaseTransaction)
	http_handler.NewHandlerCustomersTransactionCount(router, UsecaseCustomersTransactionCount)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.2:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Run on 127.0.0.2:8000")

	log.Fatal(srv.ListenAndServe())
}

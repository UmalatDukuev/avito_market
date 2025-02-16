package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"market/internal/handler"
	"market/internal/repository"
	"market/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testDB *sqlx.DB
var testRouter *gin.Engine

func TestMain(m *testing.M) {
	var err error
	testDB, err = sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	resetDatabase()

	itemRepo := repository.NewItemPostgres(testDB)
	transactionRepo := repository.NewTransactionPostgres(testDB)
	balanceService := service.NewBalanceService(transactionRepo)
	itemService := service.NewItemService(itemRepo, balanceService)
	transactionService := service.NewTransactionService(transactionRepo, balanceService)

	h := handler.NewHandler(&service.Service{
		Item:        itemService,
		Transaction: transactionService,
	})

	router := gin.Default()
	router.GET("/api/buyItem/:item", h.BuyItem)
	testRouter = router

	m.Run()
}

func resetDatabase() {
	testDB.Exec("DELETE FROM inventory")
	testDB.Exec("DELETE FROM users")
	testDB.Exec("DELETE FROM merch")
	testDB.Exec("INSERT INTO users (id, name, coins) VALUES (1, 'Alice', 100)")
	testDB.Exec("INSERT INTO merch (id, name, price) VALUES (1, 'cup', 20)")
}

func TestBuyItemIntegration(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/buyItem/cup", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedBalance int
	err := testDB.QueryRow("SELECT coins FROM users WHERE id = 1").Scan(&updatedBalance)
	assert.NoError(t, err)
	assert.Equal(t, 80, updatedBalance)

	var itemCount int
	err = testDB.QueryRow("SELECT quantity FROM inventory WHERE user_id = 1 AND merch_id = 1").Scan(&itemCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, itemCount)
}

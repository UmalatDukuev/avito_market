package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testDB *sqlx.DB
var testRouter *gin.Engine

func resetTransactionTestDB() {
	testDB.Exec("DELETE FROM transactions")
	testDB.Exec("DELETE FROM users")
	testDB.Exec("INSERT INTO users (id, name, coins) VALUES (1, 'Alice', 100), (2, 'Bob', 50)")
}

func TestSendCoinsIntegration(t *testing.T) {
	resetTransactionTestDB()

	req, _ := http.NewRequest("POST", "/api/sendCoins?from=1&to=2&amount=30", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var senderBalance, receiverBalance int
	testDB.QueryRow("SELECT coins FROM users WHERE id = 1").Scan(&senderBalance)
	testDB.QueryRow("SELECT coins FROM users WHERE id = 2").Scan(&receiverBalance)

	assert.Equal(t, 70, senderBalance)
	assert.Equal(t, 80, receiverBalance)

	var transactionCount int
	testDB.QueryRow("SELECT COUNT(*) FROM transactions WHERE from_user_id = 1 AND to_user_id = 2").Scan(&transactionCount)
	assert.Equal(t, 1, transactionCount)
}

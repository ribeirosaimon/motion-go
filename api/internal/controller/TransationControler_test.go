package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/dto"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"
	"github.com/stretchr/testify/assert"
)

func TestTransactionController_BalanceController(t *testing.T) {
	w, c, loggedUser, _, _ := configTest()
	transactionRepository := repository.NewTransactionRepository(db.Conn.GetPgsqTemplate())
	transaction := sqlDomain.Transaction{
		Value:     100,
		SessionId: loggedUser.SessionId,
		ProfileId: loggedUser.ProfileId,
	}
	transaction, err := transactionRepository.Save(transaction)
	if err != nil {
		panic(err)
	}

	NewTransactionController().Balance(c)

	var response dto.Deposit
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.Value, transaction.Value)
}

func TestTransactionController_Deposit(t *testing.T) {
	w, c, loggedUser, _, _ := configTest()
	jsonBytes, err := json.Marshal(dto.Deposit{Value: 125})
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	NewTransactionController().Deposit(c)
	var response sqlDomain.Transaction
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	transactionRepository := repository.NewTransactionRepository(db.Conn.GetPgsqTemplate())
	transactionDb, err := transactionRepository.FindByField("session_id", loggedUser.SessionId)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, transactionDb.Value, response.Value)
	assert.Equal(t, transactionDb.OperationType, sqlDomain.DEPOSIT)
}

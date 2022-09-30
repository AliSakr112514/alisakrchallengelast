package test

import (
	TransactionRepo "Challenge/internal/repositories/transactions"
	"Challenge/test/utils"
	uuid2 "github.com/google/uuid"
	"github.com/satori/go.uuid"
	"log"
	"testing"
	"time"
)

func MainTest() error {
	Db, err := utils.GetDatabaseConnection(&utils.Configs)
	if err != nil {
		log.Fatalf("The test failed during connecting to database: %s", err)
		return err
	}
	TransactionRepo.Db = Db
	return nil
}
func TestCreateTransaction(t *testing.T) {
	err := MainTest()
	if err != nil {
		t.Errorf("The test failed during connecting to database: %s", err)
	}
	testTrans := TransactionRepo.Transaction{
		Amount:   150,
		Currency: "USD",
	}
	_, err = TransactionRepo.TransactionRepo.CreateTransaction(&testTrans)
	if err != nil {
		t.Errorf("The test failed during creatin Transaction and the error is: %s", err)
	}
}

func TestGetAllTrans(t *testing.T) {
	err := MainTest()
	if err != nil {
		t.Errorf("The test failed during connecting to database: %s", err)
	}
	_, err = TransactionRepo.TransactionRepo.GetAllTransactions()
	if err != nil {
		t.Errorf("Error in getting all transactions: %s", err)
	}
}
func TestUpdateTrans(t *testing.T) {
	err := MainTest()
	if err != nil {
		t.Errorf("The test failed during connecting to database: %s", err)
	}
	testTrans := TransactionRepo.Transaction{
		// Id is related to "github.com/google/uuid" and FromStringOrNil returns "github.com/satori/go.uuid"
		Id:        uuid2.UUID(uuid.FromStringOrNil("69359037-9599-48e7-b8f2-48393c019135")),
		Amount:    150,
		Currency:  "USD",
		Createdat: time.Now().String(),
		Isactive:  true,
	}
	_, err = TransactionRepo.TransactionRepo.UpdateTransaction(&testTrans)
	if err != nil {
		t.Errorf("Error in getting all transactions: %s", err)
	}
}

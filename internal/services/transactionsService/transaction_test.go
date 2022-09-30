package TransactionService

import (
	DB "Challenge/internal/adapters/db"
	TransactionRepo "Challenge/internal/repositories/transactions"
	uuid2 "github.com/google/uuid"
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestCreateTransaction(t *testing.T) {
	DatabaseConn, error := DB.GetDatabaseConnectionForTesting()
	if error != nil {
		t.Errorf("The test failed during connecting to database: %s", error)
	}
	TransactionRepo.Db = DatabaseConn
	var trans TransactionRepo.Transaction
	trans = TransactionRepo.Transaction{
		// Id is related to "github.com/google/uuid" and FromStringOrNil returns "github.com/satori/go.uuid"
		Id:        uuid2.UUID(uuid.FromStringOrNil("69359037-9599-48e7-b8f2-48393c019135")),
		Amount:    150,
		Currency:  "USD",
		Createdat: time.Now().String(),
		Isactive:  true,
	}
	_, error = TransactionService.CreateTransactionService(&trans)
	if error != nil {
		t.Errorf("The test failed during creating transaction: %v", error)
	}
}

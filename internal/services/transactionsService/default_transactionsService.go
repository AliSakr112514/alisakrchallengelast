package TransactionService

import (
	stream "Challenge/internal/adapters/stream"
	TransactionRepo "Challenge/internal/repositories/transactions"
	"errors"
)

type DefaultTransactionService struct {
	transactionService ITransactionService
}

var TransactionService DefaultTransactionService

func (s *DefaultTransactionService) CreateTransactionService(trans *TransactionRepo.Transaction) (TransactionRepo.Transaction, error) {
	if trans.Amount < LowAmount || trans.Amount > HighAmount {
		return TransactionRepo.Transaction{}, errors.New("The amount must be between 0 and 100000")
	}
	if len(trans.Currency) > NumberOfAvailableCharacter {
		return TransactionRepo.Transaction{}, errors.New("Currency iso should be 3 char")
	}
	result, err := TransactionRepo.TransactionRepo.CreateTransaction(trans)
	stream.Produce(trans)
	return result, err
}

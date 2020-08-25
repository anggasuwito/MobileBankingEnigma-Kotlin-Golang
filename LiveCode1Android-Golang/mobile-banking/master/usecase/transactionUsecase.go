package usecase

import "banking/master/model"

type TransactionUseCase interface {
	GetAllDataTransaction() ([]*model.TransactionModel, error)
	GetTransactionById(id string) (model.TransactionModel, error)
	UpdateTransactionById(id string, DataTransaction model.TransactionModel) error
	DeleteTransactionById(id string) error
	InsertTransaction(DataTransaction model.TransactionModel) error
}

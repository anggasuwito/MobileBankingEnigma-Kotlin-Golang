package repositories

import "banking/master/model"

type TransactionRepo interface {
	GetAllDataTransaction() ([]*model.TransactionModel, error)
	GetTransactionById(id string) (model.TransactionModel, error)
	UpdateTransactionById(id string, DataTransaction model.TransactionModel) error
	DeleteTransactionById(id string) error
	InsertTransaction(DataTransaction model.TransactionModel) error
}

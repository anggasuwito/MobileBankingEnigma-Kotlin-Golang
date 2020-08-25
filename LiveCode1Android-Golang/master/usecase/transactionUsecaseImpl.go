package usecase

import (
	"banking/master/model"
	"banking/master/repositories"
)

type TransactionUsecaseImpl struct {
	TransactionRepo repositories.TransactionRepo
}

func (au TransactionUsecaseImpl) GetAllDataTransaction() ([]*model.TransactionModel, error) {
	Transaction, err := au.TransactionRepo.GetAllDataTransaction()
	if err != nil {
		return nil, err
	}
	return Transaction, nil
}
func (s TransactionUsecaseImpl) GetTransactionById(id string) (model.TransactionModel, error) {
	Transaction, err := s.TransactionRepo.GetTransactionById(id)

	return Transaction, err
}
func (s TransactionUsecaseImpl) UpdateTransactionById(id string, DataTransaction model.TransactionModel) error {

	err := s.TransactionRepo.UpdateTransactionById(id, DataTransaction)
	if err != nil {
		return err
	}
	return err
}
func (s TransactionUsecaseImpl) DeleteTransactionById(id string) error {
	err := s.TransactionRepo.DeleteTransactionById(id)
	return err
}
func (s TransactionUsecaseImpl) InsertTransaction(DataTransaction model.TransactionModel) error {

	err := s.TransactionRepo.InsertTransaction(DataTransaction)
	if err != nil {
		return err
	}
	return err
}
func InitTransactionUseCase(TransactionRepo repositories.TransactionRepo) TransactionUseCase {
	return &TransactionUsecaseImpl{TransactionRepo}
}

package usecase

import (
	"banking/master/model"
	"banking/master/repositories"
)

type UserUsecaseImpl struct {
	UserRepo repositories.UserRepo
}

func (au UserUsecaseImpl) GetAllDataUser() ([]*model.UserModel, error) {
	User, err := au.UserRepo.GetAllDataUser()
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (s UserUsecaseImpl) GetUserById(id string) (model.UserModel, error) {
	User, err := s.UserRepo.GetUserById(id)

	return User, err
}
func (s UserUsecaseImpl) UpdateUserById(id string, DataUser model.UserModel) error {
	err := s.UserRepo.UpdateUserById(id, DataUser)
	if err != nil {
		return err
	}
	return nil
}
func (s UserUsecaseImpl) DeleteUserById(id string) error {
	err := s.UserRepo.DeleteUserById(id)
	return err
}
func (s UserUsecaseImpl) InsertUser(DataUser model.UserModel) error {
	err := s.UserRepo.InsertUser(DataUser)
	if err != nil {
		return err
	}
	return nil
}
func InitUserUseCase(UserRepo repositories.UserRepo) UserUseCase {
	return &UserUsecaseImpl{UserRepo}
}

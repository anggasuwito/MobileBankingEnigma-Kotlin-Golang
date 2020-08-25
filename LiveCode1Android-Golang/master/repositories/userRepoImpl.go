package repositories

import (
	"banking/master/model"
	"database/sql"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (ar UserRepoImpl) GetAllDataUser() ([]*model.UserModel, error) {
	dataUser := []*model.UserModel{}
	query := `select * from user`
	data, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		User := model.UserModel{}
		var err = data.Scan(&User.IdUser,
			&User.Name, &User.Balance)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, &User)
	}
	return dataUser, nil
}
func (ar UserRepoImpl) GetUserById(id string) (model.UserModel, error) {
	var User model.UserModel
	query := `select * from user where id=?`
	err := ar.db.QueryRow(query, id).Scan(&User.IdUser,
		&User.Name, &User.Balance)

	return User, err
}
func (s UserRepoImpl) DeleteUserById(id string) error {
	tr, err := s.db.Begin()

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = s.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := "delete from user where id=?"
	row, err := s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		return err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return err
}
func (ar UserRepoImpl) UpdateUserById(id string, DataUser model.UserModel) error {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := `update user set name=?, balance=? where id=?`
	row, err := ar.db.Query(query, &DataUser.Name,
		&DataUser.Balance, id)
	if err != nil {
		tr.Rollback()
		return err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return err
}
func (ar UserRepoImpl) InsertUser(DataUser model.UserModel) error {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	query := "insert into user (name,balance) values (?,?)"
	row, err := ar.db.Query(query, &DataUser.Name, &DataUser.Balance)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}
func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db}
}

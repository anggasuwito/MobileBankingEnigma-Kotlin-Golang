package repositories

import (
	"banking/master/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TransactionRepoImpl struct {
	db *sql.DB
}

func (ar TransactionRepoImpl) GetAllDataTransaction() ([]*model.TransactionModel, error) {
	dataTransaction := []*model.TransactionModel{}
	query := fmt.Sprintf(`select * from transaction`)
	fmt.Print(query)
	data, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Transaction := model.TransactionModel{}
		var err = data.Scan(&Transaction.IdTransaction, &Transaction.Date, &Transaction.Amount)
		if err != nil {
			return nil, err
		}
		dataTransaction = append(dataTransaction, &Transaction)
	}
	return dataTransaction, nil
}
func (ar TransactionRepoImpl) GetTransactionById(id string) (model.TransactionModel, error) {
	var Transaction model.TransactionModel
	query := `select * from transaction where id=?`
	err := ar.db.QueryRow(query, id).Scan(&Transaction.Date, &Transaction.Amount)

	return Transaction, err
}
func (s TransactionRepoImpl) DeleteTransactionById(id string) error {
	tr, err := s.db.Begin()

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = s.GetTransactionById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := "delete from transaction where id=?"
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
func (ar TransactionRepoImpl) UpdateTransactionById(id string, DataTransaction model.TransactionModel) error {
	tr, err := ar.db.Begin()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetTransactionById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	fmt.Println(DataTransaction)
	query := `update transaction set date=?,amount=? where idn=?`
	row, err := ar.db.Query(query,
		&DataTransaction.Date, &DataTransaction.Amount, id)
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
func (ar TransactionRepoImpl) InsertTransaction(DataTransaction model.TransactionModel) error {
	tr, err := ar.db.Begin()
	datetime := time.Now()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	query := "insert into transaction (amount,date) values (?,?)"
	query2 := `update user set balance=balance-? where id=?`
	row, err := ar.db.Query(query, &DataTransaction.Amount, datetime)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	}

	_, err = ar.db.Query(query2, &DataTransaction.Amount, 1)
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
func InitTransactionRepoImpl(db *sql.DB) TransactionRepo {
	return &TransactionRepoImpl{db}
}

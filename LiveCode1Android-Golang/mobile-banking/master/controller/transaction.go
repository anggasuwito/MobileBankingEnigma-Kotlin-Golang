package controller

import (
	"banking/master/model"
	"banking/master/usecase"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	TransactionUseCase usecase.TransactionUseCase
}

func TransactionController(r *mux.Router, service usecase.TransactionUseCase) {
	TransactionHandler := TransactionHandler{service}
	r.HandleFunc("/transactions", TransactionHandler.ListTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id}", TransactionHandler.TransactionById).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id}", TransactionHandler.DeleteById).Methods(http.MethodDelete)
	r.HandleFunc("/transaction/{id}", TransactionHandler.UpdateTransaction).Methods(http.MethodPut)
	r.HandleFunc("/transaction", TransactionHandler.InsertTransaction).Methods(http.MethodPost)
}

func (s TransactionHandler) ListTransaction(w http.ResponseWriter, r *http.Request) {

	Transaction, err := s.TransactionUseCase.GetAllDataTransaction()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		byteOfTransactions, err := json.Marshal(Transaction)
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfTransactions)
	}

}
func (s TransactionHandler) TransactionById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTransaction := param["id"]
	Transaction, err := s.TransactionUseCase.GetTransactionById(idTransaction)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Id Not Found"))
		}
	} else {
		byteOfTransactions, err := json.Marshal(Transaction)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfTransactions)
	}
}

func (s TransactionHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	var Transaction model.TransactionModel
	param := mux.Vars(r)
	idTransaction := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&Transaction)
	fmt.Println(idTransaction)
	err := s.TransactionUseCase.DeleteTransactionById(idTransaction)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Deleted"))
	}

}
func (s TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var Transaction model.TransactionModel
	param := mux.Vars(r)
	idTransaction := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&Transaction)
	fmt.Println(Transaction)
	err := s.TransactionUseCase.UpdateTransactionById(idTransaction, Transaction)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Updated"))
	}
}
func (s TransactionHandler) InsertTransaction(w http.ResponseWriter, r *http.Request) {
	var Transaction model.TransactionModel
	_ = json.NewDecoder(r.Body).Decode(&Transaction)
	err := s.TransactionUseCase.InsertTransaction(Transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Insert Failed"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Insert Success"))
	}
}

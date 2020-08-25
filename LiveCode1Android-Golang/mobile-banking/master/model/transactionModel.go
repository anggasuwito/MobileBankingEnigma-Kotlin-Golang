package model

type TransactionModel struct {
	IdTransaction string `json:"id"`
	Date          string `json:"date"`
	Amount        string `json:"amount"`
}

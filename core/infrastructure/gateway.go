package todo

import (
	"context"
	"database/sql"
	"hex-arch-go/core/entities"
	"net/http"
)

// the Gateway for access to Storage
type ToDoGateway interface {
	ListRates() (int, []entities.Rates)
	ListTransactions() (int, []entities.Transactions)
}

// The Domain Logic
type ToDoLogic struct {
	St ToDoStorage
}

// List ToDo
func (t *ToDoLogic) ListRates() (int, []entities.Rates) {
	// Domain logic
	return http.StatusOK, t.St.listRatesInDb()
}

func (t *ToDoLogic) ListTransactions() (int, []entities.Transactions) {
	// Domain logic
	return http.StatusOK, t.St.listTransactionsInDb()
}

// Constructor
func NewToDoGateway(ctx context.Context, db *sql.DB) ToDoGateway {
	return &ToDoLogic{NewToDoStorage(ctx, db)}
}

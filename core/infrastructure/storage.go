package todo

import (
	"context"
	"database/sql"
	"hex-arch-go/core/entities"
	"math/rand"
	"strconv"
)

const MAX_RATE = 1.5
const MIN_RATE = 0.1

const MAX_AMOUNT = 10000
const MIN_AMOUNT = 100

// The TODO Repository
type ToDoStorage interface {
	listRatesInDb() []entities.Rates
	listTransactionsInDb() []entities.Transactions
}

// The TODO Service
type ToDoService struct {
	db  *sql.DB
	ctx context.Context
}

func (t *ToDoService) getRandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP", "JPY", "AUD", "CAD", "SEK", "RUB", "INR"}
	return currencies[rand.Intn(len(currencies))]
}

func (t *ToDoService) getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (t *ToDoService) getRandomFloat(max, min float64) float64 {
	var number = min + rand.Float64()*(max-min)

	if number < 0 {
		number = -number
	}

	return number
}

func (t *ToDoService) getRandomSku() string {
	return "T" + strconv.Itoa(t.getRandomInt(100000, 999999))
}

func (t *ToDoService) listRatesInDb() []entities.Rates {
	return []entities.Rates{
		{
			From: "EUR",
			To:   "USD",
			Rate: 1.04,
		},
		{
			From: "USD",
			To:   "EUR",
			Rate: 0.96,
		},
		{
			From: "GBP",
			To:   "EUR",
			Rate: 1.16,
		},
		{
			From: "JPY",
			To:   "USD",
			Rate: 0.0073,
		},
		{
			From: "AUD",
			To:   "INR",
			Rate: 55.29,
		},
		{
			From: "CAD",
			To:   "USD",
			Rate: 0.74,
		},
		{
			From: "SEK",
			To:   "USD",
			Rate: 0.096,
		},
		{
			From: "RUB",
			To:   "SEK",
			Rate: 0.17,
		},
		{
			From: "INR",
			To:   "EUR",
			Rate: 0.012,
		},
	}
}

func (t *ToDoService) listTransactionsInDb() []entities.Transactions {
	var transactions []entities.Transactions
	var max = t.getRandomInt(15, 100)
	var maxSku = t.getRandomInt(1, 10)
	for i := 1; i < max; i++ {
		var sku = t.getRandomSku()
		for i := 1; i < maxSku; i++ {
			var transaction entities.Transactions
			transaction.Sku = sku
			transaction.Amount = t.getRandomFloat(MAX_AMOUNT, MIN_AMOUNT)
			transaction.Currency = t.getRandomCurrency()
			transactions = append(transactions, transaction)
		}

	}

	return transactions
}

// Constructor
func NewToDoStorage(ctx context.Context, db *sql.DB) ToDoStorage {
	return &ToDoService{db: db, ctx: ctx}
}

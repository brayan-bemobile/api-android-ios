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
	currencies := []string{"EUR", "USD", "GBP", "JPY", "AUD", "CAD", "CHF", "CNY", "SEK", "NZD", "MXN", "SGD", "HKD", "NOK", "KRW", "TRY", "RUB", "INR", "BRL", "ZAR"}
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
	var rates []entities.Rates
	var max = t.getRandomInt(3, 11)
	for i := 1; i < max; i++ {
		var rate entities.Rates
		rate.From = t.getRandomCurrency()
		rate.To = t.getRandomCurrency()
		rate.Rate = t.getRandomFloat(MAX_RATE, MIN_RATE)
		rates = append(rates, rate)
	}

	return rates
}

func (t *ToDoService) listTransactionsInDb() []entities.Transactions {
	var transactions []entities.Transactions
	var max = t.getRandomInt(3, 11)
	for i := 1; i < max; i++ {
		var transaction entities.Transactions
		transaction.Sku = t.getRandomSku()
		transaction.Amount = t.getRandomFloat(MAX_AMOUNT, MIN_AMOUNT)
		transaction.Currency = t.getRandomCurrency()
		transactions = append(transactions, transaction)
	}

	return transactions
}

// Constructor
func NewToDoStorage(ctx context.Context, db *sql.DB) ToDoStorage {
	return &ToDoService{db: db, ctx: ctx}
}

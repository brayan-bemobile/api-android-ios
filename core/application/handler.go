package app_todo

// all handlers are called in /server/routes.go

import (
	"context"
	"database/sql"
	todo "hex-arch-go/core/infrastructure"

	"github.com/labstack/echo/v4"
)

// The HTTP Handler for TODO
type ToDoHTTPService struct {
	gtw todo.ToDoGateway
}

func (t *ToDoHTTPService) ListRates(c echo.Context) error {
	status, res := t.gtw.ListRates()
	return c.JSON(status, res)
}

func (t *ToDoHTTPService) ListTransactions(c echo.Context) error {
	status, res := t.gtw.ListTransactions()
	return c.JSON(status, res)
}

// Constructor
func NewToDoHTTPService(ctx context.Context, db *sql.DB) *ToDoHTTPService {
	return &ToDoHTTPService{todo.NewToDoGateway(ctx, db)}
}

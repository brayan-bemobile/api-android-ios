package server

import (
	app_todo "hex-arch-go/core/application"
)

// a TODO Routes
func (es *EchoServer) toDoRoutes() {
	// call the TODO HTTP Service
	todo := app_todo.NewToDoHTTPService(es.ctx, es.db)

	es.GET("/rates", todo.ListRates)
	es.GET("/transactions", todo.ListTransactions)
}

//func anotherRoutes()...

// All routes
func (es *EchoServer) routes() {
	es.toDoRoutes()
	//es.anotherRoutes()
}

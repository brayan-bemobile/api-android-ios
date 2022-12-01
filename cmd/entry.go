package cmd

import (
	"context"
	"hex-arch-go/env"
	"hex-arch-go/server"
)

// Start the server
func Start() {
	// App context
	ctx := context.Background()

	// env config
	_env := env.GetEnv(".env.development")

	// Run database with env config
	//db := database.NewMySQLDatabase(ctx, _env).ConnectDB() // or work with mysql

	// Run server with context, database
	//server.NewGinServer(ctx, db, _env.SERVER_PORT).Run() // with Gin for example
	server.NewEchoServer(ctx, nil, _env.SERVER_PORT).Run()
}

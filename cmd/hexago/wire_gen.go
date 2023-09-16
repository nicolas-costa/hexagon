// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"hexago/internal/application"
	"hexago/internal/infrastructure/controller"
	"hexago/internal/infrastructure/repositories"
	"hexago/internal/infrastructure/server"
)

// Injectors from wire.go:

func initialize() server.FiberServer {
	healthService := application.NewHealthService()
	healthController := controller.NewHealthController(healthService)
	coinGeckoClient := repositories.NewCoinRepository()
	coinService := application.NewCoinService(coinGeckoClient)
	coinController := controller.NewCoinController(coinService)
	fiberServer := server.NewFiberServer(healthController, coinController)
	return fiberServer
}

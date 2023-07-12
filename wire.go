//go:build wireinject
// +build wireinject

package main

import (
	"crud-barang/Config"
	"crud-barang/Controller"
	"crud-barang/Middleware"
	"crud-barang/Repository"
	"crud-barang/Route"
	"crud-barang/Services"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var barangSet = wire.NewSet(
	Repository.BarangRepositoryControllerProvider,
	Services.BarangServiceControllerProvider,
	Controller.BarangControlerControllerProvider,
)

func InitializedServer() *http.Server {
	wire.Build(
		Config.NewDB,
		validator.New,
		barangSet,
		Route.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		Middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

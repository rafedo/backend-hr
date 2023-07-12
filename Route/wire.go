//go:build wireinject
// +build wireinject

package Route

import (
	"crud-barang/Controller"
	"crud-barang/Repository"
	"crud-barang/Services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func barangDI(db *gorm.DB) *Controller.BarangControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.BarangRepositoryControllerProvider,
		Services.BarangServiceControllerProvider,
		Controller.BarangControllerControllerProvider,

		wire.Bind(new(Repository.BarangRepositoryHandler), new(*Repository.BarangRepositoryImpl)),
		wire.Bind(new(Services.BarangServiceHandler), new(*Services.BarangServiceImpl)),
		wire.Bind(new(Controller.BarangControllerHandler), new(*Controller.BarangControllerImpl)),
	),
	))
	return &Controller.BarangControllerImpl{}
}

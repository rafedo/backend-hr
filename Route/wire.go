//go:build wireinject
// +build wireinject

package Route

import (
	"muhammadiyah/Controller"
	"muhammadiyah/Repository"
	"muhammadiyah/Services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func PWMDI(db *gorm.DB) *Controller.PWMControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.PWMRepositoryControllerProvider,
		Services.PWMServiceControllerProvider,
		Repository.AnggotaRepositoryControllerProvider,
		Controller.PWMControllerControllerProvider,

		wire.Bind(new(Repository.PWMRepositoryHandler), new(*Repository.PWMRepositoryImpl)),
		wire.Bind(new(Repository.AnggotaRepositoryHandler), new(*Repository.AnggotaRepositoryImpl)),
		wire.Bind(new(Services.PWMServiceHandler), new(*Services.PWMServiceImpl)),
		wire.Bind(new(Controller.PWMControllerHandler), new(*Controller.PWMControllerImpl)),
	),
	))
	return &Controller.PWMControllerImpl{}
}

func DepartmentDI(db *gorm.DB) *Controller.DepartmentControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.DepartmentRepositoryControllerProvider,
		Services.DepartmentServiceControllerProvider,
		Controller.DepartmentControllerControllerProvider,

		wire.Bind(new(Repository.DepartmentRepositoryHandler), new(*Repository.DepartmentRepositoryImpl)),
		wire.Bind(new(Services.DepartmentServiceHandler), new(*Services.DepartmentServiceImpl)),
		wire.Bind(new(Controller.DepartmentControllerHandler), new(*Controller.DepartmentControllerImpl)),
	),
	))
	return &Controller.DepartmentControllerImpl{}
}

func PengurusDI(db *gorm.DB) *Controller.PengurusControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.PengurusRepositoryControllerProvider,
		Services.PengurusServiceControllerProvider,
		Controller.PengurusControllerControllerProvider,

		wire.Bind(new(Repository.PengurusRepositoryHandler), new(*Repository.PengurusRepositoryImpl)),
		wire.Bind(new(Services.PengurusServiceHandler), new(*Services.PengurusServiceImpl)),
		wire.Bind(new(Controller.PengurusControllerHandler), new(*Controller.PengurusControllerImpl)),
	),
	))
	return &Controller.PengurusControllerImpl{}
}
func AuthDI(db *gorm.DB) *Controller.AuthControllerImpl {
	panic(wire.Build(wire.NewSet(
		Controller.AuthControllerControllerProvider,
		Services.AuthServiceControllerProvider,
		Repository.AuthRepositoryControllerProvider,
		Repository.PengurusRepositoryControllerProvider,
		Repository.AnggotaRepositoryControllerProvider,
		wire.Bind(new(Repository.AuthRepositoryHandler), new(*Repository.AuthRepositoryImpl)),
		wire.Bind(new(Repository.PengurusRepositoryHandler), new(*Repository.PengurusRepositoryImpl)),
		wire.Bind(new(Repository.AnggotaRepositoryHandler), new(*Repository.AnggotaRepositoryImpl)),
		wire.Bind(new(Services.AuthServiceHandler), new(*Services.AuthServiceImpl)),
		wire.Bind(new(Controller.AuthControllerHandler), new(*Controller.AuthControllerImpl)),
	),
	))

	return &Controller.AuthControllerImpl{}
}

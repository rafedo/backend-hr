package Route

import (
	"muhammadiyah/Config"
	"muhammadiyah/Middleware"

	"github.com/gofiber/fiber/v2"
)

func CatatanRouter(c fiber.Router) {
	pwm := PWMDI(Config.DB)
	c.Get("/wilayah", Middleware.JwtMiddleware(), Middleware.CheckPermissions([]string{"all"}, []string{"all"}, []string{"wilayah"}), pwm.FindAllWilayah)
	c.Post("/wilayah", pwm.CreateWilayah)
	c.Patch("/wilayah", pwm.UpdateWilayah)
	c.Delete("/wilayah/:id", pwm.DeleteWilayah)

	c.Get("/daerah", pwm.FindAllDaerah)
	c.Post("/daerah", pwm.CreateDaerah)
	c.Patch("/daerah", pwm.UpdateDaerah)
	c.Delete("/daerah/:id", pwm.DeleteDaerah)

	c.Get("/cabang", pwm.FindAllCabang)
	c.Post("/cabang", pwm.CreateCabang)
	c.Patch("/cabang", pwm.UpdateCabang)
	c.Delete("/cabang/:id", pwm.DeleteCabang)

	c.Get("/ranting", pwm.FindAllRanting)
	c.Post("/ranting", pwm.CreateRanting)
	c.Patch("/ranting", pwm.UpdateRanting)
	c.Delete("/ranting/:id", pwm.DeleteRanting)

	c.Get("/member", pwm.FindAllMember)
	c.Get("/member/total", pwm.CountMember)
	c.Post("/member", pwm.CreateMember)
	c.Patch("/member", pwm.UpdateMember)
	c.Delete("/member/:id", pwm.DeleteMember)

	// Routes for finding members by IDs
	c.Get("/members/wilayah/:wilayahID", pwm.FindAllMemberByWilayahID)
	c.Get("/members/cabang/:cabangID", pwm.FindAllMemberByCabangID)
	c.Get("/members/daerah/:daerahID", pwm.FindAllMemberByDaerahID)
	c.Get("/members/ranting/:rantingID", pwm.FindAllMemberByRantingID)

	// Routes for finding entities by IDs
	c.Get("/wilayah/:wilayahID", pwm.FindWilayahByID)
	c.Get("/wilayah/:wilayahID/daerah", pwm.FindDaerahByWilayahID)
	c.Get("/wilayah/:wilayahID/daerah/:daerahID", pwm.FindDaerahByID)
	c.Get("/wilayah/:wilayahID/daerah/:daerahID/cabang", pwm.FindCabangByDaerahID)
	c.Get("/wilayah/:wilayahID/daerah/:daerahID/cabang/:cabangID", pwm.FindCabangByID)
	c.Get("/wilayah/:wilayahID/daerah/:daerahID/cabang/:cabangID/ranting", pwm.FindRantingByCabangID)
	c.Get("/wilayah/:wilayahID/daerah/:daerahID/cabang/:cabangID/ranting/:rantingID", pwm.FindRantingByID)

	dept := DepartmentDI(Config.DB)
	c.Get("/departments", dept.FindAllDepartments)
	c.Post("/departments", dept.CreateDepartment)
	c.Patch("/departments", dept.UpdateDepartment)
	c.Delete("/departments/:id", dept.DeleteDepartment)

	c.Get("/placements", dept.FindAllPlacements)
	c.Post("/placements", dept.CreatePlacement)
	c.Patch("/placements", dept.UpdatePlacement)
	c.Delete("/placements/:id", dept.DeletePlacement)

	c.Get("/positions", dept.FindAllPositions)
	c.Post("/positions", dept.CreatePosition)
	c.Patch("/positions", dept.UpdatePosition)
	c.Delete("/positions/:id", dept.DeletePosition)
	pengurus := PengurusDI(Config.DB)
	c.Post("/pengurus", pengurus.CreatePengurus)
	c.Patch("/pengurus", pengurus.UpdatePengurus)
	c.Get("/pengurus", pengurus.FindAllPengurus)
	c.Delete("/pengurus/:id", pengurus.DeletePengurus)
	c.Get("/pengurus/:id", pengurus.FindPengurusByID)
	auth := AuthDI(Config.DB)
	c.Post("/login", auth.Login)
	c.Post("/register", auth.Register)
	c.Get("/user", auth.FindAllUser)
	c.Get("/user/:userID", auth.FindUserByID)

}

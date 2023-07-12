package Route

import (
	"crud-barang/Config"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(c *fiber.App) {
	r := barangDI(Config.DB)

	c.Get("/api/barang", r.FindAll)
	c.Get("/api/barang/:barangId", r.FindById)
	c.Post("/api/barang", r.Create)
	c.Patch("/api/barang/", r.Update)
	c.Delete("/api/barang/:barangId", r.Delete)

}

package Route

import (
	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {
	newRouter := c.Group("/api")
	muhammadiyahGroup := newRouter.Group("/muhammadiyah")
	CatatanRouter(muhammadiyahGroup)
	sekolahGroup := newRouter.Group("/sekolah")
	SekolahRouter(sekolahGroup)
}

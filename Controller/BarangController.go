package Controller

import (
	"crud-barang/Constant"
	"crud-barang/Model/web"
	"crud-barang/Services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BarangControllerHandler interface {
	Create(c *fiber.Ctx) (err error)
	Update(c *fiber.Ctx) (err error)
	Delete(c *fiber.Ctx) (err error)
	FindById(c *fiber.Ctx) (err error)
	FindAll(c *fiber.Ctx) (err error)
}
type BarangControllerImpl struct {
	service Services.BarangServiceHandler
}

func BarangControllerControllerProvider(service Services.BarangServiceHandler) *BarangControllerImpl {
	return &BarangControllerImpl{
		service: service,
	}
}
func (h *BarangControllerImpl) Create(c *fiber.Ctx) (err error) {
	var (
		request    []web.BarangCreateRequest
		serviceErr *web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		fmt.Println("error at create barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(web.SuccessResponse("create succeed", nil))
}
func (h *BarangControllerImpl) Update(c *fiber.Ctx) (err error) {
	var (
		request    web.BarangUpdateRequest
		serviceErr *web.ServiceErrorDto
	)

	// Parse request body
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update barang
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		fmt.Println("error at update barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(web.SuccessResponse("update succeed", nil))
}

func (h *BarangControllerImpl) Delete(c *fiber.Ctx) error {
	// Get barang ID from URL query
	var serviceErr *web.ServiceErrorDto

	barangID, err := c.ParamsInt("barangId")

	// Convert barang ID to int

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(web.ErrorResponse("invalid barang ID", err))
	}

	// Call service to delete barang
	if serviceErr = h.service.Delete(int64(barangID)); serviceErr != nil {
		fmt.Println("error at delete barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(web.SuccessResponse("delete succeed", nil))
}

func (h *BarangControllerImpl) FindById(c *fiber.Ctx) error {
	var (
		serviceErr *web.ServiceErrorDto
		response   web.BarangResponse
		id, _      = c.ParamsInt("barangId")
	)

	if response, serviceErr = h.service.FindById(int64(id)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(web.SuccessResponse("data barang", response))
}
func (h *BarangControllerImpl) FindAll(c *fiber.Ctx) (err error) {
	var (
		serviceErr *web.ServiceErrorDto
		response   []web.BarangResponse
	)

	if response, serviceErr = h.service.FindAll(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(web.SuccessResponse("data barang", response))
}

package Controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"muhammadiyah/Constant"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"muhammadiyah/Services"
	"net/http"
	"strconv"
)

type PengurusControllerHandler interface {
	CreatePengurus(c *fiber.Ctx) (err error)
	UpdatePengurus(c *fiber.Ctx) (err error)
	FindAllPengurus(c *fiber.Ctx) (err error)
	DeletePengurus(c *fiber.Ctx) (err error)
	FindPengurusByID(c *fiber.Ctx) (err error)
}

type PengurusControllerImpl struct {
	service Services.PengurusServiceHandler
}

func PengurusControllerControllerProvider(service Services.PengurusServiceHandler) *PengurusControllerImpl {
	return &PengurusControllerImpl{
		service: service,
	}
}

func (h *PengurusControllerImpl) CreatePengurus(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreatePengurusRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreatePengurus(request); serviceErr != nil {
		fmt.Println("error at create pengurus service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create pengurus succeed", nil))
}

func (h *PengurusControllerImpl) UpdatePengurus(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdatePengurusRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdatePengurus(request); serviceErr != nil {
		fmt.Println("error at Update pengurus service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update pengurus succeed", nil))
}

func (h *PengurusControllerImpl) FindAllPengurus(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.PengurusResponse
	)

	if response, serviceErr = h.service.GetAllPengurus(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data pengurus", response))
}

func (h *PengurusControllerImpl) DeletePengurus(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeletePengurus(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted pengurus with ID %d", id), nil))
}

func (h *PengurusControllerImpl) FindPengurusByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
		response   Domain.PengurusResponse
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.GetPengurusByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data pengurus", response))
}

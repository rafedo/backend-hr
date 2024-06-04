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

type DepartmentControllerHandler interface {
	CreateDepartment(c *fiber.Ctx) (err error)
	UpdateDepartment(c *fiber.Ctx) (err error)
	FindAllDepartments(c *fiber.Ctx) (err error)
	FindDepartmentsByPenempatanID(c *fiber.Ctx) (err error)
	DeleteDepartment(c *fiber.Ctx) (err error)

	CreatePlacement(c *fiber.Ctx) (err error)
	UpdatePlacement(c *fiber.Ctx) (err error)
	FindAllPlacements(c *fiber.Ctx) (err error)
	DeletePlacement(c *fiber.Ctx) (err error)

	CreatePosition(c *fiber.Ctx) (err error)
	UpdatePosition(c *fiber.Ctx) (err error)
	FindAllPositions(c *fiber.Ctx) (err error)
	DeletePosition(c *fiber.Ctx) (err error)
}

type DepartmentControllerImpl struct {
	service Services.DepartmentServiceHandler
}

func DepartmentControllerControllerProvider(service Services.DepartmentServiceHandler) *DepartmentControllerImpl {
	return &DepartmentControllerImpl{
		service: service,
	}
}

func (h *DepartmentControllerImpl) CreateDepartment(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateDepartmentRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateDepartment(request); serviceErr != nil {
		fmt.Println("error at create department service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create department succeed", nil))
}

func (h *DepartmentControllerImpl) UpdateDepartment(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateDepartmentRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateDepartment(request); serviceErr != nil {
		fmt.Println("error at Update department service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update department succeed", nil))
}

func (h *DepartmentControllerImpl) FindAllDepartments(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.DepartmentResponse
	)

	if response, serviceErr = h.service.FindAllDepartments(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data departments", response))
}
func (h *DepartmentControllerImpl) FindDepartmentsByPenempatanID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.DepartmentInfoResponse
	)

	penempatanID, err := strconv.ParseInt(c.Params("penempatanID"), 10, 64)
	lokasiType := c.Params("lokasiType")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Wilayah ID format", err))
	}

	if response, serviceErr = h.service.FindDepartmentByWilayahID(penempatanID, lokasiType); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by Wilayah ID", response))
}

func (h *DepartmentControllerImpl) DeleteDepartment(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteDepartment(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted department with ID %d", id), nil))
}

func (h *DepartmentControllerImpl) CreatePlacement(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreatePlacementRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreatePlacement(request); serviceErr != nil {
		fmt.Println("error at create placement service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create placement succeed", nil))
}

func (h *DepartmentControllerImpl) UpdatePlacement(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdatePlacementRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdatePlacement(request); serviceErr != nil {
		fmt.Println("error at Update placement service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update placement succeed", nil))
}

func (h *DepartmentControllerImpl) FindAllPlacements(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.PlacementResponse
	)

	if response, serviceErr = h.service.FindAllPlacements(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data placements", response))
}

func (h *DepartmentControllerImpl) DeletePlacement(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeletePlacement(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted placement with ID %d", id), nil))
}

func (h *DepartmentControllerImpl) CreatePosition(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreatePositionRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreatePosition(request); serviceErr != nil {
		fmt.Println("error at create position service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create position succeed", nil))
}

func (h *DepartmentControllerImpl) UpdatePosition(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdatePositionRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdatePosition(request); serviceErr != nil {
		fmt.Println("error at Update position service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update position succeed", nil))
}

func (h *DepartmentControllerImpl) FindAllPositions(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.PositionResponse
	)

	if response, serviceErr = h.service.FindAllPositions(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data positions", response))
}

func (h *DepartmentControllerImpl) DeletePosition(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeletePosition(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted position with ID %d", id), nil))
}

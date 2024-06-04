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

type AuthControllerHandler interface {
	Login(c *fiber.Ctx) (err error)
	Register(c *fiber.Ctx) (err error)
	GetRole(c *fiber.Ctx) (err error)
}
type AuthControllerImpl struct {
	service Services.AuthServiceHandler
}

func AuthControllerControllerProvider(service Services.AuthServiceHandler) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: service,
	}
}
func (h *AuthControllerImpl) Login(c *fiber.Ctx) (err error) {
	var (
		request    Domain.LoginRequest
		serviceErr *Web.ServiceErrorDto
		response   Domain.LoginResponse
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if response, serviceErr = h.service.Login(request); serviceErr != nil {
		fmt.Println("error at login service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("login succeed", response))
}

func (h *AuthControllerImpl) Register(c *fiber.Ctx) (err error) {
	var (
		request    Domain.RegisterRequest
		response   Domain.RegisterResponse
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if response, serviceErr = h.service.Register(request); serviceErr != nil {
		fmt.Println("error at register service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusCreated).JSON(Web.SuccessResponse("registration succeed", response))
}

func (h *AuthControllerImpl) FindAllUser(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.UserResponse
	)

	if response, serviceErr = h.service.FindAllUser(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota", response))
}

func (h *AuthControllerImpl) FindUserByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
		response   Domain.User
	)

	idParam := c.Params("userID")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.GetUserByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data user", response))
}

func (h *AuthControllerImpl) GetRole(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.User
	)

	userID := c.Locals("userID")
	convertUserID, ok := userID.(int64)

	if !ok {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.GetUserByID(convertUserID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data user", response))

}

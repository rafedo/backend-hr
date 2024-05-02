package Controller

import (
	"fmt"
	"muhammadiyah/Constant"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"muhammadiyah/Services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PWMControllerHandler interface {
	CreateWilayah(c *fiber.Ctx) (err error)
	UpdateWilayah(c *fiber.Ctx) (err error)
	DeleteWilayah(c *fiber.Ctx) (err error)
	FindAllWilayah(c *fiber.Ctx) (err error)

	CreateDaerah(c *fiber.Ctx) (err error)
	UpdateDaerah(c *fiber.Ctx) (err error)
	DeleteDaerah(c *fiber.Ctx) (err error)
	FindAllDaerah(c *fiber.Ctx) (err error)

	CreateCabang(c *fiber.Ctx) (err error)
	UpdateCabang(c *fiber.Ctx) (err error)
	DeleteCabang(c *fiber.Ctx) (err error)
	FindAllCabang(c *fiber.Ctx) (err error)

	CreateRanting(c *fiber.Ctx) (err error)
	UpdateRanting(c *fiber.Ctx) (err error)
	FindAllRanting(c *fiber.Ctx) (err error)
	DeleteRanting(c *fiber.Ctx) (err error)

	FindAllMemberByWilayahID(c *fiber.Ctx) (err error)
	FindAllMemberByCabangID(c *fiber.Ctx) (err error)
	FindAllMemberByDaerahID(c *fiber.Ctx) (err error)
	FindAllMemberByRantingID(c *fiber.Ctx) (err error)

	CreateMember(c *fiber.Ctx) (err error)
	UpdateMember(c *fiber.Ctx) (err error)
	DeleteMember(c *fiber.Ctx) (err error)
	FindAllMember(c *fiber.Ctx) (err error)
	FindMemberByID(c *fiber.Ctx) (err error)

	FindWilayahByID(c *fiber.Ctx) (err error)
	FindDaerahByID(c *fiber.Ctx) (err error)
	FindDaerahByWilayahID(c *fiber.Ctx) (err error)
	FindCabangByID(c *fiber.Ctx) (err error)
	FindCabangByDaerahID(c *fiber.Ctx) (err error)
	FindRantingByID(c *fiber.Ctx) (err error)
	FindRantingByCabangID(c *fiber.Ctx) (err error)
	CountMember(c *fiber.Ctx) (err error)
}

type PWMControllerImpl struct {
	service Services.PWMServiceHandler
}

func PWMControllerControllerProvider(service Services.PWMServiceHandler) *PWMControllerImpl {
	return &PWMControllerImpl{
		service: service,
	}
}

func (h *PWMControllerImpl) CreateWilayah(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateWilayahRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateWilayah(request); serviceErr != nil {
		fmt.Println("error at create wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create wilayah succeed", nil))
}
func (h *PWMControllerImpl) UpdateWilayah(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateWilayahRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateWilayah(request); serviceErr != nil {
		fmt.Println("error at Update wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update wilayah succeed", nil))
}
func (h *PWMControllerImpl) DeleteWilayah(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteWilayah(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted wilayah with ID %d", id), nil))
}
func (h *PWMControllerImpl) FindAllWilayah(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.WilayahResponse
	)

	if response, serviceErr = h.service.FindAllWilayah(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Wilayah", response))
}

func (h *PWMControllerImpl) CreateDaerah(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateDaerahRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateDaerah(request); serviceErr != nil {
		fmt.Println("error at create wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create daerah succeed", nil))
}
func (h *PWMControllerImpl) UpdateDaerah(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateDaerahRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateDaerah(request); serviceErr != nil {
		fmt.Println("error at Update wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update daerah succeed", nil))
}
func (h *PWMControllerImpl) DeleteDaerah(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteDaerah(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted daerah with ID %d", id), nil))
}
func (h *PWMControllerImpl) FindAllDaerah(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.DaerahResponse
	)

	if response, serviceErr = h.service.FindAllDaerah(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data daerah", response))
}

func (h *PWMControllerImpl) CreateCabang(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateCabangRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateCabang(request); serviceErr != nil {
		fmt.Println("error at create wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create cabang succeed", nil))
}
func (h *PWMControllerImpl) UpdateCabang(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateCabangRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateCabang(request); serviceErr != nil {
		fmt.Println("error at Update wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update cabang succeed", nil))
}
func (h *PWMControllerImpl) DeleteCabang(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteCabang(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted cabang with ID %d", id), nil))
}
func (h *PWMControllerImpl) FindAllCabang(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.CabangResponse
	)

	if response, serviceErr = h.service.FindAllCabang(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data cabang", response))
}

func (h *PWMControllerImpl) CreateRanting(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateRantingRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateRanting(request); serviceErr != nil {
		fmt.Println("error at create wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create ranting succeed", nil))
}
func (h *PWMControllerImpl) UpdateRanting(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateRantingRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateRanting(request); serviceErr != nil {
		fmt.Println("error at Update wilayah service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Update ranting succeed", nil))
}
func (h *PWMControllerImpl) FindAllRanting(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.RantingResponse
	)

	if response, serviceErr = h.service.FindAllRanting(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data ranting", response))
}
func (h *PWMControllerImpl) DeleteRanting(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteRanting(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted ranting with ID %d", id), nil))
}

func (h *PWMControllerImpl) FindAllMemberByWilayahID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.AnggotaResponse
	)

	wilayahID, err := strconv.ParseInt(c.Params("wilayahID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Wilayah ID format", err))
	}

	if response, serviceErr = h.service.FindMemberByWilayahID(wilayahID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by Wilayah ID", response))
}
func (h *PWMControllerImpl) FindAllMemberByCabangID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.AnggotaResponse
	)

	cabangID, err := strconv.ParseInt(c.Params("cabangID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Cabang ID format", err))
	}

	if response, serviceErr = h.service.FindMemberByCabangID(cabangID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by Cabang ID", response))
}
func (h *PWMControllerImpl) FindAllMemberByDaerahID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.AnggotaResponse
	)

	daerahID, err := strconv.ParseInt(c.Params("daerahID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Daerah ID format", err))
	}

	if response, serviceErr = h.service.FindMemberByDaerahID(daerahID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by Daerah ID", response))
}
func (h *PWMControllerImpl) FindAllMemberByRantingID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.AnggotaResponse
	)

	rantingID, err := strconv.ParseInt(c.Params("rantingID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Ranting ID format", err))
	}

	if response, serviceErr = h.service.FindMemberByRantingID(rantingID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by Ranting ID", response))
}

func (h *PWMControllerImpl) CreateMember(c *fiber.Ctx) (err error) {
	var (
		request    []Domain.CreateAnggotaRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.CreateMember(request); serviceErr != nil {
		fmt.Println("error at create member service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create member succeed", nil))
}
func (h *PWMControllerImpl) UpdateMember(c *fiber.Ctx) (err error) {
	var (
		request    Domain.UpdateAnggotaRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.UpdateMember(request); serviceErr != nil {
		fmt.Println("error at update member service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("update member succeed", nil))
}
func (h *PWMControllerImpl) DeleteMember(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		id         int64
	)

	idParam := c.Params("id")
	id, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if serviceErr = h.service.DeleteMember(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse(fmt.Sprintf("Deleted member with ID %d", id), nil))
}
func (h *PWMControllerImpl) FindAllMember(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.AnggotaResponse
	)

	if response, serviceErr = h.service.FindAllMember(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota", response))
}

func (h *PWMControllerImpl) FindMemberByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.AnggotaResponse
	)

	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.FindMemberByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Anggota by ID", response))
}
func (h *PWMControllerImpl) FindWilayahByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.WilayahResponse
	)

	idParam := c.Params("wilayahID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.FindWilayahByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Wilayah by ID", response))
}
func (h *PWMControllerImpl) FindDaerahByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.DaerahResponse
	)

	idParam := c.Params("daerahID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.FindDaerahByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Daerah by ID", response))
}
func (h *PWMControllerImpl) FindDaerahByWilayahID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.DaerahResponse
	)

	wilayahID, err := strconv.ParseInt(c.Params("wilayahID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Wilayah ID format", err))
	}

	if response, serviceErr = h.service.FindDaerahByWilayahID(wilayahID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Daerah by Wilayah ID", response))
}
func (h *PWMControllerImpl) FindCabangByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.CabangResponse
	)

	idParam := c.Params("cabangID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.FindCabangByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Cabang by ID", response))
}
func (h *PWMControllerImpl) FindCabangByDaerahID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.CabangResponse
	)

	daerahID, err := strconv.ParseInt(c.Params("daerahID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Daerah ID format", err))
	}

	if response, serviceErr = h.service.FindCabangByDaerahID(daerahID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Cabang by Daerah ID", response))
}
func (h *PWMControllerImpl) FindRantingByID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.RantingResponse
	)

	idParam := c.Params("rantingID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid ID format", err))
	}

	if response, serviceErr = h.service.FindRantingByID(id); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Ranting by ID", response))
}
func (h *PWMControllerImpl) FindRantingByCabangID(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Domain.RantingResponse
	)

	cabangID, err := strconv.ParseInt(c.Params("cabangID"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Cabang ID format", err))
	}

	if response, serviceErr = h.service.FindRantingByCabangID(cabangID); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Data Ranting by Cabang ID", response))
}
func (h *PWMControllerImpl) CountMember(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Domain.CountResponse
	)

	if response, serviceErr = h.service.CountMembers(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Jumlah Anggota", response))
}

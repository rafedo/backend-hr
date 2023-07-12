package Services

import (
	"crud-barang/Constant"
	"crud-barang/Model/Database"
	"crud-barang/Model/web"
	"crud-barang/Repository"
	"net/http"
)

type (
	BarangServiceHandler interface {
		Create(request []web.BarangCreateRequest) (id int64, serviceErr *web.ServiceErrorDto)
		Update(request web.BarangUpdateRequest) (id int64, serviceErr *web.ServiceErrorDto)
		Delete(barangId int64) (serviceErr *web.ServiceErrorDto)
		FindById(barangId int64) (barang web.BarangResponse, serviceErr *web.ServiceErrorDto)
		FindAll() (barangs []web.BarangResponse, serviceErr *web.ServiceErrorDto)
	}

	BarangServiceImpl struct {
		repo Repository.BarangRepositoryHandler
	}
)

func BarangServiceControllerProvider(repo Repository.BarangRepositoryHandler) *BarangServiceImpl {
	return &BarangServiceImpl{
		repo: repo,
	}
}

func (h *BarangServiceImpl) Create(requests []web.BarangCreateRequest) (id int64, serviceErr *web.ServiceErrorDto) {
	for _, req := range requests {
		if id, err := h.repo.Save(&Database.Barang{
			Name:     req.Name,
			Category: req.Category,
			Price:    req.Price,
		}); err != nil {
			return id, web.NewCustomServiceError("Barang not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *BarangServiceImpl) Update(request web.BarangUpdateRequest) (id int64, serviceErr *web.ServiceErrorDto) {
	_, err := h.repo.FindById(request.Id)
	if err != nil {
		return 0, web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	if id, err := h.repo.Update(&Database.Barang{
		Name:     request.Name,
		Category: request.Category,
		Price:    request.Price,
	}); err != nil {
		return id, web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *BarangServiceImpl) Delete(barangID int64) (serviceErr *web.ServiceErrorDto) {
	_, err := h.repo.FindById(barangID)
	if err != nil {
		return web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	if err := h.repo.Delete(barangID); err != nil {
		return web.NewInternalServiceError(err)
	}

	return nil
}
func (h *BarangServiceImpl) FindById(barangID int64) (barang web.BarangResponse, serviceErr *web.ServiceErrorDto) {
	bar, err := h.repo.FindById(barangID)
	if err != nil {
		return web.BarangResponse{}, web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	return Constant.ToBarangResponse(bar), nil
}

func (h *BarangServiceImpl) FindAll() (barangs []web.BarangResponse, serviceErr *web.ServiceErrorDto) {
	barang, err := h.repo.FindAll()
	if err != nil {
		return []web.BarangResponse{}, web.NewInternalServiceError(err)
	}

	return Constant.ToBarangResponses(barang), nil
}

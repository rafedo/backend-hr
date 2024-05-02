package Services

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"muhammadiyah/Repository"
	"net/http"
)

type PengurusServiceHandler interface {
	CreatePengurus(request []Domain.CreatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto)
	UpdatePengurus(request Domain.UpdatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto)
	DeletePengurus(id int64) *Web.ServiceErrorDto
	GetPengurusByID(id int64) (response Domain.PengurusResponse, serviceErr *Web.ServiceErrorDto)
	GetAllPengurus() ([]Domain.PengurusResponse, *Web.ServiceErrorDto)
}

type PengurusServiceImpl struct {
	repo Repository.PengurusRepositoryHandler
}

func PengurusServiceControllerProvider(repo Repository.PengurusRepositoryHandler) *PengurusServiceImpl {
	return &PengurusServiceImpl{
		repo: repo,
	}
}

func (h *PengurusServiceImpl) CreatePengurus(request []Domain.CreatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range request {
		if id, err := h.repo.CreatePengurus(&Database.Penguru{
			AnggotaID: req.AnggotaID,
			JabatanID: req.JabatanID,
		}); err != nil {
			return id, Web.NewCustomServiceError("Pengurus not created", err, http.StatusInternalServerError)
		}

	}

	return id, nil
}
func (h *PengurusServiceImpl) UpdatePengurus(request Domain.UpdatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto) {

	if id, err := h.repo.UpdatePengurus(&Database.Penguru{
		AnggotaID: request.AnggotaID,
		JabatanID: request.JabatanID}); err != nil {
		return id, Web.NewCustomServiceError("Pengurus not updated", err, http.StatusInternalServerError)
	}
	return id, nil
}
func (service *PengurusServiceImpl) DeletePengurus(id int64) *Web.ServiceErrorDto {
	err := service.repo.DeletePengurus(id)
	if err != nil {
		return Web.NewCustomServiceError("Pengurus not deleted", err, http.StatusInternalServerError)
	}
	return nil
}
func (h *PengurusServiceImpl) GetPengurusByID(id int64) (response Domain.PengurusResponse, serviceErr *Web.ServiceErrorDto) {
	pengurus, err := h.repo.FindPengurusByID(id)
	if err != nil {
		return Domain.PengurusResponse{}, Web.NewCustomServiceError("Pengurus not found ", err, http.StatusInternalServerError)
	}
	return Web.ToPengurusResponse(pengurus), nil
}
func (h *PengurusServiceImpl) GetAllPengurus() ([]Domain.PengurusResponse, *Web.ServiceErrorDto) {
	pengurusList, err := h.repo.FindAllPengurus()
	if err != nil {
		return nil, Web.NewCustomServiceError("Pengurus not found ", err, http.StatusInternalServerError)
	}

	// Menggunakan fungsi konversi ToPengurusResponses

	return Web.ToPengurusResponses(pengurusList), nil
}

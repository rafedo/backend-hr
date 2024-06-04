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
	GetPengurusByID(id int64) (response Domain.PengurusInfoResponse, serviceErr *Web.ServiceErrorDto)
	GetAllPengurus() ([]Domain.PengurusResponse, *Web.ServiceErrorDto)

	FindPengurusByDepartementID(departementID int64) (pengurus []Domain.PengurusInfoResponse, serviceErr *Web.ServiceErrorDto)
}

type PengurusServiceImpl struct {
	pengurusRepo Repository.PengurusRepositoryHandler
}

func PengurusServiceControllerProvider(pengurusRepo Repository.PengurusRepositoryHandler) *PengurusServiceImpl {
	return &PengurusServiceImpl{
		pengurusRepo: pengurusRepo,
	}

}

func (h *PengurusServiceImpl) CreatePengurus(request []Domain.CreatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range request {
		if id, err := h.pengurusRepo.CreatePengurus(&Database.Penguru{
			AnggotaID: req.AnggotaID,
			JabatanID: req.JabatanID,
		}); err != nil {
			return id, Web.NewCustomServiceError("Pengurus not created", err, http.StatusInternalServerError)
		}

	}

	return id, nil
}
func (h *PengurusServiceImpl) UpdatePengurus(request Domain.UpdatePengurusRequest) (id int64, serviceErr *Web.ServiceErrorDto) {

	if id, err := h.pengurusRepo.UpdatePengurus(&Database.Penguru{
		AnggotaID: request.AnggotaID,
		JabatanID: request.JabatanID}); err != nil {
		return id, Web.NewCustomServiceError("Pengurus not updated", err, http.StatusInternalServerError)
	}
	return id, nil
}
func (h *PengurusServiceImpl) DeletePengurus(id int64) *Web.ServiceErrorDto {
	err := h.pengurusRepo.DeletePengurus(id)
	if err != nil {
		return Web.NewCustomServiceError("Pengurus not deleted", err, http.StatusInternalServerError)
	}
	return nil
}
func (h *PengurusServiceImpl) GetPengurusByID(id int64) (response Domain.PengurusInfoResponse, serviceErr *Web.ServiceErrorDto) {
	pengurus, err := h.pengurusRepo.FindPengurusByID(id)
	if err != nil {
		return Domain.PengurusInfoResponse{}, Web.NewCustomServiceError("Pengurus not found ", err, http.StatusInternalServerError)
	}
	return pengurus, nil
}
func (h *PengurusServiceImpl) GetAllPengurus() ([]Domain.PengurusResponse, *Web.ServiceErrorDto) {
	pengurusList, err := h.pengurusRepo.FindAllPengurus()
	if err != nil {
		return nil, Web.NewCustomServiceError("Pengurus not found ", err, http.StatusInternalServerError)
	}
	return Web.ToPengurusResponses(pengurusList), nil
}

func (h *PengurusServiceImpl) FindPengurusByDepartementID(departementID int64) (pengurus []Domain.PengurusInfoResponse, serviceErr *Web.ServiceErrorDto) {
	pengurus, err := h.pengurusRepo.FindPengurusInfoByDepartementID(departementID)
	if err != nil {
		return []Domain.PengurusInfoResponse{}, Web.NewCustomServiceError("Failed to find pengurus by wilayah ID", err, http.StatusInternalServerError)
	}

	return pengurus, nil
}

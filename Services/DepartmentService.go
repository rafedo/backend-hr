package Services

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"muhammadiyah/Repository"
	"net/http"
)

type (
	DepartmentServiceHandler interface {
		CreateDepartment(request []Domain.CreateDepartmentRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateDepartment(request Domain.UpdateDepartmentRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeleteDepartment(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllDepartments() (departments []Domain.DepartmentResponse, serviceErr *Web.ServiceErrorDto)

		CreatePlacement(request []Domain.CreatePlacementRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdatePlacement(request Domain.UpdatePlacementRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeletePlacement(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllPlacements() (placements []Domain.PlacementResponse, serviceErr *Web.ServiceErrorDto)

		CreatePosition(request []Domain.CreatePositionRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdatePosition(request Domain.UpdatePositionRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeletePosition(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllPositions() (positions []Domain.PositionResponse, serviceErr *Web.ServiceErrorDto)
	}

	DepartmentServiceImpl struct {
		repo Repository.DepartmentRepositoryHandler
	}
)

func DepartmentServiceControllerProvider(repo Repository.DepartmentRepositoryHandler) *DepartmentServiceImpl {
	return &DepartmentServiceImpl{
		repo: repo,
	}
}

func (h *DepartmentServiceImpl) CreateDepartment(request []Domain.CreateDepartmentRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range request {
		if id, err := h.repo.CreateDepartment(&Database.Departeman{
			Nama:         req.Nama,
			PenempatanID: req.PenempatanID,
		}); err != nil {
			return id, Web.NewCustomServiceError("Department not created", err, http.StatusInternalServerError)
		}

	}

	return id, nil
}

func (h *DepartmentServiceImpl) UpdateDepartment(request Domain.UpdateDepartmentRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	department := &Database.Departeman{
		ID:           request.ID,
		Nama:         request.Nama,
		PenempatanID: request.PenempatanID,
	}
	if id, err := h.repo.UpdateDepartment(department); err != nil {
		return id, Web.NewCustomServiceError("Department not updated", err, http.StatusInternalServerError)
	}
	return id, nil
}

func (h *DepartmentServiceImpl) DeleteDepartment(id int64) (serviceErr *Web.ServiceErrorDto) {
	if err := h.repo.DeleteDepartment(id); err != nil {
		return Web.NewCustomServiceError("Department not deleted", err, http.StatusInternalServerError)
	}
	return nil
}

func (h *DepartmentServiceImpl) FindAllDepartments() (departments []Domain.DepartmentResponse, serviceErr *Web.ServiceErrorDto) {
	departmentList, err := h.repo.FindAllDepartments()
	if err != nil {
		return []Domain.DepartmentResponse{}, Web.NewCustomServiceError("Error fetching departments", err, http.StatusInternalServerError)
	}
	return Web.ToDepartmentResponses(departmentList), nil
}

func (h *DepartmentServiceImpl) CreatePlacement(request []Domain.CreatePlacementRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range request {
		if id, err := h.repo.CreatePlacement(&Database.Penempatan{
			LokasiID:   req.LokasiID,
			LokasiType: req.LokasiType,
			Jenis:      req.Jenis,
		}); err != nil {
			return id, Web.NewCustomServiceError("Placement not created", err, http.StatusInternalServerError)
		}

	}

	return id, nil
}

func (h *DepartmentServiceImpl) UpdatePlacement(request Domain.UpdatePlacementRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	placement := &Database.Penempatan{
		LokasiID:   request.LokasiID,
		LokasiType: request.LokasiType,
		Jenis:      request.Jenis,
	}
	if id, err := h.repo.UpdatePlacement(placement); err != nil {
		return id, Web.NewCustomServiceError("Placement not updated", err, http.StatusInternalServerError)
	}
	return id, nil
}

func (h *DepartmentServiceImpl) DeletePlacement(id int64) (serviceErr *Web.ServiceErrorDto) {
	if err := h.repo.DeletePlacement(id); err != nil {
		return Web.NewCustomServiceError("Placement not deleted", err, http.StatusInternalServerError)
	}
	return nil
}

func (h *DepartmentServiceImpl) FindAllPlacements() (placementResponse []Domain.PlacementResponse, serviceErr *Web.ServiceErrorDto) {
	placementList, err := h.repo.FindAllPlacements()
	if err != nil {
		return []Domain.PlacementResponse{}, Web.NewCustomServiceError("Error fetching placements", err, http.StatusInternalServerError)
	}
	for _, res := range placementList {
		placementResponse = append(placementResponse, Domain.PlacementResponse{
			ID:         res.ID,
			LokasiID:   res.LokasiID,
			LokasiType: res.LokasiType,
			Jenis:      res.Jenis,
		})

	}
	return placementResponse, nil
}

func (h *DepartmentServiceImpl) CreatePosition(request []Domain.CreatePositionRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range request {
		if id, err := h.repo.CreatePosition(&Database.Jabatan{
			Nama:         req.Nama,
			DepartemenID: req.DepartemenID,
		}); err != nil {
			return id, Web.NewCustomServiceError("Position not created", err, http.StatusInternalServerError)
		}

	}
	return id, nil
}

func (h *DepartmentServiceImpl) UpdatePosition(request Domain.UpdatePositionRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	position := &Database.Jabatan{
		ID:           request.ID,
		Nama:         request.Nama,
		DepartemenID: request.DepartemenID,
	}
	if id, err := h.repo.UpdatePosition(position); err != nil {
		return id, Web.NewCustomServiceError("Position not updated", err, http.StatusInternalServerError)
	}
	return id, nil
}

func (h *DepartmentServiceImpl) DeletePosition(id int64) (serviceErr *Web.ServiceErrorDto) {
	if err := h.repo.DeletePosition(id); err != nil {
		return Web.NewCustomServiceError("Position not deleted", err, http.StatusInternalServerError)
	}
	return nil
}

func (h *DepartmentServiceImpl) FindAllPositions() (positions []Domain.PositionResponse, serviceErr *Web.ServiceErrorDto) {
	positionList, err := h.repo.FindAllPositions()
	if err != nil {
		return []Domain.PositionResponse{}, Web.NewCustomServiceError("Error fetching positions", err, http.StatusInternalServerError)
	}
	return Web.ToPositionResponses(positionList), nil
}

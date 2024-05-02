package Services

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"muhammadiyah/Repository"
	"net/http"
)

type (
	PWMServiceHandler interface {
		CreateWilayah(requests []Domain.CreateWilayahRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateWilayah(request Domain.UpdateWilayahRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeleteWilayah(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllWilayah() (wilayahResponse []Domain.WilayahResponse, serviceErr *Web.ServiceErrorDto)
		FindWilayahByID(id int64) (wilayahResponse Domain.WilayahResponse, serviceErr *Web.ServiceErrorDto)

		CreateDaerah(requests []Domain.CreateDaerahRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateDaerah(request Domain.UpdateDaerahRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		FindAllDaerah() (daerahResponse []Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto)
		DeleteDaerah(id int64) (serviceErr *Web.ServiceErrorDto)
		FindDaerahByID(id int64) (daerahResponse Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto)
		FindDaerahByWilayahID(wilayahID int64) (daerahResponse []Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto)

		CreateCabang(requests []Domain.CreateCabangRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateCabang(request Domain.UpdateCabangRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeleteCabang(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllCabang() (cabangResponse []Domain.CabangResponse, serviceErr *Web.ServiceErrorDto)
		FindCabangByID(id int64) (CabangResponse Domain.CabangResponse, serviceErr *Web.ServiceErrorDto)
		FindCabangByDaerahID(daerahID int64) (cabangResponse []Domain.CabangResponse, serviceErr *Web.ServiceErrorDto)

		CreateRanting(requests []Domain.CreateRantingRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateRanting(request Domain.UpdateRantingRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeleteRanting(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllRanting() (rantingResponse []Domain.RantingResponse, serviceErr *Web.ServiceErrorDto)
		FindRantingByID(id int64) (rantingResponse Domain.RantingResponse, serviceErr *Web.ServiceErrorDto)
		FindRantingByCabangID(cabangID int64) (rantingResponse []Domain.RantingResponse, serviceErr *Web.ServiceErrorDto)

		CreateMember(requests []Domain.CreateAnggotaRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		UpdateMember(request Domain.UpdateAnggotaRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		DeleteMember(id int64) (serviceErr *Web.ServiceErrorDto)
		FindAllMember() (anggotas []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		FindMemberByID(id int64) (memberResponse Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		FindMemberByWilayahID(wilayahID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		FindMemberByCabangID(cabangID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		FindMemberByDaerahID(daerahID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		FindMemberByRantingID(rantingID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto)
		CountMembers() (CountResponse Domain.CountResponse, serviceErr *Web.ServiceErrorDto)
	}

	PWMServiceImpl struct {
		repo Repository.PWMRepositoryHandler
	}
)

func PWMServiceControllerProvider(repo Repository.PWMRepositoryHandler) *PWMServiceImpl {
	return &PWMServiceImpl{
		repo: repo,
	}
}

func (h *PWMServiceImpl) CreateWilayah(requests []Domain.CreateWilayahRequest) (id int64, serviceErr *Web.ServiceErrorDto) {

	for _, req := range requests {
		alamatID, err := h.repo.CreateAddress(&Database.Alamat{
			Alamat:    req.Alamat,
			Kelurahan: req.Kelurahan,
			Kecamatan: req.Kecamatan,
			KabKota:   req.KabKota,
			Propinsi:  req.Propinsi,
			KodePos:   req.KodePos,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}
		ID, err := h.repo.CreateWilayah(&Database.Wilayah{
			NamaWilayah:  req.NamaWilayah,
			AlamatKantor: alamatID})
		if err != nil {
			return ID, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *PWMServiceImpl) UpdateWilayah(request Domain.UpdateWilayahRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	//_, err := h.repo.FindByID(request.ID)
	//if err != nil {
	//	return 0, Web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	//}
	id, err := h.repo.UpdateAddress(&Database.Alamat{
		ID:        request.AlamatID,
		Alamat:    request.Alamat,
		Kelurahan: request.Kelurahan,
		Kecamatan: request.Kecamatan,
		KabKota:   request.KabKota,
		Propinsi:  request.Propinsi,
		KodePos:   request.KodePos,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
	}
	id, err = h.repo.UpdateWilayah(&Database.Wilayah{
		ID:          request.ID,
		NamaWilayah: request.NamaWilayah,
	})
	if err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *PWMServiceImpl) FindAllWilayah() (wilayahResponse []Domain.WilayahResponse, serviceErr *Web.ServiceErrorDto) {
	wilayah, err := h.repo.FindAllWilayah()
	if err != nil {
		return []Domain.WilayahResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range wilayah {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.WilayahResponse{}, Web.NewInternalServiceError(err)
		}
		wilayahResponse = append(wilayahResponse, Domain.WilayahResponse{
			ID:          res.ID,
			NamaWilayah: res.NamaWilayah,
			AlamatID:    alamat.ID,
			Alamat:      alamat.Alamat,
			Kelurahan:   alamat.Kelurahan,
			Kecamatan:   alamat.Kecamatan,
			KabKota:     alamat.KabKota,
			Propinsi:    alamat.Propinsi,
			KodePos:     alamat.KodePos,
		})

	}
	return wilayahResponse, nil
}
func (h *PWMServiceImpl) DeleteWilayah(id int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.repo.DeleteWilayah(id)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
func (h *PWMServiceImpl) FindWilayahByID(id int64) (wilayahResponse Domain.WilayahResponse, serviceErr *Web.ServiceErrorDto) {
	wilayah, err := h.repo.FindWilayahByID(id)
	if err != nil {
		return Domain.WilayahResponse{}, Web.NewInternalServiceError(err)
	}
	alamat, err := h.repo.FindAddressByID(wilayah.AlamatKantor)
	if err != nil {
		return Domain.WilayahResponse{}, Web.NewInternalServiceError(err)
	}
	wilayahResponse = Domain.WilayahResponse{
		ID:          wilayah.ID,
		NamaWilayah: wilayah.NamaWilayah,
		AlamatID:    alamat.ID,
		Alamat:      alamat.Alamat,
		Kelurahan:   alamat.Kelurahan,
		Kecamatan:   alamat.Kecamatan,
		KabKota:     alamat.KabKota,
		Propinsi:    alamat.Propinsi,
		KodePos:     alamat.KodePos,
	}
	return wilayahResponse, nil
}

func (h *PWMServiceImpl) CreateDaerah(requests []Domain.CreateDaerahRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range requests {
		alamatID, err := h.repo.CreateAddress(&Database.Alamat{
			Alamat:    req.Alamat,
			Kelurahan: req.Kelurahan,
			Kecamatan: req.Kecamatan,
			KabKota:   req.KabKota,
			Propinsi:  req.Propinsi,
			KodePos:   req.KodePos,
		})
		if err != nil {
			return alamatID, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}
		id, err := h.repo.CreateDaerah(&Database.Daerah{
			NamaDaerah:   req.NamaDaerah,
			AlamatKantor: alamatID,
			WilayahID:    req.WilayahID,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *PWMServiceImpl) UpdateDaerah(request Domain.UpdateDaerahRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	if id, err := h.repo.UpdateAddress(&Database.Alamat{
		ID:        request.AlamatID,
		Alamat:    request.Alamat,
		Kelurahan: request.Kelurahan,
		Kecamatan: request.Kecamatan,
		KabKota:   request.KabKota,
		Propinsi:  request.Propinsi,
		KodePos:   request.KodePos,
	}); err != nil {
		return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
	}
	if id, err := h.repo.UpdateDaerah(&Database.Daerah{
		ID:         request.ID,
		NamaDaerah: request.NamaDaerah,
		WilayahID:  request.WilayahID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *PWMServiceImpl) DeleteDaerah(id int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.repo.DeleteDaerah(id)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
func (h *PWMServiceImpl) FindAllDaerah() (daerahResponse []Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto) {
	daerah, err := h.repo.FindAllDaerah()
	if err != nil {
		return []Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range daerah {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
		}
		daerahResponse = append(daerahResponse, Domain.DaerahResponse{
			ID:         res.ID,
			NamaDaerah: res.NamaDaerah,
			AlamatID:   alamat.ID,
			Alamat:     alamat.Alamat,
			Kelurahan:  alamat.Kelurahan,
			Kecamatan:  alamat.Kecamatan,
			KabKota:    alamat.KabKota,
			Propinsi:   alamat.Propinsi,
			KodePos:    alamat.KodePos,
		})

	}

	return daerahResponse, nil
}

func (h *PWMServiceImpl) FindDaerahByID(id int64) (daerahResponse Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto) {
	daerah, err := h.repo.FindDaerahByID(id)
	if err != nil {
		return Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
	}
	alamat, err := h.repo.FindAddressByID(daerah.AlamatKantor)
	if err != nil {
		return Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
	}
	daerahResponse = Domain.DaerahResponse{
		ID:         daerah.ID,
		NamaDaerah: daerah.NamaDaerah,
		AlamatID:   alamat.ID,
		Alamat:     alamat.Alamat,
		Kelurahan:  alamat.Kelurahan,
		Kecamatan:  alamat.Kecamatan,
		KabKota:    alamat.KabKota,
		Propinsi:   alamat.Propinsi,
		KodePos:    alamat.KodePos,
	}
	return daerahResponse, nil
}

func (h *PWMServiceImpl) FindDaerahByWilayahID(wilayahID int64) (daerahResponse []Domain.DaerahResponse, serviceErr *Web.ServiceErrorDto) {
	daerah, err := h.repo.FindDaerahByWilayahID(wilayahID)
	if err != nil {
		return []Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range daerah {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.DaerahResponse{}, Web.NewInternalServiceError(err)
		}
		daerahResponse = append(daerahResponse, Domain.DaerahResponse{
			ID:         res.ID,
			NamaDaerah: res.NamaDaerah,
			AlamatID:   alamat.ID,
			Alamat:     alamat.Alamat,
			Kelurahan:  alamat.Kelurahan,
			Kecamatan:  alamat.Kecamatan,
			KabKota:    alamat.KabKota,
			Propinsi:   alamat.Propinsi,
			KodePos:    alamat.KodePos,
		})

	}

	return daerahResponse, nil
}

func (h *PWMServiceImpl) CreateCabang(requests []Domain.CreateCabangRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range requests {
		alamatID, err := h.repo.CreateAddress(&Database.Alamat{
			Alamat:    req.Alamat,
			Kelurahan: req.Kelurahan,
			Kecamatan: req.Kecamatan,
			KabKota:   req.KabKota,
			Propinsi:  req.Propinsi,
			KodePos:   req.KodePos,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}
		id, err := h.repo.CreateCabang(&Database.Cabang{
			NamaCabang:   req.NamaCabang,
			AlamatKantor: alamatID,
			DaerahID:     req.DaerahID,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Daerah not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *PWMServiceImpl) UpdateCabang(request Domain.UpdateCabangRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	if id, err := h.repo.UpdateAddress(&Database.Alamat{
		ID:        request.AlamatID,
		Alamat:    request.Alamat,
		Kelurahan: request.Kelurahan,
		Kecamatan: request.Kecamatan,
		KabKota:   request.KabKota,
		Propinsi:  request.Propinsi,
		KodePos:   request.KodePos,
	}); err != nil {
		return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
	}
	if id, err := h.repo.UpdateCabang(&Database.Cabang{
		ID:           request.ID,
		NamaCabang:   request.NamaCabang,
		AlamatKantor: request.AlamatID,
		DaerahID:     request.DaerahID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *PWMServiceImpl) DeleteCabang(id int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.repo.DeleteCabang(id)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
func (h *PWMServiceImpl) FindAllCabang() (cabangResponse []Domain.CabangResponse, serviceErr *Web.ServiceErrorDto) {
	cabang, err := h.repo.FindAllCabang()
	if err != nil {
		return []Domain.CabangResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range cabang {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.CabangResponse{}, Web.NewInternalServiceError(err)
		}
		cabangResponse = append(cabangResponse, Domain.CabangResponse{
			ID:         res.ID,
			NamaCabang: res.NamaCabang,
			AlamatID:   alamat.ID,
			Alamat:     alamat.Alamat,
			Kelurahan:  alamat.Kelurahan,
			Kecamatan:  alamat.Kecamatan,
			KabKota:    alamat.KabKota,
			Propinsi:   alamat.Propinsi,
			KodePos:    alamat.KodePos,
		})

	}
	return cabangResponse, nil
}

func (h *PWMServiceImpl) FindCabangByID(id int64) (CabangResponse Domain.CabangResponse, serviceErr *Web.ServiceErrorDto) {
	cabang, err := h.repo.FindCabangByID(id)
	if err != nil {
		return Domain.CabangResponse{}, Web.NewInternalServiceError(err)
	}
	alamat, err := h.repo.FindAddressByID(cabang.AlamatKantor)
	if err != nil {
		return Domain.CabangResponse{}, Web.NewInternalServiceError(err)
	}
	CabangResponse = Domain.CabangResponse{
		ID:         cabang.ID,
		NamaCabang: cabang.NamaCabang,
		AlamatID:   alamat.ID,
		Alamat:     alamat.Alamat,
		Kelurahan:  alamat.Kelurahan,
		Kecamatan:  alamat.Kecamatan,
		KabKota:    alamat.KabKota,
		Propinsi:   alamat.Propinsi,
		KodePos:    alamat.KodePos,
	}
	return CabangResponse, nil
}

func (h *PWMServiceImpl) FindCabangByDaerahID(daerahID int64) (cabangResponse []Domain.CabangResponse, serviceErr *Web.ServiceErrorDto) {
	cabang, err := h.repo.FindCabangByDaerahID(daerahID)
	if err != nil {
		return []Domain.CabangResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range cabang {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.CabangResponse{}, Web.NewInternalServiceError(err)
		}
		cabangResponse = append(cabangResponse, Domain.CabangResponse{
			ID:         res.ID,
			NamaCabang: res.NamaCabang,
			AlamatID:   alamat.ID,
			Alamat:     alamat.Alamat,
			Kelurahan:  alamat.Kelurahan,
			Kecamatan:  alamat.Kecamatan,
			KabKota:    alamat.KabKota,
			Propinsi:   alamat.Propinsi,
			KodePos:    alamat.KodePos,
		})
	}
	return cabangResponse, nil
}

func (h *PWMServiceImpl) CreateRanting(requests []Domain.CreateRantingRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range requests {
		alamatID, err := h.repo.CreateAddress(&Database.Alamat{
			Alamat:    req.Alamat,
			Kelurahan: req.Kelurahan,
			Kecamatan: req.Kecamatan,
			KabKota:   req.KabKota,
			Propinsi:  req.Propinsi,
			KodePos:   req.KodePos,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}
		id, err := h.repo.CreateRanting(&Database.Ranting{
			NamaRanting:  req.NamaRanting,
			AlamatKantor: alamatID,
			CabangID:     req.CabangID,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Ranting not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *PWMServiceImpl) UpdateRanting(request Domain.UpdateRantingRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	alamatID, err := h.repo.UpdateAddress(&Database.Alamat{
		ID:        request.AlamatID,
		Alamat:    request.Alamat,
		Kelurahan: request.Kelurahan,
		Kecamatan: request.Kecamatan,
		KabKota:   request.KabKota,
		Propinsi:  request.Propinsi,
		KodePos:   request.KodePos,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
	}
	id, err = h.repo.UpdateRanting(&Database.Ranting{
		ID:           request.ID,
		NamaRanting:  request.NamaRanting,
		AlamatKantor: alamatID,
		CabangID:     request.CabangID,
	})
	if err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *PWMServiceImpl) FindAllRanting() (rantingResponse []Domain.RantingResponse, serviceErr *Web.ServiceErrorDto) {
	ranting, err := h.repo.FindAllRanting()
	if err != nil {
		return []Domain.RantingResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range ranting {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.RantingResponse{}, Web.NewInternalServiceError(err)
		}
		rantingResponse = append(rantingResponse, Domain.RantingResponse{
			ID:          res.ID,
			NamaRanting: res.NamaRanting,
			AlamatID:    alamat.ID,
			Alamat:      alamat.Alamat,
			Kelurahan:   alamat.Kelurahan,
			Kecamatan:   alamat.Kecamatan,
			KabKota:     alamat.KabKota,
			Propinsi:    alamat.Propinsi,
			KodePos:     alamat.KodePos,
		})

	}
	return rantingResponse, nil
}
func (h *PWMServiceImpl) DeleteRanting(id int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.repo.DeleteRanting(id)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
func (h *PWMServiceImpl) FindRantingByID(id int64) (rantingResponse Domain.RantingResponse, serviceErr *Web.ServiceErrorDto) {
	ranting, err := h.repo.FindRantingByID(id)
	if err != nil {
		return Domain.RantingResponse{}, Web.NewInternalServiceError(err)
	}
	alamat, err := h.repo.FindAddressByID(ranting.AlamatKantor)
	if err != nil {
		return Domain.RantingResponse{}, Web.NewInternalServiceError(err)
	}
	rantingResponse = Domain.RantingResponse{
		ID:          ranting.ID,
		NamaRanting: ranting.NamaRanting,
		AlamatID:    alamat.ID,
		Alamat:      alamat.Alamat,
		Kelurahan:   alamat.Kelurahan,
		Kecamatan:   alamat.Kecamatan,
		KabKota:     alamat.KabKota,
		Propinsi:    alamat.Propinsi,
		KodePos:     alamat.KodePos,
	}
	return rantingResponse, nil
}

func (h *PWMServiceImpl) FindRantingByCabangID(cabangID int64) (rantingResponse []Domain.RantingResponse, serviceErr *Web.ServiceErrorDto) {
	ranting, err := h.repo.FindRantingByCabangID(cabangID)
	if err != nil {
		return []Domain.RantingResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range ranting {
		alamat, err := h.repo.FindAddressByID(res.AlamatKantor)
		if err != nil {
			return []Domain.RantingResponse{}, Web.NewInternalServiceError(err)
		}
		rantingResponse = append(rantingResponse, Domain.RantingResponse{
			ID:          res.ID,
			NamaRanting: res.NamaRanting,
			AlamatID:    alamat.ID,
			Alamat:      alamat.Alamat,
			Kelurahan:   alamat.Kelurahan,
			Kecamatan:   alamat.Kecamatan,
			KabKota:     alamat.KabKota,
			Propinsi:    alamat.Propinsi,
			KodePos:     alamat.KodePos,
		})
	}
	return rantingResponse, nil
}

func (h *PWMServiceImpl) CreateMember(requests []Domain.CreateAnggotaRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range requests {
		alamatID, err := h.repo.CreateAddress(&Database.Alamat{
			Alamat:    req.Alamat,
			Kelurahan: req.Kelurahan,
			Kecamatan: req.Kecamatan,
			KabKota:   req.KabKota,
			Propinsi:  req.Propinsi,
			KodePos:   req.KodePos,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
		}
		id, err := h.repo.CreateMember(&Database.Anggotum{
			NomorKta:         req.NomorKTA,
			Ranting:          req.RantingID,
			NamaLengkap:      req.NamaLengkap,
			GelarKesarjanaan: req.GelarKesarjanaan,
			GelarLainDepan:   req.GelarLainDepan,
			TempatLahir:      req.TempatLahir,
			TanggalLahir:     req.TanggalLahir,
			JenisKelamin:     req.JenisKelamin,
			Alamat:           alamatID,
			Status:           req.Status,
		})
		if err != nil {
			return id, Web.NewCustomServiceError("Member not created", err, http.StatusNoContent)
		}
	}

	return id, nil
}

func (h *PWMServiceImpl) UpdateMember(request Domain.UpdateAnggotaRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	alamatID, err := h.repo.UpdateAddress(&Database.Alamat{
		ID:        request.AlamatID,
		Alamat:    request.Alamat,
		Kelurahan: request.Kelurahan,
		Kecamatan: request.Kecamatan,
		KabKota:   request.KabKota,
		Propinsi:  request.Propinsi,
		KodePos:   request.KodePos,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Wilayah not create", err, http.StatusNoContent)
	}
	id, err = h.repo.UpdateMember(&Database.Anggotum{
		ID:               request.ID,
		NomorKta:         request.NomorKTA,
		Ranting:          request.RantingID,
		NamaLengkap:      request.NamaLengkap,
		GelarKesarjanaan: request.GelarKesarjanaan,
		GelarLainDepan:   request.GelarLainDepan,
		TempatLahir:      request.TempatLahir,
		TanggalLahir:     request.TanggalLahir,
		JenisKelamin:     request.JenisKelamin,
		Alamat:           alamatID,
		Status:           request.Status,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Member not updated", err, http.StatusNoContent)
	}

	return id, nil
}

func (h *PWMServiceImpl) DeleteMember(id int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.repo.DeleteMember(id)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *PWMServiceImpl) FindAllMember() (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindAllMembers()
	if err != nil {
		return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		alamat, err := h.repo.FindAddressByID(res.Alamat)
		if err != nil {
			return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
		}
		memberResponse = append(memberResponse, Domain.AnggotaResponse{
			ID:               res.ID,
			NomorKTA:         res.NomorKta,
			RantingID:        res.Ranting,
			NamaLengkap:      res.NamaLengkap,
			GelarKesarjanaan: res.GelarKesarjanaan,
			GelarLainDepan:   res.GelarLainDepan,
			TempatLahir:      res.TempatLahir,
			TanggalLahir:     res.TanggalLahir,
			JenisKelamin:     res.JenisKelamin,
			AlamatID:         alamat.ID,
			Alamat:           alamat.Alamat,
			Kelurahan:        alamat.Kelurahan,
			Kecamatan:        alamat.Kecamatan,
			KabKota:          alamat.KabKota,
			Propinsi:         alamat.Propinsi,
			KodePos:          alamat.KodePos,
			Status:           res.Status,
		})

	}
	return memberResponse, nil
}

func (h *PWMServiceImpl) FindMemberByID(id int64) (memberResponse Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindMemberByID(id)
	if err != nil {
		return Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	alamat, err := h.repo.FindAddressByID(member.Alamat)
	if err != nil {
		return Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	memberResponse = Domain.AnggotaResponse{
		ID:               member.ID,
		NomorKTA:         member.NomorKta,
		RantingID:        member.Ranting,
		NamaLengkap:      member.NamaLengkap,
		GelarKesarjanaan: member.GelarKesarjanaan,
		GelarLainDepan:   member.GelarLainDepan,
		TempatLahir:      member.TempatLahir,
		TanggalLahir:     member.TanggalLahir,
		JenisKelamin:     member.JenisKelamin,
		Status:           member.Status,
		AlamatID:         member.Alamat,
		Alamat:           alamat.Alamat,
		Kelurahan:        alamat.Kelurahan,
		Kecamatan:        alamat.Kecamatan,
		KabKota:          alamat.KabKota,
		Propinsi:         alamat.Propinsi,
		KodePos:          alamat.KodePos,
	}
	return memberResponse, nil
}

func (h *PWMServiceImpl) FindMemberByWilayahID(wilayahID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindMembersByWilayahID(wilayahID)
	if err != nil {
		return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		alamat, err := h.repo.FindAddressByID(res.Alamat)
		if err != nil {
			return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
		}
		memberResponse = append(memberResponse, Domain.AnggotaResponse{
			ID:               res.ID,
			NomorKTA:         res.NomorKta,
			RantingID:        res.Ranting,
			NamaLengkap:      res.NamaLengkap,
			GelarKesarjanaan: res.GelarKesarjanaan,
			GelarLainDepan:   res.GelarLainDepan,
			TempatLahir:      res.TempatLahir,
			TanggalLahir:     res.TanggalLahir,
			JenisKelamin:     res.JenisKelamin,
			Alamat:           alamat.Alamat,
			Kelurahan:        alamat.Kelurahan,
			Kecamatan:        alamat.Kecamatan,
			KabKota:          alamat.KabKota,
			Propinsi:         alamat.Propinsi,
			KodePos:          alamat.KodePos,
		})

	}
	return memberResponse, nil
}

func (h *PWMServiceImpl) FindMemberByCabangID(cabangID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindMembersByCabangID(cabangID)
	if err != nil {
		return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		alamat, err := h.repo.FindAddressByID(res.Alamat)
		if err != nil {
			return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
		}
		memberResponse = append(memberResponse, Domain.AnggotaResponse{
			ID:               res.ID,
			NomorKTA:         res.NomorKta,
			RantingID:        res.Ranting,
			NamaLengkap:      res.NamaLengkap,
			GelarKesarjanaan: res.GelarKesarjanaan,
			GelarLainDepan:   res.GelarLainDepan,
			TempatLahir:      res.TempatLahir,
			TanggalLahir:     res.TanggalLahir,
			JenisKelamin:     res.JenisKelamin,
			Alamat:           alamat.Alamat,
			Kelurahan:        alamat.Kelurahan,
			Kecamatan:        alamat.Kecamatan,
			KabKota:          alamat.KabKota,
			Propinsi:         alamat.Propinsi,
			KodePos:          alamat.KodePos,
		})

	}
	return memberResponse, nil
}

func (h *PWMServiceImpl) FindMemberByDaerahID(daerahID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindMembersByDaerahID(daerahID)
	if err != nil {
		return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		alamat, err := h.repo.FindAddressByID(res.Alamat)
		if err != nil {
			return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
		}
		memberResponse = append(memberResponse, Domain.AnggotaResponse{
			ID:               res.ID,
			NomorKTA:         res.NomorKta,
			RantingID:        res.Ranting,
			NamaLengkap:      res.NamaLengkap,
			GelarKesarjanaan: res.GelarKesarjanaan,
			GelarLainDepan:   res.GelarLainDepan,
			TempatLahir:      res.TempatLahir,
			TanggalLahir:     res.TanggalLahir,
			JenisKelamin:     res.JenisKelamin,
			Alamat:           alamat.Alamat,
			Kelurahan:        alamat.Kelurahan,
			Kecamatan:        alamat.Kecamatan,
			KabKota:          alamat.KabKota,
			Propinsi:         alamat.Propinsi,
			KodePos:          alamat.KodePos,
		})

	}
	return memberResponse, nil
}

func (h *PWMServiceImpl) FindMemberByRantingID(rantingID int64) (memberResponse []Domain.AnggotaResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.repo.FindMembersByRantingID(rantingID)
	if err != nil {
		return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		alamat, err := h.repo.FindAddressByID(res.Alamat)
		if err != nil {
			return []Domain.AnggotaResponse{}, Web.NewInternalServiceError(err)
		}
		memberResponse = append(memberResponse, Domain.AnggotaResponse{
			ID:               res.ID,
			NomorKTA:         res.NomorKta,
			RantingID:        res.Ranting,
			NamaLengkap:      res.NamaLengkap,
			GelarKesarjanaan: res.GelarKesarjanaan,
			GelarLainDepan:   res.GelarLainDepan,
			TempatLahir:      res.TempatLahir,
			TanggalLahir:     res.TanggalLahir,
			JenisKelamin:     res.JenisKelamin,
			Alamat:           alamat.Alamat,
			Kelurahan:        alamat.Kelurahan,
			Kecamatan:        alamat.Kecamatan,
			KabKota:          alamat.KabKota,
			Propinsi:         alamat.Propinsi,
			KodePos:          alamat.KodePos,
		})

	}
	return memberResponse, nil
}
func (h *PWMServiceImpl) CountMembers() (countResponse Domain.CountResponse, serviceErr *Web.ServiceErrorDto) {
	count, err := h.repo.CountMembers()
	if err != nil {
		return Domain.CountResponse{}, Web.NewInternalServiceError(err)
	}
	countResponse = Domain.CountResponse{
		Total: count,
	}
	return countResponse, nil
}

package Domain

type CreateDepartmentRequest struct {
	Nama         string `json:"nama"`
	PenempatanID int64  `json:"penempatan_id"`
}

type UpdateDepartmentRequest struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	PenempatanID int64  `json:"penempatan_id"`
}

type DepartmentResponse struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	PenempatanID int64  `json:"penempatan_id"`
}

// CreatePlacementRequest adalah DTO untuk membuat data penempatan
type CreatePlacementRequest struct {
	LokasiID   int64  `json:"lokasi_id,omitempty"`
	LokasiType string `json:"lokasi_type,omitempty"`
	Jenis      string `json:"jenis,omitempty"`
}

// UpdatePlacementRequest adalah DTO untuk mengupdate data penempatan
type UpdatePlacementRequest struct {
	ID         int64  `json:"id"`
	LokasiID   int64  `json:"lokasi_id,omitempty"`
	LokasiType string `json:"lokasi_type,omitempty"`
	Jenis      string `json:"jenis,omitempty"`
}

// PlacementResponse adalah DTO untuk menampilkan data penempatan
type PlacementResponse struct {
	ID         int64  `json:"id"`
	LokasiID   int64  `json:"lokasi_id"`
	LokasiType string `json:"lokasi_type"`
	Jenis      string `json:"jenis"`
}

type CreatePositionRequest struct {
	Nama         string `json:"nama"`
	DepartemenID int64  `json:"departemen_id"`
}

type UpdatePositionRequest struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	DepartemenID int64  `json:"departemen_id"`
}

type PositionResponse struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	DepartemenID int64  `json:"departemen_id"`
}

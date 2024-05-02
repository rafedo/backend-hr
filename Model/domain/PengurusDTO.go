package Domain

type PengurusResponse struct {
	Id        int64 `json:"id"`
	AnggotaID int64 `json:"anggota_id"`
	JabatanID int64 `json:"jabatan_id"`
}

type CreatePengurusRequest struct {
	AnggotaID int64 `json:"anggota_id"`
	JabatanID int64 `json:"jabatan_id"`
}

type UpdatePengurusRequest struct {
	Id        int64 `json:"id"`
	AnggotaID int64 `json:"anggota_id"`
	JabatanID int64 `json:"jabatan_id"`
}

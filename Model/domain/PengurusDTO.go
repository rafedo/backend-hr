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
type PengurusInfoResponse struct {
	ID               int64  `json:"id"`
	AnggotaID        int64  `json:"anggota_id"`
	NomorKTA         int64  `json:"nomor_kta"`
	NamaLengkap      string `json:"nama_lengkap"`
	GelarKesarjanaan string `json:"gelar_kesarjanaan,omitempty"`
	GelarLainDepan   string `json:"gelar_lain_depan,omitempty"`
	Status           string `json:"status,omitempty"`
	JabatanID        int64  `json:"jabatan_id"`
	NamaJabatan      string `json:"nama_jabatan"`
	DepartemenID     int64  `json:"departemen_id"`
	NamaDepartemen   string `json:"nama_departemen,omitempty"`
	PenempatanID     int64  `json:"penempatan_id"`
	LokasiID         int64  `json:"lokasi_id"`
	LokasiType       string `json:"lokasi_type"`
	Jenis            string `json:"jenis"`
}

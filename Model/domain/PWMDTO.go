package Domain

type WilayahResponse struct {
	ID          int64  `json:"id"`
	NamaWilayah string `json:"nama_wilayah"`
	AlamatID    int64  `json:"alamat_id"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
}
type CreateWilayahRequest struct {
	NamaWilayah string `json:"nama_wilayah"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
}
type UpdateWilayahRequest struct {
	ID          int64  `json:"id"`
	NamaWilayah string `json:"nama_wilayah"`
	AlamatID    int64  `json:"alamat_id"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
}

type DaerahResponse struct {
	ID         int64  `json:"id"`
	NamaDaerah string `json:"nama_daerah"`
	AlamatID   int64  `json:"alamat_id"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	WilayahID  int64  `json:"wilayah_id"`
}
type CreateDaerahRequest struct {
	NamaDaerah string `json:"nama_daerah"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	WilayahID  int64  `json:"wilayah_id"`
}
type UpdateDaerahRequest struct {
	ID         int64  `json:"id"`
	NamaDaerah string `json:"nama_daerah"`
	AlamatID   int64  `json:"alamat_id"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	WilayahID  int64  `json:"wilayah_id"`
}

type CabangResponse struct {
	ID         int64  `json:"id"`
	NamaCabang string `json:"nama_cabang"`
	AlamatID   int64  `json:"alamat_id"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	DaerahID   int64  `json:"daerah_id"`
}
type CreateCabangRequest struct {
	NamaCabang string `json:"nama_cabang"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	DaerahID   int64  `json:"daerah_id"`
}
type UpdateCabangRequest struct {
	ID         int64  `json:"id"`
	NamaCabang string `json:"nama_cabang"`
	AlamatID   int64  `json:"alamat_id"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	KabKota    string `json:"kab_kota"`
	Propinsi   string `json:"propinsi"`
	KodePos    string `json:"kode_pos"`
	DaerahID   int64  `json:"daerah_id"`
}

type RantingResponse struct {
	ID          int64  `json:"id"`
	NamaRanting string `json:"nama_ranting"`
	AlamatID    int64  `json:"alamat_id"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
	CabangID    int64  `json:"cabang_id"`
}
type CreateRantingRequest struct {
	NamaRanting string `json:"nama_ranting"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
	CabangID    int64  `json:"cabang_id"`
}
type UpdateRantingRequest struct {
	ID          int64  `json:"id"`
	NamaRanting string `json:"nama_ranting"`
	AlamatID    int64  `json:"alamat_id"`
	Alamat      string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kab_kota"`
	Propinsi    string `json:"propinsi"`
	KodePos     string `json:"kode_pos"`
	CabangID    int64  `json:"cabang_id"`
}

type AnggotaResponse struct {
	ID               int64  `json:"id"`
	NomorKTA         int64  `json:"nomor_kta"`
	RantingID        int64  `json:"ranting_id"`
	NamaLengkap      string `json:"nama_lengkap"`
	GelarKesarjanaan string `json:"gelar_kesarjanaan"`
	GelarLainDepan   string `json:"gelar_lain_depan"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	JenisKelamin     string `json:"jenis_kelamin"`
	AlamatID         int64  `json:"alamat_id"`
	Alamat           string `json:"alamat"`
	Kelurahan        string `json:"kelurahan"`
	Kecamatan        string `json:"kecamatan"`
	KabKota          string `json:"kab_kota"`
	Propinsi         string `json:"propinsi"`
	KodePos          string `json:"kode_pos"`
	Status           string `json:"status"`
}

type CreateAnggotaRequest struct {
	NomorKTA         int64  `json:"nomor_kta"`
	RantingID        int64  `json:"ranting_id"`
	NamaLengkap      string `json:"nama_lengkap"`
	GelarKesarjanaan string `json:"gelar_kesarjanaan"`
	GelarLainDepan   string `json:"gelar_lain_depan"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	JenisKelamin     string `json:"jenis_kelamin"`
	Alamat           string `json:"alamat"`
	Kelurahan        string `json:"kelurahan"`
	Kecamatan        string `json:"kecamatan"`
	KabKota          string `json:"kab_kota"`
	Propinsi         string `json:"propinsi"`
	KodePos          string `json:"kode_pos"`
	Status           string `json:"status"`
}

type UpdateAnggotaRequest struct {
	ID               int64  `json:"id"`
	NomorKTA         int64  `json:"nomor_kta"`
	RantingID        int64  `json:"ranting_id"`
	NamaLengkap      string `json:"nama_lengkap"`
	GelarKesarjanaan string `json:"gelar_kesarjanaan"`
	GelarLainDepan   string `json:"gelar_lain_depan"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	JenisKelamin     string `json:"jenis_kelamin"`
	AlamatID         int64  `json:"alamat_id"`
	Alamat           string `json:"alamat"`
	Kelurahan        string `json:"kelurahan"`
	Kecamatan        string `json:"kecamatan"`
	KabKota          string `json:"kab_kota"`
	Propinsi         string `json:"propinsi"`
	KodePos          string `json:"kode_pos"`
	Status           string `json:"status"`
}
type CountResponse struct {
	Total int64 `json:"total"`
}

package web

type BarangResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
}

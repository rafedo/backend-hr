package web

type BarangCreateRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Category string `validate:"required" json:"category"`
	Price    int64  `validate:"required" json:"price"`
}

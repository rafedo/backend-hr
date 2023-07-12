package web

type BarangUpdateRequest struct {
	Id       int64  `validate:"required"`
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Category string `validate:"required" json:"category"`
	Price    int64  `validate:"required" json:"price"`
}

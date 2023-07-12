package Constant

import (
	"crud-barang/Model/Database"
	"crud-barang/Model/web"
)

func ToBarangResponse(barang Database.Barang) web.BarangResponse {
	return web.BarangResponse{
		Id:       barang.ID,
		Name:     barang.Name,
		Category: barang.Category,
		Price:    barang.Price,
	}
}

func ToBarangResponses(barangs []Database.Barang) []web.BarangResponse {
	var barangResponses []web.BarangResponse
	for _, barang := range barangs {
		barangResponses = append(barangResponses, ToBarangResponse(barang))
	}
	return barangResponses
}

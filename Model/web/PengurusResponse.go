package Web

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
)

func ToPengurusResponse(pengurus Database.Penguru) Domain.PengurusResponse {
	return Domain.PengurusResponse{
		Id:        pengurus.ID,
		AnggotaID: pengurus.AnggotaID,
		JabatanID: pengurus.JabatanID,
	}
}

func ToPengurusResponses(pengurus []Database.Penguru) []Domain.PengurusResponse {
	var pengurusResponses []Domain.PengurusResponse
	for _, pengurusResponse := range pengurus {
		pengurusResponses = append(pengurusResponses, ToPengurusResponse(pengurusResponse))
	}
	return pengurusResponses
}

package Web

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
)

func ToDepartmentResponse(entity *Database.Departeman) *Domain.DepartmentResponse {
	if entity == nil {
		return nil
	}

	return &Domain.DepartmentResponse{
		ID:           entity.ID,
		Nama:         entity.Nama,
		PenempatanID: entity.PenempatanID,
	}
}

func ToPositionResponse(entity *Database.Jabatan) *Domain.PositionResponse {
	if entity == nil {
		return nil
	}

	return &Domain.PositionResponse{
		ID:           entity.ID,
		Nama:         entity.Nama,
		DepartemenID: entity.DepartemenID,
	}
}

func ToDepartmentResponses(entities []Database.Departeman) []Domain.DepartmentResponse {
	var responses []Domain.DepartmentResponse
	for _, entity := range entities {
		response := ToDepartmentResponse(&entity)
		responses = append(responses, *response)
	}
	return responses
}

func ToPositionResponses(entities []Database.Jabatan) []Domain.PositionResponse {
	var responses []Domain.PositionResponse
	for _, entity := range entities {
		response := ToPositionResponse(&entity)
		responses = append(responses, *response)
	}
	return responses
}

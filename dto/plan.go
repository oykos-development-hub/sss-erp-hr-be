package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type PlanDTO struct {
	Name string `json:"name" validate:"required"`
	Year string `json:"year" validate:"required"`
}

type PlanResponseDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PlanFilter struct {
	Page *int    `json:"page"`
	Size *int    `json:"size"`
	Year *string `json:"year"`
}

func (dto PlanDTO) ToPlan() *data.Plan {
	return &data.Plan{
		Name: dto.Name,
		Year: dto.Year,
	}
}

func ToPlanResponseDTO(data data.Plan) PlanResponseDTO {
	return PlanResponseDTO{
		ID:        data.ID,
		Name:      data.Name,
		Year:      data.Year,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToPlanListResponseDTO(plans []*data.Plan) []PlanResponseDTO {
	dtoList := make([]PlanResponseDTO, len(plans))
	for i, x := range plans {
		dtoList[i] = ToPlanResponseDTO(*x)
	}
	return dtoList
}

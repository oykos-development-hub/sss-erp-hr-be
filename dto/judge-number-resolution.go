package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type JudgeNumberResolutionDTO struct {
	Active       bool   `json:"active"`
	SerialNumber string `json:"serial_number" validate:"required"`
}

type JudgeNumberResolutionResponseDTO struct {
	ID           int       `json:"id"`
	Active       bool      `json:"active"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetJudgeNumberResolutionInputDTO struct {
	Page     *int  `json:"page" validate:"omitempty"`
	PageSize *int  `json:"page_size" validate:"omitempty"`
	Active   *bool `json:"active"`
}

func (dto JudgeNumberResolutionDTO) ToJudgeNumberResolution() *data.JudgeNumberResolution {
	return &data.JudgeNumberResolution{
		Active:       dto.Active,
		SerialNumber: dto.SerialNumber,
	}
}

func ToJudgeNumberResolutionResponseDTO(data data.JudgeNumberResolution) JudgeNumberResolutionResponseDTO {
	return JudgeNumberResolutionResponseDTO{
		ID:           data.ID,
		Active:       data.Active,
		SerialNumber: data.SerialNumber,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToJudgeNumberResolutionListResponseDTO(judge_number_resolutions []*data.JudgeNumberResolution) []JudgeNumberResolutionResponseDTO {
	dtoList := make([]JudgeNumberResolutionResponseDTO, len(judge_number_resolutions))
	for i, x := range judge_number_resolutions {
		dtoList[i] = ToJudgeNumberResolutionResponseDTO(*x)
	}
	return dtoList
}

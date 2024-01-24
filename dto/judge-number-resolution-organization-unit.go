package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetJudgeNumberResolutionOrganizationUnitInputDTO struct {
	Page               *int `json:"page" validate:"omitempty"`
	PageSize           *int `json:"page_size" validate:"omitempty"`
	ResolutionID       *int `json:"resolution_id"`
	OrganizationUnitID *int `json:"organization_unit_id"`
}

type JudgeNumberResolutionOrganizationUnitDTO struct {
	ID                 int `json:"id,omitempty"`
	Resoultion         int `json:"resolution_id"`
	OrganizationUnitID int `json:"organization_unit_id"`
	NumberOfJudges     int `json:"number_of_judges"`
	NumberOfPresidents int `json:"number_of_presidents"`
}

type JudgeNumberResolutionOrganizationUnitResponseDTO struct {
	ID                 int       `json:"id"`
	Resoultion         int       `json:"resolution_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	NumberOfJudges     int       `json:"number_of_judges"`
	NumberOfPresidents int       `json:"number_of_presidents"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (dto JudgeNumberResolutionOrganizationUnitDTO) ToJudgeNumberResolutionOrganizationUnit() *data.JudgeNumberResolutionOrganizationUnit {
	return &data.JudgeNumberResolutionOrganizationUnit{
		Resoultion:         dto.Resoultion,
		OrganizationUnitID: dto.OrganizationUnitID,
		NumberOfJudges:     dto.NumberOfJudges,
		NumberOfPresidents: dto.NumberOfPresidents,
	}
}

func ToJudgeNumberResolutionOrganizationUnitResponseDTO(data data.JudgeNumberResolutionOrganizationUnit) JudgeNumberResolutionOrganizationUnitResponseDTO {
	return JudgeNumberResolutionOrganizationUnitResponseDTO{
		ID:                 data.ID,
		Resoultion:         data.Resoultion,
		OrganizationUnitID: data.OrganizationUnitID,
		NumberOfJudges:     data.NumberOfJudges,
		NumberOfPresidents: data.NumberOfPresidents,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToJudgeNumberResolutionOrganizationUnitListResponseDTO(judge_number_resolution_organization_units []*data.JudgeNumberResolutionOrganizationUnit) []JudgeNumberResolutionOrganizationUnitResponseDTO {
	dtoList := make([]JudgeNumberResolutionOrganizationUnitResponseDTO, len(judge_number_resolution_organization_units))
	for i, x := range judge_number_resolution_organization_units {
		dtoList[i] = ToJudgeNumberResolutionOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

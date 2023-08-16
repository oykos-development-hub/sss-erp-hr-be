package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateJobPositionsInOrganizationUnitsDTO struct {
	SystematizationID        int `json:"systematization_id" validate:"required"`
	ParentOrganizationUnitID int `json:"parent_organization_unit_id" validate:"required"`
	JobPositionID            int `json:"job_position_id" validate:"required"`
	AvailableSlots           int `json:"available_slots" validate:"required"`
}

type GetJobPositionsInOrganizationUnitsDTO struct {
	Page               *int `json:"page" validate:"omitempty"`
	PageSize           *int `json:"page_size" validate:"omitempty"`
	OrganizationUnitID *int `json:"organization_unit_id" validate:"omitempty"`
	JobPositionID      *int `json:"job_position_id" validate:"omitempty"`
}

type JobPositionsInOrganizationUnitsResponseDTO struct {
	ID                       int       `json:"id"`
	SystematizationID        int       `json:"systematization_id"`
	ParentOrganizationUnitID int       `json:"parent_organization_unit_id"`
	JobPositionID            int       `json:"job_position_id"`
	AvailableSlots           int       `json:"available_slots"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (dto CreateJobPositionsInOrganizationUnitsDTO) ToJobPositionsInOrganizationUnits() *data.JobPositionsInOrganizationUnits {
	return &data.JobPositionsInOrganizationUnits{
		SystematizationID:        dto.SystematizationID,
		ParentOrganizationUnitID: dto.ParentOrganizationUnitID,
		JobPositionID:            dto.JobPositionID,
		AvailableSlots:           dto.AvailableSlots,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}
}

func ToJobPositionsInOrganizationUnitsResponseDTO(data data.JobPositionsInOrganizationUnits) JobPositionsInOrganizationUnitsResponseDTO {
	return JobPositionsInOrganizationUnitsResponseDTO{
		ID:                       data.ID,
		SystematizationID:        data.SystematizationID,
		ParentOrganizationUnitID: data.ParentOrganizationUnitID,
		JobPositionID:            data.JobPositionID,
		AvailableSlots:           data.AvailableSlots,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}
}

func ToJobPositionsInOrganizationUnitsListResponseDTO(job_positions_in_organization_units []*data.JobPositionsInOrganizationUnits) []JobPositionsInOrganizationUnitsResponseDTO {
	dtoList := make([]JobPositionsInOrganizationUnitsResponseDTO, len(job_positions_in_organization_units))
	for i, x := range job_positions_in_organization_units {
		dtoList[i] = ToJobPositionsInOrganizationUnitsResponseDTO(*x)
	}
	return dtoList
}

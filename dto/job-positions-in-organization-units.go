package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateJobPositionsInOrganizationUnitsDTO struct {
	SystematizationID        int     `json:"systematization_id" validate:"required"`
	ParentOrganizationUnitID int     `json:"parent_organization_unit_id" validate:"required"`
	JobPositionID            int     `json:"job_position_id" validate:"required"`
	ParentJobPositionID      *int    `json:"parent_job_position_id,omitempty"`
	SystemPermissionID       *int    `json:"system_permission_id,omitempty"`
	Description              string  `json:"description" validate:"required"`
	SerialNumber             string  `json:"serial_number" validate:"required"`
	AvailableSlots           int     `json:"available_slots" validate:"required"`
	Requirements             *string `json:"requirements" validate:"omitempty"`
	Icon                     *string `json:"icon" validate:"omitempty"`
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
	ParentJobPositionID      *int      `json:"parent_job_position_id"`
	SystemPermissionID       *int      `json:"system_permission_id"`
	Description              string    `json:"description"`
	SerialNumber             string    `json:"serial_number"`
	AvailableSlots           int       `json:"available_slots"`
	Requirements             *string   `json:"requirements"`
	Icon                     *string   `json:"icon" validate:"omitempty"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (dto CreateJobPositionsInOrganizationUnitsDTO) ToJobPositionsInOrganizationUnits() *data.JobPositionsInOrganizationUnits {
	return &data.JobPositionsInOrganizationUnits{
		SystematizationID:        dto.SystematizationID,
		ParentOrganizationUnitID: dto.ParentOrganizationUnitID,
		JobPositionID:            dto.JobPositionID,
		ParentJobPositionID:      dto.ParentJobPositionID,
		SystemPermissionID:       dto.SystemPermissionID,
		Description:              dto.Description,
		SerialNumber:             dto.SerialNumber,
		AvailableSlots:           dto.AvailableSlots,
		Requirements:             dto.Requirements,
		Icon:                     dto.Icon,
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
		ParentJobPositionID:      data.ParentJobPositionID,
		SystemPermissionID:       data.SystemPermissionID,
		Description:              data.Description,
		SerialNumber:             data.SerialNumber,
		AvailableSlots:           data.AvailableSlots,
		Requirements:             data.Requirements,
		Icon:                     data.Icon,
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

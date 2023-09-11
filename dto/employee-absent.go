package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetEmployeeAbsentsInputDTO struct {
	Date *time.Time `json:"date"`
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

type EmployeeAbsentDTO struct {
	AbsentTypeID             int       `json:"absent_type_id" validate:"required"`
	UserProfileID            int       `json:"user_profile_id" validate:"required"`
	TargetOrganizationUnitID *int      `json:"target_organization_unit_id" validate:"omitempty"`
	Description              *string   `json:"description" validate:"omitempty"`
	DateOfStart              time.Time `json:"date_of_start" validate:"required,datetime"`
	DateOfEnd                time.Time `json:"date_of_end" validate:"required,datetime"`
	Location                 *string   `json:"location" validate:"omitempty"`
	FileID                   *int      `json:"file_id" validate:"omitempty"`
}

type EmployeeAbsentResponseDTO struct {
	ID                       int       `json:"id"`
	AbsentTypeID             int       `json:"absent_type_id"`
	UserProfileID            int       `json:"user_profile_id"`
	TargetOrganizationUnitID *int      `json:"target_organization_unit_id"`
	Description              *string   `json:"description"`
	DateOfStart              time.Time `json:"date_of_start"`
	DateOfEnd                time.Time `json:"date_of_end"`
	Location                 *string   `json:"location"`
	FileID                   *int      `json:"file_id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (dto EmployeeAbsentDTO) ToEmployeeAbsent() *data.EmployeeAbsent {
	return &data.EmployeeAbsent{
		AbsentTypeID:             dto.AbsentTypeID,
		UserProfileID:            dto.UserProfileID,
		TargetOrganizationUnitID: dto.TargetOrganizationUnitID,
		Description:              dto.Description,
		DateOfStart:              dto.DateOfStart,
		DateOfEnd:                dto.DateOfEnd,
		Location:                 dto.Location,
		FileID:                   dto.FileID,
	}
}

func ToEmployeeAbsentResponseDTO(data data.EmployeeAbsent) EmployeeAbsentResponseDTO {
	return EmployeeAbsentResponseDTO{
		ID:                       data.ID,
		AbsentTypeID:             data.AbsentTypeID,
		UserProfileID:            data.UserProfileID,
		TargetOrganizationUnitID: data.TargetOrganizationUnitID,
		Description:              data.Description,
		DateOfStart:              data.DateOfStart,
		DateOfEnd:                data.DateOfEnd,
		Location:                 data.Location,
		FileID:                   data.FileID,
		CreatedAt:                data.CreatedAt,
		UpdatedAt:                data.UpdatedAt,
	}
}

func ToEmployeeAbsentListResponseDTO(employeeabsents []*data.EmployeeAbsent) []EmployeeAbsentResponseDTO {
	dtoList := make([]EmployeeAbsentResponseDTO, len(employeeabsents))
	for i, x := range employeeabsents {
		dtoList[i] = ToEmployeeAbsentResponseDTO(*x)
	}
	return dtoList
}

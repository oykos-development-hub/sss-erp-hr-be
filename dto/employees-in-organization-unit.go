package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetEmployeesInOrganizationUnitInput struct {
	PositionInOrganizationUnit *int  `json:"position_in_organization_unit_id" validate:"omitempty"`
	Active                     *bool `json:"active" validate:"omitempty"`
}

type EmployeesInOrganizationUnitDTO struct {
	UserAccountId                int  `json:"user_account_id" validate:"required"`
	UserProfileId                int  `json:"user_profile_id" validate:"required"`
	PositionInOrganizationUnitId int  `json:"position_in_organization_unit_id" validate:"required"`
	Active                       bool `json:"active"`
}

type EmployeesInOrganizationUnitResponseDTO struct {
	ID                           int       `json:"id"`
	UserAccountId                int       `json:"user_account_id"`
	UserProfileId                int       `json:"user_profile_id"`
	PositionInOrganizationUnitId int       `json:"position_in_organization_unit_id"`
	Active                       bool      `json:"active"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}

func (dto EmployeesInOrganizationUnitDTO) ToEmployeesInOrganizationUnit() *data.EmployeesInOrganizationUnit {
	return &data.EmployeesInOrganizationUnit{
		UserAccountId:                dto.UserAccountId,
		UserProfileId:                dto.UserProfileId,
		PositionInOrganizationUnitId: dto.PositionInOrganizationUnitId,
		Active:                       dto.Active,
	}
}

func ToEmployeesInOrganizationUnitResponseDTO(data data.EmployeesInOrganizationUnit) EmployeesInOrganizationUnitResponseDTO {
	return EmployeesInOrganizationUnitResponseDTO{
		ID:                           data.ID,
		UserAccountId:                data.UserAccountId,
		UserProfileId:                data.UserProfileId,
		PositionInOrganizationUnitId: data.PositionInOrganizationUnitId,
		Active:                       data.Active,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
	}
}

func ToEmployeesInOrganizationUnitListResponseDTO(employeesinorganizationunits []*data.EmployeesInOrganizationUnit) []EmployeesInOrganizationUnitResponseDTO {
	dtoList := make([]EmployeesInOrganizationUnitResponseDTO, len(employeesinorganizationunits))
	for i, x := range employeesinorganizationunits {
		dtoList[i] = ToEmployeesInOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

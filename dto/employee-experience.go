package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

// EmployeeExperienceDTO struct
type EmployeeExperienceDTO struct {
	UserProfileID             int       `json:"user_profile_id" validate:"required"`
	Relevant                  bool      `json:"relevant"`
	OrganizationUnit          *string   `json:"organization_unit" validate:"omitempty,required_without=organization_unit_id"`
	OrganizationUnitID        *int      `json:"organization_unit_id" validate:"omitempty,required_without=organization_unit"`
	AmountOfExperience        *int      `json:"amount_of_experience"`
	AmountOfInsuredExperience *int      `json:"amount_of_insured_experience"`
	DateOfStart               time.Time `json:"date_of_start" validate:"required"`
	DateOfEnd                 time.Time `json:"date_of_end" validate:"required"`
	FileID                    int       `json:"reference_file_id"`
}

type EmployeeExperienceResponseDTO struct {
	ID                        int        `json:"id"`
	UserProfileID             int        `json:"user_profile_id"`
	Relevant                  bool       `json:"relevant"`
	OrganizationUnit          *string    `json:"organization_unit"`
	OrganizationUnitID        *int       `json:"organization_unit_id"`
	AmountOfExperience        *int       `json:"amount_of_experience"`
	AmountOfInsuredExperience *int       `json:"amount_of_insured_experience"`
	DateOfStart               time.Time  `json:"date_of_start"`
	DateOfEnd                 time.Time  `json:"date_of_end"`
	FileID                    int        `json:"reference_file_id"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
}

func (dto EmployeeExperienceDTO) ToEmployeeExperience() *data.EmployeeExperience {
	return &data.EmployeeExperience{
		UserProfileID:             dto.UserProfileID,
		Relevant:                  dto.Relevant,
		OrganizationUnit:          dto.OrganizationUnit,
		OrganizationUnitID:        dto.OrganizationUnitID,
		AmountOfExperience:        dto.AmountOfExperience,
		AmountOfInsuredExperience: dto.AmountOfInsuredExperience,
		DateOfStart:               dto.DateOfStart,
		DateOfEnd:                 dto.DateOfEnd,
		FileID:                    dto.FileID,
	}
}

func ToEmployeeExperienceResponseDTO(data data.EmployeeExperience) EmployeeExperienceResponseDTO {
	return EmployeeExperienceResponseDTO{
		ID:                        data.ID,
		UserProfileID:             data.UserProfileID,
		Relevant:                  data.Relevant,
		OrganizationUnit:          data.OrganizationUnit,
		OrganizationUnitID:        data.OrganizationUnitID,
		AmountOfExperience:        data.AmountOfExperience,
		AmountOfInsuredExperience: data.AmountOfInsuredExperience,
		DateOfStart:               data.DateOfStart,
		DateOfEnd:                 data.DateOfEnd,
		FileID:                    data.FileID,
		CreatedAt:                 &data.CreatedAt,
		UpdatedAt:                 &data.UpdatedAt,
	}
}

func ToEmployeeExperienceListResponseDTO(employeeexperiences []*data.EmployeeExperience) []EmployeeExperienceResponseDTO {
	dtoList := make([]EmployeeExperienceResponseDTO, len(employeeexperiences))
	for i, x := range employeeexperiences {
		dtoList[i] = ToEmployeeExperienceResponseDTO(*x)
	}
	return dtoList
}

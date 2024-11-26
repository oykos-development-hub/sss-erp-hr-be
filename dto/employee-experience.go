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
	YearsOfExperience         *int      `json:"years_of_experience"`
	YearsOfInsuredExperience  *int      `json:"years_of_insured_experience"`
	MonthsOfExperience        *int      `json:"months_of_experience"`
	MonthsOfInsuredExperience *int      `json:"months_of_insured_experience"`
	DaysOfExperience          *int      `json:"days_of_experience"`
	DaysOfInsuredExperience   *int      `json:"days_of_insured_experience"`
	DateOfStart               time.Time `json:"date_of_start" validate:"required"`
	DateOfEnd                 time.Time `json:"date_of_end" validate:"required"`
	FileIDs                   []int64   `json:"file_ids"`
}

type EmployeeExperienceResponseDTO struct {
	ID                        int        `json:"id"`
	UserProfileID             int        `json:"user_profile_id"`
	Relevant                  bool       `json:"relevant"`
	OrganizationUnit          *string    `json:"organization_unit"`
	OrganizationUnitID        *int       `json:"organization_unit_id"`
	YearsOfExperience         *int       `json:"years_of_experience"`
	YearsOfInsuredExperience  *int       `json:"years_of_insured_experience"`
	MonthsOfExperience        *int       `json:"months_of_experience"`
	MonthsOfInsuredExperience *int       `json:"months_of_insured_experience"`
	DaysOfExperience          *int       `json:"days_of_experience"`
	DaysOfInsuredExperience   *int       `json:"days_of_insured_experience"`
	DateOfStart               time.Time  `json:"date_of_start"`
	DateOfEnd                 time.Time  `json:"date_of_end"`
	FileIDs                   []int64    `json:"file_ids"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
}

func (dto EmployeeExperienceDTO) ToEmployeeExperience() *data.EmployeeExperience {
	return &data.EmployeeExperience{
		UserProfileID:             dto.UserProfileID,
		Relevant:                  dto.Relevant,
		OrganizationUnit:          dto.OrganizationUnit,
		OrganizationUnitID:        dto.OrganizationUnitID,
		YearsOfExperience:         dto.YearsOfExperience,
		YearsOfInsuredExperience:  dto.YearsOfInsuredExperience,
		MonthsOfExperience:        dto.MonthsOfExperience,
		MonthsOfInsuredExperience: dto.MonthsOfInsuredExperience,
		DaysOfExperience:          dto.DaysOfExperience,
		DaysOfInsuredExperience:   dto.DaysOfInsuredExperience,
		DateOfStart:               dto.DateOfStart,
		DateOfEnd:                 dto.DateOfEnd,
		FileIDs:                   dto.FileIDs,
	}
}

func ToEmployeeExperienceResponseDTO(data data.EmployeeExperience) EmployeeExperienceResponseDTO {
	return EmployeeExperienceResponseDTO{
		ID:                        data.ID,
		UserProfileID:             data.UserProfileID,
		Relevant:                  data.Relevant,
		OrganizationUnit:          data.OrganizationUnit,
		OrganizationUnitID:        data.OrganizationUnitID,
		YearsOfExperience:         data.YearsOfExperience,
		YearsOfInsuredExperience:  data.YearsOfInsuredExperience,
		MonthsOfExperience:        data.MonthsOfExperience,
		MonthsOfInsuredExperience: data.MonthsOfInsuredExperience,
		DaysOfExperience:          data.DaysOfExperience,
		DaysOfInsuredExperience:   data.DaysOfInsuredExperience,
		DateOfStart:               data.DateOfStart,
		DateOfEnd:                 data.DateOfEnd,
		FileIDs:                   data.FileIDs,
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

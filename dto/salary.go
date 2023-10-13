package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type SalaryDTO struct {
	UserProfileID       int    `json:"user_profile_id"`
	OrganizationUnitID  int    `json:"organization_unit_id"`
	BenefitedTrack      *bool  `json:"benefited_track"`
	WithoutRaise        *bool  `json:"without_raise"`
	InsuranceBasis      string `json:"insurance_basis"`
	SalaryRank          string `json:"salary_rank"`
	DailyWorkHours      string `json:"daily_work_hours"`
	WeeklyWorkHours     string `json:"weekly_work_hours"`
	EducationRank       string `json:"education_rank"`
	EducationNaming     string `json:"education_naming"`
	ObligationReduction string `json:"obligation_reduction"`
	UserResoultionID    *int   `json:"user_resolution_id,omitempty"`
}

type SalaryResponseDTO struct {
	ID                  int       `json:"id"`
	UserProfileID       int       `json:"user_profile_id"`
	OrganizationUnitID  int       `json:"organization_unit_id"`
	BenefitedTrack      *bool     `json:"benefited_track"`
	WithoutRaise        *bool     `json:"without_raise"`
	InsuranceBasis      string    `json:"insurance_basis"`
	SalaryRank          string    `json:"salary_rank"`
	DailyWorkHours      string    `json:"daily_work_hours"`
	WeeklyWorkHours     string    `json:"weekly_work_hours"`
	EducationRank       string    `json:"education_rank"`
	EducationNaming     string    `json:"education_naming"`
	ObligationReduction string    `json:"obligation_reduction"`
	UserResoultionID    *int      `json:"user_resolution_id,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (dto SalaryDTO) ToSalary() *data.Salary {
	return &data.Salary{
		UserProfileID:       dto.UserProfileID,
		OrganizationUnitID:  dto.OrganizationUnitID,
		BenefitedTrack:      dto.BenefitedTrack,
		WithoutRaise:        dto.WithoutRaise,
		InsuranceBasis:      dto.InsuranceBasis,
		SalaryRank:          dto.SalaryRank,
		DailyWorkHours:      dto.DailyWorkHours,
		WeeklyWorkHours:     dto.WeeklyWorkHours,
		EducationRank:       dto.EducationRank,
		EducationNaming:     dto.EducationNaming,
		ObligationReduction: dto.ObligationReduction,
		UserResoultionID:    dto.UserResoultionID,
	}
}

func ToSalaryResponseDTO(data data.Salary) SalaryResponseDTO {
	return SalaryResponseDTO{
		ID:                  data.ID,
		UserProfileID:       data.UserProfileID,
		OrganizationUnitID:  data.OrganizationUnitID,
		BenefitedTrack:      data.BenefitedTrack,
		WithoutRaise:        data.WithoutRaise,
		InsuranceBasis:      data.InsuranceBasis,
		SalaryRank:          data.SalaryRank,
		DailyWorkHours:      data.DailyWorkHours,
		WeeklyWorkHours:     data.WeeklyWorkHours,
		EducationRank:       data.EducationRank,
		EducationNaming:     data.EducationNaming,
		ObligationReduction: data.ObligationReduction,
		UserResoultionID:    data.UserResoultionID,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
}

func ToSalaryListResponseDTO(salaries []*data.Salary) []SalaryResponseDTO {
	dtoList := make([]SalaryResponseDTO, len(salaries))
	for i, x := range salaries {
		dtoList[i] = ToSalaryResponseDTO(*x)
	}
	return dtoList
}

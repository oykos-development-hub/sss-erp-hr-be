package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type SalaryDTO struct {
	UserProfileID    int    `json:"user_profile_id" validate:"required"`
	BenefitedTrack   *bool  `json:"benefited_track" validate:"required"`
	WithoutRaise     *bool  `json:"without_raise" validate:"required"`
	InsuranceBasis   string `json:"insurance_basis" validate:"required"`
	SalaryRank       string `json:"salary_rank" validate:"required"`
	DailyWorkHours   string `json:"daily_work_hours" validate:"required"`
	WeeklyWorkHours  string `json:"weekly_work_hours" validate:"required"`
	EducationRank    string `json:"education_rank" validate:"required"`
	EducationNaming  string `json:"education_naming" validate:"required"`
	UserResoultionID *int   `json:"user_resolution_id,omitempty" validate:"omitempty"`
}

type SalaryResponseDTO struct {
	ID               int       `json:"id"`
	UserProfileID    int       `json:"user_profile_id"`
	BenefitedTrack   *bool     `json:"benefited_track"`
	WithoutRaise     *bool     `json:"without_raise"`
	InsuranceBasis   string    `json:"insurance_basis"`
	SalaryRank       string    `json:"salary_rank"`
	DailyWorkHours   string    `json:"daily_work_hours"`
	WeeklyWorkHours  string    `json:"weekly_work_hours"`
	EducationRank    string    `json:"education_rank"`
	EducationNaming  string    `json:"education_naming"`
	UserResoultionID *int      `json:"user_resolution_id,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (dto SalaryDTO) ToSalary() *data.Salary {
	return &data.Salary{
		UserProfileID:    dto.UserProfileID,
		BenefitedTrack:   dto.BenefitedTrack,
		WithoutRaise:     dto.WithoutRaise,
		InsuranceBasis:   dto.InsuranceBasis,
		SalaryRank:       dto.SalaryRank,
		DailyWorkHours:   dto.DailyWorkHours,
		WeeklyWorkHours:  dto.WeeklyWorkHours,
		EducationRank:    dto.EducationRank,
		EducationNaming:  dto.EducationNaming,
		UserResoultionID: dto.UserResoultionID,
	}
}

func ToSalaryResponseDTO(data data.Salary) SalaryResponseDTO {
	return SalaryResponseDTO{
		ID:               data.ID,
		UserProfileID:    data.UserProfileID,
		BenefitedTrack:   data.BenefitedTrack,
		WithoutRaise:     data.WithoutRaise,
		InsuranceBasis:   data.InsuranceBasis,
		SalaryRank:       data.SalaryRank,
		DailyWorkHours:   data.DailyWorkHours,
		WeeklyWorkHours:  data.WeeklyWorkHours,
		EducationRank:    data.EducationRank,
		EducationNaming:  data.EducationNaming,
		UserResoultionID: data.UserResoultionID,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func ToSalaryListResponseDTO(salaries []*data.Salary) []SalaryResponseDTO {
	dtoList := make([]SalaryResponseDTO, len(salaries))
	for i, x := range salaries {
		dtoList[i] = ToSalaryResponseDTO(*x)
	}
	return dtoList
}

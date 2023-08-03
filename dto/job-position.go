package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateJobPositionDTO struct {
	Title            string  `json:"title" validate:"required,min=2"`
	Abbreviation     string  `json:"abbreviation" validate:"required,min=2,max=8"`
	SerialNumber     string  `json:"serial_number" validate:"required"`
	Description      *string `json:"description" validate:"omitempty"`
	Requirements     string  `json:"requirements" validate:"required"`
	IsJudge          *bool   `json:"is_judge" validate:"required"`
	IsJudgePresident *bool   `json:"is_judge_president" validate:"required"`
	Color            *string `json:"color" validate:"omitempty"`
	Icon             *string `json:"icon" validate:"omitempty"`
}

type UpdateJobPositionDTO struct {
	Title            *string `json:"title" validate:"min=2"`
	Abbreviation     *string `json:"abbreviation" validate:"min=2,max=4"`
	SerialNumber     *string `json:"serial_number"`
	Description      *string `json:"description" validate:"omitempty"`
	Requirements     *string `json:"requirements"`
	IsJudge          *bool   `json:"is_judge"`
	IsJudgePresident *bool   `json:"is_judge_president"`
	Color            *string `json:"color" validate:"omitempty"`
	Icon             *string `json:"icon" validate:"omitempty"`
}

type JobPositionResponseDTO struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Abbreviation     string    `json:"abbreviation"`
	SerialNumber     string    `json:"serial_number"`
	Description      *string   `json:"description"`
	Requirements     string    `json:"requirements"`
	IsJudge          *bool     `json:"is_judge"`
	IsJudgePresident *bool     `json:"is_judge_president"`
	Color            *string   `json:"color"`
	Icon             *string   `json:"icon"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type GetJobPositionsDTO struct {
	Page     *int    `json:"page" validate:"omitempty"`
	PageSize *int    `json:"page_size" validate:"omitempty"`
	IsJudge  *bool   `json:"is_judge"`
	Search   *string `json:"search"`
}

func (dto CreateJobPositionDTO) ToJobPosition() *data.JobPosition {
	return &data.JobPosition{
		Title:            dto.Title,
		Abbreviation:     dto.Abbreviation,
		SerialNumber:     dto.SerialNumber,
		Description:      dto.Description,
		Requirements:     dto.Requirements,
		IsJudge:          dto.IsJudge,
		IsJudgePresident: dto.IsJudgePresident,
		Color:            dto.Color,
		Icon:             dto.Icon,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (dto UpdateJobPositionDTO) ToJobPosition(data *data.JobPosition) {
	if dto.Title != nil {
		data.Title = *dto.Title
	}
	if dto.Abbreviation != nil {
		data.Abbreviation = *dto.Abbreviation
	}
	if dto.SerialNumber != nil {
		data.SerialNumber = *dto.SerialNumber
	}
	if dto.Requirements != nil {
		data.Requirements = *dto.Requirements
	}
	if dto.IsJudge != nil {
		data.IsJudge = dto.IsJudge
	}
	if dto.IsJudgePresident != nil {
		data.IsJudgePresident = dto.IsJudgePresident
	}
	data.Description = dto.Description
	data.Color = dto.Color
	data.Icon = dto.Icon
	data.UpdatedAt = time.Now()
}

func ToJobPositionResponseDTO(data data.JobPosition) JobPositionResponseDTO {
	return JobPositionResponseDTO{
		ID:               data.ID,
		Title:            data.Title,
		Abbreviation:     data.Abbreviation,
		SerialNumber:     data.SerialNumber,
		Description:      data.Description,
		Requirements:     data.Requirements,
		IsJudge:          data.IsJudge,
		IsJudgePresident: data.IsJudgePresident,
		Color:            data.Color,
		Icon:             data.Icon,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func ToJobPositionListResponseDTO(job_positions []*data.JobPosition) []JobPositionResponseDTO {
	dtoList := make([]JobPositionResponseDTO, len(job_positions))
	for i, x := range job_positions {
		dtoList[i] = ToJobPositionResponseDTO(*x)
	}
	return dtoList
}

package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type TenderTypeDTO struct {
	Title            string  `json:"title" validate:"required"`
	Value            *string `json:"value" validate:"omitempty"`
	Abbreviation     *string `json:"abbreviation" validate:"omitempty"`
	IsJudge          bool    `json:"is_judge"`
	IsJudgePresident bool    `json:"is_judge_president"`
	Description      *string `json:"description" validate:"omitempty"`
	Color            *string `json:"color" validate:"omitempty"`
	Icon             *string `json:"icon" validate:"omitempty"`
}

type TenderTypeResponseDTO struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Abbreviation     *string   `json:"abbreviation"`
	Value            *string   `json:"value"`
	IsJudge          bool      `json:"is_judge"`
	IsJudgePresident bool      `json:"is_judge_president"`
	Description      *string   `json:"description"`
	Color            *string   `json:"color"`
	Icon             *string   `json:"icon"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (dto TenderTypeDTO) ToTenderType() *data.TenderType {
	return &data.TenderType{
		Title:            dto.Title,
		Value:            dto.Value,
		Abbreviation:     dto.Abbreviation,
		IsJudge:          dto.IsJudge,
		IsJudgePresident: dto.IsJudgePresident,
		Description:      dto.Description,
		Color:            dto.Color,
		Icon:             dto.Icon,
	}
}

func ToTenderTypeResponseDTO(data data.TenderType) TenderTypeResponseDTO {
	return TenderTypeResponseDTO{
		ID:               data.ID,
		Title:            data.Title,
		Abbreviation:     data.Abbreviation,
		Value:            data.Value,
		IsJudge:          data.IsJudge,
		IsJudgePresident: data.IsJudgePresident,
		Description:      data.Description,
		Color:            data.Color,
		Icon:             data.Icon,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func ToTenderTypeListResponseDTO(tendertypes []*data.TenderType) []TenderTypeResponseDTO {
	dtoList := make([]TenderTypeResponseDTO, len(tendertypes))
	for i, x := range tendertypes {
		dtoList[i] = ToTenderTypeResponseDTO(*x)
	}
	return dtoList
}

type GetTenderTypeInputDTO struct {
	Search *string `json:"search"`
}

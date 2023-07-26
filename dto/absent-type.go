package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type AbsentTypeDTO struct {
	ParentID          *int    `json:"parent_id" validate:"omitempty"`
	Title             string  `json:"title" validate:"required,min=2"`
	Abbreviation      string  `json:"abbreviation" validate:"required,min=2,max=6"`
	AccountingDaysOff bool    `json:"accounting_days_off" validate:"required"`
	Relocation        bool    `json:"relocation" validate:"required"`
	Description       *string `json:"description" validate:"omitempty,min=2"`
	Color             *string `json:"color" validate:"omitempty,min=2"`
	Icon              *string `json:"icon" validate:"omitempty,min=2"`
}

type AbsentTypeResponseDTO struct {
	ID                int       `json:"id"`
	ParentID          *int      `json:"parent_id"`
	Title             string    `json:"title"`
	Abbreviation      string    `json:"abbreviation"`
	AccountingDaysOff bool      `json:"accounting_days_off"`
	Relocation        bool      `json:"relocation"`
	Description       *string   `json:"description"`
	Color             *string   `json:"color"`
	Icon              *string   `json:"icon"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (dto AbsentTypeDTO) ToAbsentType() *data.AbsentType {
	return &data.AbsentType{
		ParentID:          dto.ParentID,
		Title:             dto.Title,
		Relocation:        dto.Relocation,
		AccountingDaysOff: dto.AccountingDaysOff,
		Abbreviation:      dto.Abbreviation,
		Description:       dto.Description,
		Color:             dto.Color,
		Icon:              dto.Icon,
	}
}

func ToAbsentTypeResponseDTO(data data.AbsentType) AbsentTypeResponseDTO {
	return AbsentTypeResponseDTO{
		ID:                data.ID,
		ParentID:          data.ParentID,
		Title:             data.Title,
		Abbreviation:      data.Abbreviation,
		AccountingDaysOff: data.AccountingDaysOff,
		Relocation:        data.Relocation,
		Description:       data.Description,
		Color:             data.Color,
		Icon:              data.Icon,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func ToAbsentTypeListResponseDTO(absenttypes []*data.AbsentType) []AbsentTypeResponseDTO {
	dtoList := make([]AbsentTypeResponseDTO, len(absenttypes))
	for i, x := range absenttypes {
		dtoList[i] = ToAbsentTypeResponseDTO(*x)
	}
	return dtoList
}

type GetAbesntTypeDTO struct {
	Page *int `json:"page" validate:"omitempty"`
	Size *int `json:"size" validate:"omitempty"`
}

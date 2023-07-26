package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateOrganizationUnitDTO struct {
	ParentID       *int    `json:"parent_id,omitempty"`
	Title          string  `json:"title" validate:"required,min=2"`
	Abbreviation   string  `json:"abbreviation" validate:"required,min=2,max=4"`
	NumberOfJudges *int    `json:"number_of_judges"`
	Color          *string `json:"color" validate:"omitempty,min=2"`
	Icon           *string `json:"icon" validate:"omitempty,min=2"`
	Address        *string `json:"address" validate:"omitempty,min=2"`
	Description    *string `json:"description" validate:"omitempty,min=2"`
	FolderID       *int    `json:"folder_id,omitempty"`
}

type UpdateOrganizationUnitDTO struct {
	ParentID       *int    `json:"parent_id,omitempty"`
	Title          *string `json:"title" validate:"min=2"`
	Abbreviation   *string `json:"abbreviation" validate:"min=2,max=4"`
	NumberOfJudges *int    `json:"number_of_judges"`
	Color          *string `json:"color" validate:"omitempty,min=2"`
	Icon           *string `json:"icon" validate:"omitempty,min=2"`
	Address        *string `json:"address" validate:"omitempty,min=2"`
	Description    *string `json:"description" validate:"omitempty,min=2"`
	FolderID       *int    `json:"folder_id,omitempty"`
}

type OrganizationUnitResponseDTO struct {
	ID             int       `json:"id"`
	ParentID       *int      `json:"parent_id"`
	Title          string    `json:"title"`
	Abbreviation   string    `json:"abbreviation"`
	NumberOfJudges *int      `json:"number_of_judges"`
	Color          *string   `json:"color"`
	Icon           *string   `json:"icon"`
	Address        *string   `json:"address"`
	Description    *string   `json:"description"`
	FolderID       *int      `json:"folder_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetOrganizationUnitsDTO struct {
	Page     *int    `json:"page" validate:"omitempty"`
	PageSize *int    `json:"page_size" validate:"omitempty"`
	ParentID *int    `json:"parent_id" validate:"omitempty"`
	Search   *string `json:"search" validate:"omitempty"`
}

func (dto CreateOrganizationUnitDTO) ToOrganizationUnit() *data.OrganizationUnit {
	var pid *int
	if dto.ParentID != nil && *dto.ParentID != 0 {
		pid = dto.ParentID
	}
	return &data.OrganizationUnit{
		ParentID:       pid,
		Title:          dto.Title,
		Abbreviation:   dto.Abbreviation,
		NumberOfJudges: dto.NumberOfJudges,
		Color:          dto.Color,
		Icon:           dto.Icon,
		Address:        dto.Address,
		Description:    dto.Description,
		FolderID:       dto.FolderID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (dto UpdateOrganizationUnitDTO) ToOrganizationUnit(data *data.OrganizationUnit) {
	data.ParentID = dto.ParentID
	if dto.Title != nil {
		data.Title = *dto.Title
	}
	if dto.Abbreviation != nil {
		data.Abbreviation = *dto.Abbreviation
	}
	data.NumberOfJudges = dto.NumberOfJudges
	data.Color = dto.Color
	data.Icon = dto.Icon
	data.Address = dto.Address
	data.Description = dto.Description
	data.FolderID = dto.FolderID

	data.UpdatedAt = time.Now()
}

func ToOrganizationUnitResponseDTO(data data.OrganizationUnit) OrganizationUnitResponseDTO {
	return OrganizationUnitResponseDTO{
		ID:             data.ID,
		ParentID:       data.ParentID,
		Title:          data.Title,
		Abbreviation:   data.Abbreviation,
		NumberOfJudges: data.NumberOfJudges,
		Color:          data.Color,
		Icon:           data.Icon,
		Address:        data.Address,
		Description:    data.Description,
		FolderID:       data.FolderID,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
}

func ToOrganizationUnitListResponseDTO(organization_units []*data.OrganizationUnit) []OrganizationUnitResponseDTO {
	dtoList := make([]OrganizationUnitResponseDTO, len(organization_units))
	for i, x := range organization_units {
		dtoList[i] = ToOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

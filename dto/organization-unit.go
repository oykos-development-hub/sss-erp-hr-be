package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateOrganizationUnitDTO struct {
	ParentID       *int    `json:"parent_id,omitempty"`
	Title          string  `json:"title" validate:"required"`
	Pib            *string `json:"pib"`
	Abbreviation   *string `json:"abbreviation"`
	NumberOfJudges *int    `json:"number_of_judges"`
	Color          *string `json:"color" validate:"omitempty"`
	Icon           *string `json:"icon" validate:"omitempty"`
	City           *string `json:"city"`
	OrderID        *int    `json:"order_id"`
	Address        *string `json:"address" validate:"omitempty"`
	Description    *string `json:"description" validate:"omitempty"`
	FolderID       *int    `json:"folder_id,omitempty"`
}

type UpdateOrganizationUnitDTO struct {
	ParentID       *int    `json:"parent_id,omitempty"`
	Title          *string `json:"title"`
	Abbreviation   *string `json:"abbreviation"`
	Pib            *string `json:"pib"`
	NumberOfJudges *int    `json:"number_of_judges"`
	Color          *string `json:"color" validate:"omitempty"`
	Icon           *string `json:"icon" validate:"omitempty"`
	City           *string `json:"city"`
	OrderID        *int    `json:"order_id"`
	Address        *string `json:"address" validate:"omitempty"`
	Description    *string `json:"description" validate:"omitempty"`
	FolderID       *int    `json:"folder_id,omitempty"`
}

type OrganizationUnitResponseDTO struct {
	ID             int       `json:"id"`
	ParentID       *int      `json:"parent_id"`
	Title          string    `json:"title"`
	Pib            *string   `json:"pib"`
	Abbreviation   *string   `json:"abbreviation"`
	NumberOfJudges *int      `json:"number_of_judges"`
	Color          *string   `json:"color"`
	OrderID        *int      `json:"order_id"`
	City           *string   `json:"city"`
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
	IsParent *bool   `json:"is_parent"`
	Search   *string `json:"search" validate:"omitempty"`
}

func (dto CreateOrganizationUnitDTO) ToOrganizationUnit() *data.OrganizationUnit {
	var pid *int
	if dto.ParentID != nil && *dto.ParentID != 0 {
		pid = dto.ParentID
	}
	return &data.OrganizationUnit{
		ParentID:       pid,
		Pib:            dto.Pib,
		Title:          dto.Title,
		OrderID:        dto.OrderID,
		Abbreviation:   dto.Abbreviation,
		NumberOfJudges: dto.NumberOfJudges,
		City:           dto.City,
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
		data.Abbreviation = dto.Abbreviation
	}
	data.Pib = dto.Pib
	data.NumberOfJudges = dto.NumberOfJudges
	data.Color = dto.Color
	data.Icon = dto.Icon
	data.Address = dto.Address
	data.City = dto.City
	data.OrderID = dto.OrderID
	data.Description = dto.Description
	data.FolderID = dto.FolderID

	data.UpdatedAt = time.Now()
}

func ToOrganizationUnitResponseDTO(data data.OrganizationUnit) OrganizationUnitResponseDTO {
	return OrganizationUnitResponseDTO{
		ID:             data.ID,
		ParentID:       data.ParentID,
		Title:          data.Title,
		Pib:            data.Pib,
		Abbreviation:   data.Abbreviation,
		NumberOfJudges: data.NumberOfJudges,
		City:           data.City,
		OrderID:        data.OrderID,
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

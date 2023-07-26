package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type CreateSystematizationDTO struct {
	UserProfileID      int       `json:"user_profile_id"  validate:"required"`
	OrganizationUnitID int       `json:"organization_unit_id" validate:"required"`
	Description        string    `json:"description" validate:"required,min=2"`
	SerialNumber       string    `json:"serial_number" validate:"required"`
	Active             bool      `json:"active"`
	DateOfActivation   time.Time `json:"date_of_activation" validate:"required"`
	FileId             *int      `json:"file_id,omitempty"`
}

type UpdateSystematizationDTO struct {
	UserProfileID      *int       `json:"user_profile_id"  validate:"required"`
	OrganizationUnitID *int       `json:"organization_unit_id" validate:"required"`
	Description        *string    `json:"description" validate:"required,min=2"`
	SerialNumber       *string    `json:"serial_number" validate:"required"`
	Active             *bool      `json:"active"`
	DateOfActivation   *time.Time `json:"date_of_activation" validate:"required"`
	FileId             *int       `json:"file_id"`
}

type SystematizationResponseDTO struct {
	ID                 int       `json:"id"`
	UserProfileID      int       `json:"user_profile_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Description        string    `json:"description"`
	SerialNumber       string    `json:"serial_number"`
	Active             bool      `json:"active"`
	DateOfActivation   time.Time `json:"date_of_activation"`
	FileId             *int      `json:"file_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type GetSystematizationsDTO struct {
	Page               *int  `json:"page" validate:"omitempty"`
	PageSize           *int  `json:"page_size" validate:"omitempty"`
	OrganizationUnitID *int  `json:"organization_unit_id" validate:"omitempty"`
	Active             *bool `json:"active" validate:"omitempty"`
}

func (dto CreateSystematizationDTO) ToSystematization() *data.Systematization {
	return &data.Systematization{
		UserProfileID:      dto.UserProfileID,
		OrganizationUnitID: dto.OrganizationUnitID,
		Description:        dto.Description,
		SerialNumber:       dto.SerialNumber,
		Active:             dto.Active,
		DateOfActivation:   dto.DateOfActivation,
		FileId:             dto.FileId,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

func (dto UpdateSystematizationDTO) ToSystematization(data *data.Systematization) {
	if dto.UserProfileID != nil {
		data.UserProfileID = *dto.UserProfileID
	}
	if dto.OrganizationUnitID != nil {
		data.OrganizationUnitID = *dto.OrganizationUnitID
	}
	if dto.Description != nil {
		data.Description = *dto.Description
	}
	if dto.SerialNumber != nil {
		data.SerialNumber = *dto.SerialNumber
	}
	if dto.Active != nil {
		data.Active = *dto.Active
	}
	if dto.DateOfActivation != nil {
		data.DateOfActivation = *dto.DateOfActivation
	}
	if dto.FileId != nil {
		data.FileId = dto.FileId
	}

	data.UpdatedAt = time.Now()
}

func ToSystematizationResponseDTO(data data.Systematization) SystematizationResponseDTO {
	return SystematizationResponseDTO{
		ID:                 data.ID,
		UserProfileID:      data.UserProfileID,
		OrganizationUnitID: data.OrganizationUnitID,
		Description:        data.Description,
		SerialNumber:       data.SerialNumber,
		Active:             data.Active,
		DateOfActivation:   data.DateOfActivation,
		FileId:             data.FileId,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToSystematizationListResponseDTO(systematizations []*data.Systematization) []SystematizationResponseDTO {
	dtoList := make([]SystematizationResponseDTO, len(systematizations))
	for i, x := range systematizations {
		dtoList[i] = ToSystematizationResponseDTO(*x)
	}
	return dtoList
}

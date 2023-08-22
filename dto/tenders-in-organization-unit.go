package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetTendersInputDTO struct {
	Page *int `json:"page" validate:"omitempty"`
	Size *int `json:"size" validate:"omitempty"`
}

type TendersInOrganizationUnitDTO struct {
	PositionInOrganizationUnitID *int      `json:"position_in_organization_unit_id"`
	OrganizationUnitID           int       `json:"organization_unit_id"`
	Type                         int       `json:"type" validate:"required"`
	DateOfStart                  time.Time `json:"date_of_start" validate:"required"`
	DateOfEnd                    time.Time `json:"date_of_end" validate:"required"`
	Description                  string    `json:"description"`
	SerialNumber                 string    `json:"serial_number" validate:"required"`
	AvailableSlots               int       `json:"available_slots" validate:"required"`
	FileID                       *int      `json:"file_id"`
}

type TendersInOrganizationUnitResponseDTO struct {
	ID                           int       `json:"id"`
	PositionInOrganizationUnitID *int      `json:"position_in_organization_unit_id"`
	OrganizationUnitID           int       `json:"organization_unit_id"`
	Type                         int       `json:"type"`
	DateOfStart                  time.Time `json:"date_of_start"`
	DateOfEnd                    time.Time `json:"date_of_end"`
	Description                  string    `json:"description"`
	SerialNumber                 string    `json:"serial_number"`
	AvailableSlots               int       `json:"available_slots"`
	FileID                       *int      `json:"file_id"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}

func (dto TendersInOrganizationUnitDTO) ToTendersInOrganizationUnit() *data.TendersInOrganizationUnit {
	return &data.TendersInOrganizationUnit{
		PositionInOrganizationUnitID: dto.PositionInOrganizationUnitID,
		OrganizationUnitID:           dto.OrganizationUnitID,
		Type:                         dto.Type,
		DateOfStart:                  dto.DateOfStart,
		DateOfEnd:                    dto.DateOfEnd,
		Description:                  dto.Description,
		SerialNumber:                 dto.SerialNumber,
		AvailableSlots:               dto.AvailableSlots,
		FileID:                       dto.FileID,
	}
}

func ToTendersInOrganizationUnitResponseDTO(data data.TendersInOrganizationUnit) TendersInOrganizationUnitResponseDTO {
	return TendersInOrganizationUnitResponseDTO{
		ID:                           data.ID,
		PositionInOrganizationUnitID: data.PositionInOrganizationUnitID,
		OrganizationUnitID:           data.OrganizationUnitID,
		Type:                         data.Type,
		DateOfStart:                  data.DateOfStart,
		DateOfEnd:                    data.DateOfEnd,
		Description:                  data.Description,
		SerialNumber:                 data.SerialNumber,
		AvailableSlots:               data.AvailableSlots,
		FileID:                       data.FileID,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
	}
}

func ToTendersInOrganizationUnitListResponseDTO(tendersinorganizationunits []*data.TendersInOrganizationUnit) []TendersInOrganizationUnitResponseDTO {
	dtoList := make([]TendersInOrganizationUnitResponseDTO, len(tendersinorganizationunits))
	for i, x := range tendersinorganizationunits {
		dtoList[i] = ToTendersInOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

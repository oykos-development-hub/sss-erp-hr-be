package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type TenderApplicationsInOrganizationUnitDTO struct {
	JobTenderID   int  `json:"job_tender_id" validate:"required"`
	UserProfileID int  `json:"user_profile_id" validate:"required"`
	Active        bool `json:"active" validate:"required"`
	FileID        int  `json:"file_id"`
}

type TenderApplicationsInOrganizationUnitResponseDTO struct {
	ID            int       `json:"id"`
	JobTenderID   int       `json:"job_tender_id"`
	UserProfileID int       `json:"user_profile_id"`
	Active        bool      `json:"active"`
	FileID        int       `json:"file_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (dto TenderApplicationsInOrganizationUnitDTO) ToTenderApplicationsInOrganizationUnit() *data.TenderApplicationsInOrganizationUnit {
	return &data.TenderApplicationsInOrganizationUnit{
		JobTenderID:   dto.JobTenderID,
		UserProfileID: dto.UserProfileID,
		Active:        dto.Active,
		FileID:        dto.FileID,
	}
}

func ToTenderApplicationsInOrganizationUnitResponseDTO(data data.TenderApplicationsInOrganizationUnit) TenderApplicationsInOrganizationUnitResponseDTO {
	return TenderApplicationsInOrganizationUnitResponseDTO{
		ID:            data.ID,
		JobTenderID:   data.JobTenderID,
		UserProfileID: data.UserProfileID,
		Active:        data.Active,
		FileID:        data.FileID,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func ToTenderApplicationsInOrganizationUnitListResponseDTO(tenderapplicationsinorganizationunits []*data.TenderApplicationsInOrganizationUnit) []TenderApplicationsInOrganizationUnitResponseDTO {
	dtoList := make([]TenderApplicationsInOrganizationUnitResponseDTO, len(tenderapplicationsinorganizationunits))
	for i, x := range tenderapplicationsinorganizationunits {
		dtoList[i] = ToTenderApplicationsInOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

type GetTenderApplicationsInputDTO struct {
	Page *int `json:"page" validate:"omitempty"`
	Size *int `json:"size" validate:"omitempty"`
}

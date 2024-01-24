package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type TenderApplicationType string

const (
	Internal TenderApplicationType = "internal"
	External TenderApplicationType = "external"
)

type TenderApplicationsInOrganizationUnitDTO struct {
	JobTenderID                    int                   `json:"job_tender_id" validate:"required"`
	UserProfileID                  *int                  `json:"user_profile_id" `
	Active                         bool                  `json:"active"`
	Type                           TenderApplicationType `json:"type" validate:"required"`
	FirstName                      *string               `json:"first_name"`
	LastName                       *string               `json:"last_name"`
	Nationality                    *string               `json:"citizenship"`
	DateOfBirth                    *time.Time            `json:"date_of_birth"`
	DateOfApplication              *time.Time            `json:"date_of_application"`
	OfficialPersonalDocumentNumber *string               `json:"official_personal_document_number"`
	Evaluation                     *string               `json:"evaluation"`
	Status                         string                `json:"status"`
	FileID                         *int                  `json:"file_id"`
}

type TenderApplicationsInOrganizationUnitResponseDTO struct {
	ID                             int                   `json:"id"`
	JobTenderID                    int                   `json:"job_tender_id"`
	UserProfileID                  *int                  `json:"user_profile_id"`
	Active                         bool                  `json:"active"`
	Type                           TenderApplicationType `json:"type"`
	FirstName                      *string               `json:"first_name"`
	LastName                       *string               `json:"last_name"`
	Nationality                    *string               `json:"citizenship"`
	DateOfBirth                    *time.Time            `json:"date_of_birth"`
	DateOfApplication              *time.Time            `json:"date_of_application"`
	OfficialPersonalDocumentNumber *string               `json:"official_personal_document_number"`
	Evaluation                     *string               `json:"evaluation"`
	Status                         string                `json:"status"`
	FileID                         *int                  `json:"file_id"`
	CreatedAt                      time.Time             `json:"created_at"`
	UpdatedAt                      time.Time             `json:"updated_at"`
}

func (dto TenderApplicationsInOrganizationUnitDTO) ToTenderApplicationsInOrganizationUnit() *data.TenderApplicationsInOrganizationUnit {
	return &data.TenderApplicationsInOrganizationUnit{
		JobTenderID:                    dto.JobTenderID,
		UserProfileID:                  dto.UserProfileID,
		IsInternal:                     dto.Type == Internal,
		Active:                         dto.Active,
		FirstName:                      dto.FirstName,
		LastName:                       dto.LastName,
		Status:                         dto.Status,
		Evaluation:                     dto.Evaluation,
		OfficialPersonalDocumentNumber: dto.OfficialPersonalDocumentNumber,
		DateOfBirth:                    dto.DateOfBirth,
		Nationality:                    dto.Nationality,
		DateOfApplication:              dto.DateOfApplication,
		FileID:                         dto.FileID,
	}
}

func ToTenderApplicationsInOrganizationUnitResponseDTO(data data.TenderApplicationsInOrganizationUnit) TenderApplicationsInOrganizationUnitResponseDTO {
	var applicationType TenderApplicationType
	if data.IsInternal {
		applicationType = Internal
	} else {
		applicationType = External
	}

	return TenderApplicationsInOrganizationUnitResponseDTO{
		ID:                             data.ID,
		JobTenderID:                    data.JobTenderID,
		UserProfileID:                  data.UserProfileID,
		Active:                         data.Active,
		Type:                           applicationType,
		FirstName:                      data.FirstName,
		LastName:                       data.LastName,
		Status:                         data.Status,
		Evaluation:                     data.Evaluation,
		OfficialPersonalDocumentNumber: data.OfficialPersonalDocumentNumber,
		DateOfBirth:                    data.DateOfBirth,
		DateOfApplication:              data.DateOfApplication,
		FileID:                         data.FileID,
		Nationality:                    data.Nationality,
		CreatedAt:                      data.CreatedAt,
		UpdatedAt:                      data.UpdatedAt,
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
	JobTenderID   *int    `json:"job_tender_id"`
	UserProfileID *int    `json:"user_profile_id"`
	Search        *string `json:"search"`
}

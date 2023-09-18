package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionsOfOrganizationUnitDTO struct {
	Name                            *string    `json:"name"`
	RevisionTypeID                  *int       `json:"revision_type_id"`
	RevisorUserProfileID            *int       `json:"revisor_user_profile_id"`
	RevisorUserProfile              *string    `json:"revisor_user_profile"`
	InternalOrganizationUnitID      *int       `json:"internal_organization_unit_id"`
	ExternalOrganizationUnitID      *int       `json:"external_organization_unit_id"`
	ResponsibleUserProfileID        *int       `json:"responsible_user_profile_id"`
	ResponsibleUserProfile          *string    `json:"responsible_user_profile"`
	ImplementationUserProfileID     *int       `json:"implementation_user_profile_id"`
	ImplementationUserProfile       *string    `json:"implementation_user_profile"`
	Title                           *string    `json:"title"`
	PlanedYear                      *string    `json:"planned_year"`
	PlannedQuarter                  *string    `json:"planned_quarter"`
	SerialNumber                    *string    `json:"serial_number"`
	Priority                        *string    `json:"priority"`
	DateOfRevision                  *time.Time `json:"date_of_revision"`
	DateOfAcceptance                *time.Time `json:"date_of_acceptance"`
	DateOfRejection                 *time.Time `json:"date_of_rejection"`
	ImplementationSuggestion        *string    `json:"implementation_suggestion"`
	ImplementationMonthSpan         *string    `json:"implementation_month_span"`
	DateOfImplementation            *time.Time `json:"date_of_implementation"`
	StateOfImplementation           *string    `json:"state_of_implementation"`
	ImplementationFailedDescription *string    `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string    `json:"second_implementation_month_span"`
	SecondDateOfRevision            *time.Time `json:"second_date_of_revision"`
	RefDocument                     *string    `json:"ref_document"`
	FileID                          *int       `json:"file_id"`
}

type RevisionsOfOrganizationUnitResponseDTO struct {
	ID                              int        `json:"id"`
	Name                            *string    `json:"name"`
	RevisionTypeID                  *int       `json:"revision_type_id"`
	RevisorUserProfileID            *int       `json:"revisor_user_profile_id"`
	RevisorUserProfile              *string    `json:"revisor_user_profile"`
	InternalOrganizationUnitID      *int       `json:"internal_organization_unit_id"`
	ExternalOrganizationUnitID      *int       `json:"external_organization_unit_id"`
	ResponsibleUserProfileID        *int       `json:"responsible_user_profile_id"`
	ResponsibleUserProfile          *string    `json:"responsible_user_profile"`
	ImplementationUserProfileID     *int       `json:"implementation_user_profile_id"`
	ImplementationUserProfile       *string    `json:"implementation_user_profile"`
	Title                           *string    `json:"title"`
	PlanedYear                      *string    `json:"planned_year"`
	PlannedQuarter                  *string    `json:"planned_quarter"`
	SerialNumber                    *string    `json:"serial_number"`
	Priority                        *string    `json:"priority"`
	DateOfRevision                  *time.Time `json:"date_of_revision"`
	DateOfAcceptance                *time.Time `json:"date_of_acceptance"`
	DateOfRejection                 *time.Time `json:"date_of_rejection"`
	ImplementationSuggestion        *string    `json:"implementation_suggestion"`
	ImplementationMonthSpan         *string    `json:"implementation_month_span"`
	DateOfImplementation            *time.Time `json:"date_of_implementation"`
	StateOfImplementation           *string    `json:"state_of_implementation"`
	ImplementationFailedDescription *string    `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string    `json:"second_implementation_month_span"`
	SecondDateOfRevision            *time.Time `json:"second_date_of_revision"`
	FileID                          *int       `json:"file_id"`
	RefDocument                     *string    `json:"ref_document"`
	CreatedAt                       time.Time  `json:"created_at"`
	UpdatedAt                       time.Time  `json:"updated_at"`
}

func (dto RevisionsOfOrganizationUnitDTO) ToRevisionsOfOrganizationUnit() *data.RevisionsOfOrganizationUnit {
	return &data.RevisionsOfOrganizationUnit{
		Name:                            dto.Name,
		RevisionTypeID:                  dto.RevisionTypeID,
		RevisorUserProfileID:            dto.RevisorUserProfileID,
		RevisorUserProfile:              dto.RevisorUserProfile,
		InternalOrganizationUnitID:      dto.InternalOrganizationUnitID,
		ExternalOrganizationUnitID:      dto.ExternalOrganizationUnitID,
		ResponsibleUserProfileID:        dto.ResponsibleUserProfileID,
		ResponsibleUserProfile:          dto.ResponsibleUserProfile,
		ImplementationUserProfileID:     dto.ImplementationUserProfileID,
		ImplementationUserProfile:       dto.ImplementationUserProfile,
		Title:                           dto.Title,
		PlanedYear:                      dto.PlanedYear,
		PlannedQuarter:                  dto.PlannedQuarter,
		SerialNumber:                    dto.SerialNumber,
		Priority:                        dto.Priority,
		DateOfRevision:                  dto.DateOfRevision,
		DateOfAcceptance:                dto.DateOfAcceptance,
		DateOfRejection:                 dto.DateOfRejection,
		ImplementationSuggestion:        dto.ImplementationSuggestion,
		ImplementationMonthSpan:         dto.ImplementationMonthSpan,
		DateOfImplementation:            dto.DateOfImplementation,
		StateOfImplementation:           dto.StateOfImplementation,
		ImplementationFailedDescription: dto.ImplementationFailedDescription,
		SecondImplementationMonthSpan:   dto.SecondImplementationMonthSpan,
		SecondDateOfRevision:            dto.SecondDateOfRevision,
		RefDocument:                     dto.RefDocument,
		FileID:                          dto.FileID,
	}
}

func ToRevisionsOfOrganizationUnitResponseDTO(data data.RevisionsOfOrganizationUnit) RevisionsOfOrganizationUnitResponseDTO {
	return RevisionsOfOrganizationUnitResponseDTO{
		ID:                              data.ID,
		Name:                            data.Name,
		RevisionTypeID:                  data.RevisionTypeID,
		RevisorUserProfileID:            data.RevisorUserProfileID,
		RevisorUserProfile:              data.RevisorUserProfile,
		InternalOrganizationUnitID:      data.InternalOrganizationUnitID,
		ExternalOrganizationUnitID:      data.ExternalOrganizationUnitID,
		ResponsibleUserProfileID:        data.ResponsibleUserProfileID,
		ResponsibleUserProfile:          data.ResponsibleUserProfile,
		ImplementationUserProfileID:     data.ImplementationUserProfileID,
		ImplementationUserProfile:       data.ImplementationUserProfile,
		Title:                           data.Title,
		PlanedYear:                      data.PlanedYear,
		PlannedQuarter:                  data.PlannedQuarter,
		SerialNumber:                    data.SerialNumber,
		Priority:                        data.Priority,
		DateOfRevision:                  data.DateOfRevision,
		DateOfAcceptance:                data.DateOfAcceptance,
		DateOfRejection:                 data.DateOfRejection,
		ImplementationSuggestion:        data.ImplementationSuggestion,
		ImplementationMonthSpan:         data.ImplementationMonthSpan,
		DateOfImplementation:            data.DateOfImplementation,
		StateOfImplementation:           data.StateOfImplementation,
		ImplementationFailedDescription: data.ImplementationFailedDescription,
		SecondImplementationMonthSpan:   data.SecondImplementationMonthSpan,
		SecondDateOfRevision:            data.SecondDateOfRevision,
		FileID:                          data.FileID,
		RefDocument:                     data.RefDocument,
		CreatedAt:                       data.CreatedAt,
		UpdatedAt:                       data.UpdatedAt,
	}
}

func ToRevisionsOfOrganizationUnitListResponseDTO(revisionsoforganizationunits []*data.RevisionsOfOrganizationUnit) []RevisionsOfOrganizationUnitResponseDTO {
	dtoList := make([]RevisionsOfOrganizationUnitResponseDTO, len(revisionsoforganizationunits))
	for i, x := range revisionsoforganizationunits {
		dtoList[i] = ToRevisionsOfOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

type GetRevisionsInputDTO struct {
	Page                       *int `json:"page" validate:"omitempty"`
	Size                       *int `json:"size" validate:"omitempty"`
	InternalOrganizationUnitID *int `json:"internal_organization_unit_id" validate:"omitempty"`
	ExternalOrganizationUnitID *int `json:"external_organization_unit_id" validate:"omitempty"`
	RevisorUserProfileID       *int `json:"revisor_user_profile_id" validate:"omitempty"`
}

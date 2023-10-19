package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionsInOrganizationUnitDTO struct {
	ID                 int `json:"id"`
	RevisionID         int `json:"revision_id"`
	OrganizationUnitID int `json:"organization_unit_id"`
}

type RevisionsInOrganizationUnitResponseDTO struct {
	ID                 int       `json:"id"`
	RevisionID         int       `json:"revision_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type RevisionOrgUnitFilter struct {
	RevisionID         *int `json:"revision_id"`
	OrganizationUnitID *int `json:"organization_unit_id"`
}

func (dto RevisionsInOrganizationUnitDTO) ToRevisionsInOrganizationUnit() *data.RevisionsInOrganizationUnit {
	return &data.RevisionsInOrganizationUnit{
		ID:                 dto.ID,
		RevisionID:         dto.RevisionID,
		OrganizationUnitID: dto.OrganizationUnitID,
	}
}

func ToRevisionsInOrganizationUnitResponseDTO(data data.RevisionsInOrganizationUnit) RevisionsInOrganizationUnitResponseDTO {
	return RevisionsInOrganizationUnitResponseDTO{
		ID:                 data.ID,
		RevisionID:         data.RevisionID,
		OrganizationUnitID: data.OrganizationUnitID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToRevisionsInOrganizationUnitListResponseDTO(revisions_in_organization_units []*data.RevisionsInOrganizationUnit) []RevisionsInOrganizationUnitResponseDTO {
	dtoList := make([]RevisionsInOrganizationUnitResponseDTO, len(revisions_in_organization_units))
	for i, x := range revisions_in_organization_units {
		dtoList[i] = ToRevisionsInOrganizationUnitResponseDTO(*x)
	}
	return dtoList
}

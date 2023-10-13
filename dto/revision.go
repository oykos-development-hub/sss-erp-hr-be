package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisonFilter struct {
	Page                    *int `json:"page"`
	Size                    *int `json:"size"`
	Revisor                 *int `json:"revisor"`
	RevisionType            *int `json:"revision_type_id"`
	InternalRevisionSubject *int `json:"internal_revision_subject"`
	PlanID                  *int `json:"plan_id"`
}

type RevisionDTO struct {
	Title                   string    `json:"title" validate:"required"`
	PlanID                  int       `json:"plan_id" validate:"required"`
	SerialNumber            string    `json:"serial_number" validate:"required"`
	DateOfRevision          time.Time `json:"date_of_revision" validate:"required"`
	RevisionQuartal         string    `json:"revision_quartal" validate:"required"`
	InternalRevisionSubject *int      `json:"internal_revision_subject"`
	ExternalRevisionSubject *int      `json:"external_revision_subject"`
	Revisor                 int       `json:"revisor" validate:"required"`
	RevisionType            int       `json:"revision_type" validate:"required"`
	FileID                  *int      `json:"file_id"`
}

type RevisionResponseDTO struct {
	ID                      int       `json:"id"`
	Title                   string    `json:"title"`
	PlanID                  int       `json:"plan_id"`
	SerialNumber            string    `json:"serial_number"`
	DateOfRevision          time.Time `json:"date_of_revision"`
	RevisionQuartal         string    `json:"revision_quartal"`
	InternalRevisionSubject *int      `json:"internal_revision_subject"`
	ExternalRevisionSubject *int      `json:"external_revision_subject"`
	Revisor                 int       `json:"revisor"`
	RevisionType            int       `json:"revision_type"`
	FileID                  *int      `json:"file_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func (dto RevisionDTO) ToRevision() *data.Revision {
	return &data.Revision{
		Title:                   dto.Title,
		PlanID:                  dto.PlanID,
		SerialNumber:            dto.SerialNumber,
		DateOfRevision:          dto.DateOfRevision,
		RevisionQuartal:         dto.RevisionQuartal,
		InternalRevisionSubject: dto.InternalRevisionSubject,
		ExternalRevisionSubject: dto.ExternalRevisionSubject,
		Revisor:                 dto.Revisor,
		RevisionType:            dto.RevisionType,
		FileID:                  dto.FileID,
	}
}

func ToRevisionResponseDTO(data data.Revision) RevisionResponseDTO {
	return RevisionResponseDTO{
		ID:                      data.ID,
		Title:                   data.Title,
		PlanID:                  data.PlanID,
		SerialNumber:            data.SerialNumber,
		DateOfRevision:          data.DateOfRevision,
		RevisionQuartal:         data.RevisionQuartal,
		InternalRevisionSubject: data.InternalRevisionSubject,
		ExternalRevisionSubject: data.ExternalRevisionSubject,
		Revisor:                 data.Revisor,
		RevisionType:            data.RevisionType,
		FileID:                  data.FileID,
		CreatedAt:               data.CreatedAt,
		UpdatedAt:               data.UpdatedAt,
	}
}

func ToRevisionListResponseDTO(revisions []*data.Revision) []RevisionResponseDTO {
	dtoList := make([]RevisionResponseDTO, len(revisions))
	for i, x := range revisions {
		dtoList[i] = ToRevisionResponseDTO(*x)
	}
	return dtoList
}

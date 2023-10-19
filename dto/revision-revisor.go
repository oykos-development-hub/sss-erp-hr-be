package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionRevisorDTO struct {
	ID         int `json:"id"`
	RevisionID int `json:"revision_id"`
	RevisorID  int `json:"revisor_id"`
}

type RevisionRevisorResponseDTO struct {
	ID         int       `json:"id"`
	RevisionID int       `json:"revision_id"`
	RevisorID  int       `json:"revisor_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RevisionRevisorFilter struct {
	RevisionID *int `json:"revision_id"`
	RevisorID  *int `json:"revisor_id"`
}

func (dto RevisionRevisorDTO) ToRevisionRevisor() *data.RevisionRevisor {
	return &data.RevisionRevisor{
		ID:         dto.ID,
		RevisionID: dto.RevisionID,
		RevisorID:  dto.RevisorID,
	}
}

func ToRevisionRevisorResponseDTO(data data.RevisionRevisor) RevisionRevisorResponseDTO {
	return RevisionRevisorResponseDTO{
		ID:         data.ID,
		RevisionID: data.RevisionID,
		RevisorID:  data.RevisorID,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func ToRevisionRevisorListResponseDTO(revision_revisors []*data.RevisionRevisor) []RevisionRevisorResponseDTO {
	dtoList := make([]RevisionRevisorResponseDTO, len(revision_revisors))
	for i, x := range revision_revisors {
		dtoList[i] = ToRevisionRevisorResponseDTO(*x)
	}
	return dtoList
}

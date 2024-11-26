package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionTipImplementationsFilter struct {
	TipID *int `json:"tip_id"`
}

type RevisionTipImplementationDTO struct {
	ID                     int        `json:"id,omitempty"`
	TipID                  int        `json:"tip_id"`
	Status                 string     `json:"status"`
	NewDueDate             *int       `json:"new_due_date"`
	NewDateOfExecution     *time.Time `json:"new_date_of_execution"`
	ReasonsForNonExecuting *string    `json:"reasons_for_non_executing"`
	RevisorID              *int       `json:"revisor_id"`
	Documents              string     `json:"documents"`
	FileIDs                []int64    `json:"file_ids"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

type RevisionTipImplementationResponseDTO struct {
	ID                     int        `json:"id,omitempty"`
	TipID                  int        `json:"tip_id"`
	Status                 string     `json:"status"`
	NewDueDate             *int       `json:"new_due_date"`
	NewDateOfExecution     *time.Time `json:"new_date_of_execution"`
	ReasonsForNonExecuting *string    `json:"reasons_for_non_executing"`
	RevisorID              *int       `json:"revisor_id"`
	Documents              string     `json:"documents"`
	FileIDs                []int64    `json:"file_ids"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

func (dto RevisionTipImplementationDTO) ToRevisionTipImplementation() *data.RevisionTipImplementation {
	return &data.RevisionTipImplementation{
		TipID:                  dto.TipID,
		Status:                 dto.Status,
		NewDueDate:             dto.NewDueDate,
		NewDateOfExecution:     dto.NewDateOfExecution,
		ReasonsForNonExecuting: dto.ReasonsForNonExecuting,
		RevisorID:              dto.RevisorID,
		Documents:              dto.Documents,
		FileIDs:                dto.FileIDs,
	}
}

func ToRevisionTipImplementationResponseDTO(data data.RevisionTipImplementation) RevisionTipImplementationResponseDTO {
	return RevisionTipImplementationResponseDTO{
		ID:                     data.ID,
		TipID:                  data.TipID,
		Status:                 data.Status,
		NewDueDate:             data.NewDueDate,
		NewDateOfExecution:     data.NewDateOfExecution,
		ReasonsForNonExecuting: data.ReasonsForNonExecuting,
		RevisorID:              data.RevisorID,
		Documents:              data.Documents,
		FileIDs:                data.FileIDs,
		UpdatedAt:              data.UpdatedAt,
		CreatedAt:              data.CreatedAt,
	}
}

func ToRevisionTipImplementationListResponseDTO(revisiontipimplementations []*data.RevisionTipImplementation) []RevisionTipImplementationResponseDTO {
	dtoList := make([]RevisionTipImplementationResponseDTO, len(revisiontipimplementations))
	for i, x := range revisiontipimplementations {
		dtoList[i] = ToRevisionTipImplementationResponseDTO(*x)
	}
	return dtoList
}

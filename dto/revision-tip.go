package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionTipFilter struct {
	RevisionID *int `json:"revision_id"`
	Page       *int `json:"page"`
	Size       *int `json:"size"`
}

type RevisionTipDTO struct {
	RevisionID             int        `json:"revision_id" validate:"required"`
	UserProfileID          *int       `json:"user_profile_id"`
	DateOfAccept           *time.Time `json:"date_of_accept"`
	DueDate                int        `json:"due_date" validate:"required"`
	DateOfReject           *time.Time `json:"date_of_reject"`
	DateOfExecution        *time.Time `json:"date_of_execution"`
	Recommendation         string     `json:"recommendation" validate:"required"`
	Status                 *string    `json:"status"`
	Documents              *string    `json:"documents"`
	ResponsiblePerson      *string    `json:"responsible_person"`
	ReasonsForNonExecuting *string    `json:"reasons_for_non_executing"`
	FileID                 *int       `json:"file_id"`
}

type RevisionTipResponseDTO struct {
	ID                     int        `json:"id"`
	RevisionID             int        `json:"revision_id"`
	UserProfileID          *int       `json:"user_profile_id"`
	DateOfAccept           *time.Time `json:"date_of_accept"`
	DueDate                int        `json:"due_date"`
	DateOfReject           *time.Time `json:"date_of_reject"`
	DateOfExecution        *time.Time `json:"date_of_execution"`
	Recommendation         string     `json:"recommendation"`
	Status                 *string    `json:"status"`
	Documents              *string    `json:"documents"`
	ReasonsForNonExecuting *string    `json:"reasons_for_non_executing"`
	FileID                 *int       `json:"file_id"`
	ResponsiblePerson      *string    `json:"responsible_person"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

func (dto RevisionTipDTO) ToRevisionTip() *data.RevisionTip {
	return &data.RevisionTip{
		UserProfileID:          dto.UserProfileID,
		RevisionID:             dto.RevisionID,
		DateOfAccept:           dto.DateOfAccept,
		DueDate:                dto.DueDate,
		DateOfReject:           dto.DateOfReject,
		DateOfExecution:        dto.DateOfExecution,
		Recommendation:         dto.Recommendation,
		Status:                 dto.Status,
		Documents:              dto.Documents,
		ResponsiblePerson:      dto.ResponsiblePerson,
		ReasonsForNonExecuting: dto.ReasonsForNonExecuting,
		FileID:                 dto.FileID,
	}
}

func ToRevisionTipResponseDTO(data data.RevisionTip) RevisionTipResponseDTO {
	return RevisionTipResponseDTO{
		ID:                     data.ID,
		RevisionID:             data.RevisionID,
		UserProfileID:          data.UserProfileID,
		DateOfAccept:           data.DateOfAccept,
		DueDate:                data.DueDate,
		DateOfReject:           data.DateOfReject,
		DateOfExecution:        data.DateOfExecution,
		Recommendation:         data.Recommendation,
		Status:                 data.Status,
		Documents:              data.Documents,
		ReasonsForNonExecuting: data.ReasonsForNonExecuting,
		ResponsiblePerson:      data.ResponsiblePerson,
		FileID:                 data.FileID,
		CreatedAt:              data.CreatedAt,
		UpdatedAt:              data.UpdatedAt,
	}
}

func ToRevisionTipListResponseDTO(revision_tips []*data.RevisionTip) []RevisionTipResponseDTO {
	dtoList := make([]RevisionTipResponseDTO, len(revision_tips))
	for i, x := range revision_tips {
		dtoList[i] = ToRevisionTipResponseDTO(*x)
	}
	return dtoList
}

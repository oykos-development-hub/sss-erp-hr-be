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
	RevisionID        int        `json:"revision_id" validate:"required"`
	UserProfileID     *int       `json:"user_profile_id"`
	DateOfAccept      *time.Time `json:"date_of_accept"`
	DueDate           int        `json:"due_date" validate:"required"`
	RevisionPriority  *string    `json:"revision_priority"`
	EndDate           *time.Time `json:"end_date"`
	DateOfReject      *time.Time `json:"date_of_reject"`
	DateOfExecution   *time.Time `json:"date_of_execution"`
	Recommendation    string     `json:"recommendation" validate:"required"`
	Status            string     `json:"status"`
	ResponsiblePerson *string    `json:"responsible_person"`
	FileIDs           []int64    `json:"file_ids"`
}

type RevisionTipResponseDTO struct {
	ID                int        `json:"id"`
	RevisionID        int        `json:"revision_id" validate:"required"`
	UserProfileID     *int       `json:"user_profile_id"`
	DateOfAccept      *time.Time `json:"date_of_accept"`
	DueDate           int        `json:"due_date" validate:"required"`
	RevisionPriority  *string    `json:"revision_priority"`
	EndDate           *time.Time `json:"end_date"`
	DateOfReject      *time.Time `json:"date_of_reject"`
	DateOfExecution   *time.Time `json:"date_of_execution"`
	Recommendation    string     `json:"recommendation" validate:"required"`
	Status            string     `json:"status"`
	ResponsiblePerson *string    `json:"responsible_person"`
	FileIDs           []int64    `json:"file_ids"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (dto RevisionTipDTO) ToRevisionTip() *data.RevisionTip {
	if dto.Status == "" {
		dto.Status = data.StatusTipImplNotFinished
	}

	return &data.RevisionTip{
		UserProfileID:     dto.UserProfileID,
		RevisionID:        dto.RevisionID,
		DateOfAccept:      dto.DateOfAccept,
		DueDate:           dto.DueDate,
		DateOfReject:      dto.DateOfReject,
		EndDate:           dto.EndDate,
		DateOfExecution:   dto.DateOfExecution,
		RevisionPriority:  dto.RevisionPriority,
		Recommendation:    dto.Recommendation,
		Status:            dto.Status,
		ResponsiblePerson: dto.ResponsiblePerson,
		FileIDs:           dto.FileIDs,
	}
}

func ToRevisionTipResponseDTO(data data.RevisionTip) RevisionTipResponseDTO {
	return RevisionTipResponseDTO{
		ID:                data.ID,
		RevisionID:        data.RevisionID,
		UserProfileID:     data.UserProfileID,
		DateOfAccept:      data.DateOfAccept,
		DueDate:           data.DueDate,
		EndDate:           data.EndDate,
		DateOfReject:      data.DateOfReject,
		RevisionPriority:  data.RevisionPriority,
		DateOfExecution:   data.DateOfExecution,
		Recommendation:    data.Recommendation,
		Status:            data.Status,
		ResponsiblePerson: data.ResponsiblePerson,
		FileIDs:           data.FileIDs,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func ToRevisionTipListResponseDTO(revision_tips []*data.RevisionTip) []RevisionTipResponseDTO {
	dtoList := make([]RevisionTipResponseDTO, len(revision_tips))
	for i, x := range revision_tips {
		dtoList[i] = ToRevisionTipResponseDTO(*x)
	}
	return dtoList
}

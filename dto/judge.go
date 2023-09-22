package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type JudgeFilter struct {
	UserProfileID      *int `json:"user_profile_id"`
	OrganizationUnitID *int `json:"organization_unit_id"`
	ResolutionID       *int `json:"resolution_id"`
	Page               *int `json:"page"`
	Size               *int `json:"size"`
}

type JudgeDTO struct {
	UserProfileID      int  `json:"user_profile_id"`
	OrganizationUnitID int  `json:"organization_unit_id"`
	ResolutionID       int  `json:"resolution_id"`
	IsPresident        bool `json:"is_president"`
}

type JudgeResponseDTO struct {
	ID                 int       `json:"id"`
	UserProfileID      int       `json:"user_profile_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	ResolutionID       int       `json:"resolution_id"`
	IsPresident        bool      `json:"is_president"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (dto JudgeDTO) ToJudge() *data.Judge {
	return &data.Judge{
		UserProfileID:      dto.UserProfileID,
		OrganizationUnitID: dto.OrganizationUnitID,
		ResolutionID:       dto.ResolutionID,
		IsPresident:        dto.IsPresident,
	}
}

func ToJudgeResponseDTO(data data.Judge) JudgeResponseDTO {
	return JudgeResponseDTO{
		ID:                 data.ID,
		UserProfileID:      data.UserProfileID,
		OrganizationUnitID: data.OrganizationUnitID,
		ResolutionID:       data.ResolutionID,
		IsPresident:        data.IsPresident,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToJudgeListResponseDTO(judges []*data.Judge) []JudgeResponseDTO {
	dtoList := make([]JudgeResponseDTO, len(judges))
	for i, x := range judges {
		dtoList[i] = ToJudgeResponseDTO(*x)
	}
	return dtoList
}

package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type ForeignerDTO struct {
	UserProfileId                   int        `json:"user_profile_id" validate:"required"`
	WorkPermitNumber                string     `json:"work_permit_number" validate:"required"`
	WorkPermitIssuer                string     `json:"work_permit_issuer" validate:"required"`
	WorkPermitDateOfStart           time.Time  `json:"work_permit_date_of_start" validate:"required"`
	WorkPermitDateOfEnd             *time.Time `json:"work_permit_date_of_end" validate:"omitempty"`
	WorkPermitIndefiniteLength      *bool      `json:"work_permit_indefinite_length" validate:"required"`
	ResidencePermitDateOfStart      time.Time  `json:"residence_permit_date_of_start" validate:"required"`
	ResidencePermitDateOfEnd        *time.Time `json:"residence_permit_date_of_end" validate:"omitempty"`
	ResidencePermitIndefiniteLength *bool      `json:"residence_permit_indefinite_length" validate:"required"`
	CountryOfOrigin                 string     `json:"country_of_origin" validate:"required"`
	WorkPermitFileId                *int       `json:"work_permit_file_id" validate:"omitempty"`
	ResidencePermitFileId           *int       `json:"residence_permit_file_id" validate:"omitempty"`
}

type ForeignerResponseDTO struct {
	ID                              int        `json:"id" db:"id,omitempty"`
	UserProfileId                   int        `json:"user_profile_id"`
	WorkPermitNumber                string     `json:"work_permit_number"`
	WorkPermitIssuer                string     `json:"work_permit_issuer"`
	WorkPermitDateOfStart           time.Time  `json:"work_permit_date_of_start"`
	WorkPermitDateOfEnd             *time.Time `json:"work_permit_date_of_end"`
	WorkPermitIndefiniteLength      *bool      `json:"work_permit_indefinite_length"`
	ResidencePermitDateOfStart      time.Time  `json:"residence_permit_date_of_start"`
	ResidencePermitDateOfEnd        *time.Time `json:"residence_permit_date_of_end"`
	ResidencePermitIndefiniteLength *bool      `json:"residence_permit_indefinite_length"`
	CountryOfOrigin                 string     `json:"country_of_origin"`
	WorkPermitFileId                *int       `json:"work_permit_file_id"`
	ResidencePermitFileId           *int       `json:"residence_permit_file_id"`
	CreatedAt                       time.Time  `json:"created_at"`
	UpdatedAt                       time.Time  `json:"updated_at"`
}

func (dto ForeignerDTO) ToForeigner() *data.Foreigner {
	return &data.Foreigner{
		UserProfileId:                   dto.UserProfileId,
		WorkPermitNumber:                dto.WorkPermitNumber,
		WorkPermitIssuer:                dto.WorkPermitIssuer,
		WorkPermitDateOfStart:           dto.WorkPermitDateOfStart,
		WorkPermitDateOfEnd:             dto.WorkPermitDateOfEnd,
		WorkPermitIndefiniteLength:      dto.WorkPermitIndefiniteLength,
		ResidencePermitDateOfStart:      dto.ResidencePermitDateOfStart,
		ResidencePermitDateOfEnd:        dto.ResidencePermitDateOfEnd,
		ResidencePermitIndefiniteLength: dto.ResidencePermitIndefiniteLength,
		CountryOfOrigin:                 dto.CountryOfOrigin,
		WorkPermitFileId:                dto.WorkPermitFileId,
		ResidencePermitFileId:           dto.ResidencePermitFileId,
	}
}

func ToForeignerResponseDTO(data data.Foreigner) ForeignerResponseDTO {
	return ForeignerResponseDTO{
		ID:                              data.ID,
		UserProfileId:                   data.UserProfileId,
		WorkPermitNumber:                data.WorkPermitNumber,
		WorkPermitIssuer:                data.WorkPermitIssuer,
		WorkPermitDateOfStart:           data.WorkPermitDateOfStart,
		WorkPermitDateOfEnd:             data.WorkPermitDateOfEnd,
		WorkPermitIndefiniteLength:      data.WorkPermitIndefiniteLength,
		ResidencePermitDateOfStart:      data.ResidencePermitDateOfStart,
		ResidencePermitDateOfEnd:        data.ResidencePermitDateOfEnd,
		ResidencePermitIndefiniteLength: data.ResidencePermitIndefiniteLength,
		CountryOfOrigin:                 data.CountryOfOrigin,
		WorkPermitFileId:                data.WorkPermitFileId,
		ResidencePermitFileId:           data.ResidencePermitFileId,
		CreatedAt:                       data.CreatedAt,
		UpdatedAt:                       data.UpdatedAt,
	}
}

func ToForeignerListResponseDTO(foreigners []*data.Foreigner) []ForeignerResponseDTO {
	dtoList := make([]ForeignerResponseDTO, len(foreigners))
	for i, x := range foreigners {
		dtoList[i] = ToForeignerResponseDTO(*x)
	}
	return dtoList
}

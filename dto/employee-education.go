package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type EducationInput struct {
	UserProfileID int  `json:"user_profile_id" validate:"required"`
	TypeID        *int `json:"type_id"`
	SubTypeID     *int `json:"sub_type_id"`
}

type EmployeeEducationDTO struct {
	UserProfileID       int        `json:"user_profile_id" validate:"required"`
	TypeID              int        `json:"type_id" validate:"required"`
	DateOfCertification *time.Time `json:"date_of_certification" validate:"omitempty,datetime"`
	Price               *int       `json:"price" validate:"omitempty,numeric"`
	DateOfStart         *time.Time `json:"date_of_start" validate:"omitempty,datetime"`
	DateOfEnd           *time.Time `json:"date_of_end" validate:"omitempty,datetime"`
	AcademicTitle       *string    `json:"academic_title" validate:"omitempty"`
	ExpertiseLevel      *string    `json:"expertise_level" validate:"omitempty"`
	CertificateIssuer   *string    `json:"certificate_issuer" validate:"omitempty"`
	Description         *string    `json:"description" validate:"omitempty"`
	Title               *string    `json:"title" validate:"omitempty"`
	FileId              *int       `json:"file_id" validate:"omitempty"`
}

type EmployeeEducationResponseDTO struct {
	ID                  int        `json:"id"`
	UserProfileID       int        `json:"user_profile_id"`
	TypeID              int        `json:"type_id"`
	DateOfCertification *time.Time `json:"date_of_certification"`
	Price               *int       `json:"price"`
	DateOfStart         *time.Time `json:"date_of_start"`
	DateOfEnd           *time.Time `json:"date_of_end"`
	AcademicTitle       *string    `json:"academic_title"`
	ExpertiseLevel      *string    `json:"expertise_level"`
	CertificateIssuer   *string    `json:"certificate_issuer"`
	Title               *string    `json:"title"`
	Description         *string    `json:"description"`
	FileID              *int       `json:"file_id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (dto EmployeeEducationDTO) ToEmployeeEducation() *data.EmployeeEducation {
	return &data.EmployeeEducation{
		UserProfileID:       dto.UserProfileID,
		TypeID:              dto.TypeID,
		DateOfCertification: dto.DateOfCertification,
		Price:               dto.Price,
		DateOfStart:         dto.DateOfStart,
		DateOfEnd:           dto.DateOfEnd,
		AcademicTitle:       dto.AcademicTitle,
		ExpertiseLevel:      dto.ExpertiseLevel,
		CertificateIssuer:   dto.CertificateIssuer,
		Title:               dto.Title,
		Description:         dto.Description,
		FileId:              dto.FileId,
	}
}

func ToEmployeeEducationResponseDTO(data data.EmployeeEducation) EmployeeEducationResponseDTO {
	return EmployeeEducationResponseDTO{
		ID:                  data.ID,
		UserProfileID:       data.UserProfileID,
		TypeID:              data.TypeID,
		DateOfCertification: data.DateOfCertification,
		Price:               data.Price,
		DateOfStart:         data.DateOfStart,
		DateOfEnd:           data.DateOfEnd,
		AcademicTitle:       data.AcademicTitle,
		ExpertiseLevel:      data.ExpertiseLevel,
		CertificateIssuer:   data.CertificateIssuer,
		Description:         data.Description,
		FileID:              data.FileId,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
}

func ToEmployeeEducationListResponseDTO(employeeeducations []*data.EmployeeEducation) []EmployeeEducationResponseDTO {
	dtoList := make([]EmployeeEducationResponseDTO, len(employeeeducations))
	for i, x := range employeeeducations {
		dtoList[i] = ToEmployeeEducationResponseDTO(*x)
	}
	return dtoList
}

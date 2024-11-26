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
	DateOfCertification *time.Time `json:"date_of_certification"`
	Price               *float32   `json:"price"`
	DateOfStart         *time.Time `json:"date_of_start"`
	DateOfEnd           *time.Time `json:"date_of_end"`
	AcademicTitle       *string    `json:"academic_title"`
	ExpertiseLevel      *string    `json:"expertise_level"`
	Score               *string    `json:"score"`
	CertificateIssuer   *string    `json:"certificate_issuer"`
	Description         *string    `json:"description"`
	Title               *string    `json:"title"`
	FileIDs             []int64    `json:"file_ids"`
}

type EmployeeEducationResponseDTO struct {
	ID                  int        `json:"id"`
	UserProfileID       int        `json:"user_profile_id"`
	TypeID              int        `json:"type_id"`
	DateOfCertification *time.Time `json:"date_of_certification"`
	Price               *float32   `json:"price"`
	DateOfStart         *time.Time `json:"date_of_start"`
	DateOfEnd           *time.Time `json:"date_of_end"`
	AcademicTitle       *string    `json:"academic_title"`
	ExpertiseLevel      *string    `json:"expertise_level"`
	Score               *string    `json:"score"`
	CertificateIssuer   *string    `json:"certificate_issuer"`
	Title               *string    `json:"title"`
	Description         *string    `json:"description"`
	FileIDs             []int64    `json:"file_ids"`
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
		Score:               dto.Score,
		CertificateIssuer:   dto.CertificateIssuer,
		Title:               dto.Title,
		Description:         dto.Description,
		FileIDs:             dto.FileIDs,
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
		Score:               data.Score,
		Description:         data.Description,
		FileIDs:             data.FileIDs,
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

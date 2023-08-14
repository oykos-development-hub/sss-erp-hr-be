package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type EmployeeFamilyMemberDTO struct {
	UserProfileID        int       `json:"user_profile_id" validate:"required"`
	FirstName            string    `json:"first_name" validate:"required"`
	MiddleName           *string   `json:"middle_name"`
	LastName             string    `json:"last_name" validate:"required"`
	BirthLastName        *string   `json:"birth_last_name"`
	FatherName           string    `json:"father_name" validate:"required"`
	MotherName           string    `json:"mother_name" validate:"required"`
	MotherBirthLastName  string    `json:"mother_birth_last_name" validate:"required"`
	DateOfBirth          time.Time `json:"date_of_birth" validate:"required"`
	CountryOfBirth       string    `json:"country_of_birth" validate:"required"`
	CityOfBirth          string    `json:"city_of_birth" validate:"required"`
	Nationality          string    `json:"nationality" validate:"required"`
	NationalMinority     *string   `json:"national_minority"`
	Citizenship          string    `json:"citizenship" validate:"required"`
	Address              string    `json:"address" validate:"required"`
	OfficialPersonalID   string    `json:"official_personal_id" validate:"required"`
	Gender               string    `json:"gender" validate:"required"`
	EmployeeRelationship string    `json:"employee_relationship" validate:"required"`
	InsuranceCoverage    string    `json:"insurance_coverage"`
	HandicappedPerson    bool      `json:"handicapped_person"`
}

type EmployeeFamilyMemberResponseDTO struct {
	ID                   int       `json:"id"`
	UserProfileID        int       `json:"user_profile_id"`
	FirstName            string    `json:"first_name"`
	MiddleName           *string   `json:"middle_name"`
	LastName             string    `json:"last_name"`
	BirthLastName        *string   `json:"birth_last_name"`
	FatherName           string    `json:"father_name"`
	MotherName           string    `json:"mother_name"`
	MotherBirthLastName  string    `json:"mother_birth_last_name"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	CountryOfBirth       string    `json:"country_of_birth"`
	CityOfBirth          string    `json:"city_of_birth"`
	Nationality          string    `json:"nationality"`
	NationalMinority     *string   `json:"national_minority"`
	Citizenship          string    `json:"citizenship"`
	Address              string    `json:"address"`
	OfficialPersonalID   string    `json:"official_personal_id"`
	Gender               string    `json:"gender"`
	EmployeeRelationship string    `json:"employee_relationship"`
	InsuranceCoverage    string    `json:"insurance_coverage"`
	HandicappedPerson    bool      `json:"handicapped_person"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (dto EmployeeFamilyMemberDTO) ToEmployeeFamilyMember() *data.EmployeeFamilyMember {
	return &data.EmployeeFamilyMember{
		UserProfileID:        dto.UserProfileID,
		FirstName:            dto.FirstName,
		MiddleName:           dto.MiddleName,
		LastName:             dto.LastName,
		BirthLastName:        dto.BirthLastName,
		FatherName:           dto.FatherName,
		MotherName:           dto.MotherName,
		MotherBirthLastName:  dto.MotherBirthLastName,
		DateOfBirth:          dto.DateOfBirth,
		CountryOfBirth:       dto.CountryOfBirth,
		CityOfBirth:          dto.CityOfBirth,
		Nationality:          dto.Nationality,
		NationalMinority:     dto.NationalMinority,
		Citizenship:          dto.Citizenship,
		Address:              dto.Address,
		OfficialPersonalID:   dto.OfficialPersonalID,
		Gender:               dto.Gender,
		EmployeeRelationship: dto.EmployeeRelationship,
		InsuranceCoverage:    dto.InsuranceCoverage,
		HandicappedPerson:    dto.HandicappedPerson,
	}
}

func ToEmployeeFamilyMemberResponseDTO(data data.EmployeeFamilyMember) EmployeeFamilyMemberResponseDTO {
	return EmployeeFamilyMemberResponseDTO{
		ID:                   data.ID,
		UserProfileID:        data.UserProfileID,
		FirstName:            data.FirstName,
		MiddleName:           data.MiddleName,
		LastName:             data.LastName,
		BirthLastName:        data.BirthLastName,
		FatherName:           data.FatherName,
		MotherName:           data.MotherName,
		MotherBirthLastName:  data.MotherBirthLastName,
		DateOfBirth:          data.DateOfBirth,
		CountryOfBirth:       data.CountryOfBirth,
		CityOfBirth:          data.CityOfBirth,
		Nationality:          data.Nationality,
		NationalMinority:     data.NationalMinority,
		Citizenship:          data.Citizenship,
		Address:              data.Address,
		OfficialPersonalID:   data.OfficialPersonalID,
		Gender:               data.Gender,
		EmployeeRelationship: data.EmployeeRelationship,
		InsuranceCoverage:    data.InsuranceCoverage,
		HandicappedPerson:    data.HandicappedPerson,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
	}
}

func ToEmployeeFamilyMemberListResponseDTO(employeefamilymembers []*data.EmployeeFamilyMember) []EmployeeFamilyMemberResponseDTO {
	dtoList := make([]EmployeeFamilyMemberResponseDTO, len(employeefamilymembers))
	for i, x := range employeefamilymembers {
		dtoList[i] = ToEmployeeFamilyMemberResponseDTO(*x)
	}
	return dtoList
}

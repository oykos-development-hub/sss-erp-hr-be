package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type UserProfileDTO struct {
	UserAccountId             int       `json:"user_account_id"`
	FirstName                 string    `json:"first_name" validate:"required"`
	MiddleName                *string   `json:"middle_name"`
	LastName                  string    `json:"last_name" validate:"required"`
	BirthLastName             *string   `json:"birth_last_name"`
	FatherName                string    `json:"father_name" validate:"required"`
	MotherName                string    `json:"mother_name" validate:"required"`
	MotherBirthLastName       *string   `json:"mother_birth_last_name"`
	DateOfBirth               JSONTime  `json:"date_of_birth" validate:"required"`
	CountryOfBirth            string    `json:"country_of_birth" validate:"required"`
	CityOfBirth               string    `json:"city_of_birth"`
	Nationality               string    `json:"nationality" validate:"required"`
	NationalMinority          *string   `json:"national_minority"`
	Citizenship               string    `json:"citizenship" validate:"required"`
	Address                   string    `json:"address" validate:"required"`
	BankAccount               *string   `json:"bank_account"`
	BankName                  *string   `json:"bank_name,omitempty"`
	OfficialPersonalID        string    `json:"official_personal_id" validate:"required"`
	OfficialPersonalDocNumber string    `json:"official_personal_document_number" validate:"required"`
	OfficialPersonalDocIssuer string    `json:"official_personal_document_issuer" validate:"required"`
	Gender                    string    `json:"gender" validate:"required"`
	SingleParent              *bool     `json:"single_parent" validate:"required"`
	HousingDone               *bool     `json:"housing_done" validate:"required"`
	HousingDescription        string    `json:"housing_description"`
	MartialStatus             string    `json:"martial_status"`
	DateOfTakingOath          *JSONTime `json:"date_of_taking_oath"`
	DateOfBecomingJudge       *string   `json:"date_of_becoming_judge"`
	RevisorRole               *bool     `json:"revisor_role"`
	EngagementTypeID          *int      `json:"engagement_type_id" validate:"omitempty"`
	ActiveContract            *bool     `db:"active_contract"`
}

type UserProfileResponseDTO struct {
	ID                        int        `json:"id"`
	UserAccountId             int        `json:"user_account_id"`
	FirstName                 string     `json:"first_name"`
	MiddleName                *string    `json:"middle_name"`
	LastName                  string     `json:"last_name"`
	BirthLastName             *string    `json:"birth_last_name"`
	FatherName                string     `json:"father_name"`
	MotherName                string     `json:"mother_name"`
	MotherBirthLastName       *string    `json:"mother_birth_last_name"`
	DateOfBirth               time.Time  `json:"date_of_birth"`
	CountryOfBirth            string     `json:"country_of_birth"`
	CityOfBirth               string     `json:"city_of_birth"`
	Nationality               string     `json:"nationality"`
	NationalMinority          *string    `json:"national_minority"`
	Citizenship               string     `json:"citizenship"`
	Address                   string     `json:"address"`
	BankAccount               *string    `json:"bank_account"`
	BankName                  *string    `json:"bank_name"`
	OfficialPersonalID        string     `json:"official_personal_id"`
	OfficialPersonalDocNumber string     `json:"official_personal_document_number"`
	OfficialPersonalDocIssuer string     `json:"official_personal_document_issuer"`
	Gender                    string     `json:"gender"`
	SingleParent              *bool      `json:"single_parent"`
	HousingDone               *bool      `json:"housing_done"`
	HousingDescription        string     `json:"housing_description"`
	MartialStatus             string     `json:"martial_status"`
	DateOfTakingOath          *time.Time `json:"date_of_taking_oath"`
	DateOfBecomingJudge       *string    `json:"date_of_becoming_judge"`
	RevisorRole               *bool      `json:"revisor_role"`
	EngagementTypeID          *int       `json:"engagement_type_id"`
	ActiveContract            *bool      `db:"active_contract"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}

func (dto *UserProfileDTO) ToUserProfile() *data.UserProfile {
	return &data.UserProfile{
		UserAccountId:             dto.UserAccountId,
		FirstName:                 dto.FirstName,
		MiddleName:                dto.MiddleName,
		LastName:                  dto.LastName,
		BirthLastName:             dto.BirthLastName,
		FatherName:                dto.FatherName,
		MotherName:                dto.MotherName,
		MotherBirthLastName:       dto.MotherBirthLastName,
		DateOfBirth:               time.Time(dto.DateOfBirth),
		CountryOfBirth:            dto.CountryOfBirth,
		CityOfBirth:               dto.CityOfBirth,
		Nationality:               dto.Nationality,
		NationalMinority:          dto.NationalMinority,
		Citizenship:               dto.Citizenship,
		Address:                   dto.Address,
		BankAccount:               dto.BankAccount,
		BankName:                  dto.BankName,
		OfficialPersonalID:        dto.OfficialPersonalID,
		OfficialPersonalDocNumber: dto.OfficialPersonalDocNumber,
		OfficialPersonalDocIssuer: dto.OfficialPersonalDocIssuer,
		Gender:                    dto.Gender,
		SingleParent:              dto.SingleParent,
		HousingDone:               dto.HousingDone,
		HousingDescription:        dto.HousingDescription,
		MartialStatus:             dto.MartialStatus,
		DateOfTakingOath:          (*time.Time)(dto.DateOfTakingOath),
		DateOfBecomingJudge:       dto.DateOfBecomingJudge,
		RevisorRole:               dto.RevisorRole,
		EngagementTypeID:          dto.EngagementTypeID,
		ActiveContract:            dto.ActiveContract,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}
}

func ToUserProfileListResponseDTO(userprofiles []*data.UserProfile) []UserProfileResponseDTO {
	dtoList := make([]UserProfileResponseDTO, len(userprofiles))
	for i, x := range userprofiles {
		dtoList[i] = ToUserProfileResponseDTO(*x)
	}
	return dtoList
}

func ToUserProfileResponseDTO(data data.UserProfile) UserProfileResponseDTO {
	return UserProfileResponseDTO{
		ID:                        data.ID,
		UserAccountId:             data.UserAccountId,
		FirstName:                 data.FirstName,
		MiddleName:                data.MiddleName,
		LastName:                  data.LastName,
		BirthLastName:             data.BirthLastName,
		FatherName:                data.FatherName,
		MotherName:                data.MotherName,
		MotherBirthLastName:       data.MotherBirthLastName,
		DateOfBirth:               data.DateOfBirth,
		CountryOfBirth:            data.CountryOfBirth,
		CityOfBirth:               data.CityOfBirth,
		Nationality:               data.Nationality,
		NationalMinority:          data.NationalMinority,
		Citizenship:               data.Citizenship,
		Address:                   data.Address,
		BankAccount:               data.BankAccount,
		BankName:                  data.BankName,
		OfficialPersonalID:        data.OfficialPersonalID,
		OfficialPersonalDocNumber: data.OfficialPersonalDocNumber,
		OfficialPersonalDocIssuer: data.OfficialPersonalDocIssuer,
		Gender:                    data.Gender,
		SingleParent:              data.SingleParent,
		HousingDone:               data.HousingDone,
		HousingDescription:        data.HousingDescription,
		MartialStatus:             data.MartialStatus,
		DateOfTakingOath:          data.DateOfTakingOath,
		DateOfBecomingJudge:       data.DateOfBecomingJudge,
		RevisorRole:               data.RevisorRole,
		EngagementTypeID:          data.EngagementTypeID,
		ActiveContract:            data.ActiveContract,
		CreatedAt:                 data.CreatedAt,
		UpdatedAt:                 data.UpdatedAt,
	}
}

type GetProfilesInputDTO struct {
	Page      *int  `json:"page" validate:"omitempty"`
	Size      *int  `json:"size" validate:"omitempty"`
	IsRevisor *bool `json:"is_revisor" validate:"omitempty"`
	AccountID *int  `json:"account_id" validate:"omitempty"`
}

package dto

import (
	"fmt"
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetEmployeeContracts struct {
	Active *bool `json:"active"`
}

type EmployeeContractDTO struct {
	UserProfileID     int        `json:"user_profile_id"`
	ContractTypeID    int        `json:"contract_type_id"`
	Abbreviation      *string    `json:"abbreviation"`
	Description       *string    `json:"description"`
	Active            bool       `json:"active"`
	SerialNumber      *string    `json:"serial_number"`
	NetSalary         *string    `json:"net_salary"`
	GrossSalary       *string    `json:"gross_salary"`
	BankAccount       *string    `json:"bank_account"`
	BankName          *string    `json:"bank_name"`
	DateOfSignature   *time.Time `json:"date_of_signature"`
	DateOfEligibility *time.Time `json:"date_of_eligibility"`
	DateOfStart       time.Time  `json:"date_of_start"`
	DateOfEnd         *time.Time `json:"date_of_end"`
	FileID            *int       `json:"file_id"`
}

func (dto EmployeeContractDTO) ToEmployeeContract() *data.EmployeeContract {
	fmt.Println(dto.UserProfileID)
	return &data.EmployeeContract{
		UserProfileID:     dto.UserProfileID,
		ContractTypeID:    dto.ContractTypeID,
		Abbreviation:      dto.Abbreviation,
		Description:       dto.Description,
		Active:            dto.Active,
		SerialNumber:      dto.SerialNumber,
		NetSalary:         dto.NetSalary,
		GrossSalary:       dto.GrossSalary,
		BankAccount:       dto.BankAccount,
		BankName:          dto.BankName,
		DateOfSignature:   dto.DateOfSignature,
		DateOfEligibility: dto.DateOfEligibility,
		DateOfStart:       dto.DateOfStart,
		DateOfEnd:         dto.DateOfEnd,
		FileID:            dto.FileID,
	}
}

type EmployeeContractResponseDTO struct {
	ID                int        `json:"id"`
	UserProfileID     int        `json:"user_profile_id"`
	ContractTypeID    int        `json:"contract_type_id"`
	Abbreviation      *string    `json:"abbreviation"`
	Description       *string    `json:"description"`
	Active            bool       `json:"active"`
	SerialNumber      *string    `json:"serial_number"`
	NetSalary         *string    `json:"net_salary"`
	GrossSalary       *string    `json:"gross_salary"`
	BankAccount       *string    `json:"bank_account"`
	BankName          *string    `json:"bank_name"`
	DateOfSignature   *time.Time `json:"date_of_signature"`
	DateOfEligibility *time.Time `json:"date_of_eligibility"`
	DateOfStart       time.Time  `json:"date_of_start"`
	DateOfEnd         *time.Time `json:"date_of_end"`
	FileID            *int       `json:"file_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func ToEmployeeContractResponseDTO(data data.EmployeeContract) EmployeeContractResponseDTO {
	return EmployeeContractResponseDTO{
		ID:                data.ID,
		UserProfileID:     data.UserProfileID,
		ContractTypeID:    data.ContractTypeID,
		Abbreviation:      data.Abbreviation,
		Description:       data.Description,
		Active:            data.Active,
		SerialNumber:      data.SerialNumber,
		NetSalary:         data.NetSalary,
		GrossSalary:       data.GrossSalary,
		BankAccount:       data.BankAccount,
		BankName:          data.BankName,
		DateOfSignature:   data.DateOfSignature,
		DateOfEligibility: data.DateOfEligibility,
		DateOfStart:       data.DateOfStart,
		DateOfEnd:         data.DateOfEnd,
		FileID:            data.FileID,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func ToEmployeeContractListResponseDTO(employeecontract []*data.EmployeeContract) []EmployeeContractResponseDTO {
	dtoList := make([]EmployeeContractResponseDTO, len(employeecontract))
	for i, x := range employeecontract {
		dtoList[i] = ToEmployeeContractResponseDTO(*x)
	}
	return dtoList
}

package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type EmployeeContractServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeContract
}

func NewEmployeeContractServiceImpl(app *celeritas.Celeritas, repo data.EmployeeContract) EmployeeContractService {
	return &EmployeeContractServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeContractServiceImpl) CreateEmployeeContract(input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error) {
	data := input.ToEmployeeContract()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToEmployeeContractResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeContractServiceImpl) UpdateEmployeeContract(id int, input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error) {
	data := input.ToEmployeeContract()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	updatedData, _ := h.repo.Get(id)

	response := dto.ToEmployeeContractResponseDTO(*updatedData)

	return &response, nil
}

func (h *EmployeeContractServiceImpl) DeleteEmployeeContract(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

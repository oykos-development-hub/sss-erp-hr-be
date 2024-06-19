package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *EmployeeContractServiceImpl) CreateEmployeeContract(ctx context.Context, input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error) {
	data := input.ToEmployeeContract()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee contract insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee contract get")
	}

	res := dto.ToEmployeeContractResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeContractServiceImpl) UpdateEmployeeContract(ctx context.Context, id int, input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error) {
	data := input.ToEmployeeContract()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee contract update")
	}

	updatedData, _ := h.repo.Get(id)

	response := dto.ToEmployeeContractResponseDTO(*updatedData)

	return &response, nil
}

func (h *EmployeeContractServiceImpl) DeleteEmployeeContract(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee contract delete")
	}

	return nil
}

package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeeResolutionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeResolution
}

func NewEmployeeResolutionServiceImpl(app *celeritas.Celeritas, repo data.EmployeeResolution) EmployeeResolutionService {
	return &EmployeeResolutionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeResolutionServiceImpl) CreateEmployeeResolution(input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error) {
	data := input.ToEmployeeResolution()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToEmployeeResolutionResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeResolutionServiceImpl) UpdateEmployeeResolution(id int, input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error) {
	data := input.ToEmployeeResolution()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToEmployeeResolutionResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeResolutionServiceImpl) DeleteEmployeeResolution(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *EmployeeResolutionServiceImpl) GetEmployeeResolution(id int) (*dto.EmployeeResolutionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToEmployeeResolutionResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeResolutionServiceImpl) GetEmployeeResolutionList(userProfileID int, input dto.GetResolutionListInputDTO) ([]dto.EmployeeResolutionResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": userProfileID,
	}
	if input.From != nil {
		cond["date_of_end >= "] = input.From
	}
	if input.To != nil {
		cond["date_of_start <= "] = input.To
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToEmployeeResolutionListResponseDTO(data)

	return response, nil
}

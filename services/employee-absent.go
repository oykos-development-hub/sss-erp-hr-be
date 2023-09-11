package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeeAbsentServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeAbsent
}

func NewEmployeeAbsentServiceImpl(app *celeritas.Celeritas, repo data.EmployeeAbsent) EmployeeAbsentService {
	return &EmployeeAbsentServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeAbsentServiceImpl) CreateEmployeeAbsent(input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error) {
	data := input.ToEmployeeAbsent()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToEmployeeAbsentResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeAbsentServiceImpl) UpdateEmployeeAbsent(id int, input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error) {
	data := input.ToEmployeeAbsent()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToEmployeeAbsentResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeAbsentServiceImpl) DeleteEmployeeAbsent(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *EmployeeAbsentServiceImpl) GetEmployeeAbsentList(userProfileID int, input dto.GetEmployeeAbsentsInputDTO) ([]dto.EmployeeAbsentResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": userProfileID,
	}
	if input.Date != nil {
		cond["date_of_start <="] = input.Date
		cond["date_of_end >="] = input.Date
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
	response := dto.ToEmployeeAbsentListResponseDTO(data)

	return response, nil
}

func (h *EmployeeAbsentServiceImpl) GetAbsent(id int) (*dto.EmployeeAbsentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToEmployeeAbsentResponseDTO(*data)

	return &response, nil
}

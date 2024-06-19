package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *EmployeeAbsentServiceImpl) CreateEmployeeAbsent(ctx context.Context, input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error) {
	data := input.ToEmployeeAbsent()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee absent insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee absent get")
	}

	res := dto.ToEmployeeAbsentResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeAbsentServiceImpl) UpdateEmployeeAbsent(ctx context.Context, id int, input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error) {
	data := input.ToEmployeeAbsent()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee absent update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee absent get")
	}

	response := dto.ToEmployeeAbsentResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeAbsentServiceImpl) DeleteEmployeeAbsent(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee absent delete")
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
		return nil, newErrors.Wrap(err, "repo employee absent get all")
	}
	response := dto.ToEmployeeAbsentListResponseDTO(data)

	return response, nil
}

func (h *EmployeeAbsentServiceImpl) GetAbsent(id int) (*dto.EmployeeAbsentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee absent get")
	}
	response := dto.ToEmployeeAbsentResponseDTO(*data)

	return &response, nil
}

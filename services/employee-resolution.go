package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *EmployeeResolutionServiceImpl) CreateEmployeeResolution(ctx context.Context, input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error) {
	data := input.ToEmployeeResolution()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee resolution insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee resolution get")
	}

	res := dto.ToEmployeeResolutionResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeResolutionServiceImpl) UpdateEmployeeResolution(ctx context.Context, id int, input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error) {
	data := input.ToEmployeeResolution()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee resolution update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee resolution get")
	}

	response := dto.ToEmployeeResolutionResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeResolutionServiceImpl) DeleteEmployeeResolution(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee resolution delete")
	}

	return nil
}

func (h *EmployeeResolutionServiceImpl) GetEmployeeResolution(id int) (*dto.EmployeeResolutionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee resolution resolution")
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
		return nil, newErrors.Wrap(err, "repo employee resolution get all")
	}
	response := dto.ToEmployeeResolutionListResponseDTO(data)

	return response, nil
}

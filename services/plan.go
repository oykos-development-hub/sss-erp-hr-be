package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type PlanServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Plan
}

func NewPlanServiceImpl(app *celeritas.Celeritas, repo data.Plan) PlanService {
	return &PlanServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *PlanServiceImpl) CreatePlan(ctx context.Context, input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}

	res := dto.ToPlanResponseDTO(*data)

	return &res, nil
}

func (h *PlanServiceImpl) UpdatePlan(ctx context.Context, id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}

	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) DeletePlan(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo plan delete")
	}

	return nil
}

func (h *PlanServiceImpl) GetPlan(id int) (*dto.PlanResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}
	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) GetPlanList(input dto.PlanFilter) ([]dto.PlanResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.Year != nil {
		cond["year"] = input.Year
	}
	data, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo plan get all")
	}
	response := dto.ToPlanListResponseDTO(data)

	return response, total, nil
}

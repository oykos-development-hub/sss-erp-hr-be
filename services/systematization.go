package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

type SystematizationServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Systematization
}

func NewSystematizationServiceImpl(app *celeritas.Celeritas, repo data.Systematization) SystematizationService {
	return &SystematizationServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *SystematizationServiceImpl) CreateSystematization(ctx context.Context, input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data := input.ToSystematization()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo systematization insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo systematization get")
	}

	res := dto.ToSystematizationResponseDTO(*data)

	return &res, nil
}

func (h *SystematizationServiceImpl) UpdateSystematization(ctx context.Context, id int, input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data := input.ToSystematization()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo systematization update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo systematization get")
	}

	response := dto.ToSystematizationResponseDTO(*data)

	return &response, nil
}

func (h *SystematizationServiceImpl) DeleteSystematization(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {

		return newErrors.Wrap(err, "repo systematization delete")
	}

	return nil
}

func (h *SystematizationServiceImpl) GetSystematization(id int) (*dto.SystematizationResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo systematization get")
	}
	response := dto.ToSystematizationResponseDTO(*data)

	return &response, nil
}

func (h *SystematizationServiceImpl) GetSystematizationList(input dto.GetSystematizationsDTO) ([]dto.SystematizationResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if input.Active != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"active": *input.Active})
	}

	if input.OrganizationUnitID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"organization_unit_id =": *input.OrganizationUnitID})
	}

	if input.Search != nil {
		likeCondition := fmt.Sprintf("%%%s%%", *input.Search)
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"serial_number ILIKE": likeCondition})
	}

	if input.Year != nil {
		year, err := strconv.Atoi(*input.Year)
		if err == nil {
			startDate := fmt.Sprintf("%d-01-01", year)
			endDate := fmt.Sprintf("%d-12-31", year)
			// Pretpostavljajući da koristite SQL biblioteku koja podržava placeholder-e u ovom obliku
			conditionAndExp = up.And(conditionAndExp, &up.Cond{"date_of_activation >=": startDate})
			conditionAndExp = up.And(conditionAndExp, &up.Cond{"date_of_activation <=": endDate})
		}
	}

	res, total, err := h.repo.GetAll(input.Page, input.PageSize, conditionAndExp)
	if err != nil {

		return nil, nil, newErrors.Wrap(err, "repo systematization get all")
	}
	response := dto.ToSystematizationListResponseDTO(res)

	return response, total, nil
}

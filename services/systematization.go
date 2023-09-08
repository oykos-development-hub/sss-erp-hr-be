package services

import (
	"fmt"
	"strconv"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"
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

func (h *SystematizationServiceImpl) CreateSystematization(input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data := input.ToSystematization()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToSystematizationResponseDTO(*data)

	return &res, nil
}

func (h *SystematizationServiceImpl) UpdateSystematization(id int, input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToSystematization()

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToSystematizationResponseDTO(*data)

	return &response, nil
}

func (h *SystematizationServiceImpl) DeleteSystematization(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *SystematizationServiceImpl) GetSystematization(id int) (*dto.SystematizationResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
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
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToSystematizationListResponseDTO(res)

	return response, total, nil
}

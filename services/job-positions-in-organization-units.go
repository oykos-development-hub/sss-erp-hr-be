package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type JobPositionsInOrganizationUnitsServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.JobPositionsInOrganizationUnits
}

func NewJobPositionsInOrganizationUnitsServiceImpl(app *celeritas.Celeritas, repo data.JobPositionsInOrganizationUnits) JobPositionsInOrganizationUnitsService {
	return &JobPositionsInOrganizationUnitsServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *JobPositionsInOrganizationUnitsServiceImpl) CreateJobPositionsInOrganizationUnits(input dto.CreateJobPositionsInOrganizationUnitsDTO) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error) {
	data := input.ToJobPositionsInOrganizationUnits()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToJobPositionsInOrganizationUnitsResponseDTO(*data)

	return &res, nil
}

func (h *JobPositionsInOrganizationUnitsServiceImpl) UpdateJobPositionsInOrganizationUnits(input dto.CreateJobPositionsInOrganizationUnitsDTO) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error) {
	data := input.ToJobPositionsInOrganizationUnits()

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(input.Id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToJobPositionsInOrganizationUnitsResponseDTO(*data)

	return &res, nil
}

func (h *JobPositionsInOrganizationUnitsServiceImpl) DeleteJobPositionsInOrganizationUnits(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *JobPositionsInOrganizationUnitsServiceImpl) GetJobPositionInOrganziationUnitById(id int) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToJobPositionsInOrganizationUnitsResponseDTO(*data)

	return &response, nil
}

func (h *JobPositionsInOrganizationUnitsServiceImpl) GetJobPositionsInOrganizationUnitsList(data dto.GetJobPositionsInOrganizationUnitsDTO) ([]dto.JobPositionsInOrganizationUnitsResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if data.JobPositionID != nil && *data.JobPositionID > 0 {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"job_position_id =": *data.JobPositionID})
	}

	if data.SystematizationID != nil && *data.SystematizationID > 0 {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"systematization_id =": *data.SystematizationID})
	}

	if data.OrganizationUnitID != nil && *data.OrganizationUnitID > 0 {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"parent_organization_unit_id =": *data.OrganizationUnitID})
	}

	res, total, err := h.repo.GetAll(data.Page, data.PageSize, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToJobPositionsInOrganizationUnitsListResponseDTO(res)

	return response, total, nil
}

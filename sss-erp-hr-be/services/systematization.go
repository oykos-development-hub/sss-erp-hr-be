package services

import (
	"github.com/oykos-development-hub/celeritas"
	"github.com/upper/db/v4"
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

func (h *SystematizationServiceImpl) CreateSystematization(input dto.CreateSystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data := input.ToSystematization()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	if data.Active {
		err = h.checkSystematizationActiveStatuses(data.ID, data.OrganizationUnitID)
		if err != nil {
			return nil, errors.ErrInternalServer
		}
	}

	res := dto.ToSystematizationResponseDTO(*data)

	return &res, nil
}

func (h *SystematizationServiceImpl) UpdateSystematization(id int, input dto.UpdateSystematizationDTO) (*dto.SystematizationResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToSystematization(data)

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	err = h.checkSystematizationActiveStatuses(data.ID, data.OrganizationUnitID)
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
	cond := up.Cond{}
	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = input.OrganizationUnitID
	}
	if input.Active != nil {
		cond["active"] = input.Active
	}
	res, total, err := h.repo.GetAll(input.Page, input.PageSize, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToSystematizationListResponseDTO(res)

	return response, total, nil
}

func (h *SystematizationServiceImpl) checkSystematizationActiveStatuses(id int, organizationUnitID int) error {
	condition := db.Cond{"id <>": id, "active =": true, "organization_unit_id =": organizationUnitID}
	res, total, err := h.repo.GetAll(nil, nil, &condition)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}
	if *total > 0 {
		for _, systematization := range res {
			systematization.Active = false
			err := h.repo.Update(*systematization)
			if err != nil {
				return errors.ErrInternalServer
			}
		}
	}
	return nil
}

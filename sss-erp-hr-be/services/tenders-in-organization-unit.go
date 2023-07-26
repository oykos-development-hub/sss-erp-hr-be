package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type TendersInOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.TendersInOrganizationUnit
}

func NewTendersInOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.TendersInOrganizationUnit) TendersInOrganizationUnitService {
	return &TendersInOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *TendersInOrganizationUnitServiceImpl) CreateTendersInOrganizationUnit(input dto.TendersInOrganizationUnitDTO) (*dto.TendersInOrganizationUnitResponseDTO, error) {
	data := input.ToTendersInOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToTendersInOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *TendersInOrganizationUnitServiceImpl) UpdateTendersInOrganizationUnit(id int, input dto.TendersInOrganizationUnitDTO) (*dto.TendersInOrganizationUnitResponseDTO, error) {
	data := input.ToTendersInOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToTendersInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *TendersInOrganizationUnitServiceImpl) DeleteTendersInOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *TendersInOrganizationUnitServiceImpl) GetTendersInOrganizationUnit(id int) (*dto.TendersInOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToTendersInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *TendersInOrganizationUnitServiceImpl) GetTendersInOrganizationUnitList(input dto.GetTendersInputDTO) ([]dto.TendersInOrganizationUnitResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.IsActive != nil {
		cond["active"] = input.IsActive
	}
	data, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToTendersInOrganizationUnitListResponseDTO(data)

	return response, total, nil
}

package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type TenderApplicationsInOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.TenderApplicationsInOrganizationUnit
}

func NewTenderApplicationsInOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.TenderApplicationsInOrganizationUnit) TenderApplicationsInOrganizationUnitService {
	return &TenderApplicationsInOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *TenderApplicationsInOrganizationUnitServiceImpl) CreateTenderApplicationsInOrganizationUnit(input dto.TenderApplicationsInOrganizationUnitDTO) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error) {
	data := input.ToTenderApplicationsInOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToTenderApplicationsInOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *TenderApplicationsInOrganizationUnitServiceImpl) UpdateTenderApplicationsInOrganizationUnit(id int, input dto.TenderApplicationsInOrganizationUnitDTO) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error) {
	data := input.ToTenderApplicationsInOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToTenderApplicationsInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *TenderApplicationsInOrganizationUnitServiceImpl) DeleteTenderApplicationsInOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *TenderApplicationsInOrganizationUnitServiceImpl) GetTenderApplicationsInOrganizationUnit(id int) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToTenderApplicationsInOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *TenderApplicationsInOrganizationUnitServiceImpl) GetTenderApplicationsInOrganizationUnitList(input dto.GetTenderApplicationsInputDTO) ([]dto.TenderApplicationsInOrganizationUnitResponseDTO, *uint64, error) {
	data, total, err := h.repo.GetAll(input.Page, input.Size, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToTenderApplicationsInOrganizationUnitListResponseDTO(data)

	return response, total, nil
}

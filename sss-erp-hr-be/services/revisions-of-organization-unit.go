package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionsOfOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RevisionsOfOrganizationUnit
}

func NewRevisionsOfOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.RevisionsOfOrganizationUnit) RevisionsOfOrganizationUnitService {
	return &RevisionsOfOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionsOfOrganizationUnitServiceImpl) CreateRevisionsOfOrganizationUnit(input dto.RevisionsOfOrganizationUnitDTO) (*dto.RevisionsOfOrganizationUnitResponseDTO, error) {
	data := input.ToRevisionsOfOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRevisionsOfOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) UpdateRevisionsOfOrganizationUnit(id int, input dto.RevisionsOfOrganizationUnitDTO) (*dto.RevisionsOfOrganizationUnitResponseDTO, error) {
	data := input.ToRevisionsOfOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		h.App.InfoLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRevisionsOfOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) DeleteRevisionsOfOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) GetRevisionsOfOrganizationUnit(id int) (*dto.RevisionsOfOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRevisionsOfOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) GetRevisionsOfOrganizationUnitList(input *dto.GetRevisionsInputDTO) ([]dto.RevisionsOfOrganizationUnitResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.InternalOrganizationUnitID != nil {
		cond["internal_organization_unit_id"] = *input.InternalOrganizationUnitID
	}
	if input.ExternalOrganizationUnitID != nil {
		cond["external_organization_unit_id"] = *input.ExternalOrganizationUnitID
	}
	if input.RevisorUserProfileID != nil {
		cond["revisor_user_profile_id"] = *input.RevisorUserProfileID
	}
	res, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToRevisionsOfOrganizationUnitListResponseDTO(res)

	return response, total, nil
}

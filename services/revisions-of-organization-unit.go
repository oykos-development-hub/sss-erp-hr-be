package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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
		return nil, newErrors.Wrap(err, "repo revisions of organization unit create")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revisions of organization unit get")
	}

	res := dto.ToRevisionsOfOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) UpdateRevisionsOfOrganizationUnit(id int, input dto.RevisionsOfOrganizationUnitDTO) (*dto.RevisionsOfOrganizationUnitResponseDTO, error) {
	data := input.ToRevisionsOfOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revisions of organization unit update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revisions of organization unit get")
	}

	response := dto.ToRevisionsOfOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) DeleteRevisionsOfOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {

		return newErrors.Wrap(err, "repo revisions of organization unit delete")
	}

	return nil
}

func (h *RevisionsOfOrganizationUnitServiceImpl) GetRevisionsOfOrganizationUnit(id int) (*dto.RevisionsOfOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo revisions of organization unit get")
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

		return nil, nil, newErrors.Wrap(err, "repo revisions of organization unit get all")
	}
	response := dto.ToRevisionsOfOrganizationUnitListResponseDTO(res)

	return response, total, nil
}

package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type JudgeNumberResolutionOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.JudgeNumberResolutionOrganizationUnit
}

func NewJudgeNumberResolutionOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.JudgeNumberResolutionOrganizationUnit) JudgeNumberResolutionOrganizationUnitService {
	return &JudgeNumberResolutionOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *JudgeNumberResolutionOrganizationUnitServiceImpl) CreateJudgeNumberResolutionOrganizationUnit(input dto.JudgeNumberResolutionOrganizationUnitDTO) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error) {
	data := input.ToJudgeNumberResolutionOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge number resolution organization unit insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge number resolution organization unit get")
	}

	res := dto.ToJudgeNumberResolutionOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *JudgeNumberResolutionOrganizationUnitServiceImpl) UpdateJudgeNumberResolutionOrganizationUnit(id int, input dto.JudgeNumberResolutionOrganizationUnitDTO) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error) {
	data := input.ToJudgeNumberResolutionOrganizationUnit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge number resolution organization unit update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge number resolution organization unit get")
	}

	response := dto.ToJudgeNumberResolutionOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *JudgeNumberResolutionOrganizationUnitServiceImpl) DeleteJudgeNumberResolutionOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo judge number resolution organization unit delete")
	}

	return nil
}

func (h *JudgeNumberResolutionOrganizationUnitServiceImpl) GetJudgeNumberResolutionOrganizationUnit(id int) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge number resolution organization unit get")
	}
	response := dto.ToJudgeNumberResolutionOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *JudgeNumberResolutionOrganizationUnitServiceImpl) GetJudgeNumberResolutionOrganizationUnitList(input dto.GetJudgeNumberResolutionOrganizationUnitInputDTO) ([]dto.JudgeNumberResolutionOrganizationUnitResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.ResolutionID != nil {
		cond["resolution_id"] = input.ResolutionID
	}
	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = input.OrganizationUnitID
	}
	data, total, err := h.repo.GetAll(input.Page, input.PageSize, &cond)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo judge number resolution organization unit get all")
	}

	response := dto.ToJudgeNumberResolutionOrganizationUnitListResponseDTO(data)

	return response, total, nil
}

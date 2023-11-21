package services

import (
	"fmt"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type OrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.OrganizationUnit
}

func NewOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.OrganizationUnit) OrganizationUnitService {
	return &OrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *OrganizationUnitServiceImpl) CreateOrganizationUnit(input dto.CreateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error) {
	data := input.ToOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *OrganizationUnitServiceImpl) UpdateOrganizationUnit(id int, input dto.UpdateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToOrganizationUnit(data)

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitServiceImpl) DeleteOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *OrganizationUnitServiceImpl) GetOrganizationUnit(id int) (*dto.OrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitServiceImpl) GetOrganizationUnitList(data dto.GetOrganizationUnitsDTO) ([]dto.OrganizationUnitResponseDTO, *uint64, error) {
	var conditionAndExp *up.AndExpr
	if data.Search != nil && *data.Search != "" {
		likeCondition := fmt.Sprintf("%%%s%%", *data.Search)
		conditionAndExp = up.And(
			up.Or(
				up.Cond{"title ILIKE": likeCondition},
				up.Cond{"description ILIKE": likeCondition},
				up.Cond{"abbreviation ILIKE": likeCondition},
				up.Cond{"address ILIKE": likeCondition},
			),
		)
	}

	if data.ParentID != nil && *data.ParentID > 0 {
		if conditionAndExp == nil {
			conditionAndExp = &up.AndExpr{}
		}
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"parent_id": *data.ParentID})
	}

	if data.IsParent != nil && *data.IsParent {
		if conditionAndExp == nil {
			conditionAndExp = &up.AndExpr{}
		}
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"parent_id IS NOT NULL": *data.IsParent})
	}

	res, total, err := h.repo.GetAll(data.Page, data.PageSize, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToOrganizationUnitListResponseDTO(res)

	return response, total, nil
}

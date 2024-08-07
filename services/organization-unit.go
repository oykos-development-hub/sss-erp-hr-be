package services

import (
	"context"
	"fmt"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *OrganizationUnitServiceImpl) CreateOrganizationUnit(ctx context.Context, input dto.CreateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error) {
	data := input.ToOrganizationUnit()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit get")
	}

	res := dto.ToOrganizationUnitResponseDTO(*data)

	return &res, nil
}

func (h *OrganizationUnitServiceImpl) UpdateOrganizationUnit(ctx context.Context, id int, input dto.UpdateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToOrganizationUnit(data)

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit update")
	}

	response := dto.ToOrganizationUnitResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitServiceImpl) DeleteOrganizationUnit(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo organization unit delete")
	}

	return nil
}

func (h *OrganizationUnitServiceImpl) GetOrganizationUnit(id int) (*dto.OrganizationUnitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit get")
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
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"parent_id": nil})
	}

	if data.Active == nil && *data.Active {
		if conditionAndExp == nil {
			conditionAndExp = &up.AndExpr{}
		}
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"active": true})
	} else if !*data.Active {
		if conditionAndExp == nil {
			conditionAndExp = &up.AndExpr{}
		}
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"active": false})
	}

	res, total, err := h.repo.GetAll(data.Page, data.PageSize, conditionAndExp)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo organization unit get all")
	}
	response := dto.ToOrganizationUnitListResponseDTO(res)

	return response, total, nil
}

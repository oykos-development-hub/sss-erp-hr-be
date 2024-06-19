package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type TenderTypeServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.TenderType
}

func NewTenderTypeServiceImpl(app *celeritas.Celeritas, repo data.TenderType) TenderTypeService {
	return &TenderTypeServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *TenderTypeServiceImpl) CreateTenderType(input dto.TenderTypeDTO) (*dto.TenderTypeResponseDTO, error) {
	data := input.ToTenderType()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo tender type insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo tender type get")
	}

	res := dto.ToTenderTypeResponseDTO(*data)

	return &res, nil
}

func (h *TenderTypeServiceImpl) UpdateTenderType(id int, input dto.TenderTypeDTO) (*dto.TenderTypeResponseDTO, error) {
	data := input.ToTenderType()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo tender type update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo tender type get")
	}

	response := dto.ToTenderTypeResponseDTO(*data)

	return &response, nil
}

func (h *TenderTypeServiceImpl) DeleteTenderType(id int) error {
	err := h.repo.Delete(id)
	if err != nil {

		return newErrors.Wrap(err, "repo tender type delete")
	}

	return nil
}

func (h *TenderTypeServiceImpl) GetTenderType(id int) (*dto.TenderTypeResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo tender type get")
	}
	response := dto.ToTenderTypeResponseDTO(*data)

	return &response, nil
}

func (h *TenderTypeServiceImpl) GetTenderTypeList(input dto.GetTenderTypeInputDTO) ([]dto.TenderTypeResponseDTO, error) {
	var cond []up.LogicalExpr
	var combinedCond *up.AndExpr
	if input.Search != nil {
		search := "%" + *input.Search + "%"
		isJudgeCond := up.Or(
			up.Cond{"title ILIKE": search},
			up.Cond{"abbreviation ILIKE": search},
			up.Cond{"description ILIKE": search},
		)
		cond = append(cond, isJudgeCond)
	}
	if len(cond) > 0 {
		combinedCond = up.And(cond...)
	}
	data, err := h.repo.GetAll(combinedCond)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo tender type get all")
	}
	response := dto.ToTenderTypeListResponseDTO(data)

	return response, nil
}

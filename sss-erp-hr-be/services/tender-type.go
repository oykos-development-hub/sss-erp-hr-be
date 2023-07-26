package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	"github.com/upper/db/v4"
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
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToTenderTypeResponseDTO(*data)

	return &res, nil
}

func (h *TenderTypeServiceImpl) UpdateTenderType(id int, input dto.TenderTypeDTO) (*dto.TenderTypeResponseDTO, error) {
	data := input.ToTenderType()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToTenderTypeResponseDTO(*data)

	return &response, nil
}

func (h *TenderTypeServiceImpl) DeleteTenderType(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *TenderTypeServiceImpl) GetTenderType(id int) (*dto.TenderTypeResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
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
			db.Cond{"title ILIKE": search},
			db.Cond{"abbreviation ILIKE": search},
			db.Cond{"description ILIKE": search},
		)
		cond = append(cond, isJudgeCond)
	}
	if len(cond) > 0 {
		combinedCond = up.And(cond...)
	}
	data, err := h.repo.GetAll(combinedCond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToTenderTypeListResponseDTO(data)

	return response, nil
}

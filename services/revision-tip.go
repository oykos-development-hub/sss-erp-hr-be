package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionTipServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RevisionTip
}

func NewRevisionTipServiceImpl(app *celeritas.Celeritas, repo data.RevisionTip) RevisionTipService {
	return &RevisionTipServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionTipServiceImpl) CreateRevisionTip(input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error) {
	data := input.ToRevisionTip()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRevisionTipResponseDTO(*data)

	return &res, nil
}

func (h *RevisionTipServiceImpl) UpdateRevisionTip(id int, input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error) {
	data := input.ToRevisionTip()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRevisionTipResponseDTO(*data)

	return &response, nil
}

func (h *RevisionTipServiceImpl) DeleteRevisionTip(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RevisionTipServiceImpl) GetRevisionTip(id int) (*dto.RevisionTipResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRevisionTipResponseDTO(*data)

	return &response, nil
}

func (h *RevisionTipServiceImpl) GetRevisionTipList(input dto.RevisionTipFilter) ([]dto.RevisionTipResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.RevisionID != nil {
		cond["revision_id"] = input.RevisionID
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToRevisionTipListResponseDTO(data)

	return response, total, nil
}

package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Revision
}

func NewRevisionServiceImpl(app *celeritas.Celeritas, repo data.Revision) RevisionService {
	return &RevisionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionServiceImpl) CreateRevision(input dto.RevisionDTO) (*dto.RevisionResponseDTO, error) {
	data := input.ToRevision()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRevisionResponseDTO(*data)

	return &res, nil
}

func (h *RevisionServiceImpl) UpdateRevision(id int, input dto.RevisionDTO) (*dto.RevisionResponseDTO, error) {
	data := input.ToRevision()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRevisionResponseDTO(*data)

	return &response, nil
}

func (h *RevisionServiceImpl) DeleteRevision(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RevisionServiceImpl) GetRevision(id int) (*dto.RevisionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRevisionResponseDTO(*data)

	return &response, nil
}

func (h *RevisionServiceImpl) GetRevisionList(input dto.RevisonFilter) ([]dto.RevisionResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.RevisionType != nil {
		cond["revison_type"] = input.RevisionType
	}
	if input.InternalRevisionSubject != nil {
		cond["internal_revision_subject"] = input.InternalRevisionSubject
	}
	if input.Revisor != nil {
		cond["revisor"] = input.Revisor
	}
	if input.PlanID != nil {
		cond["plan_id"] = input.PlanID
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToRevisionListResponseDTO(data)

	return response, total, nil
}

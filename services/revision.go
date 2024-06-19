package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *RevisionServiceImpl) CreateRevision(ctx context.Context, input dto.RevisionDTO) (*dto.RevisionResponseDTO, error) {
	data := input.ToRevision()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision get")
	}

	res := dto.ToRevisionResponseDTO(*data)

	return &res, nil
}

func (h *RevisionServiceImpl) UpdateRevision(ctx context.Context, id int, input dto.RevisionDTO) (*dto.RevisionResponseDTO, error) {
	data := input.ToRevision()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision get")
	}

	response := dto.ToRevisionResponseDTO(*data)

	return &response, nil
}

func (h *RevisionServiceImpl) DeleteRevision(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo revision delete")
	}

	return nil
}

func (h *RevisionServiceImpl) GetRevision(id int) (*dto.RevisionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision get")
	}
	response := dto.ToRevisionResponseDTO(*data)

	return &response, nil
}

func (h *RevisionServiceImpl) GetRevisionList(input dto.RevisonFilter) ([]dto.RevisionResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.RevisionType != nil {
		cond["revision_type_id"] = input.RevisionType
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

		return nil, nil, newErrors.Wrap(err, "repo revision get all")
	}
	response := dto.ToRevisionListResponseDTO(data)

	return response, total, nil
}

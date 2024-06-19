package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *RevisionTipServiceImpl) CreateRevisionTip(ctx context.Context, input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error) {
	data := input.ToRevisionTip()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision tip insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision tip get")
	}

	res := dto.ToRevisionTipResponseDTO(*data)

	return &res, nil
}

func (h *RevisionTipServiceImpl) UpdateRevisionTip(ctx context.Context, id int, input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error) {
	data := input.ToRevisionTip()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision tip update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision tip get")
	}

	response := dto.ToRevisionTipResponseDTO(*data)

	return &response, nil
}

func (h *RevisionTipServiceImpl) DeleteRevisionTip(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo revision tip delete")
	}

	return nil
}

func (h *RevisionTipServiceImpl) GetRevisionTip(id int) (*dto.RevisionTipResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo revision tip get")
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
		return nil, nil, newErrors.Wrap(err, "repo revision tip get all")
	}
	response := dto.ToRevisionTipListResponseDTO(data)

	return response, total, nil
}

package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
	datapkg "gitlab.sudovi.me/erp/hr-ms-api/data"
)

type RevisionTipImplementationServiceImpl struct {
	App     *celeritas.Celeritas
	repo    data.RevisionTipImplementation
	tipRepo data.RevisionTip
}

// zavrsi kad se izmijeni ili doda da se provjeri da li je status zavrsen, ako jeste onda update na zavrsen. isto treba staviti da kad se tip kreira da se defaulno setuje na nesprovedena

func NewRevisionTipImplementationServiceImpl(app *celeritas.Celeritas, repo data.RevisionTipImplementation, tipRepo data.RevisionTip) RevisionTipImplementationService {
	return &RevisionTipImplementationServiceImpl{
		App:     app,
		repo:    repo,
		tipRepo: tipRepo,
	}
}

func (h *RevisionTipImplementationServiceImpl) CreateRevisionTipImplementation(ctx context.Context, input dto.RevisionTipImplementationDTO) (*dto.RevisionTipImplementationResponseDTO, error) {
	data := input.ToRevisionTipImplementation()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}

	res := dto.ToRevisionTipImplementationResponseDTO(*data)

	if res.Status == datapkg.StatusTipImplFinished {
		tipData, err := h.tipRepo.Get(id)
		if err != nil {
			return nil, newErrors.Wrap(err, "tip repo get")
		}

		tipData.Status = datapkg.StatusTipImplFinished

		err = h.tipRepo.Update(ctx, *tipData)
		if err != nil {
			return nil, newErrors.Wrap(err, "tip repo update")
		}
	}

	return &res, nil
}

func (h *RevisionTipImplementationServiceImpl) UpdateRevisionTipImplementation(ctx context.Context, id int, input dto.RevisionTipImplementationDTO) (*dto.RevisionTipImplementationResponseDTO, error) {
	data := input.ToRevisionTipImplementation()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}

	res := dto.ToRevisionTipImplementationResponseDTO(*data)

	if res.Status == datapkg.StatusTipImplFinished {
		tipData, err := h.tipRepo.Get(id)
		if err != nil {
			return nil, newErrors.Wrap(err, "repo get")
		}

		tipData.Status = datapkg.StatusTipImplFinished

		err = h.tipRepo.Update(ctx, *tipData)
		if err != nil {
			return nil, newErrors.Wrap(err, "tip repo update")
		}
	}

	return &res, nil
}

func (h *RevisionTipImplementationServiceImpl) DeleteRevisionTipImplementation(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo delete")
	}

	return nil
}

func (h *RevisionTipImplementationServiceImpl) GetRevisionTipImplementation(id int) (*dto.RevisionTipImplementationResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}
	response := dto.ToRevisionTipImplementationResponseDTO(*data)

	return &response, nil
}

func (h *RevisionTipImplementationServiceImpl) GetRevisionTipImplementationList(input dto.RevisionTipImplementationsFilter) ([]dto.RevisionTipImplementationResponseDTO, error) {
	cond := up.Cond{}
	if input.TipID != nil {
		cond["tip_id"] = input.TipID
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get all")
	}
	response := dto.ToRevisionTipImplementationListResponseDTO(data)

	return response, nil
}

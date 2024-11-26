package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RevisionTipImplementationServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RevisionTipImplementation
}

func NewRevisionTipImplementationServiceImpl(app *celeritas.Celeritas, repo data.RevisionTipImplementation) RevisionTipImplementationService {
	return &RevisionTipImplementationServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RevisionTipImplementationServiceImpl) CreateRevisionTipImplementation(input dto.RevisionTipImplementationDTO) (*dto.RevisionTipImplementationResponseDTO, error) {
	data := input.ToRevisionTipImplementation()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}

	res := dto.ToRevisionTipImplementationResponseDTO(*data)

	return &res, nil
}

func (h *RevisionTipImplementationServiceImpl) UpdateRevisionTipImplementation(id int, input dto.RevisionTipImplementationDTO) (*dto.RevisionTipImplementationResponseDTO, error) {
	data := input.ToRevisionTipImplementation()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}

	response := dto.ToRevisionTipImplementationResponseDTO(*data)

	return &response, nil
}

func (h *RevisionTipImplementationServiceImpl) DeleteRevisionTipImplementation(id int) error {
	err := h.repo.Delete(id)
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

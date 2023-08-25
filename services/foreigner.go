package services

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ForeignerServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Foreigner
}

func NewForeignerServiceImpl(app *celeritas.Celeritas, repo data.Foreigner) ForeignerService {
	return &ForeignerServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ForeignerServiceImpl) CreateForeigner(input dto.ForeignerDTO) (*dto.ForeignerResponseDTO, error) {
	data := input.ToForeigner()
	data.CreatedAt = time.Now()
	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	res := dto.ToForeignerResponseDTO(*data)

	return &res, nil
}

func (h *ForeignerServiceImpl) UpdateForeigner(id int, input dto.ForeignerDTO) (*dto.ForeignerResponseDTO, error) {
	data := input.ToForeigner()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	response := dto.ToForeignerResponseDTO(*data)

	return &response, nil
}

func (h *ForeignerServiceImpl) DeleteForeigner(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ForeignerServiceImpl) GetForeigner(id int) (*dto.ForeignerResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToForeignerResponseDTO(*data)

	return &response, nil
}

func (h *ForeignerServiceImpl) GetForeignerList(id int) ([]dto.ForeignerResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": id,
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToForeignerListResponseDTO(data)

	return response, nil
}

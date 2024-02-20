package services

import (
	"fmt"
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type UserNormFulfilmentServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.UserNormFulfilment
}

func NewUserNormFulfilmentServiceImpl(app *celeritas.Celeritas, repo data.UserNormFulfilment) UserNormFulfilmentService {
	return &UserNormFulfilmentServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *UserNormFulfilmentServiceImpl) CreateUserNormFulfilment(input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error) {
	fmt.Println(input.DateOfEvaluation.String())
	data := input.ToUserNormFulfilment()
	fmt.Println(data.DateOfEvaluation.String())

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &res, nil
}

func (h *UserNormFulfilmentServiceImpl) UpdateUserNormFulfilment(id int, input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error) {
	data := input.ToUserNormFulfilment()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &response, nil
}

func (h *UserNormFulfilmentServiceImpl) DeleteUserNormFulfilment(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *UserNormFulfilmentServiceImpl) GetUserNormFulfilment(id int) (*dto.UserNormFulfilmentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &response, nil
}

func (h *UserNormFulfilmentServiceImpl) GetUserNormFulfilmentList(userProfileID int, input dto.GetUserNormFulfilmentListInput) ([]dto.UserNormFulfilmentResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": userProfileID,
	}
	if input.NormYear != nil {
		start := time.Date(*input.NormYear, time.January, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(*input.NormYear, time.December, 31, 23, 59, 59, 999999999, time.UTC)

		cond["norm_start_date"] = up.Gte(start)
		cond["norm_start_date"] = up.Lte(end)
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserNormFulfilmentListResponseDTO(data)

	return response, nil
}

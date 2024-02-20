package services

import (
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
	data := input.ToUserNormFulfilment()

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
	var (
		combinedCond *up.AndExpr
		conditions   []up.LogicalExpr
	)

	if input.NormYear != nil && *input.NormYear != 0 {
		startOfYear := time.Date(*input.NormYear, time.January, 1, 0, 0, 0, 0, time.UTC)
		endOfYear := time.Date(*input.NormYear, time.December, 31, 23, 59, 59, 999999999, time.UTC)

		dateCond := up.And(
			up.Cond{"norm_start_date >=": startOfYear},
			up.Cond{"norm_start_date <=": endOfYear},
		)
		conditions = append(conditions, dateCond)
	}

	cond := up.Cond{
		"user_profile_id": userProfileID,
	}

	conditions = append(conditions, cond)

	if len(conditions) > 0 {
		combinedCond = up.And(conditions...)
	}

	data, err := h.repo.GetAll(combinedCond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserNormFulfilmentListResponseDTO(data)

	return response, nil
}

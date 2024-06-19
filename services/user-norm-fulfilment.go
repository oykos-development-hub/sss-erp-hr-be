package services

import (
	"context"
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *UserNormFulfilmentServiceImpl) CreateUserNormFulfilment(ctx context.Context, input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error) {
	data := input.ToUserNormFulfilment()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user norm fulfilment create")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user norm fulfilment get")
	}

	res := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &res, nil
}

func (h *UserNormFulfilmentServiceImpl) UpdateUserNormFulfilment(ctx context.Context, id int, input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error) {
	data := input.ToUserNormFulfilment()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user norm fulfilment update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user norm fulfilment get")
	}

	response := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &response, nil
}

func (h *UserNormFulfilmentServiceImpl) DeleteUserNormFulfilment(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {

		return newErrors.Wrap(err, "repo user norm fulfilment delete")
	}

	return nil
}

func (h *UserNormFulfilmentServiceImpl) GetUserNormFulfilment(id int) (*dto.UserNormFulfilmentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo user norm fulfilment get")
	}
	response := dto.ToUserNormFulfilmentResponseDTO(*data)

	return &response, nil
}

func (h *UserNormFulfilmentServiceImpl) GetUserNormFulfilmentList(userProfileID int, input dto.GetUserNormFulfilmentListInput) ([]dto.UserNormFulfilmentResponseDTO, error) {
	conditionAndExp := &up.AndExpr{}

	if input.NormYear != nil && *input.NormYear > 0 {
		startOfYear := time.Date(*input.NormYear, time.January, 1, 0, 0, 0, 0, time.UTC)
		endOfYear := time.Date(*input.NormYear, time.December, 31, 23, 59, 59, 999999999, time.UTC)

		conditionAndExp = up.And(conditionAndExp, &up.Cond{"norm_start_date": up.Gte(startOfYear)})
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"norm_start_date": up.Lte(endOfYear)})
	}

	conditionAndExp = up.And(conditionAndExp, &up.Cond{"user_profile_id": userProfileID})

	data, err := h.repo.GetAll(conditionAndExp)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo user norm fulfilment get all")
	}

	response := dto.ToUserNormFulfilmentListResponseDTO(data)

	return response, nil
}

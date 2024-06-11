package services

import (
	"context"
	"fmt"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type JobPositionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.JobPosition
}

func NewJobPositionServiceImpl(app *celeritas.Celeritas, repo data.JobPosition) JobPositionService {
	return &JobPositionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *JobPositionServiceImpl) CreateJobPosition(ctx context.Context, input dto.CreateJobPositionDTO) (*dto.JobPositionResponseDTO, error) {
	data := input.ToJobPosition()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToJobPositionResponseDTO(*data)

	return &res, nil
}

func (h *JobPositionServiceImpl) UpdateJobPosition(ctx context.Context, id int, input dto.UpdateJobPositionDTO) (*dto.JobPositionResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToJobPosition(data)

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToJobPositionResponseDTO(*data)

	return &response, nil
}

func (h *JobPositionServiceImpl) DeleteJobPosition(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *JobPositionServiceImpl) GetJobPosition(id int) (*dto.JobPositionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToJobPositionResponseDTO(*data)

	return &response, nil
}

func (h *JobPositionServiceImpl) GetJobPositionList(data dto.GetJobPositionsDTO) ([]dto.JobPositionResponseDTO, *uint64, error) {
	var (
		combinedCond *up.AndExpr
		conditions   []up.LogicalExpr
	)

	if data.IsJudge != nil {
		if *data.IsJudge {
			isJudgeCond := up.Or(
				up.Cond{"is_judge": true},
				up.Cond{"is_judge_president": true},
			)
			conditions = append(conditions, isJudgeCond)
		} else {
			isJudgeCond := up.And(
				up.Cond{"is_judge": false},
				up.Cond{"is_judge_president": false},
			)
			conditions = append(conditions, isJudgeCond)
		}
	}

	if data.Search != nil && *data.Search != "" {
		likeCondition := fmt.Sprintf("%%%s%%", *data.Search)
		searchCond := up.Or(
			up.Cond{"title ILIKE": likeCondition},
			up.Cond{"description ILIKE": likeCondition},
			up.Cond{"abbreviation ILIKE": likeCondition},
			up.Cond{"serial_number ILIKE": likeCondition},
		)
		conditions = append(conditions, searchCond)
	}

	if len(conditions) > 0 {
		combinedCond = up.And(conditions...)
	}

	res, total, err := h.repo.GetAll(data.Page, data.PageSize, combinedCond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToJobPositionListResponseDTO(res)

	return response, total, nil
}

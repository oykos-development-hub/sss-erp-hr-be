package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EvaluationServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Evaluation
}

func NewEvaluationServiceImpl(app *celeritas.Celeritas, repo data.Evaluation) EvaluationService {
	return &EvaluationServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EvaluationServiceImpl) CreateEvaluation(ctx context.Context, input dto.EvaluationDTO) (*dto.EvaluationResponseDTO, error) {
	data := input.ToEvaluation()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation get")
	}

	res := dto.ToEvaluationResponseDTO(*data)

	return &res, nil
}

func (h *EvaluationServiceImpl) UpdateEvaluation(ctx context.Context, id int, input dto.EvaluationDTO) (*dto.EvaluationResponseDTO, error) {
	data, _ := h.repo.Get(id)
	createdAt := data.CreatedAt
	data = input.ToEvaluation()
	data.ID = id
	data.CreatedAt = createdAt

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation get")
	}

	response := dto.ToEvaluationResponseDTO(*data)

	return &response, nil
}

func (h *EvaluationServiceImpl) DeleteEvaluation(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo evaluation delete")
	}

	return nil
}

func (h *EvaluationServiceImpl) GetEvaluation(id int) (*dto.EvaluationResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation get")
	}
	response := dto.ToEvaluationResponseDTO(*data)

	return &response, nil
}

func (h *EvaluationServiceImpl) GetEmployeesEvaluationList(id int) ([]dto.EvaluationResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": id,
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation get all")
	}
	response := dto.ToEvaluationListResponseDTO(data)

	return response, nil
}

func (h *EvaluationServiceImpl) GetEvaluationList(input dto.GetEvaluationListInputDTO) ([]dto.EvaluationResponseDTO, error) {
	cond := up.Cond{}
	if input.IsJudge != nil && *input.IsJudge {
		cond["decision_number !="] = ""
		cond["decision_number !="] = "0"
		cond["decision_number IS NOT"] = nil
	}
	if input.Score != nil && *input.Score != "" {
		cond["score"] = *input.Score
	}
	if input.ReasonForEvaluation != nil && *input.ReasonForEvaluation != "" {
		cond["reason_for_evaluation"] = *input.ReasonForEvaluation
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo evaluation get all")
	}
	response := dto.ToEvaluationListResponseDTO(data)

	return response, nil
}

package dto

import (
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetEvaluationListInputDTO struct {
	IsJudge             *bool   `json:"is_judge"`
	Score               *string `json:"score"`
	ReasonForEvaluation *string `json:"reason_for_evaluation"`
}

type EvaluationDTO struct {
	UserProfileID       int        `json:"user_profile_id"  validate:"required"`
	EvaluationTypeID    int        `json:"evaluation_type_id"  validate:"required"`
	Score               string     `json:"score"`
	DateOfEvaluation    *time.Time `json:"date_of_evaluation"  validate:"required"`
	Evaluator           string     `json:"evaluator"`
	IsRelevant          *bool      `json:"is_relevant" validate:"required"`
	FileIDs             []int64    `json:"file_ids"  validate:"omitempty"`
	ReasonForEvaluation *string    `json:"reason_for_evaluation"`
	EvaluationPeriod    *string    `json:"evaluation_period"`
	DecisionNumber      *string    `json:"decision_number"`
}

type EvaluationResponseDTO struct {
	ID                  int        `json:"id"`
	UserProfileID       int        `json:"user_profile_id"`
	EvaluationTypeID    int        `json:"evaluation_type_id"`
	Score               string     `json:"score"`
	DateOfEvaluation    *time.Time `json:"date_of_evaluation"`
	Evaluator           string     `json:"evaluator"`
	IsRelevant          *bool      `json:"is_relevant"`
	ReasonForEvaluation *string    `json:"reason_for_evaluation"`
	EvaluationPeriod    *string    `json:"evaluation_period"`
	DecisionNumber      *string    `json:"decision_number"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	FileIDs             []int64    `json:"file_ids"`
}

func (dto EvaluationDTO) ToEvaluation() *data.Evaluation {
	return &data.Evaluation{
		UserProfileID:       dto.UserProfileID,
		EvaluationTypeID:    dto.EvaluationTypeID,
		Score:               dto.Score,
		DateOfEvaluation:    dto.DateOfEvaluation,
		Evaluator:           dto.Evaluator,
		IsRelevant:          dto.IsRelevant,
		FileIDs:             dto.FileIDs,
		ReasonForEvaluation: dto.ReasonForEvaluation,
		EvaluationPeriod:    dto.EvaluationPeriod,
		DecisionNumber:      dto.DecisionNumber,
	}
}

func ToEvaluationResponseDTO(data data.Evaluation) EvaluationResponseDTO {
	return EvaluationResponseDTO{
		ID:                  data.ID,
		UserProfileID:       data.UserProfileID,
		EvaluationTypeID:    data.EvaluationTypeID,
		Score:               data.Score,
		DateOfEvaluation:    data.DateOfEvaluation,
		Evaluator:           data.Evaluator,
		IsRelevant:          data.IsRelevant,
		ReasonForEvaluation: data.ReasonForEvaluation,
		EvaluationPeriod:    data.EvaluationPeriod,
		DecisionNumber:      data.DecisionNumber,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
		FileIDs:             data.FileIDs,
	}
}

func ToEvaluationListResponseDTO(evaluations []*data.Evaluation) []EvaluationResponseDTO {
	dtoList := make([]EvaluationResponseDTO, len(evaluations))
	for i, x := range evaluations {
		dtoList[i] = ToEvaluationResponseDTO(*x)
	}
	return dtoList
}

package dto

import (
	"fmt"
	"time"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
)

type GetUserNormFulfilmentListInput struct {
	NormYear *int `json:"norm_year"`
}

type UserNormFulfilmentDTO struct {
	UserProfileID            int       `json:"user_profile_id" validate:"required"`
	Topic                    string    `json:"topic" validate:"required"`
	Title                    string    `json:"title" validate:"required"`
	NumberOfNormDecrease     int       `json:"number_of_norm_decrease" validate:"required"`
	NumberOfItems            int       `json:"number_of_items" validate:"required"`
	NumberOfItemsSolved      int       `json:"number_of_items_solved" validate:"required"`
	EvaluationID             *int      `json:"evaluation_id"`
	DateOfEvaluation         time.Time `json:"date_of_evaluation"`
	DateOfEvaluationValidity time.Time `json:"date_of_evaluation_validity"`
	NormStartDate            string    `json:"norm_start_date"`
	NormEndDate              string    `json:"norm_end_date"`
	RelocationID             *int      `json:"relocation_id"`
	FileID                   *int      `json:"file_id"`
}

type UserNormFulfilmentResponseDTO struct {
	ID                         int       `json:"id"`
	UserProfileID              int       `json:"user_profile_id"`
	Topic                      string    `json:"topic"`
	Title                      string    `json:"title"`
	PercentageOfNormDecrease   float32   `json:"percentage_of_norm_decrease"`
	NumberOfNormDecrease       int       `json:"number_of_norm_decrease"`
	NumberOfItems              int       `json:"number_of_items"`
	PercentageOfNormFulfilment float32   `json:"percentage_of_norm_fulfilment"`
	NumberOfItemsSolved        int       `json:"number_of_items_solved"`
	EvaluationID               *int      `json:"evaluation_id"`
	DateOfEvaluation           time.Time `json:"date_of_evaluation"`
	DateOfEvaluationValidity   time.Time `json:"date_of_evaluation_validity"`
	NormStartDate              string    `json:"norm_start_date"`
	NormEndDate                string    `json:"norm_end_date"`
	RelocationID               *int      `json:"relocation_id"`
	FileID                     *int      `json:"file_id"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (dto UserNormFulfilmentDTO) ToUserNormFulfilment() *data.UserNormFulfilment {
	fmt.Println(dto.NormStartDate, dto.NormEndDate)
	return &data.UserNormFulfilment{
		UserProfileID:            dto.UserProfileID,
		Topic:                    dto.Topic,
		Title:                    dto.Title,
		NumberOfNormDecrease:     dto.NumberOfNormDecrease,
		NumberOfItems:            dto.NumberOfItems,
		NumberOfItemsSolved:      dto.NumberOfItemsSolved,
		NormEndDate:              dto.NormEndDate,
		NormStartDate:            dto.NormStartDate,
		EvaluationID:             dto.EvaluationID,
		DateOfEvaluation:         dto.DateOfEvaluation,
		DateOfEvaluationValidity: dto.DateOfEvaluationValidity,
		RelocationID:             dto.RelocationID,
		FileID:                   dto.FileID,
	}
}

func ToUserNormFulfilmentResponseDTO(data data.UserNormFulfilment) UserNormFulfilmentResponseDTO {
	percentageOfNormDecrease := float32(data.NumberOfNormDecrease) / float32(data.NumberOfItems) * 100
	percentageOfNormFulfilment := float32(data.NumberOfItemsSolved) / float32(data.NumberOfItems) * 100

	return UserNormFulfilmentResponseDTO{
		ID:                         data.ID,
		UserProfileID:              data.UserProfileID,
		Topic:                      data.Topic,
		Title:                      data.Title,
		PercentageOfNormDecrease:   percentageOfNormDecrease,
		NumberOfNormDecrease:       data.NumberOfNormDecrease,
		NumberOfItems:              data.NumberOfItems,
		PercentageOfNormFulfilment: percentageOfNormFulfilment,
		NumberOfItemsSolved:        data.NumberOfItemsSolved,
		EvaluationID:               data.EvaluationID,
		DateOfEvaluation:           data.DateOfEvaluation,
		DateOfEvaluationValidity:   data.DateOfEvaluationValidity,
		NormEndDate:                data.NormEndDate,
		NormStartDate:              data.NormStartDate,
		RelocationID:               data.RelocationID,
		FileID:                     data.FileID,
		CreatedAt:                  data.CreatedAt,
		UpdatedAt:                  data.UpdatedAt,
	}
}

func ToUserNormFulfilmentListResponseDTO(usernormfulfilments []*data.UserNormFulfilment) []UserNormFulfilmentResponseDTO {
	dtoList := make([]UserNormFulfilmentResponseDTO, len(usernormfulfilments))
	for i, x := range usernormfulfilments {
		dtoList[i] = ToUserNormFulfilmentResponseDTO(*x)
	}
	return dtoList
}

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

type SalaryServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Salary
}

func NewSalaryServiceImpl(app *celeritas.Celeritas, repo data.Salary) SalaryService {
	return &SalaryServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *SalaryServiceImpl) CreateSalary(ctx context.Context, input dto.SalaryDTO) (*dto.SalaryResponseDTO, error) {
	data := input.ToSalary()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo salary insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo salary get")
	}

	res := dto.ToSalaryResponseDTO(*data)

	return &res, nil
}

func (h *SalaryServiceImpl) UpdateSalary(ctx context.Context, id int, input dto.SalaryDTO) (*dto.SalaryResponseDTO, error) {
	curr, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo salary get")
	}

	data := input.ToSalary()
	data.ID = id
	data.UpdatedAt = time.Now()
	data.CreatedAt = curr.CreatedAt

	err = h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo salary update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo salary get")
	}

	response := dto.ToSalaryResponseDTO(*data)

	return &response, nil
}

func (h *SalaryServiceImpl) DeleteSalary(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {

		return newErrors.Wrap(err, "repo salary delete")
	}

	return nil
}

func (h *SalaryServiceImpl) GetSalary(id int) (*dto.SalaryResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo salary get")
	}
	response := dto.ToSalaryResponseDTO(*data)

	return &response, nil
}

func (h *SalaryServiceImpl) GetSalaryList(id int) ([]dto.SalaryResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": id,
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo salary get all")
	}
	response := dto.ToSalaryListResponseDTO(data)

	return response, nil
}

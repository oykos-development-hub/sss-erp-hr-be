package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeeExperienceServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeExperience
}

func NewEmployeeExperienceServiceImpl(app *celeritas.Celeritas, repo data.EmployeeExperience) EmployeeExperienceService {
	return &EmployeeExperienceServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeExperienceServiceImpl) CreateEmployeeExperience(input dto.EmployeeExperienceDTO) (*dto.EmployeeExperienceResponseDTO, error) {
	data := input.ToEmployeeExperience()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee experience insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee experience get")
	}

	res := dto.ToEmployeeExperienceResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeExperienceServiceImpl) UpdateEmployeeExperience(id int, input dto.EmployeeExperienceDTO) (*dto.EmployeeExperienceResponseDTO, error) {
	data := input.ToEmployeeExperience()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee experience update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee experience get")
	}

	response := dto.ToEmployeeExperienceResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeExperienceServiceImpl) DeleteEmployeeExperience(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee experience delete")
	}

	return nil
}

func (h *EmployeeExperienceServiceImpl) GetEmployeeExperienceList(userProfileID int) ([]dto.EmployeeExperienceResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": userProfileID,
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee experience get all")
	}
	response := dto.ToEmployeeExperienceListResponseDTO(data)

	return response, nil
}

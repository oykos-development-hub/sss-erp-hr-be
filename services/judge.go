package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type JudgeServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Judge
}

func NewJudgeServiceImpl(app *celeritas.Celeritas, repo data.Judge) JudgeService {
	return &JudgeServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *JudgeServiceImpl) CreateJudge(input dto.JudgeDTO) (*dto.JudgeResponseDTO, error) {
	data := input.ToJudge()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge get")
	}

	res := dto.ToJudgeResponseDTO(*data)

	return &res, nil
}

func (h *JudgeServiceImpl) UpdateJudge(id int, input dto.JudgeDTO) (*dto.JudgeResponseDTO, error) {
	data := input.ToJudge()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge get")
	}

	response := dto.ToJudgeResponseDTO(*data)

	return &response, nil
}

func (h *JudgeServiceImpl) DeleteJudge(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo judge delete")
	}

	return nil
}

func (h *JudgeServiceImpl) GetJudge(id int) (*dto.JudgeResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo judge get")
	}
	response := dto.ToJudgeResponseDTO(*data)

	return &response, nil
}

func (h *JudgeServiceImpl) GetJudgeList(input dto.JudgeFilter) ([]dto.JudgeResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.UserProfileID != nil {
		cond["user_profile_id"] = input.UserProfileID
	}
	if input.ResolutionID != nil {
		cond["resolution_id"] = input.ResolutionID
	}
	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = input.OrganizationUnitID
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo judge get all")
	}
	response := dto.ToJudgeListResponseDTO(data)

	return response, total, nil
}

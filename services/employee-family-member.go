package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeeFamilyMemberServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeeFamilyMember
}

func NewEmployeeFamilyMemberServiceImpl(app *celeritas.Celeritas, repo data.EmployeeFamilyMember) EmployeeFamilyMemberService {
	return &EmployeeFamilyMemberServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeeFamilyMemberServiceImpl) CreateEmployeeFamilyMember(input dto.EmployeeFamilyMemberDTO) (*dto.EmployeeFamilyMemberResponseDTO, error) {
	data := input.ToEmployeeFamilyMember()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member get")
	}

	res := dto.ToEmployeeFamilyMemberResponseDTO(*data)

	return &res, nil
}

func (h *EmployeeFamilyMemberServiceImpl) UpdateEmployeeFamilyMember(id int, input dto.EmployeeFamilyMemberDTO) (*dto.EmployeeFamilyMemberResponseDTO, error) {
	data := input.ToEmployeeFamilyMember()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member get")
	}

	response := dto.ToEmployeeFamilyMemberResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeFamilyMemberServiceImpl) DeleteEmployeeFamilyMember(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee family member delete")
	}

	return nil
}

func (h *EmployeeFamilyMemberServiceImpl) GetEmployeeFamilyMember(id int) (*dto.EmployeeFamilyMemberResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member get")
	}
	response := dto.ToEmployeeFamilyMemberResponseDTO(*data)

	return &response, nil
}

func (h *EmployeeFamilyMemberServiceImpl) GetEmployeeFamilyMemberList(id int) ([]dto.EmployeeFamilyMemberResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id": id,
	}
	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee family member get all")
	}
	response := dto.ToEmployeeFamilyMemberListResponseDTO(data)

	return response, nil
}

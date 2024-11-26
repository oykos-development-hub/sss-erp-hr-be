package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type EmployeesInOrganizationUnitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.EmployeesInOrganizationUnit
}

func NewEmployeesInOrganizationUnitServiceImpl(app *celeritas.Celeritas, repo data.EmployeesInOrganizationUnit) EmployeesInOrganizationUnitService {
	return &EmployeesInOrganizationUnitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *EmployeesInOrganizationUnitServiceImpl) CreateEmployeesInOrganizationUnit(input dto.EmployeesInOrganizationUnitDTO) (*dto.EmployeesInOrganizationUnitResponseDTO, error) {
	cond := up.Cond{
		"user_profile_id":                  input.UserProfileId,
		"position_in_organization_unit_id": input.PositionInOrganizationUnitId,
	}
	result, _ := h.repo.GetAll(&cond)
	if len(result) > 0 {
		return nil, newErrors.Wrap(errors.ErrResourceExists, "repo employee in organization unit get all")
	}

	data := input.ToEmployeesInOrganizationUnit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit insert")
	}

	data, _ = h.repo.Get(id)

	/*	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit get")
	}*/

	var response dto.EmployeesInOrganizationUnitResponseDTO

	if data != nil {
		response = dto.ToEmployeesInOrganizationUnitResponseDTO(*data)
	}
	return &response, nil
}

func (h *EmployeesInOrganizationUnitServiceImpl) DeleteEmployeesInOrganizationUnit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee in organization unit delete")
	}

	return nil
}

func (h *EmployeesInOrganizationUnitServiceImpl) DeleteEmployeesInOrganizationUnitByID(id int) error {
	err := h.repo.DeleteByID(id)
	if err != nil {
		return newErrors.Wrap(err, "repo employee in organization unit delete by id")
	}

	return nil
}

func (h *EmployeesInOrganizationUnitServiceImpl) GetEmployeesInOrganizationUnitByEmployee(id int) (*dto.EmployeesInOrganizationUnitResponseDTO, error) {

	data, err := h.repo.GetEmployeeInActiveSystematization(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit get all")
	}

	if data != nil {
		response := dto.ToEmployeesInOrganizationUnitResponseDTO(*data)
		return &response, nil
	}

	return nil, nil
}

func (h *EmployeesInOrganizationUnitServiceImpl) GetEmployeesInOrganizationUnitList(data dto.GetEmployeesInOrganizationUnitInput) ([]dto.EmployeesInOrganizationUnitResponseDTO, error) {
	condition := up.Cond{}
	if data.PositionInOrganizationUnit != nil {
		condition["position_in_organization_unit_id"] = data.PositionInOrganizationUnit
	}
	if data.Active != nil {
		condition["active"] = data.Active
	}
	res, err := h.repo.GetAll(&condition)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit get all")
	}
	response := dto.ToEmployeesInOrganizationUnitListResponseDTO(res)

	return response, nil
}

func (h *EmployeesInOrganizationUnitServiceImpl) UpdateEmployeesInOrganizationUnit(id int, input dto.EmployeesInOrganizationUnitDTO) (*dto.EmployeesInOrganizationUnitResponseDTO, error) {
	data := input.ToEmployeesInOrganizationUnit()
	data.ID = id

	cond := up.Cond{
		"user_profile_id":                  data.UserProfileId,
		"position_in_organization_unit_id": data.PositionInOrganizationUnitId,
	}
	result, _ := h.repo.GetAll(&cond)
	if len(result) > 0 && result[0].ID != id {
		return nil, newErrors.Wrap(errors.ErrResourceExists, "repo employee in organization unit get all")
	}

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit update")
	}

	data, _ = h.repo.Get(id)

	/*	if err != nil {
		return nil, newErrors.Wrap(err, "repo employee in organization unit get")
	}*/

	var response dto.EmployeesInOrganizationUnitResponseDTO

	if data != nil {
		response = dto.ToEmployeesInOrganizationUnitResponseDTO(*data)
	}

	return &response, nil
}

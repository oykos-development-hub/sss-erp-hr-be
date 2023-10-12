package handlers

import (
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"
	"gitlab.sudovi.me/erp/hr-ms-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// EmployeesInOrganizationUnitHandler is a concrete type that implements EmployeesInOrganizationUnitHandler
type employeesinorganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.EmployeesInOrganizationUnitService
}

// NewEmployeesInOrganizationUnitHandler initializes a new EmployeesInOrganizationUnitHandler with its dependencies
func NewEmployeesInOrganizationUnitHandler(app *celeritas.Celeritas, employeesinorganizationunitService services.EmployeesInOrganizationUnitService) EmployeesInOrganizationUnitHandler {
	return &employeesinorganizationunitHandlerImpl{
		App:     app,
		service: employeesinorganizationunitService,
	}
}

func (h *employeesinorganizationunitHandlerImpl) CreateEmployeesInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.EmployeesInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateEmployeesInOrganizationUnit(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeesInOrganizationUnit created successfuly", res)
}

func (h *employeesinorganizationunitHandlerImpl) DeleteEmployeesInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "position_in_organization_unit_id"))

	err := h.service.DeleteEmployeesInOrganizationUnit(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeesInOrganizationUnit deleted successfuly")
}

func (h *employeesinorganizationunitHandlerImpl) DeleteEmployeesInOrganizationUnitByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteEmployeesInOrganizationUnitByID(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeesInOrganizationUnit deleted successfuly")
}

func (h *employeesinorganizationunitHandlerImpl) GetEmployeesInOrganizationUnitByEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetEmployeesInOrganizationUnitByEmployee(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *employeesinorganizationunitHandlerImpl) GetEmployeesInOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetEmployeesInOrganizationUnitInput
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetEmployeesInOrganizationUnitList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *employeesinorganizationunitHandlerImpl) UpdateJobPositionInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.EmployeesInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateEmployeesInOrganizationUnit(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeesInOrganizationUnit updated successfuly", res)
}

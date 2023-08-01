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

// EmployeeAbsentHandler is a concrete type that implements EmployeeAbsentHandler
type employeeabsentHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.EmployeeAbsentService
}

// NewEmployeeAbsentHandler initializes a new EmployeeAbsentHandler with its dependencies
func NewEmployeeAbsentHandler(app *celeritas.Celeritas, employeeabsentService services.EmployeeAbsentService) EmployeeAbsentHandler {
	return &employeeabsentHandlerImpl{
		App:     app,
		service: employeeabsentService,
	}
}

func (h *employeeabsentHandlerImpl) CreateEmployeeAbsent(w http.ResponseWriter, r *http.Request) {
	var input dto.EmployeeAbsentDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateEmployeeAbsent(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeAbsent created successfuly", res)
}

func (h *employeeabsentHandlerImpl) UpdateEmployeeAbsent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.EmployeeAbsentDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateEmployeeAbsent(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeAbsent updated successfuly", res)
}

func (h *employeeabsentHandlerImpl) DeleteEmployeeAbsent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteEmployeeAbsent(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeeAbsent deleted successfuly")
}

func (h *employeeabsentHandlerImpl) GetEmployeeAbsentList(w http.ResponseWriter, r *http.Request) {
	userProfileId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.GetEmployeeAbsentsInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetEmployeeAbsentList(userProfileId, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *employeeabsentHandlerImpl) GetAbsentById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetAbsent(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

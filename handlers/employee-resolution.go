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

// EmployeeResolutionHandler is a concrete type that implements EmployeeResolutionHandler
type employeeresolutionHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.EmployeeResolutionService
}

// NewEmployeeResolutionHandler initializes a new EmployeeResolutionHandler with its dependencies
func NewEmployeeResolutionHandler(app *celeritas.Celeritas, employeeresolutionService services.EmployeeResolutionService) EmployeeResolutionHandler {
	return &employeeresolutionHandlerImpl{
		App:     app,
		service: employeeresolutionService,
	}
}

func (h *employeeresolutionHandlerImpl) CreateEmployeeResolution(w http.ResponseWriter, r *http.Request) {
	var input dto.EmployeeResolutionDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateEmployeeResolution(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeResolution created successfuly", res)
}

func (h *employeeresolutionHandlerImpl) UpdateEmployeeResolution(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.EmployeeResolutionDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateEmployeeResolution(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeResolution updated successfuly", res)
}

func (h *employeeresolutionHandlerImpl) DeleteEmployeeResolution(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteEmployeeResolution(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeeResolution deleted successfuly")
}

func (h *employeeresolutionHandlerImpl) GetEmployeeResolutionList(w http.ResponseWriter, r *http.Request) {
	userProfileID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.GetResolutionListInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetEmployeeResolutionList(userProfileID, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *employeeresolutionHandlerImpl) GetEmployeeResolution(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetEmployeeResolution(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

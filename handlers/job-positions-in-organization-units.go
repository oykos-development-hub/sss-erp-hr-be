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

// JobPositionsInOrganizationUnitsHandler is a concrete type that implements JobPositionHandler
type jobPositionsInOrganizationUnitsHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.JobPositionsInOrganizationUnitsService
}

// NewJobPositionHandler initializes a new JobPositionHandler with its dependencies
func NewJobPositionsInOrganizationUnitsHandler(app *celeritas.Celeritas, jobPositionsInOrganizationUnitsService services.JobPositionsInOrganizationUnitsService) JobPositionsInOrganizationUnitsHandler {
	return &jobPositionsInOrganizationUnitsHandlerImpl{
		App:     app,
		service: jobPositionsInOrganizationUnitsService,
	}
}

func (h *jobPositionsInOrganizationUnitsHandlerImpl) CreateJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateJobPositionsInOrganizationUnitsDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateJobPositionsInOrganizationUnits(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Job Position in organization unit created successfuly", res)
}

func (h *jobPositionsInOrganizationUnitsHandlerImpl) UpdateJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateJobPositionsInOrganizationUnitsDTO
	_ = h.App.ReadJSON(w, r, &input)

	input.Id, _ = strconv.Atoi(chi.URLParam(r, "id"))

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateJobPositionsInOrganizationUnits(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Job Position in organization unit updated successfuly", res)

}

func (h *jobPositionsInOrganizationUnitsHandlerImpl) DeleteJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteJobPositionsInOrganizationUnits(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Job Position in organization unit deleted successfuly")
}

func (h *jobPositionsInOrganizationUnitsHandlerImpl) GetJobPositionsInOrganizationUnitsById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetJobPositionInOrganziationUnitById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *jobPositionsInOrganizationUnitsHandlerImpl) GetJobPositionsInOrganizationUnitsList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJobPositionsInOrganizationUnitsDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetJobPositionsInOrganizationUnitsList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

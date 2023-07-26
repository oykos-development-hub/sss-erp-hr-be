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

// EmployeeExperienceHandler is a concrete type that implements EmployeeExperienceHandler
type employeeexperienceHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.EmployeeExperienceService
}

// NewEmployeeExperienceHandler initializes a new EmployeeExperienceHandler with its dependencies
func NewEmployeeExperienceHandler(app *celeritas.Celeritas, employeeexperienceService services.EmployeeExperienceService) EmployeeExperienceHandler {
	return &employeeexperienceHandlerImpl{
		App:     app,
		service: employeeexperienceService,
	}
}

func (h *employeeexperienceHandlerImpl) CreateEmployeeExperience(w http.ResponseWriter, r *http.Request) {
	var input dto.EmployeeExperienceDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateEmployeeExperience(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeExperience created successfuly", res)
}

func (h *employeeexperienceHandlerImpl) UpdateEmployeeExperience(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.EmployeeExperienceDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateEmployeeExperience(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeExperience updated successfuly", res)
}

func (h *employeeexperienceHandlerImpl) DeleteEmployeeExperience(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteEmployeeExperience(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeeExperience deleted successfuly")
}

func (h *employeeexperienceHandlerImpl) GetEmployeeExperienceList(w http.ResponseWriter, r *http.Request) {
	userProfileID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetEmployeeExperienceList(userProfileID)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

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

// RevisionTipImplementationHandler is a concrete type that implements RevisionTipImplementationHandler
type revisiontipimplementationHandlerImpl struct {
	App             *celeritas.Celeritas
	service         services.RevisionTipImplementationService
	errorLogService services.ErrorLogService
}

// NewRevisionTipImplementationHandler initializes a new RevisionTipImplementationHandler with its dependencies
func NewRevisionTipImplementationHandler(app *celeritas.Celeritas, revisiontipimplementationService services.RevisionTipImplementationService, errorLogService services.ErrorLogService) RevisionTipImplementationHandler {
	return &revisiontipimplementationHandlerImpl{
		App:             app,
		service:         revisiontipimplementationService,
		errorLogService: errorLogService,
	}
}

func (h *revisiontipimplementationHandlerImpl) CreateRevisionTipImplementation(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionTipImplementationDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateRevisionTipImplementation(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionTipImplementation created successfuly", res)
}

func (h *revisiontipimplementationHandlerImpl) UpdateRevisionTipImplementation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RevisionTipImplementationDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateRevisionTipImplementation(id, input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionTipImplementation updated successfuly", res)
}

func (h *revisiontipimplementationHandlerImpl) DeleteRevisionTipImplementation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRevisionTipImplementation(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RevisionTipImplementation deleted successfuly")
}

func (h *revisiontipimplementationHandlerImpl) GetRevisionTipImplementationById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRevisionTipImplementation(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *revisiontipimplementationHandlerImpl) GetRevisionTipImplementationList(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionTipImplementationsFilter
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetRevisionTipImplementationList(input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

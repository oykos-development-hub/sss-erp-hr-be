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

// RevisionsInOrganizationUnitHandler is a concrete type that implements RevisionsInOrganizationUnitHandler
type revisionsinorganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.RevisionsInOrganizationUnitService
}

// NewRevisionsInOrganizationUnitHandler initializes a new RevisionsInOrganizationUnitHandler with its dependencies
func NewRevisionsInOrganizationUnitHandler(app *celeritas.Celeritas, revisionsinorganizationunitService services.RevisionsInOrganizationUnitService) RevisionsInOrganizationUnitHandler {
	return &revisionsinorganizationunitHandlerImpl{
		App:     app,
		service: revisionsinorganizationunitService,
	}
}

func (h *revisionsinorganizationunitHandlerImpl) CreateRevisionsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionsInOrganizationUnitDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
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

	res, err := h.service.CreateRevisionsInOrganizationUnit(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionsInOrganizationUnit created successfuly", res)
}

func (h *revisionsinorganizationunitHandlerImpl) UpdateRevisionsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RevisionsInOrganizationUnitDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
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

	res, err := h.service.UpdateRevisionsInOrganizationUnit(id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionsInOrganizationUnit updated successfuly", res)
}

func (h *revisionsinorganizationunitHandlerImpl) DeleteRevisionsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRevisionsInOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RevisionsInOrganizationUnit deleted successfuly")
}

func (h *revisionsinorganizationunitHandlerImpl) GetRevisionsInOrganizationUnitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRevisionsInOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *revisionsinorganizationunitHandlerImpl) GetRevisionsInOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionOrgUnitFilter
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetRevisionsInOrganizationUnitList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

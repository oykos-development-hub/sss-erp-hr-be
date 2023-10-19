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

// RevisionRevisorHandler is a concrete type that implements RevisionRevisorHandler
type revisionrevisorHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.RevisionRevisorService
}

// NewRevisionRevisorHandler initializes a new RevisionRevisorHandler with its dependencies
func NewRevisionRevisorHandler(app *celeritas.Celeritas, revisionrevisorService services.RevisionRevisorService) RevisionRevisorHandler {
	return &revisionrevisorHandlerImpl{
		App:     app,
		service: revisionrevisorService,
	}
}

func (h *revisionrevisorHandlerImpl) CreateRevisionRevisor(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionRevisorDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateRevisionRevisor(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionRevisor created successfuly", res)
}

func (h *revisionrevisorHandlerImpl) UpdateRevisionRevisor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RevisionRevisorDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateRevisionRevisor(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionRevisor updated successfuly", res)
}

func (h *revisionrevisorHandlerImpl) DeleteRevisionRevisor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRevisionRevisor(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RevisionRevisor deleted successfuly")
}

func (h *revisionrevisorHandlerImpl) GetRevisionRevisorById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRevisionRevisor(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *revisionrevisorHandlerImpl) GetRevisionRevisorList(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionRevisorFilter
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetRevisionRevisorList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

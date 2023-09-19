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

// RevisionTipHandler is a concrete type that implements RevisionTipHandler
type revisiontipHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.RevisionTipService
}

// NewRevisionTipHandler initializes a new RevisionTipHandler with its dependencies
func NewRevisionTipHandler(app *celeritas.Celeritas, revisiontipService services.RevisionTipService) RevisionTipHandler {
	return &revisiontipHandlerImpl{
		App:     app,
		service: revisiontipService,
	}
}

func (h *revisiontipHandlerImpl) CreateRevisionTip(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionTipDTO
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

	res, err := h.service.CreateRevisionTip(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionTip created successfuly", res)
}

func (h *revisiontipHandlerImpl) UpdateRevisionTip(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RevisionTipDTO
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

	res, err := h.service.UpdateRevisionTip(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionTip updated successfuly", res)
}

func (h *revisiontipHandlerImpl) DeleteRevisionTip(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRevisionTip(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RevisionTip deleted successfuly")
}

func (h *revisiontipHandlerImpl) GetRevisionTipById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRevisionTip(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *revisiontipHandlerImpl) GetRevisionTipList(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionTipFilter
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetRevisionTipList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

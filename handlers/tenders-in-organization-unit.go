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

// TendersInOrganizationUnitHandler is a concrete type that implements TendersInOrganizationUnitHandler
type tendersinorganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.TendersInOrganizationUnitService
}

// NewTendersInOrganizationUnitHandler initializes a new TendersInOrganizationUnitHandler with its dependencies
func NewTendersInOrganizationUnitHandler(app *celeritas.Celeritas, tendersinorganizationunitService services.TendersInOrganizationUnitService) TendersInOrganizationUnitHandler {
	return &tendersinorganizationunitHandlerImpl{
		App:     app,
		service: tendersinorganizationunitService,
	}
}

func (h *tendersinorganizationunitHandlerImpl) CreateTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.TendersInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateTendersInOrganizationUnit(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "TendersInOrganizationUnit created successfuly", res)
}

func (h *tendersinorganizationunitHandlerImpl) UpdateTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.TendersInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateTendersInOrganizationUnit(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "TendersInOrganizationUnit updated successfuly", res)
}

func (h *tendersinorganizationunitHandlerImpl) DeleteTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteTendersInOrganizationUnit(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "TendersInOrganizationUnit deleted successfuly")
}

func (h *tendersinorganizationunitHandlerImpl) GetTendersInOrganizationUnitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetTendersInOrganizationUnit(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *tendersinorganizationunitHandlerImpl) GetTendersInOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetTendersInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetTendersInOrganizationUnitList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

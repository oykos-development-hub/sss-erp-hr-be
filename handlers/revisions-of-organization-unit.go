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

// RevisionsOfOrganizationUnitHandler is a concrete type that implements RevisionsOfOrganizationUnitHandler
type revisionsoforganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.RevisionsOfOrganizationUnitService
}

// NewRevisionsOfOrganizationUnitHandler initializes a new RevisionsOfOrganizationUnitHandler with its dependencies
func NewRevisionsOfOrganizationUnitHandler(app *celeritas.Celeritas, revisionsoforganizationunitService services.RevisionsOfOrganizationUnitService) RevisionsOfOrganizationUnitHandler {
	return &revisionsoforganizationunitHandlerImpl{
		App:     app,
		service: revisionsoforganizationunitService,
	}
}

func (h *revisionsoforganizationunitHandlerImpl) CreateRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.RevisionsOfOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateRevisionsOfOrganizationUnit(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionsOfOrganizationUnit created successfuly", res)
}

func (h *revisionsoforganizationunitHandlerImpl) UpdateRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RevisionsOfOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateRevisionsOfOrganizationUnit(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RevisionsOfOrganizationUnit updated successfuly", res)
}

func (h *revisionsoforganizationunitHandlerImpl) DeleteRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRevisionsOfOrganizationUnit(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RevisionsOfOrganizationUnit deleted successfuly")
}

func (h *revisionsoforganizationunitHandlerImpl) GetRevisionsOfOrganizationUnitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRevisionsOfOrganizationUnit(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *revisionsoforganizationunitHandlerImpl) GetRevisionsOfOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetRevisionsInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetRevisionsOfOrganizationUnitList(&input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

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

// TenderApplicationsInOrganizationUnitHandler is a concrete type that implements TenderApplicationsInOrganizationUnitHandler
type tenderapplicationsinorganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.TenderApplicationsInOrganizationUnitService
}

// NewTenderApplicationsInOrganizationUnitHandler initializes a new TenderApplicationsInOrganizationUnitHandler with its dependencies
func NewTenderApplicationsInOrganizationUnitHandler(app *celeritas.Celeritas, tenderapplicationsinorganizationunitService services.TenderApplicationsInOrganizationUnitService) TenderApplicationsInOrganizationUnitHandler {
	return &tenderapplicationsinorganizationunitHandlerImpl{
		App:     app,
		service: tenderapplicationsinorganizationunitService,
	}
}

func (h *tenderapplicationsinorganizationunitHandlerImpl) CreateTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.TenderApplicationsInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateTenderApplicationsInOrganizationUnit(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "TenderApplicationsInOrganizationUnit created successfuly", res)
}

func (h *tenderapplicationsinorganizationunitHandlerImpl) UpdateTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.TenderApplicationsInOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateTenderApplicationsInOrganizationUnit(id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "TenderApplicationsInOrganizationUnit updated successfuly", res)
}

func (h *tenderapplicationsinorganizationunitHandlerImpl) DeleteTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteTenderApplicationsInOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "TenderApplicationsInOrganizationUnit deleted successfuly")
}

func (h *tenderapplicationsinorganizationunitHandlerImpl) GetTenderApplicationsInOrganizationUnitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetTenderApplicationsInOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *tenderapplicationsinorganizationunitHandlerImpl) GetTenderApplicationsInOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetTenderApplicationsInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetTenderApplicationsInOrganizationUnitList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

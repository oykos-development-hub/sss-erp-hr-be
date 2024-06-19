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

// JudgeNumberResolutionOrganizationUnitHandler is a concrete type that implements JudgeNumberResolutionOrganizationUnitHandler
type judgenumberresolutionorganizationunitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.JudgeNumberResolutionOrganizationUnitService
}

// NewJudgeNumberResolutionOrganizationUnitHandler initializes a new JudgeNumberResolutionOrganizationUnitHandler with its dependencies
func NewJudgeNumberResolutionOrganizationUnitHandler(app *celeritas.Celeritas, judgenumberresolutionorganizationunitService services.JudgeNumberResolutionOrganizationUnitService) JudgeNumberResolutionOrganizationUnitHandler {
	return &judgenumberresolutionorganizationunitHandlerImpl{
		App:     app,
		service: judgenumberresolutionorganizationunitService,
	}
}

func (h *judgenumberresolutionorganizationunitHandlerImpl) CreateJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	var input dto.JudgeNumberResolutionOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateJudgeNumberResolutionOrganizationUnit(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "JudgeNumberResolutionOrganizationUnit created successfuly", res)
}

func (h *judgenumberresolutionorganizationunitHandlerImpl) UpdateJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.JudgeNumberResolutionOrganizationUnitDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateJudgeNumberResolutionOrganizationUnit(id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "JudgeNumberResolutionOrganizationUnit updated successfuly", res)
}

func (h *judgenumberresolutionorganizationunitHandlerImpl) DeleteJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteJudgeNumberResolutionOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "JudgeNumberResolutionOrganizationUnit deleted successfuly")
}

func (h *judgenumberresolutionorganizationunitHandlerImpl) GetJudgeNumberResolutionOrganizationUnitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetJudgeNumberResolutionOrganizationUnit(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *judgenumberresolutionorganizationunitHandlerImpl) GetJudgeNumberResolutionOrganizationUnitList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJudgeNumberResolutionOrganizationUnitInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetJudgeNumberResolutionOrganizationUnitList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

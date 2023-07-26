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

// EmployeeFamilyMemberHandler is a concrete type that implements EmployeeFamilyMemberHandler
type employeefamilymemberHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.EmployeeFamilyMemberService
}

// NewEmployeeFamilyMemberHandler initializes a new EmployeeFamilyMemberHandler with its dependencies
func NewEmployeeFamilyMemberHandler(app *celeritas.Celeritas, employeefamilymemberService services.EmployeeFamilyMemberService) EmployeeFamilyMemberHandler {
	return &employeefamilymemberHandlerImpl{
		App:     app,
		service: employeefamilymemberService,
	}
}

func (h *employeefamilymemberHandlerImpl) CreateEmployeeFamilyMember(w http.ResponseWriter, r *http.Request) {
	var input dto.EmployeeFamilyMemberDTO
	err := h.App.ReadJSON(w, r, &input)
	h.App.ErrorLog.Println(err)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateEmployeeFamilyMember(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeFamilyMember created successfuly", res)
}

func (h *employeefamilymemberHandlerImpl) UpdateEmployeeFamilyMember(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.EmployeeFamilyMemberDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateEmployeeFamilyMember(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeFamilyMember updated successfuly", res)
}

func (h *employeefamilymemberHandlerImpl) DeleteEmployeeFamilyMember(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteEmployeeFamilyMember(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "EmployeeFamilyMember deleted successfuly")
}

func (h *employeefamilymemberHandlerImpl) GetEmployeeFamilyMemberById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetEmployeeFamilyMember(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *employeefamilymemberHandlerImpl) GetEmployeeFamilyMemberList(w http.ResponseWriter, r *http.Request) {
	userProfileID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetEmployeeFamilyMemberList(userProfileID)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

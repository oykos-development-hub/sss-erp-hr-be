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

// UserProfileHandler is a concrete type that implements UserProfileHandler
type userprofileHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.UserProfileService
}

// NewUserProfileHandler initializes a new UserProfileHandler with its dependencies
func NewUserProfileHandler(app *celeritas.Celeritas, userprofileService services.UserProfileService) UserProfileHandler {
	return &userprofileHandlerImpl{
		App:     app,
		service: userprofileService,
	}
}

func (h *userprofileHandlerImpl) CreateUserProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UserProfileDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateUserProfile(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "UserProfile created successfuly", res)
}

func (h *userprofileHandlerImpl) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.UserProfileDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateUserProfile(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "UserProfile updated successfuly", res)
}

func (h *userprofileHandlerImpl) DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteUserProfile(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "UserProfile deleted successfuly")
}

func (h *userprofileHandlerImpl) GetUserProfileById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetUserProfile(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *userprofileHandlerImpl) GetUserProfileList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetProfilesInputDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetUserProfileList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

func (h *userprofileHandlerImpl) GetContracts(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var input dto.GetEmployeeContracts
	_ = h.App.ReadJSON(w, r, &input)

	res, err := h.service.GetContracts(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "EmployeeContract fetched successfuly", res)
}

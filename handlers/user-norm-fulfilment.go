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

// UserNormFulfilmentHandler is a concrete type that implements UserNormFulfilmentHandler
type usernormfulfilmentHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.UserNormFulfilmentService
}

// NewUserNormFulfilmentHandler initializes a new UserNormFulfilmentHandler with its dependencies
func NewUserNormFulfilmentHandler(app *celeritas.Celeritas, usernormfulfilmentService services.UserNormFulfilmentService) UserNormFulfilmentHandler {
	return &usernormfulfilmentHandlerImpl{
		App:     app,
		service: usernormfulfilmentService,
	}
}

func (h *usernormfulfilmentHandlerImpl) CreateUserNormFulfilment(w http.ResponseWriter, r *http.Request) {
	var input dto.UserNormFulfilmentDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateUserNormFulfilment(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "UserNormFulfilment created successfuly", res)
}

func (h *usernormfulfilmentHandlerImpl) UpdateUserNormFulfilment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.UserNormFulfilmentDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateUserNormFulfilment(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "UserNormFulfilment updated successfuly", res)
}

func (h *usernormfulfilmentHandlerImpl) DeleteUserNormFulfilment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteUserNormFulfilment(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "UserNormFulfilment deleted successfuly")
}

func (h *usernormfulfilmentHandlerImpl) GetUserNormFulfilmentById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetUserNormFulfilment(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *usernormfulfilmentHandlerImpl) GetUserNormFulfilmentList(w http.ResponseWriter, r *http.Request) {
	userProfileID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var input dto.GetUserNormFulfilmentListInput
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetUserNormFulfilmentList(userProfileID, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

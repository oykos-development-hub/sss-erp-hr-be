package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type UserProfileServiceImpl struct {
	App          *celeritas.Celeritas
	repo         data.UserProfile
	contractRepo data.EmployeeContract
}

func NewUserProfileServiceImpl(app *celeritas.Celeritas, repo data.UserProfile, contractRepo data.EmployeeContract) UserProfileService {
	return &UserProfileServiceImpl{
		App:          app,
		repo:         repo,
		contractRepo: contractRepo,
	}
}

func (h *UserProfileServiceImpl) CreateUserProfile(input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error) {
	userData := input.ToUserProfile()

	id, err := h.repo.Insert(*userData)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	user, err := h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToUserProfileResponseDTO(*user)

	return &res, nil
}

func (h *UserProfileServiceImpl) UpdateUserProfile(id int, input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error) {
	userData := input.ToUserProfile()
	userData.ID = id

	err := h.repo.Update(*userData)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	user, err := h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserProfileResponseDTO(*user)

	return &response, nil
}

func (h *UserProfileServiceImpl) DeleteUserProfile(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *UserProfileServiceImpl) GetUserProfile(id int) (*dto.UserProfileResponseDTO, error) {
	user, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserProfileResponseDTO(*user)

	return &response, nil
}

func (h *UserProfileServiceImpl) GetUserProfileList(data dto.GetProfilesInputDTO) ([]dto.UserProfileResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if data.IsRevisor != nil {
		cond["revisor_role"] = data.IsRevisor
	}
	if data.AccountID != nil {
		cond["user_account_id"] = data.AccountID
	}
	res, total, err := h.repo.GetAll(nil, nil, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}

	response := dto.ToUserProfileListResponseDTO(res)

	return response, total, nil
}

func (h *UserProfileServiceImpl) GetContracts(id int, input dto.GetEmployeeContracts) ([]dto.EmployeeContractResponseDTO, error) {
	cond := up.Cond{}
	if input.Active != nil {
		cond["active"] = *input.Active
	}
	contracts, err := h.contractRepo.GetByUserProfileId(id, &cond)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToEmployeeContractListResponseDTO(contracts)

	return response, nil
}

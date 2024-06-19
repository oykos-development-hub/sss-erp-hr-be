package services

import (
	"context"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"

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

func (h *UserProfileServiceImpl) CreateUserProfile(ctx context.Context, input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error) {
	userData := input.ToUserProfile()

	id, err := h.repo.Insert(ctx, *userData)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user profile create")
	}

	user, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user profile get")
	}

	res := dto.ToUserProfileResponseDTO(*user)

	return &res, nil
}

func (h *UserProfileServiceImpl) UpdateUserProfile(ctx context.Context, id int, input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error) {
	userData := input.ToUserProfile()
	userData.ID = id

	err := h.repo.Update(ctx, *userData)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user profile update")
	}
	user, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user profile get")
	}

	response := dto.ToUserProfileResponseDTO(*user)

	return &response, nil
}

func (h *UserProfileServiceImpl) DeleteUserProfile(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {

		return newErrors.Wrap(err, "repo user profile delete")
	}

	return nil
}

func (h *UserProfileServiceImpl) GetUserProfile(id int) (*dto.UserProfileResponseDTO, error) {
	user, err := h.repo.Get(id)
	if err != nil {

		return nil, newErrors.Wrap(err, "repo user profile get")
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

		return nil, nil, newErrors.Wrap(err, "repo user profile get all")
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
		return nil, newErrors.Wrap(err, "repo employee contracts get")
	}

	response := dto.ToEmployeeContractListResponseDTO(contracts)

	return response, nil
}

func (h *UserProfileServiceImpl) GetRevisors() ([]*data.Revisor, error) {
	revisors, err := h.repo.GetRevisors()
	if err != nil {
		return nil, newErrors.Wrap(err, "repo user profile get revisors")
	}

	return revisors, nil
}

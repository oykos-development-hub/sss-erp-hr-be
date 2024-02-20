package services

import (
	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type OrganizationUnitService interface {
	CreateOrganizationUnit(input dto.CreateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error)
	UpdateOrganizationUnit(id int, input dto.UpdateOrganizationUnitDTO) (*dto.OrganizationUnitResponseDTO, error)
	DeleteOrganizationUnit(id int) error
	GetOrganizationUnit(id int) (*dto.OrganizationUnitResponseDTO, error)
	GetOrganizationUnitList(data dto.GetOrganizationUnitsDTO) ([]dto.OrganizationUnitResponseDTO, *uint64, error)
}

type SystematizationService interface {
	CreateSystematization(input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error)
	UpdateSystematization(id int, input dto.SystematizationDTO) (*dto.SystematizationResponseDTO, error)
	DeleteSystematization(id int) error
	GetSystematization(id int) (*dto.SystematizationResponseDTO, error)
	GetSystematizationList(data dto.GetSystematizationsDTO) ([]dto.SystematizationResponseDTO, *uint64, error)
}

type JobPositionService interface {
	CreateJobPosition(input dto.CreateJobPositionDTO) (*dto.JobPositionResponseDTO, error)
	UpdateJobPosition(id int, input dto.UpdateJobPositionDTO) (*dto.JobPositionResponseDTO, error)
	DeleteJobPosition(id int) error
	GetJobPosition(id int) (*dto.JobPositionResponseDTO, error)
	GetJobPositionList(data dto.GetJobPositionsDTO) ([]dto.JobPositionResponseDTO, *uint64, error)
}

type JobPositionsInOrganizationUnitsService interface {
	CreateJobPositionsInOrganizationUnits(input dto.CreateJobPositionsInOrganizationUnitsDTO) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error)
	UpdateJobPositionsInOrganizationUnits(input dto.CreateJobPositionsInOrganizationUnitsDTO) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error)
	DeleteJobPositionsInOrganizationUnits(id int) error
	GetJobPositionInOrganziationUnitById(id int) (*dto.JobPositionsInOrganizationUnitsResponseDTO, error)
	GetJobPositionsInOrganizationUnitsList(data dto.GetJobPositionsInOrganizationUnitsDTO) ([]dto.JobPositionsInOrganizationUnitsResponseDTO, *uint64, error)
}

type UserProfileService interface {
	CreateUserProfile(input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error)
	UpdateUserProfile(id int, input dto.UserProfileDTO) (*dto.UserProfileResponseDTO, error)
	DeleteUserProfile(id int) error
	GetUserProfile(id int) (*dto.UserProfileResponseDTO, error)
	GetUserProfileList(data dto.GetProfilesInputDTO) ([]dto.UserProfileResponseDTO, *uint64, error)
	GetContracts(id int, input dto.GetEmployeeContracts) ([]dto.EmployeeContractResponseDTO, error)
	GetRevisors() ([]*data.Revisor, error)
}

type EmployeeContractService interface {
	CreateEmployeeContract(input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error)
	UpdateEmployeeContract(id int, input dto.EmployeeContractDTO) (*dto.EmployeeContractResponseDTO, error)
	DeleteEmployeeContract(id int) error
}

type EmployeesInOrganizationUnitService interface {
	CreateEmployeesInOrganizationUnit(input dto.EmployeesInOrganizationUnitDTO) (*dto.EmployeesInOrganizationUnitResponseDTO, error)
	DeleteEmployeesInOrganizationUnit(id int) error
	DeleteEmployeesInOrganizationUnitByID(id int) error
	GetEmployeesInOrganizationUnitByEmployee(id int) (*dto.EmployeesInOrganizationUnitResponseDTO, error)
	GetEmployeesInOrganizationUnitList(dto.GetEmployeesInOrganizationUnitInput) ([]dto.EmployeesInOrganizationUnitResponseDTO, error)
	UpdateEmployeesInOrganizationUnit(id int, input dto.EmployeesInOrganizationUnitDTO) (*dto.EmployeesInOrganizationUnitResponseDTO, error)
}

type EmployeeExperienceService interface {
	CreateEmployeeExperience(input dto.EmployeeExperienceDTO) (*dto.EmployeeExperienceResponseDTO, error)
	UpdateEmployeeExperience(id int, input dto.EmployeeExperienceDTO) (*dto.EmployeeExperienceResponseDTO, error)
	DeleteEmployeeExperience(id int) error
	GetEmployeeExperienceList(userProfileID int) ([]dto.EmployeeExperienceResponseDTO, error)
}

type EmployeeEducationService interface {
	CreateEmployeeEducation(input dto.EmployeeEducationDTO) (*dto.EmployeeEducationResponseDTO, error)
	UpdateEmployeeEducation(id int, input dto.EmployeeEducationDTO) (*dto.EmployeeEducationResponseDTO, error)
	DeleteEmployeeEducation(id int) error
	GetEmployeeEducationList(input dto.EducationInput) ([]dto.EmployeeEducationResponseDTO, error)
}

type EmployeeResolutionService interface {
	CreateEmployeeResolution(input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error)
	UpdateEmployeeResolution(id int, input dto.EmployeeResolutionDTO) (*dto.EmployeeResolutionResponseDTO, error)
	DeleteEmployeeResolution(id int) error
	GetEmployeeResolutionList(userProfileID int, input dto.GetResolutionListInputDTO) ([]dto.EmployeeResolutionResponseDTO, error)
	GetEmployeeResolution(id int) (*dto.EmployeeResolutionResponseDTO, error)
}

type AbsentTypeService interface {
	CreateAbsentType(input dto.AbsentTypeDTO) (*dto.AbsentTypeResponseDTO, error)
	UpdateAbsentType(id int, input dto.AbsentTypeDTO) (*dto.AbsentTypeResponseDTO, error)
	DeleteAbsentType(id int) error
	GetAbsentType(id int) (*dto.AbsentTypeResponseDTO, error)
	GetAbsentTypeList(dto.GetAbesntTypeDTO) ([]dto.AbsentTypeResponseDTO, *uint64, error)
}

type EvaluationService interface {
	CreateEvaluation(input dto.EvaluationDTO) (*dto.EvaluationResponseDTO, error)
	UpdateEvaluation(id int, input dto.EvaluationDTO) (*dto.EvaluationResponseDTO, error)
	DeleteEvaluation(id int) error
	GetEvaluation(id int) (*dto.EvaluationResponseDTO, error)
	GetEmployeesEvaluationList(id int) ([]dto.EvaluationResponseDTO, error)
	GetEvaluationList(dto.GetEvaluationListInputDTO) ([]dto.EvaluationResponseDTO, error)
}

type ForeignerService interface {
	CreateForeigner(input dto.ForeignerDTO) (*dto.ForeignerResponseDTO, error)
	UpdateForeigner(id int, input dto.ForeignerDTO) (*dto.ForeignerResponseDTO, error)
	DeleteForeigner(id int) error
	GetForeigner(id int) (*dto.ForeignerResponseDTO, error)
	GetForeignerList(id int) ([]dto.ForeignerResponseDTO, error)
}

type EmployeeFamilyMemberService interface {
	CreateEmployeeFamilyMember(input dto.EmployeeFamilyMemberDTO) (*dto.EmployeeFamilyMemberResponseDTO, error)
	UpdateEmployeeFamilyMember(id int, input dto.EmployeeFamilyMemberDTO) (*dto.EmployeeFamilyMemberResponseDTO, error)
	DeleteEmployeeFamilyMember(id int) error
	GetEmployeeFamilyMember(id int) (*dto.EmployeeFamilyMemberResponseDTO, error)
	GetEmployeeFamilyMemberList(id int) ([]dto.EmployeeFamilyMemberResponseDTO, error)
}

type SalaryService interface {
	CreateSalary(input dto.SalaryDTO) (*dto.SalaryResponseDTO, error)
	UpdateSalary(id int, input dto.SalaryDTO) (*dto.SalaryResponseDTO, error)
	DeleteSalary(id int) error
	GetSalary(id int) (*dto.SalaryResponseDTO, error)
	GetSalaryList(id int) ([]dto.SalaryResponseDTO, error)
}

type UserNormFulfilmentService interface {
	CreateUserNormFulfilment(input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error)
	UpdateUserNormFulfilment(id int, input dto.UserNormFulfilmentDTO) (*dto.UserNormFulfilmentResponseDTO, error)
	DeleteUserNormFulfilment(id int) error
	GetUserNormFulfilment(id int) (*dto.UserNormFulfilmentResponseDTO, error)
	GetUserNormFulfilmentList(userProfileID int, input dto.GetUserNormFulfilmentListInput) ([]dto.UserNormFulfilmentResponseDTO, error)
}

type EmployeeAbsentService interface {
	CreateEmployeeAbsent(input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error)
	UpdateEmployeeAbsent(id int, input dto.EmployeeAbsentDTO) (*dto.EmployeeAbsentResponseDTO, error)
	DeleteEmployeeAbsent(id int) error
	GetAbsent(id int) (*dto.EmployeeAbsentResponseDTO, error)
	GetEmployeeAbsentList(userProfileID int, input dto.GetEmployeeAbsentsInputDTO) ([]dto.EmployeeAbsentResponseDTO, error)
}

type RevisionsOfOrganizationUnitService interface {
	CreateRevisionsOfOrganizationUnit(input dto.RevisionsOfOrganizationUnitDTO) (*dto.RevisionsOfOrganizationUnitResponseDTO, error)
	UpdateRevisionsOfOrganizationUnit(id int, input dto.RevisionsOfOrganizationUnitDTO) (*dto.RevisionsOfOrganizationUnitResponseDTO, error)
	DeleteRevisionsOfOrganizationUnit(id int) error
	GetRevisionsOfOrganizationUnit(id int) (*dto.RevisionsOfOrganizationUnitResponseDTO, error)
	GetRevisionsOfOrganizationUnitList(input *dto.GetRevisionsInputDTO) ([]dto.RevisionsOfOrganizationUnitResponseDTO, *uint64, error)
}

type TendersInOrganizationUnitService interface {
	CreateTendersInOrganizationUnit(input dto.TendersInOrganizationUnitDTO) (*dto.TendersInOrganizationUnitResponseDTO, error)
	UpdateTendersInOrganizationUnit(id int, input dto.TendersInOrganizationUnitDTO) (*dto.TendersInOrganizationUnitResponseDTO, error)
	DeleteTendersInOrganizationUnit(id int) error
	GetTendersInOrganizationUnit(id int) (*dto.TendersInOrganizationUnitResponseDTO, error)
	GetTendersInOrganizationUnitList(input dto.GetTendersInputDTO) ([]dto.TendersInOrganizationUnitResponseDTO, *uint64, error)
}

type TenderApplicationsInOrganizationUnitService interface {
	CreateTenderApplicationsInOrganizationUnit(input dto.TenderApplicationsInOrganizationUnitDTO) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error)
	UpdateTenderApplicationsInOrganizationUnit(id int, input dto.TenderApplicationsInOrganizationUnitDTO) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error)
	DeleteTenderApplicationsInOrganizationUnit(id int) error
	GetTenderApplicationsInOrganizationUnit(id int) (*dto.TenderApplicationsInOrganizationUnitResponseDTO, error)
	GetTenderApplicationsInOrganizationUnitList(input dto.GetTenderApplicationsInputDTO) ([]dto.TenderApplicationsInOrganizationUnitResponseDTO, *uint64, error)
}

type TenderTypeService interface {
	CreateTenderType(input dto.TenderTypeDTO) (*dto.TenderTypeResponseDTO, error)
	UpdateTenderType(id int, input dto.TenderTypeDTO) (*dto.TenderTypeResponseDTO, error)
	DeleteTenderType(id int) error
	GetTenderType(id int) (*dto.TenderTypeResponseDTO, error)
	GetTenderTypeList(input dto.GetTenderTypeInputDTO) ([]dto.TenderTypeResponseDTO, error)
}

type JudgeNumberResolutionService interface {
	CreateJudgeNumberResolution(input dto.JudgeNumberResolutionDTO) (*dto.JudgeNumberResolutionResponseDTO, error)
	UpdateJudgeNumberResolution(id int, input dto.JudgeNumberResolutionDTO) (*dto.JudgeNumberResolutionResponseDTO, error)
	DeleteJudgeNumberResolution(id int) error
	GetJudgeNumberResolution(id int) (*dto.JudgeNumberResolutionResponseDTO, error)
	GetJudgeNumberResolutionList(input dto.GetJudgeNumberResolutionInputDTO) ([]dto.JudgeNumberResolutionResponseDTO, *uint64, error)
}

type JudgeNumberResolutionOrganizationUnitService interface {
	CreateJudgeNumberResolutionOrganizationUnit(input dto.JudgeNumberResolutionOrganizationUnitDTO) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error)
	UpdateJudgeNumberResolutionOrganizationUnit(id int, input dto.JudgeNumberResolutionOrganizationUnitDTO) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error)
	DeleteJudgeNumberResolutionOrganizationUnit(id int) error
	GetJudgeNumberResolutionOrganizationUnit(id int) (*dto.JudgeNumberResolutionOrganizationUnitResponseDTO, error)
	GetJudgeNumberResolutionOrganizationUnitList(input dto.GetJudgeNumberResolutionOrganizationUnitInputDTO) ([]dto.JudgeNumberResolutionOrganizationUnitResponseDTO, *uint64, error)
}

type PlanService interface {
	CreatePlan(input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	UpdatePlan(id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	DeletePlan(id int) error
	GetPlan(id int) (*dto.PlanResponseDTO, error)
	GetPlanList(input dto.PlanFilter) ([]dto.PlanResponseDTO, *uint64, error)
}

type RevisionService interface {
	CreateRevision(input dto.RevisionDTO) (*dto.RevisionResponseDTO, error)
	UpdateRevision(id int, input dto.RevisionDTO) (*dto.RevisionResponseDTO, error)
	DeleteRevision(id int) error
	GetRevision(id int) (*dto.RevisionResponseDTO, error)
	GetRevisionList(filter dto.RevisonFilter) ([]dto.RevisionResponseDTO, *uint64, error)
}

type RevisionTipService interface {
	CreateRevisionTip(input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error)
	UpdateRevisionTip(id int, input dto.RevisionTipDTO) (*dto.RevisionTipResponseDTO, error)
	DeleteRevisionTip(id int) error
	GetRevisionTip(id int) (*dto.RevisionTipResponseDTO, error)
	GetRevisionTipList(input dto.RevisionTipFilter) ([]dto.RevisionTipResponseDTO, *uint64, error)
}

type JudgeService interface {
	CreateJudge(input dto.JudgeDTO) (*dto.JudgeResponseDTO, error)
	UpdateJudge(id int, input dto.JudgeDTO) (*dto.JudgeResponseDTO, error)
	DeleteJudge(id int) error
	GetJudge(id int) (*dto.JudgeResponseDTO, error)
	GetJudgeList(input dto.JudgeFilter) ([]dto.JudgeResponseDTO, *uint64, error)
}

type RevisionsInOrganizationUnitService interface {
	CreateRevisionsInOrganizationUnit(input dto.RevisionsInOrganizationUnitDTO) (*dto.RevisionsInOrganizationUnitResponseDTO, error)
	UpdateRevisionsInOrganizationUnit(id int, input dto.RevisionsInOrganizationUnitDTO) (*dto.RevisionsInOrganizationUnitResponseDTO, error)
	DeleteRevisionsInOrganizationUnit(id int) error
	GetRevisionsInOrganizationUnit(id int) (*dto.RevisionsInOrganizationUnitResponseDTO, error)
	GetRevisionsInOrganizationUnitList(input dto.RevisionOrgUnitFilter) ([]dto.RevisionsInOrganizationUnitResponseDTO, error)
}

type RevisionRevisorService interface {
	CreateRevisionRevisor(input dto.RevisionRevisorDTO) (*dto.RevisionRevisorResponseDTO, error)
	UpdateRevisionRevisor(id int, input dto.RevisionRevisorDTO) (*dto.RevisionRevisorResponseDTO, error)
	DeleteRevisionRevisor(id int) error
	GetRevisionRevisor(id int) (*dto.RevisionRevisorResponseDTO, error)
	GetRevisionRevisorList(input dto.RevisionRevisorFilter) ([]dto.RevisionRevisorResponseDTO, error)
}

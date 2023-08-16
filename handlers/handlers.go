package handlers

import (
	"net/http"
)

type Handlers struct {
	OrganizationUnitHandler                      OrganizationUnitHandler
	JobPositionHandler                           JobPositionHandler
	JobPositionsInOrganizationUnitsHandler       JobPositionsInOrganizationUnitsHandler
	SystematizationHandler                       SystematizationHandler
	UserProfileHandler                           UserProfileHandler
	EmployeeContractHandler                      EmployeeContractHandler
	EmployeesInOrganizationUnitHandler           EmployeesInOrganizationUnitHandler
	EmployeeExperienceHandler                    EmployeeExperienceHandler
	EmployeeEducationHandler                     EmployeeEducationHandler
	EmployeeResolutionHandler                    EmployeeResolutionHandler
	AbsentTypeHandler                            AbsentTypeHandler
	EvaluationHandler                            EvaluationHandler
	ForeignerHandler                             ForeignerHandler
	EmployeeFamilyMemberHandler                  EmployeeFamilyMemberHandler
	SalaryHandler                                SalaryHandler
	UserNormFulfilmentHandler                    UserNormFulfilmentHandler
	EmployeeAbsentHandler                        EmployeeAbsentHandler
	RevisionsOfOrganizationUnitHandler           RevisionsOfOrganizationUnitHandler
	TendersInOrganizationUnitHandler             TendersInOrganizationUnitHandler
	TenderApplicationsInOrganizationUnitHandler  TenderApplicationsInOrganizationUnitHandler
	TenderTypeHandler                            TenderTypeHandler
	JudgeNumberResolutionHandler                 JudgeNumberResolutionHandler
	JudgeNumberResolutionOrganizationUnitHandler JudgeNumberResolutionOrganizationUnitHandler
}

type EngagementTypeHandler interface {
	CreateEngagementType(w http.ResponseWriter, r *http.Request)
	UpdateEngagementType(w http.ResponseWriter, r *http.Request)
	DeleteEngagementType(w http.ResponseWriter, r *http.Request)
	GetEngagementTypeById(w http.ResponseWriter, r *http.Request)
	GetEngagementTypeList(w http.ResponseWriter, r *http.Request)
}

type EmployeeContractTypeHandler interface {
	CreateEmployeeContractType(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeContractType(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeContractType(w http.ResponseWriter, r *http.Request)
	GetEmployeeContractTypeById(w http.ResponseWriter, r *http.Request)
	GetEmployeeContractTypeList(w http.ResponseWriter, r *http.Request)
}

type OrganizationUnitHandler interface {
	CreateOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateOrganizationUnit(w http.ResponseWriter, r *http.Request)
	DeleteOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitById(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitList(w http.ResponseWriter, r *http.Request)
}

type JobPositionHandler interface {
	CreateJobPosition(w http.ResponseWriter, r *http.Request)
	UpdateJobPosition(w http.ResponseWriter, r *http.Request)
	DeleteJobPosition(w http.ResponseWriter, r *http.Request)
	GetJobPositionById(w http.ResponseWriter, r *http.Request)
	GetJobPositionList(w http.ResponseWriter, r *http.Request)
}

type JobPositionsInOrganizationUnitsHandler interface {
	CreateJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request)
	DeleteJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request)
	UpdateJobPositionsInOrganizationUnits(w http.ResponseWriter, r *http.Request)
	GetJobPositionsInOrganizationUnitsById(w http.ResponseWriter, r *http.Request)
	GetJobPositionsInOrganizationUnitsList(w http.ResponseWriter, r *http.Request)
}

type SystematizationHandler interface {
	CreateSystematization(w http.ResponseWriter, r *http.Request)
	UpdateSystematization(w http.ResponseWriter, r *http.Request)
	DeleteSystematization(w http.ResponseWriter, r *http.Request)
	GetSystematizationById(w http.ResponseWriter, r *http.Request)
	GetSystematizationList(w http.ResponseWriter, r *http.Request)
}

type UserProfileHandler interface {
	CreateUserProfile(w http.ResponseWriter, r *http.Request)
	UpdateUserProfile(w http.ResponseWriter, r *http.Request)
	DeleteUserProfile(w http.ResponseWriter, r *http.Request)
	GetUserProfileById(w http.ResponseWriter, r *http.Request)
	GetUserProfileList(w http.ResponseWriter, r *http.Request)
	GetContracts(w http.ResponseWriter, r *http.Request)
}

type EmployeeContractHandler interface {
	CreateEmployeeContract(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeContract(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeContract(w http.ResponseWriter, r *http.Request)
}

type EmployeesInOrganizationUnitHandler interface {
	CreateEmployeesInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetEmployeesInOrganizationUnitList(w http.ResponseWriter, r *http.Request)
	GetEmployeesInOrganizationUnitByEmployee(w http.ResponseWriter, r *http.Request)
	DeleteEmployeesInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateJobPositionInOrganizationUnit(w http.ResponseWriter, r *http.Request)
}

type EmployeeExperienceHandler interface {
	CreateEmployeeExperience(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeExperience(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeExperience(w http.ResponseWriter, r *http.Request)
	GetEmployeeExperienceList(w http.ResponseWriter, r *http.Request)
}

type EmployeeEducationHandler interface {
	CreateEmployeeEducation(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeEducation(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeEducation(w http.ResponseWriter, r *http.Request)
	GetEmployeeEducationList(w http.ResponseWriter, r *http.Request)
}

type EmployeeResolutionHandler interface {
	CreateEmployeeResolution(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeResolution(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeResolution(w http.ResponseWriter, r *http.Request)
	GetEmployeeResolutionList(w http.ResponseWriter, r *http.Request)
	GetEmployeeResolution(w http.ResponseWriter, r *http.Request)
}

type AbsentTypeHandler interface {
	CreateAbsentType(w http.ResponseWriter, r *http.Request)
	UpdateAbsentType(w http.ResponseWriter, r *http.Request)
	DeleteAbsentType(w http.ResponseWriter, r *http.Request)
	GetAbsentTypeById(w http.ResponseWriter, r *http.Request)
	GetAbsentTypeList(w http.ResponseWriter, r *http.Request)
}

type EvaluationHandler interface {
	CreateEvaluation(w http.ResponseWriter, r *http.Request)
	UpdateEvaluation(w http.ResponseWriter, r *http.Request)
	DeleteEvaluation(w http.ResponseWriter, r *http.Request)
	GetEvaluationById(w http.ResponseWriter, r *http.Request)
	GetEvaluationList(w http.ResponseWriter, r *http.Request)
}

type ForeignerHandler interface {
	CreateForeigner(w http.ResponseWriter, r *http.Request)
	UpdateForeigner(w http.ResponseWriter, r *http.Request)
	DeleteForeigner(w http.ResponseWriter, r *http.Request)
	GetForeignerById(w http.ResponseWriter, r *http.Request)
	GetForeignerList(w http.ResponseWriter, r *http.Request)
}

type EmployeeFamilyMemberHandler interface {
	CreateEmployeeFamilyMember(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeFamilyMember(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeFamilyMember(w http.ResponseWriter, r *http.Request)
	GetEmployeeFamilyMemberById(w http.ResponseWriter, r *http.Request)
	GetEmployeeFamilyMemberList(w http.ResponseWriter, r *http.Request)
}

type SalaryHandler interface {
	CreateSalary(w http.ResponseWriter, r *http.Request)
	UpdateSalary(w http.ResponseWriter, r *http.Request)
	DeleteSalary(w http.ResponseWriter, r *http.Request)
	GetSalaryById(w http.ResponseWriter, r *http.Request)
	GetSalaryList(w http.ResponseWriter, r *http.Request)
}

type UserNormFulfilmentHandler interface {
	CreateUserNormFulfilment(w http.ResponseWriter, r *http.Request)
	UpdateUserNormFulfilment(w http.ResponseWriter, r *http.Request)
	DeleteUserNormFulfilment(w http.ResponseWriter, r *http.Request)
	GetUserNormFulfilmentById(w http.ResponseWriter, r *http.Request)
	GetUserNormFulfilmentList(w http.ResponseWriter, r *http.Request)
}

type EmployeeAbsentHandler interface {
	CreateEmployeeAbsent(w http.ResponseWriter, r *http.Request)
	UpdateEmployeeAbsent(w http.ResponseWriter, r *http.Request)
	DeleteEmployeeAbsent(w http.ResponseWriter, r *http.Request)
	GetEmployeeAbsentList(w http.ResponseWriter, r *http.Request)
	GetAbsentById(w http.ResponseWriter, r *http.Request)
}

type RevisionsOfOrganizationUnitHandler interface {
	CreateRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request)
	DeleteRevisionsOfOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetRevisionsOfOrganizationUnitById(w http.ResponseWriter, r *http.Request)
	GetRevisionsOfOrganizationUnitList(w http.ResponseWriter, r *http.Request)
}

type TendersInOrganizationUnitHandler interface {
	CreateTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	DeleteTendersInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetTendersInOrganizationUnitById(w http.ResponseWriter, r *http.Request)
	GetTendersInOrganizationUnitList(w http.ResponseWriter, r *http.Request)
}

type TenderApplicationsInOrganizationUnitHandler interface {
	CreateTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	DeleteTenderApplicationsInOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetTenderApplicationsInOrganizationUnitById(w http.ResponseWriter, r *http.Request)
	GetTenderApplicationsInOrganizationUnitList(w http.ResponseWriter, r *http.Request)
}

type TenderTypeHandler interface {
	CreateTenderType(w http.ResponseWriter, r *http.Request)
	UpdateTenderType(w http.ResponseWriter, r *http.Request)
	DeleteTenderType(w http.ResponseWriter, r *http.Request)
	GetTenderTypeById(w http.ResponseWriter, r *http.Request)
	GetTenderTypeList(w http.ResponseWriter, r *http.Request)
}

type JudgeNumberResolutionHandler interface {
	CreateJudgeNumberResolution(w http.ResponseWriter, r *http.Request)
	UpdateJudgeNumberResolution(w http.ResponseWriter, r *http.Request)
	DeleteJudgeNumberResolution(w http.ResponseWriter, r *http.Request)
	GetJudgeNumberResolutionById(w http.ResponseWriter, r *http.Request)
	GetJudgeNumberResolutionList(w http.ResponseWriter, r *http.Request)
}

type JudgeNumberResolutionOrganizationUnitHandler interface {
	CreateJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request)
	UpdateJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request)
	DeleteJudgeNumberResolutionOrganizationUnit(w http.ResponseWriter, r *http.Request)
	GetJudgeNumberResolutionOrganizationUnitById(w http.ResponseWriter, r *http.Request)
	GetJudgeNumberResolutionOrganizationUnitList(w http.ResponseWriter, r *http.Request)
}

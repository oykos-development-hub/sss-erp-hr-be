package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/hr-ms-api/data"
	"gitlab.sudovi.me/erp/hr-ms-api/handlers"
	"gitlab.sudovi.me/erp/hr-ms-api/middleware"

	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/hr-ms-api/services"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/hr-ms-api"

	models := data.New(cel.DB.Pool)

	ErrorLogService := services.NewErrorLogServiceImpl(cel, models.ErrorLog)
	ErrorLogHandler := handlers.NewErrorLogHandler(cel, ErrorLogService)

	OrganizationUnitService := services.NewOrganizationUnitServiceImpl(cel, models.OrganizationUnit)
	OrganizationUnitHandler := handlers.NewOrganizationUnitHandler(cel, OrganizationUnitService, ErrorLogService)

	UserProfileService := services.NewUserProfileServiceImpl(cel, models.UserProfile, models.EmployeeContract)
	UserProfileHandler := handlers.NewUserProfileHandler(cel, UserProfileService, ErrorLogService)

	EmployeeContractService := services.NewEmployeeContractServiceImpl(cel, models.EmployeeContract)
	EmployeeContractHandler := handlers.NewEmployeeContractHandler(cel, EmployeeContractService, ErrorLogService)

	SystematizationService := services.NewSystematizationServiceImpl(cel, models.Systematization)
	SystematizationHandler := handlers.NewSystematizationHandler(cel, SystematizationService, ErrorLogService)

	JobPositionService := services.NewJobPositionServiceImpl(cel, models.JobPosition)
	JobPositionHandler := handlers.NewJobPositionHandler(cel, JobPositionService, ErrorLogService)

	JobPositionsInOrganizationUnitsService := services.NewJobPositionsInOrganizationUnitsServiceImpl(cel, models.JobPositionsInOrganizationUnits)
	JobPositionsInOrganizationUnitsHandler := handlers.NewJobPositionsInOrganizationUnitsHandler(cel, JobPositionsInOrganizationUnitsService, ErrorLogService)

	EmployeesInOrganizationUnitService := services.NewEmployeesInOrganizationUnitServiceImpl(cel, models.EmployeesInOrganizationUnit)
	EmployeesInOrganizationUnitHandler := handlers.NewEmployeesInOrganizationUnitHandler(cel, EmployeesInOrganizationUnitService, ErrorLogService)

	EmployeeExperienceService := services.NewEmployeeExperienceServiceImpl(cel, models.EmployeeExperience)
	EmployeeExperienceHandler := handlers.NewEmployeeExperienceHandler(cel, EmployeeExperienceService, ErrorLogService)

	EmployeeEducationService := services.NewEmployeeEducationServiceImpl(cel, models.EmployeeEducation)
	EmployeeEducationHandler := handlers.NewEmployeeEducationHandler(cel, EmployeeEducationService, ErrorLogService)

	EmployeeResolutionService := services.NewEmployeeResolutionServiceImpl(cel, models.EmployeeResolution)
	EmployeeResolutionHandler := handlers.NewEmployeeResolutionHandler(cel, EmployeeResolutionService, ErrorLogService)

	AbsentTypeService := services.NewAbsentTypeServiceImpl(cel, models.AbsentType)
	AbsentTypeHandler := handlers.NewAbsentTypeHandler(cel, AbsentTypeService, ErrorLogService)

	EvaluationService := services.NewEvaluationServiceImpl(cel, models.Evaluation)
	EvaluationHandler := handlers.NewEvaluationHandler(cel, EvaluationService, ErrorLogService)

	ForeignerService := services.NewForeignerServiceImpl(cel, models.Foreigner)
	ForeignerHandler := handlers.NewForeignerHandler(cel, ForeignerService, ErrorLogService)

	UserNormFulfilmentService := services.NewUserNormFulfilmentServiceImpl(cel, models.UserNormFulfilment)
	UserNormFulfilmentHandler := handlers.NewUserNormFulfilmentHandler(cel, UserNormFulfilmentService, ErrorLogService)

	EmployeeFamilyMemberService := services.NewEmployeeFamilyMemberServiceImpl(cel, models.EmployeeFamilyMember)
	EmployeeFamilyMemberHandler := handlers.NewEmployeeFamilyMemberHandler(cel, EmployeeFamilyMemberService, ErrorLogService)

	SalaryService := services.NewSalaryServiceImpl(cel, models.Salary)
	SalaryHandler := handlers.NewSalaryHandler(cel, SalaryService, ErrorLogService)

	EmployeeAbsentService := services.NewEmployeeAbsentServiceImpl(cel, models.EmployeeAbsent)
	EmployeeAbsentHandler := handlers.NewEmployeeAbsentHandler(cel, EmployeeAbsentService, ErrorLogService)

	RevisionsOfOrganizationUnitService := services.NewRevisionsOfOrganizationUnitServiceImpl(cel, models.RevisionsOfOrganizationUnit)
	RevisionsOfOrganizationUnitHandler := handlers.NewRevisionsOfOrganizationUnitHandler(cel, RevisionsOfOrganizationUnitService, ErrorLogService)

	TendersInOrganizationUnitService := services.NewTendersInOrganizationUnitServiceImpl(cel, models.TendersInOrganizationUnit)
	TendersInOrganizationUnitHandler := handlers.NewTendersInOrganizationUnitHandler(cel, TendersInOrganizationUnitService, ErrorLogService)

	TenderApplicationsInOrganizationUnitService := services.NewTenderApplicationsInOrganizationUnitServiceImpl(cel, models.TenderApplicationsInOrganizationUnit)
	TenderApplicationsInOrganizationUnitHandler := handlers.NewTenderApplicationsInOrganizationUnitHandler(cel, TenderApplicationsInOrganizationUnitService, ErrorLogService)

	TenderTypeService := services.NewTenderTypeServiceImpl(cel, models.TenderType)
	TenderTypeHandler := handlers.NewTenderTypeHandler(cel, TenderTypeService, ErrorLogService)

	JudgeNumberResolutionService := services.NewJudgeNumberResolutionServiceImpl(cel, models.JudgeNumberResolution)
	JudgeNumberResolutionHandler := handlers.NewJudgeNumberResolutionHandler(cel, JudgeNumberResolutionService, ErrorLogService)

	JudgeNumberResolutionOrganizationUnitService := services.NewJudgeNumberResolutionOrganizationUnitServiceImpl(cel, models.JudgeNumberResolutionOrganizationUnit)
	JudgeNumberResolutionOrganizationUnitHandler := handlers.NewJudgeNumberResolutionOrganizationUnitHandler(cel, JudgeNumberResolutionOrganizationUnitService, ErrorLogService)

	PlanService := services.NewPlanServiceImpl(cel, models.Plan)
	PlanHandler := handlers.NewPlanHandler(cel, PlanService, ErrorLogService)

	RevisionService := services.NewRevisionServiceImpl(cel, models.Revision)
	RevisionHandler := handlers.NewRevisionHandler(cel, RevisionService, ErrorLogService)

	RevisionTipService := services.NewRevisionTipServiceImpl(cel, models.RevisionTip)
	RevisionTipHandler := handlers.NewRevisionTipHandler(cel, RevisionTipService, ErrorLogService)

	JudgeService := services.NewJudgeServiceImpl(cel, models.Judge)
	JudgeHandler := handlers.NewJudgeHandler(cel, JudgeService, ErrorLogService)

	RevisionsInOrganizationUnitService := services.NewRevisionsInOrganizationUnitServiceImpl(cel, models.RevisionsInOrganizationUnit)
	RevisionsInOrganizationUnitHandler := handlers.NewRevisionsInOrganizationUnitHandler(cel, RevisionsInOrganizationUnitService, ErrorLogService)

	RevisionRevisorService := services.NewRevisionRevisorServiceImpl(cel, models.RevisionRevisor)
	RevisionRevisorHandler := handlers.NewRevisionRevisorHandler(cel, RevisionRevisorService, ErrorLogService)

	LogService := services.NewLogServiceImpl(cel, models.Log)
	LogHandler := handlers.NewLogHandler(cel, LogService, ErrorLogService)

	RevisionTipImplementationService := services.NewRevisionTipImplementationServiceImpl(cel, models.RevisionTipImplementation)
	RevisionTipImplementationHandler := handlers.NewRevisionTipImplementationHandler(cel, RevisionTipImplementationService, ErrorLogService)

	myHandlers := &handlers.Handlers{
		OrganizationUnitHandler:                      OrganizationUnitHandler,
		JobPositionHandler:                           JobPositionHandler,
		JobPositionsInOrganizationUnitsHandler:       JobPositionsInOrganizationUnitsHandler,
		SystematizationHandler:                       SystematizationHandler,
		UserProfileHandler:                           UserProfileHandler,
		EmployeeContractHandler:                      EmployeeContractHandler,
		EmployeesInOrganizationUnitHandler:           EmployeesInOrganizationUnitHandler,
		EmployeeExperienceHandler:                    EmployeeExperienceHandler,
		EmployeeEducationHandler:                     EmployeeEducationHandler,
		EmployeeResolutionHandler:                    EmployeeResolutionHandler,
		AbsentTypeHandler:                            AbsentTypeHandler,
		EvaluationHandler:                            EvaluationHandler,
		ForeignerHandler:                             ForeignerHandler,
		UserNormFulfilmentHandler:                    UserNormFulfilmentHandler,
		EmployeeFamilyMemberHandler:                  EmployeeFamilyMemberHandler,
		SalaryHandler:                                SalaryHandler,
		EmployeeAbsentHandler:                        EmployeeAbsentHandler,
		RevisionsOfOrganizationUnitHandler:           RevisionsOfOrganizationUnitHandler,
		TendersInOrganizationUnitHandler:             TendersInOrganizationUnitHandler,
		TenderApplicationsInOrganizationUnitHandler:  TenderApplicationsInOrganizationUnitHandler,
		TenderTypeHandler:                            TenderTypeHandler,
		JudgeNumberResolutionHandler:                 JudgeNumberResolutionHandler,
		JudgeNumberResolutionOrganizationUnitHandler: JudgeNumberResolutionOrganizationUnitHandler,
		PlanHandler:                                  PlanHandler,
		RevisionHandler:                              RevisionHandler,
		RevisionTipHandler:                           RevisionTipHandler,
		JudgeHandler:                                 JudgeHandler,
		RevisionsInOrganizationUnitHandler:           RevisionsInOrganizationUnitHandler,
		RevisionRevisorHandler:                       RevisionRevisorHandler,
		LogHandler:                                   LogHandler,
		ErrorLogHandler:                              ErrorLogHandler,
		RevisionTipImplementationHandler:             RevisionTipImplementationHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	cel.Validator().V = setupValidator(cel.Validator().V)

	return cel
}

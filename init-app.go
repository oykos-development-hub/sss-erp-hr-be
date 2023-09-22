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

	OrganizationUnitService := services.NewOrganizationUnitServiceImpl(cel, models.OrganizationUnit)
	OrganizationUnitHandler := handlers.NewOrganizationUnitHandler(cel, OrganizationUnitService)

	UserProfileService := services.NewUserProfileServiceImpl(cel, models.UserProfile, models.EmployeeContract)
	UserProfileHandler := handlers.NewUserProfileHandler(cel, UserProfileService)

	EmployeeContractService := services.NewEmployeeContractServiceImpl(cel, models.EmployeeContract)
	EmployeeContractHandler := handlers.NewEmployeeContractHandler(cel, EmployeeContractService)

	SystematizationService := services.NewSystematizationServiceImpl(cel, models.Systematization)
	SystematizationHandler := handlers.NewSystematizationHandler(cel, SystematizationService)

	JobPositionService := services.NewJobPositionServiceImpl(cel, models.JobPosition)
	JobPositionHandler := handlers.NewJobPositionHandler(cel, JobPositionService)

	JobPositionsInOrganizationUnitsService := services.NewJobPositionsInOrganizationUnitsServiceImpl(cel, models.JobPositionsInOrganizationUnits)
	JobPositionsInOrganizationUnitsHandler := handlers.NewJobPositionsInOrganizationUnitsHandler(cel, JobPositionsInOrganizationUnitsService)

	EmployeesInOrganizationUnitService := services.NewEmployeesInOrganizationUnitServiceImpl(cel, models.EmployeesInOrganizationUnit)
	EmployeesInOrganizationUnitHandler := handlers.NewEmployeesInOrganizationUnitHandler(cel, EmployeesInOrganizationUnitService)

	EmployeeExperienceService := services.NewEmployeeExperienceServiceImpl(cel, models.EmployeeExperience)
	EmployeeExperienceHandler := handlers.NewEmployeeExperienceHandler(cel, EmployeeExperienceService)

	EmployeeEducationService := services.NewEmployeeEducationServiceImpl(cel, models.EmployeeEducation)
	EmployeeEducationHandler := handlers.NewEmployeeEducationHandler(cel, EmployeeEducationService)

	EmployeeResolutionService := services.NewEmployeeResolutionServiceImpl(cel, models.EmployeeResolution)
	EmployeeResolutionHandler := handlers.NewEmployeeResolutionHandler(cel, EmployeeResolutionService)

	AbsentTypeService := services.NewAbsentTypeServiceImpl(cel, models.AbsentType)
	AbsentTypeHandler := handlers.NewAbsentTypeHandler(cel, AbsentTypeService)

	EvaluationService := services.NewEvaluationServiceImpl(cel, models.Evaluation)
	EvaluationHandler := handlers.NewEvaluationHandler(cel, EvaluationService)

	ForeignerService := services.NewForeignerServiceImpl(cel, models.Foreigner)
	ForeignerHandler := handlers.NewForeignerHandler(cel, ForeignerService)

	UserNormFulfilmentService := services.NewUserNormFulfilmentServiceImpl(cel, models.UserNormFulfilment)
	UserNormFulfilmentHandler := handlers.NewUserNormFulfilmentHandler(cel, UserNormFulfilmentService)

	EmployeeFamilyMemberService := services.NewEmployeeFamilyMemberServiceImpl(cel, models.EmployeeFamilyMember)
	EmployeeFamilyMemberHandler := handlers.NewEmployeeFamilyMemberHandler(cel, EmployeeFamilyMemberService)

	SalaryService := services.NewSalaryServiceImpl(cel, models.Salary)
	SalaryHandler := handlers.NewSalaryHandler(cel, SalaryService)

	EmployeeAbsentService := services.NewEmployeeAbsentServiceImpl(cel, models.EmployeeAbsent)
	EmployeeAbsentHandler := handlers.NewEmployeeAbsentHandler(cel, EmployeeAbsentService)

	RevisionsOfOrganizationUnitService := services.NewRevisionsOfOrganizationUnitServiceImpl(cel, models.RevisionsOfOrganizationUnit)
	RevisionsOfOrganizationUnitHandler := handlers.NewRevisionsOfOrganizationUnitHandler(cel, RevisionsOfOrganizationUnitService)

	TendersInOrganizationUnitService := services.NewTendersInOrganizationUnitServiceImpl(cel, models.TendersInOrganizationUnit)
	TendersInOrganizationUnitHandler := handlers.NewTendersInOrganizationUnitHandler(cel, TendersInOrganizationUnitService)

	TenderApplicationsInOrganizationUnitService := services.NewTenderApplicationsInOrganizationUnitServiceImpl(cel, models.TenderApplicationsInOrganizationUnit)
	TenderApplicationsInOrganizationUnitHandler := handlers.NewTenderApplicationsInOrganizationUnitHandler(cel, TenderApplicationsInOrganizationUnitService)

	TenderTypeService := services.NewTenderTypeServiceImpl(cel, models.TenderType)
	TenderTypeHandler := handlers.NewTenderTypeHandler(cel, TenderTypeService)

	JudgeNumberResolutionService := services.NewJudgeNumberResolutionServiceImpl(cel, models.JudgeNumberResolution)
	JudgeNumberResolutionHandler := handlers.NewJudgeNumberResolutionHandler(cel, JudgeNumberResolutionService)

	JudgeNumberResolutionOrganizationUnitService := services.NewJudgeNumberResolutionOrganizationUnitServiceImpl(cel, models.JudgeNumberResolutionOrganizationUnit)
	JudgeNumberResolutionOrganizationUnitHandler := handlers.NewJudgeNumberResolutionOrganizationUnitHandler(cel, JudgeNumberResolutionOrganizationUnitService)

		
	PlanService := services.NewPlanServiceImpl(cel, models.Plan)
	PlanHandler := handlers.NewPlanHandler(cel, PlanService)

		
	RevisionService := services.NewRevisionServiceImpl(cel, models.Revision)
	RevisionHandler := handlers.NewRevisionHandler(cel, RevisionService)

		
	RevisionTipService := services.NewRevisionTipServiceImpl(cel, models.RevisionTip)
	RevisionTipHandler := handlers.NewRevisionTipHandler(cel, RevisionTipService)

		
	JudgeService := services.NewJudgeServiceImpl(cel, models.Judge)
	JudgeHandler := handlers.NewJudgeHandler(cel, JudgeService)

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
		PlanHandler: PlanHandler,
		RevisionHandler: RevisionHandler,
		RevisionTipHandler: RevisionTipHandler,
		JudgeHandler: JudgeHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	cel.Validator().V = setupValidator(cel.Validator().V)

	return cel
}

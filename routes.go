package main

import (
	"gitlab.sudovi.me/erp/hr-ms-api/handlers"
	"gitlab.sudovi.me/erp/hr-ms-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	//api
	app.Routes.Route("/api", func(rt chi.Router) {

		// Organization units
		rt.Post("/organization-units", handlers.OrganizationUnitHandler.CreateOrganizationUnit)
		rt.Get("/organization-units/{id}", handlers.OrganizationUnitHandler.GetOrganizationUnitById)
		rt.Get("/organization-units", handlers.OrganizationUnitHandler.GetOrganizationUnitList)
		rt.Put("/organization-units/{id}", handlers.OrganizationUnitHandler.UpdateOrganizationUnit)
		rt.Delete("/organization-units/{id}", handlers.OrganizationUnitHandler.DeleteOrganizationUnit)

		// Job positions
		rt.Post("/job-positions", handlers.JobPositionHandler.CreateJobPosition)
		rt.Get("/job-positions/{id}", handlers.JobPositionHandler.GetJobPositionById)
		rt.Get("/job-positions", handlers.JobPositionHandler.GetJobPositionList)
		rt.Put("/job-positions/{id}", handlers.JobPositionHandler.UpdateJobPosition)
		rt.Delete("/job-positions/{id}", handlers.JobPositionHandler.DeleteJobPosition)

		rt.Get("/job-positions-in-organization-units", handlers.JobPositionsInOrganizationUnitsHandler.GetJobPositionsInOrganizationUnitsList)
		rt.Post("/job-positions-in-organization-units", handlers.JobPositionsInOrganizationUnitsHandler.CreateJobPositionsInOrganizationUnits)
		rt.Delete("/job-positions-in-organization-units/{id}", handlers.JobPositionsInOrganizationUnitsHandler.DeleteJobPositionsInOrganizationUnits)
		rt.Put("/job-positions-in-organization-units/{id}", handlers.JobPositionsInOrganizationUnitsHandler.UpdateJobPositionsInOrganizationUnits)
		rt.Get("/job-positions-in-organization-units/{id}", handlers.JobPositionsInOrganizationUnitsHandler.GetJobPositionsInOrganizationUnitsById)

		// Systematizations
		rt.Post("/systematizations", handlers.SystematizationHandler.CreateSystematization)
		rt.Get("/systematizations/{id}", handlers.SystematizationHandler.GetSystematizationById)
		rt.Get("/systematizations", handlers.SystematizationHandler.GetSystematizationList)
		rt.Put("/systematizations/{id}", handlers.SystematizationHandler.UpdateSystematization)
		rt.Delete("/systematizations/{id}", handlers.SystematizationHandler.DeleteSystematization)

		rt.Post("/user-profiles", handlers.UserProfileHandler.CreateUserProfile)
		rt.Get("/user-profiles/{id}", handlers.UserProfileHandler.GetUserProfileById)
		rt.Get("/user-profiles", handlers.UserProfileHandler.GetUserProfileList)
		rt.Get("/user-profiles/{id}/contracts", handlers.UserProfileHandler.GetContracts)
		rt.Put("/user-profiles/{id}", handlers.UserProfileHandler.UpdateUserProfile)
		rt.Delete("/user-profiles/{id}", handlers.UserProfileHandler.DeleteUserProfile)

		rt.Post("/employee-contracts", handlers.EmployeeContractHandler.CreateEmployeeContract)
		rt.Put("/employee-contracts/{id}", handlers.EmployeeContractHandler.UpdateEmployeeContract)
		rt.Delete("/employee-contracts/{id}", handlers.EmployeeContractHandler.DeleteEmployeeContract)

		rt.Post("/employees-in-organization-units", handlers.EmployeesInOrganizationUnitHandler.CreateEmployeesInOrganizationUnit)
		rt.Get("/user-profiles/{id}/employee-in-organization-unit", handlers.EmployeesInOrganizationUnitHandler.GetEmployeesInOrganizationUnitByEmployee)
		rt.Get("/employees-in-organization-units", handlers.EmployeesInOrganizationUnitHandler.GetEmployeesInOrganizationUnitList)
		rt.Delete("/employees-in-organization-units/{position_in_organization_unit_id}", handlers.EmployeesInOrganizationUnitHandler.DeleteEmployeesInOrganizationUnit)
		rt.Put("/employees-in-organization-units/{id}", handlers.EmployeesInOrganizationUnitHandler.UpdateJobPositionInOrganizationUnit)

		rt.Post("/employee-experiences", handlers.EmployeeExperienceHandler.CreateEmployeeExperience)
		rt.Get("/user-profiles/{id}/experiences", handlers.EmployeeExperienceHandler.GetEmployeeExperienceList)
		rt.Put("/employee-experiences/{id}", handlers.EmployeeExperienceHandler.UpdateEmployeeExperience)
		rt.Delete("/employee-experiences/{id}", handlers.EmployeeExperienceHandler.DeleteEmployeeExperience)

		rt.Post("/employee-educations", handlers.EmployeeEducationHandler.CreateEmployeeEducation)
		rt.Get("/employee-educations", handlers.EmployeeEducationHandler.GetEmployeeEducationList)
		rt.Put("/employee-educations/{id}", handlers.EmployeeEducationHandler.UpdateEmployeeEducation)
		rt.Delete("/employee-educations/{id}", handlers.EmployeeEducationHandler.DeleteEmployeeEducation)

		rt.Post("/employee-resolutions", handlers.EmployeeResolutionHandler.CreateEmployeeResolution)
		rt.Get("/user-profiles/{id}/resolutions", handlers.EmployeeResolutionHandler.GetEmployeeResolutionList)
		rt.Get("/employee-resolutions/{id}", handlers.EmployeeResolutionHandler.GetEmployeeResolution)
		rt.Put("/employee-resolutions/{id}", handlers.EmployeeResolutionHandler.UpdateEmployeeResolution)
		rt.Delete("/employee-resolutions/{id}", handlers.EmployeeResolutionHandler.DeleteEmployeeResolution)

		rt.Post("/absent-types", handlers.AbsentTypeHandler.CreateAbsentType)
		rt.Get("/absent-types/{id}", handlers.AbsentTypeHandler.GetAbsentTypeById)
		rt.Get("/absent-types", handlers.AbsentTypeHandler.GetAbsentTypeList)
		rt.Put("/absent-types/{id}", handlers.AbsentTypeHandler.UpdateAbsentType)
		rt.Delete("/absent-types/{id}", handlers.AbsentTypeHandler.DeleteAbsentType)

		rt.Post("/evaluations", handlers.EvaluationHandler.CreateEvaluation)
		rt.Get("/user-profiles/{id}/evaluations", handlers.EvaluationHandler.GetEvaluationList)
		rt.Put("/evaluations/{id}", handlers.EvaluationHandler.UpdateEvaluation)
		rt.Get("/evaluations/{id}", handlers.EvaluationHandler.GetEvaluationById)
		rt.Delete("/evaluations/{id}", handlers.EvaluationHandler.DeleteEvaluation)

		rt.Post("/foreigners", handlers.ForeignerHandler.CreateForeigner)
		rt.Get("/foreigners/{id}", handlers.ForeignerHandler.GetForeignerById)
		rt.Get("/user-profiles/{id}/foreigners", handlers.ForeignerHandler.GetForeignerList)
		rt.Put("/foreigners/{id}", handlers.ForeignerHandler.UpdateForeigner)
		rt.Delete("/foreigners/{id}", handlers.ForeignerHandler.DeleteForeigner)

		rt.Post("/employee-family-members", handlers.EmployeeFamilyMemberHandler.CreateEmployeeFamilyMember)
		rt.Get("/employee-family-members/{id}", handlers.EmployeeFamilyMemberHandler.GetEmployeeFamilyMemberById)
		rt.Get("/user-profiles/{id}/family-members", handlers.EmployeeFamilyMemberHandler.GetEmployeeFamilyMemberList)
		rt.Put("/employee-family-members/{id}", handlers.EmployeeFamilyMemberHandler.UpdateEmployeeFamilyMember)
		rt.Delete("/employee-family-members/{id}", handlers.EmployeeFamilyMemberHandler.DeleteEmployeeFamilyMember)

		rt.Post("/salaries", handlers.SalaryHandler.CreateSalary)
		rt.Get("/salaries/{id}", handlers.SalaryHandler.GetSalaryById)
		rt.Get("/user-profiles/{id}/salaries", handlers.SalaryHandler.GetSalaryList)
		rt.Put("/salaries/{id}", handlers.SalaryHandler.UpdateSalary)
		rt.Delete("/salaries/{id}", handlers.SalaryHandler.DeleteSalary)

		rt.Post("/user-norms", handlers.UserNormFulfilmentHandler.CreateUserNormFulfilment)
		rt.Get("/user-norms/{id}", handlers.UserNormFulfilmentHandler.GetUserNormFulfilmentById)
		rt.Get("/user-profiles/{id}/norms", handlers.UserNormFulfilmentHandler.GetUserNormFulfilmentList)
		rt.Put("/user-norms/{id}", handlers.UserNormFulfilmentHandler.UpdateUserNormFulfilment)
		rt.Delete("/user-norms/{id}", handlers.UserNormFulfilmentHandler.DeleteUserNormFulfilment)

		rt.Post("/employee-absents", handlers.EmployeeAbsentHandler.CreateEmployeeAbsent)
		rt.Put("/employee-absents/{id}", handlers.EmployeeAbsentHandler.UpdateEmployeeAbsent)
		rt.Get("/user-profiles/{id}/absents", handlers.EmployeeAbsentHandler.GetEmployeeAbsentList)
		rt.Get("/employee-absents/{id}", handlers.EmployeeAbsentHandler.GetAbsentById)
		rt.Delete("/employee-absents/{id}", handlers.EmployeeAbsentHandler.DeleteEmployeeAbsent)

		rt.Post("/revisions-of-organization-units", handlers.RevisionsOfOrganizationUnitHandler.CreateRevisionsOfOrganizationUnit)
		rt.Get("/revisions-of-organization-units/{id}", handlers.RevisionsOfOrganizationUnitHandler.GetRevisionsOfOrganizationUnitById)
		rt.Get("/revisions-of-organization-units", handlers.RevisionsOfOrganizationUnitHandler.GetRevisionsOfOrganizationUnitList)
		rt.Put("/revisions-of-organization-units/{id}", handlers.RevisionsOfOrganizationUnitHandler.UpdateRevisionsOfOrganizationUnit)
		rt.Delete("/revisions-of-organization-units/{id}", handlers.RevisionsOfOrganizationUnitHandler.DeleteRevisionsOfOrganizationUnit)

		rt.Post("/tenders-in-organization-units", handlers.TendersInOrganizationUnitHandler.CreateTendersInOrganizationUnit)
		rt.Get("/tenders-in-organization-units/{id}", handlers.TendersInOrganizationUnitHandler.GetTendersInOrganizationUnitById)
		rt.Get("/tenders-in-organization-units", handlers.TendersInOrganizationUnitHandler.GetTendersInOrganizationUnitList)
		rt.Put("/tenders-in-organization-units/{id}", handlers.TendersInOrganizationUnitHandler.UpdateTendersInOrganizationUnit)
		rt.Delete("/tenders-in-organization-units/{id}", handlers.TendersInOrganizationUnitHandler.DeleteTendersInOrganizationUnit)

		rt.Post("/tender-applications-in-organization-units", handlers.TenderApplicationsInOrganizationUnitHandler.CreateTenderApplicationsInOrganizationUnit)
		rt.Get("/tender-applications-in-organization-units/{id}", handlers.TenderApplicationsInOrganizationUnitHandler.GetTenderApplicationsInOrganizationUnitById)
		rt.Get("/tender-applications-in-organization-units", handlers.TenderApplicationsInOrganizationUnitHandler.GetTenderApplicationsInOrganizationUnitList)
		rt.Put("/tender-applications-in-organization-units/{id}", handlers.TenderApplicationsInOrganizationUnitHandler.UpdateTenderApplicationsInOrganizationUnit)
		rt.Delete("/tender-applications-in-organization-units/{id}", handlers.TenderApplicationsInOrganizationUnitHandler.DeleteTenderApplicationsInOrganizationUnit)

		rt.Post("/tender-types", handlers.TenderTypeHandler.CreateTenderType)
		rt.Get("/tender-types/{id}", handlers.TenderTypeHandler.GetTenderTypeById)
		rt.Get("/tender-types", handlers.TenderTypeHandler.GetTenderTypeList)
		rt.Put("/tender-types/{id}", handlers.TenderTypeHandler.UpdateTenderType)
		rt.Delete("/tender-types/{id}", handlers.TenderTypeHandler.DeleteTenderType)

		rt.Post("/judge-number-resolutions", handlers.JudgeNumberResolutionHandler.CreateJudgeNumberResolution)
		rt.Get("/judge-number-resolutions/{id}", handlers.JudgeNumberResolutionHandler.GetJudgeNumberResolutionById)
		rt.Get("/judge-number-resolutions", handlers.JudgeNumberResolutionHandler.GetJudgeNumberResolutionList)
		rt.Put("/judge-number-resolutions/{id}", handlers.JudgeNumberResolutionHandler.UpdateJudgeNumberResolution)
		rt.Delete("/judge-number-resolutions/{id}", handlers.JudgeNumberResolutionHandler.DeleteJudgeNumberResolution)

		rt.Post("/judge-number-resolution-organization-units", handlers.JudgeNumberResolutionOrganizationUnitHandler.CreateJudgeNumberResolutionOrganizationUnit)
		rt.Get("/judge-number-resolution-organization-units/{id}", handlers.JudgeNumberResolutionOrganizationUnitHandler.GetJudgeNumberResolutionOrganizationUnitById)
		rt.Get("/judge-number-resolution-organization-units", handlers.JudgeNumberResolutionOrganizationUnitHandler.GetJudgeNumberResolutionOrganizationUnitList)
		rt.Put("/judge-number-resolution-organization-units/{id}", handlers.JudgeNumberResolutionOrganizationUnitHandler.UpdateJudgeNumberResolutionOrganizationUnit)
		rt.Delete("/judge-number-resolution-organization-units/{id}", handlers.JudgeNumberResolutionOrganizationUnitHandler.DeleteJudgeNumberResolutionOrganizationUnit)

		rt.Post("/plans", handlers.PlanHandler.CreatePlan)
		rt.Get("/plans/{id}", handlers.PlanHandler.GetPlanById)
		rt.Get("/plans", handlers.PlanHandler.GetPlanList)
		rt.Put("/plans/{id}", handlers.PlanHandler.UpdatePlan)
		rt.Delete("/plans/{id}", handlers.PlanHandler.DeletePlan)

		rt.Post("/revisions", handlers.RevisionHandler.CreateRevision)
		rt.Get("/revisions/{id}", handlers.RevisionHandler.GetRevisionById)
		rt.Get("/revisions", handlers.RevisionHandler.GetRevisionList)
		rt.Put("/revisions/{id}", handlers.RevisionHandler.UpdateRevision)
		rt.Delete("/revisions/{id}", handlers.RevisionHandler.DeleteRevision)

		rt.Post("/revision-tips", handlers.RevisionTipHandler.CreateRevisionTip)
		rt.Get("/revision-tips/{id}", handlers.RevisionTipHandler.GetRevisionTipById)
		rt.Get("/revision-tips", handlers.RevisionTipHandler.GetRevisionTipList)
		rt.Put("/revision-tips/{id}", handlers.RevisionTipHandler.UpdateRevisionTip)
		rt.Delete("/revision-tips/{id}", handlers.RevisionTipHandler.DeleteRevisionTip)
	})

	return app.Routes
}

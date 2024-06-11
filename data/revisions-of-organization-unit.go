package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type RevisionsOfOrganizationUnit struct {
	ID                              int        `db:"id,omitempty"`
	Name                            *string    `db:"name"`
	RevisionTypeID                  *int       `db:"revision_type_id"`
	RevisorUserProfileID            *int       `db:"revisor_user_profile_id"`
	RevisorUserProfile              *string    `db:"revisor_user_profile"`
	InternalOrganizationUnitID      *int       `db:"internal_organization_unit_id"`
	ExternalOrganizationUnitID      *int       `db:"external_organization_unit_id"`
	ResponsibleUserProfileID        *int       `db:"responsible_user_profile_id"`
	ResponsibleUserProfile          *string    `db:"responsible_user_profile"`
	ImplementationUserProfileID     *int       `db:"implementation_user_profile_id"`
	ImplementationUserProfile       *string    `db:"implementation_user_profile"`
	Title                           *string    `db:"title"`
	PlanedYear                      *string    `db:"planned_year"`
	PlannedQuarter                  *string    `db:"planned_quarter"`
	SerialNumber                    *string    `db:"serial_number"`
	Priority                        *string    `db:"priority"`
	DateOfRevision                  *time.Time `db:"date_of_revision"`
	DateOfAcceptance                *time.Time `db:"date_of_acceptance"`
	DateOfRejection                 *time.Time `db:"date_of_rejection"`
	ImplementationSuggestion        *string    `db:"implementation_suggestion"`
	ImplementationMonthSpan         *string    `db:"implementation_month_span"`
	DateOfImplementation            *time.Time `db:"date_of_implementation"`
	StateOfImplementation           *string    `db:"state_of_implementation"`
	ImplementationFailedDescription *string    `db:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string    `db:"second_implementation_month_span"`
	SecondDateOfRevision            *time.Time `db:"second_date_of_revision"`
	FileID                          *int       `db:"file_id"`
	RefDocument                     *string    `db:"ref_document"`
	CreatedAt                       time.Time  `db:"created_at,omitempty"`
	UpdatedAt                       time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *RevisionsOfOrganizationUnit) Table() string {
	return "revisions_of_organization_units"
}

// GetAll gets all records from the database, using Upper
func (t *RevisionsOfOrganizationUnit) GetAll(page *int, pageSize *int, condition *up.Cond) ([]*RevisionsOfOrganizationUnit, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*RevisionsOfOrganizationUnit
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *RevisionsOfOrganizationUnit) Get(id int) (*RevisionsOfOrganizationUnit, error) {
	var one RevisionsOfOrganizationUnit
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *RevisionsOfOrganizationUnit) Update(m RevisionsOfOrganizationUnit) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *RevisionsOfOrganizationUnit) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *RevisionsOfOrganizationUnit) Insert(m RevisionsOfOrganizationUnit) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

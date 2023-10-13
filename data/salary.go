package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Salary struct
type Salary struct {
	ID                  int       `db:"id,omitempty"`
	UserProfileID       int       `db:"user_profile_id"`
	BenefitedTrack      *bool     `db:"benefited_track"`
	WithoutRaise        *bool     `db:"without_raise"`
	InsuranceBasis      string    `db:"insurance_basis"`
	SalaryRank          string    `db:"salary_rank"`
	DailyWorkHours      string    `db:"daily_work_hours"`
	WeeklyWorkHours     string    `db:"weekly_work_hours"`
	EducationRank       string    `db:"education_rank"`
	EducationNaming     string    `db:"education_naming"`
	ObligationReduction string    `db:"obligation_reduction"`
	UserResoultionID    *int      `db:"user_resolution_id,omitempty"`
	OrganizationUnitID  int       `db:"organization_unit_id"`
	CreatedAt           time.Time `db:"created_at,omitempty"`
	UpdatedAt           time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Salary) Table() string {
	return "salaries"
}

// GetAll gets all records from the database, using upper
func (t *Salary) GetAll(condition *up.Cond) ([]*Salary, error) {
	collection := upper.Collection(t.Table())
	var all []*Salary
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Salary) Get(id int) (*Salary, error) {
	var one Salary
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Salary) Update(m Salary) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Salary) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Salary) Insert(m Salary) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

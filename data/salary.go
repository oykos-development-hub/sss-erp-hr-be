package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
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

// GetAll gets all records from the database, using Upper
func (t *Salary) GetAll(condition *up.Cond) ([]*Salary, error) {
	collection := Upper.Collection(t.Table())
	var all []*Salary
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper order by")
	}

	return all, err
}

// Get gets one record from the database, by id, using Upper
func (t *Salary) Get(id int) (*Salary, error) {
	var one Salary
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Salary) Update(ctx context.Context, m Salary) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return newErrors.Wrap(err, "upper update")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Salary) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return newErrors.Wrap(err, "upper delete")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Salary) Insert(ctx context.Context, m Salary) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return 0, newErrors.Wrap(err, "context get user id")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return newErrors.Wrap(err, "upper insert")
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

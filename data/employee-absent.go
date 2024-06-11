package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
)

// EmployeeAbsent struct
type EmployeeAbsent struct {
	ID                       int       `db:"id,omitempty"`
	AbsentTypeID             int       `db:"absent_type_id"`
	UserProfileID            int       `db:"user_profile_id"`
	TargetOrganizationUnitID *int      `db:"target_organization_unit_id"`
	Description              *string   `db:"description"`
	DateOfStart              time.Time `db:"date_of_start"`
	DateOfEnd                time.Time `db:"date_of_end"`
	Location                 *string   `db:"location"`
	FileID                   *int      `db:"file_id"`
	CreatedAt                time.Time `db:"created_at,omitempty"`
	UpdatedAt                time.Time `db:"updated_at,omitempty"`
}

// Table returns the table name
func (t *EmployeeAbsent) Table() string {
	return "employee_absents"
}

// GetAll gets all records from the database, using Upper
func (t *EmployeeAbsent) GetAll(condition *up.Cond) ([]*EmployeeAbsent, error) {
	collection := Upper.Collection(t.Table())
	var all []*EmployeeAbsent
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

// Get gets one record from the database, by id, using Upper
func (t *EmployeeAbsent) Get(id int) (*EmployeeAbsent, error) {
	var one EmployeeAbsent
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *EmployeeAbsent) Update(ctx context.Context, m EmployeeAbsent) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *EmployeeAbsent) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *EmployeeAbsent) Insert(ctx context.Context, m EmployeeAbsent) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

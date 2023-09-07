package data

import (
	"time"

	up "github.com/upper/db/v4"
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

// GetAll gets all records from the database, using upper
func (t *EmployeeAbsent) GetAll(condition *up.Cond) ([]*EmployeeAbsent, error) {
	collection := upper.Collection(t.Table())
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

// Get gets one record from the database, by id, using upper
func (t *EmployeeAbsent) Get(id int) (*EmployeeAbsent, error) {
	var one EmployeeAbsent
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *EmployeeAbsent) Update(m EmployeeAbsent) error {
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
func (t *EmployeeAbsent) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeeAbsent) Insert(m EmployeeAbsent) (int, error) {
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

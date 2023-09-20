package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// EmployeesInOrganizationUnit struct
type EmployeesInOrganizationUnit struct {
	ID                           int       `db:"id,omitempty"`
	UserAccountId                int       `db:"user_account_id"`
	UserProfileId                int       `db:"user_profile_id"`
	PositionInOrganizationUnitId int       `db:"position_in_organization_unit_id"`
	Active                       bool      `db:"active"`
	CreatedAt                    time.Time `db:"created_at"`
	UpdatedAt                    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeesInOrganizationUnit) Table() string {
	return "employees_in_organization_units"
}

// GetAll gets all records from the database, using upper
func (t *EmployeesInOrganizationUnit) GetAll(condition *up.Cond) ([]*EmployeesInOrganizationUnit, error) {
	collection := upper.Collection(t.Table())
	var all []*EmployeesInOrganizationUnit
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("updated_at desc").All(&all)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *EmployeesInOrganizationUnit) Get(id int) (*EmployeesInOrganizationUnit, error) {
	var one EmployeesInOrganizationUnit
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *EmployeesInOrganizationUnit) Update(m EmployeesInOrganizationUnit) error {
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
func (t *EmployeesInOrganizationUnit) Delete(id int) error {
	collection := upper.Collection(t.Table())
	var res up.Result
	condition := up.Cond{
		"position_in_organization_unit_id": id,
	}

	res = collection.Find(condition)

	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeesInOrganizationUnit) Insert(m EmployeesInOrganizationUnit) (int, error) {
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

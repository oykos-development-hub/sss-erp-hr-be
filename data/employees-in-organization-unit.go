package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

// EmployeesInOrganizationUnit struct
type EmployeesInOrganizationUnit struct {
	ID                           int       `db:"id,omitempty"`
	UserAccountId                int       `db:"user_account_id"`
	UserProfileId                int       `db:"user_profile_id"`
	PositionInOrganizationUnitId int       `db:"position_in_organization_unit_id"`
	Active                       bool      `db:"active"`
	CreatedAt                    time.Time `db:"created_at,omitempty"`
	UpdatedAt                    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeesInOrganizationUnit) Table() string {
	return "employees_in_organization_units"
}

// GetAll gets all records from the database, using Upper
func (t *EmployeesInOrganizationUnit) GetAll(condition *up.Cond) ([]*EmployeesInOrganizationUnit, error) {
	collection := Upper.Collection(t.Table())
	var all []*EmployeesInOrganizationUnit
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, newErrors.Wrap(err, "upper order by")
		}
	}

	return all, err
}

// Get gets one record from the database, by id, using Upper
func (t *EmployeesInOrganizationUnit) Get(id int) (*EmployeesInOrganizationUnit, error) {
	var one EmployeesInOrganizationUnit
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *EmployeesInOrganizationUnit) Update(m EmployeesInOrganizationUnit) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *EmployeesInOrganizationUnit) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	var res up.Result
	condition := up.Cond{
		"position_in_organization_unit_id": id,
	}

	res = collection.Find(condition)

	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

func (t *EmployeesInOrganizationUnit) DeleteByID(id int) error {
	collection := Upper.Collection(t.Table())
	var res up.Result
	condition := up.Cond{
		"id": id,
	}

	res = collection.Find(condition)

	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *EmployeesInOrganizationUnit) Insert(m EmployeesInOrganizationUnit) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}

func (t *EmployeesInOrganizationUnit) GetEmployeeInActiveSystematization(id int) (*EmployeesInOrganizationUnit, error) {
	var response EmployeesInOrganizationUnit

	query1 := `select id, position_in_organization_unit_id, active from employees_in_organization_units e,
	           job_positions_in_organization_units u, systematizations s 
			   where s.id = u.systematization_id and e.position_in_organization_unit_id = u.id 
			   and s.active = 2 and e.user_profile_id = $1;
`

	rows1, err := Upper.SQL().Query(query1, id)
	if err != nil {
		return &response, newErrors.Wrap(err, "upper exec")
	}
	defer rows1.Close()

	for rows1.Next() {
		err = rows1.Scan(&response.ID, &response.PositionInOrganizationUnitId, &response.Active)
		response.UserProfileId = id

		if err != nil {
			return &response, newErrors.Wrap(err, "upper scan")
		}
	}

	return &response, nil
}

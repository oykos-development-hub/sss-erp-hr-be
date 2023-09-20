package data

import (
	"fmt"
	"time"

	up "github.com/upper/db/v4"
)

// EmployeeContract struct
type EmployeeContract struct {
	ID                 int        `db:"id,omitempty"`
	UserProfileID      int        `db:"user_profile_id"`
	ContractTypeID     int        `db:"contract_type_id"`
	OrganizationUnitID int        `db:"organization_unit_id"`
	DepartmentID       *int       `db:"organization_unit_department_id"`
	Abbreviation       *string    `db:"abbreviation"`
	Description        *string    `db:"description"`
	Active             bool       `db:"active"`
	SerialNumber       *string    `db:"serial_number"`
	NetSalary          *string    `db:"net_salary"`
	GrossSalary        *string    `db:"gross_salary"`
	BankAccount        *string    `db:"bank_account"`
	BankName           *string    `db:"bank_name"`
	DateOfSignature    *time.Time `db:"date_of_signature"`
	DateOfEligibility  *time.Time `db:"date_of_eligibility"`
	DateOfStart        time.Time  `db:"date_of_start"`
	DateOfEnd          *time.Time `db:"date_of_end"`
	FileID             *int       `db:"file_id"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeContract) Table() string {
	return "employee_contracts"
}

// Delete deletes a record from the database by id, using upper
func (t *EmployeeContract) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (t *EmployeeContract) GetByUserProfileId(userProfileId int, condition *up.Cond) ([]*EmployeeContract, error) {
	var employeeContracts []*EmployeeContract

	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"user_profile_id =": userProfileId})

	if condition != nil {
		res = res.And(*condition)
	}

	err := res.OrderBy("updated_at desc").All(&employeeContracts)
	if err != nil {
		fmt.Println(err)
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	return employeeContracts, nil
}

// Delete deletes a records from the database for user id, using upper
func (t *EmployeeContract) DeleteForUser(userID int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"user_profile_id": userID})
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeeContract) Insert(m EmployeeContract) (int, error) {
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

// Update updates a record in the database, using upper
func (t *EmployeeContract) Update(c EmployeeContract) error {
	c.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(c.ID)
	err := res.Update(&c)
	if err != nil {
		return err
	}
	return nil
}

// Get gets one record from the database, by id, using upper
func (t *EmployeeContract) Get(id int) (*EmployeeContract, error) {
	var one EmployeeContract
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}

	return &one, nil
}

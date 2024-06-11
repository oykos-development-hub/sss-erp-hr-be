package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// EmployeeFamilyMember struct
type EmployeeFamilyMember struct {
	ID                   int       `db:"id,omitempty"`
	UserProfileID        int       `db:"user_profile_id"`
	FirstName            string    `db:"first_name"`
	MiddleName           *string   `db:"middle_name"`
	LastName             string    `db:"last_name"`
	BirthLastName        *string   `db:"birth_last_name"`
	FatherName           string    `db:"father_name"`
	MotherName           string    `db:"mother_name"`
	MotherBirthLastName  string    `db:"mother_birth_last_name"`
	DateOfBirth          time.Time `db:"date_of_birth"`
	CountryOfBirth       string    `db:"country_of_birth"`
	CityOfBirth          string    `db:"city_of_birth"`
	Nationality          string    `db:"nationality"`
	NationalMinority     *string   `db:"national_minority"`
	Citizenship          string    `db:"citizenship"`
	Address              string    `db:"address"`
	OfficialPersonalID   string    `db:"official_personal_id"`
	Gender               string    `db:"gender"`
	EmployeeRelationship string    `db:"employee_relationship"`
	InsuranceCoverage    string    `db:"insurance_coverage"`
	HandicappedPerson    bool      `db:"handicapped_person"`
	CreatedAt            time.Time `db:"created_at,omitempty"`
	UpdatedAt            time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeFamilyMember) Table() string {
	return "employee_family_members"
}

// GetAll gets all records from the database, using Upper
func (t *EmployeeFamilyMember) GetAll(condition *up.Cond) ([]*EmployeeFamilyMember, error) {
	collection := Upper.Collection(t.Table())
	var all []*EmployeeFamilyMember
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
func (t *EmployeeFamilyMember) Get(id int) (*EmployeeFamilyMember, error) {
	var one EmployeeFamilyMember
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *EmployeeFamilyMember) Update(m EmployeeFamilyMember) error {
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
func (t *EmployeeFamilyMember) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *EmployeeFamilyMember) Insert(m EmployeeFamilyMember) (int, error) {
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

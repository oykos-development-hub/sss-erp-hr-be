package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// EmployeeEducation struct
type EmployeeEducation struct {
	ID                  int        `db:"id,omitempty"`
	UserProfileID       int        `db:"user_profile_id"`
	TypeID              int        `db:"type_id"`
	DateOfCertification *time.Time `db:"date_of_certification"`
	Price               *float32   `db:"price"`
	DateOfStart         *time.Time `db:"date_of_start"`
	DateOfEnd           *time.Time `db:"date_of_end"`
	AcademicTitle       *string    `db:"academic_title"`
	ExpertiseLevel      *string    `db:"expertise_level"`
	CertificateIssuer   *string    `db:"certificate_issuer"`
	Title               *string    `db:"title"`
	Description         *string    `db:"description"`
	FileId              *int       `db:"file_id"`
	CreatedAt           time.Time  `db:"created_at,omitempty"`
	UpdatedAt           time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeEducation) Table() string {
	return "employee_educations"
}

// GetAll gets all records from the database, using upper
func (t *EmployeeEducation) GetAll(condition *up.Cond) ([]*EmployeeEducation, error) {
	collection := upper.Collection(t.Table())
	var all []*EmployeeEducation
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("updated_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *EmployeeEducation) Get(id int) (*EmployeeEducation, error) {
	var one EmployeeEducation
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *EmployeeEducation) Update(m EmployeeEducation) error {
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
func (t *EmployeeEducation) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeeEducation) Insert(m EmployeeEducation) (int, error) {
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

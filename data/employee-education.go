package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
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
	Score               *string    `db:"score"`
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

// GetAll gets all records from the database, using Upper
func (t *EmployeeEducation) GetAll(condition *up.Cond) ([]*EmployeeEducation, error) {
	collection := Upper.Collection(t.Table())
	var all []*EmployeeEducation
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
func (t *EmployeeEducation) Get(id int) (*EmployeeEducation, error) {
	var one EmployeeEducation
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *EmployeeEducation) Update(m EmployeeEducation) error {
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
func (t *EmployeeEducation) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *EmployeeEducation) Insert(m EmployeeEducation) (int, error) {
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

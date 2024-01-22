package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// EmployeeResolution struct
type EmployeeResolution struct {
	ID                int        `db:"id,omitempty"`
	ResolutionTypeID  int        `db:"resolution_type_id"`
	UserProfileID     int        `db:"user_profile_id"`
	ResolutionPurpose *string    `db:"resolution_purpose"`
	DateOfStart       time.Time  `db:"date_of_start"`
	DateOfEnd         *time.Time `db:"date_of_end"`
	Year              int        `db:"year"`
	Value             *string    `db:"value"`
	IsAffect          *bool      `db:"is_affect"`
	FileID            *int       `db:"file_id"`
	CreatedAt         time.Time  `db:"created_at,omitempty"`
	UpdatedAt         time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *EmployeeResolution) Table() string {
	return "employee_resolutions"
}

// GetAll gets all records from the database, using upper
func (t *EmployeeResolution) GetAll(condition *up.Cond) ([]*EmployeeResolution, error) {
	collection := upper.Collection(t.Table())
	var all []*EmployeeResolution
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
func (t *EmployeeResolution) Get(id int) (*EmployeeResolution, error) {
	var one EmployeeResolution
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *EmployeeResolution) Update(m EmployeeResolution) error {
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
func (t *EmployeeResolution) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *EmployeeResolution) Insert(m EmployeeResolution) (int, error) {
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

package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Systematization struct
type Systematization struct {
	ID                 int        `db:"id,omitempty"`
	UserProfileID      int        `db:"user_profile_id"`
	OrganizationUnitID int        `db:"organization_unit_id"`
	Description        string     `db:"description"`
	SerialNumber       string     `db:"serial_number"`
	Active             bool       `db:"active"`
	DateOfActivation   *time.Time `db:"date_of_activation"`
	FileId             *int       `db:"file_id"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Systematization) Table() string {
	return "systematizations"
}

// GetAll gets all records from the database, using upper
func (t *Systematization) GetAll(page *int, pageSize *int, condition *up.AndExpr) ([]*Systematization, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Systematization
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Systematization) Get(id int) (*Systematization, error) {
	var one Systematization
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Systematization) Update(m Systematization) error {
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
func (t *Systematization) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Systematization) Insert(m Systematization) (int, error) {
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

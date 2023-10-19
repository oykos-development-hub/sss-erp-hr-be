package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// RevisionsInOrganizationUnit struct
type RevisionsInOrganizationUnit struct {
	ID                 int       `db:"id,omitempty"`
	RevisionID         int       `db:"revision_id"`
	OrganizationUnitID int       `db:"organization_unit_id"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *RevisionsInOrganizationUnit) Table() string {
	return "revisions_in_organization_units"
}

// GetAll gets all records from the database, using upper
func (t *RevisionsInOrganizationUnit) GetAll(condition *up.Cond) ([]*RevisionsInOrganizationUnit, error) {
	collection := upper.Collection(t.Table())
	var all []*RevisionsInOrganizationUnit
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *RevisionsInOrganizationUnit) Get(id int) (*RevisionsInOrganizationUnit, error) {
	var one RevisionsInOrganizationUnit
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *RevisionsInOrganizationUnit) Update(m RevisionsInOrganizationUnit) error {
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
func (t *RevisionsInOrganizationUnit) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *RevisionsInOrganizationUnit) Insert(m RevisionsInOrganizationUnit) (int, error) {
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

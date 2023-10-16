package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Revision struct
type Revision struct {
	ID                      int       `db:"id,omitempty"`
	Title                   string    `db:"title"`
	PlanID                  int       `db:"plan_id"`
	SerialNumber            string    `db:"serial_number"`
	DateOfRevision          time.Time `db:"date_of_revision"`
	RevisionQuartal         string    `db:"revision_quartal"`
	InternalRevisionSubject []int     `db:"internal_revision_subject"`
	ExternalRevisionSubject *int      `db:"external_revision_subject"`
	Revisor                 []int     `db:"revisor_id"`
	RevisionType            int       `db:"revision_type"`
	FileID                  *int      `db:"file_id"`
	CreatedAt               time.Time `db:"created_at,omitempty"`
	UpdatedAt               time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Revision) Table() string {
	return "revisions"
}

// GetAll gets all records from the database, using upper
func (t *Revision) GetAll(Page *int, Size *int, condition *up.Cond) ([]*Revision, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Revision
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()

	if err != nil {
		return nil, nil, err
	}

	if Page != nil && Size != nil {
		res = paginateResult(res, *Page, *Size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Revision) Get(id int) (*Revision, error) {
	var one Revision
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Revision) Update(m Revision) error {
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
func (t *Revision) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Revision) Insert(m Revision) (int, error) {
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

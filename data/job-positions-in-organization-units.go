package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// JobPosition struct
type JobPositionsInOrganizationUnits struct {
	ID                       int       `db:"id,omitempty"`
	SystematizationID        int       `db:"systematization_id"`
	ParentOrganizationUnitID int       `db:"parent_organization_unit_id"`
	JobPositionID            int       `db:"job_position_id"`
	AvailableSlots           int       `db:"available_slots"`
	CreatedAt                time.Time `db:"created_at"`
	UpdatedAt                time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JobPositionsInOrganizationUnits) Table() string {
	return "job_positions_in_organization_units"
}

// Delete deletes a record from the database by id, using upper
func (t *JobPositionsInOrganizationUnits) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Get gets one record from the database, by id, using upper
func (t *JobPositionsInOrganizationUnits) Get(id int) (*JobPositionsInOrganizationUnits, error) {
	var one JobPositionsInOrganizationUnits
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Insert inserts a model into the database, using upper
func (t *JobPositionsInOrganizationUnits) Insert(m JobPositionsInOrganizationUnits) (int, error) {
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

// GetAll gets all records from the database, using upper
func (t *JobPositionsInOrganizationUnits) GetAll(page *int, pageSize *int, condition *up.Cond) ([]*JobPositionsInOrganizationUnits, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*JobPositionsInOrganizationUnits
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

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

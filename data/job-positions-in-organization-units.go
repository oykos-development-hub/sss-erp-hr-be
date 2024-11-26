package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

// JobPosition struct
type JobPositionsInOrganizationUnits struct {
	ID                       int       `db:"id,omitempty"`
	SystematizationID        int       `db:"systematization_id"`
	ParentOrganizationUnitID int       `db:"parent_organization_unit_id"`
	JobPositionID            int       `db:"job_position_id"`
	AvailableSlots           int       `db:"available_slots"`
	Requirements             *string   `db:"requirements"`
	Description              *string   `db:"description"`
	CreatedAt                time.Time `db:"created_at,omitempty"`
	UpdatedAt                time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JobPositionsInOrganizationUnits) Table() string {
	return "job_positions_in_organization_units"
}

// Delete deletes a record from the database by id, using Upper
func (t *JobPositionsInOrganizationUnits) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Get gets one record from the database, by id, using Upper
func (t *JobPositionsInOrganizationUnits) Get(id int) (*JobPositionsInOrganizationUnits, error) {
	var one JobPositionsInOrganizationUnits
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Insert inserts a model into the database, using Upper
func (t *JobPositionsInOrganizationUnits) Insert(m JobPositionsInOrganizationUnits) (int, error) {
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

func (t *JobPositionsInOrganizationUnits) Update(m JobPositionsInOrganizationUnits) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// GetAll gets all records from the database, using Upper
func (t *JobPositionsInOrganizationUnits) GetAll(page *int, pageSize *int, condition *up.AndExpr) ([]*JobPositionsInOrganizationUnits, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*JobPositionsInOrganizationUnits
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.OrderBy("id").All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order by")
	}

	return all, &total, err
}

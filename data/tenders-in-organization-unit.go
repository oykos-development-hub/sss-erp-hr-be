package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// TendersInOrganizationUnit struct
type TendersInOrganizationUnit struct {
	ID                           int       `db:"id,omitempty"`
	PositionInOrganizationUnitID int       `db:"position_in_organization_unit_id"`
	Active                       bool      `db:"active"`
	Type                         int       `db:"type"`
	DateOfStart                  time.Time `db:"date_of_start"`
	DateOfEnd                    time.Time `db:"date_of_end"`
	Description                  string    `db:"description"`
	SerialNumber                 string    `db:"serial_number"`
	AvailableSlots               int       `db:"available_slots"`
	FileID                       int       `db:"file_id"`
	CreatedAt                    time.Time `db:"created_at"`
	UpdatedAt                    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *TendersInOrganizationUnit) Table() string {
	return "tenders_in_organization_units"
}

// GetAll gets all records from the database, using upper
func (t *TendersInOrganizationUnit) GetAll(page *int, size *int, condition *up.Cond) ([]*TendersInOrganizationUnit, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*TendersInOrganizationUnit
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

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *TendersInOrganizationUnit) Get(id int) (*TendersInOrganizationUnit, error) {
	var one TendersInOrganizationUnit
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *TendersInOrganizationUnit) Update(m TendersInOrganizationUnit) error {
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
func (t *TendersInOrganizationUnit) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *TendersInOrganizationUnit) Insert(m TendersInOrganizationUnit) (int, error) {
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

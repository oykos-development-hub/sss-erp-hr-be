package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// AbsentType struct
type AbsentType struct {
	ID                int       `db:"id,omitempty"`
	ParentID          *int      `db:"parent_id"`
	Title             string    `db:"title"`
	Abbreviation      string    `db:"abbreviation"`
	AccountingDaysOff bool      `db:"accounting_days_off"`
	Relocation        bool      `db:"relocation"`
	Description       *string   `db:"description,omitempty"`
	Color             *string   `db:"color,omitempty"`
	Icon              *string   `db:"icon,omitempty"`
	CreatedAt         time.Time `db:"created_at,omitempty"`
	UpdatedAt         time.Time `db:"updated_at,omitempty"`
}

// Table returns the table name
func (t *AbsentType) Table() string {
	return "absent_types"
}

// GetAll gets all records from the database, using upper
func (t *AbsentType) GetAll(page *int, size *int, condition *up.Cond) ([]*AbsentType, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*AbsentType
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

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *AbsentType) Get(id int) (*AbsentType, error) {
	var one AbsentType
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *AbsentType) Update(m AbsentType) error {
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
func (t *AbsentType) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *AbsentType) Insert(m AbsentType) (int, error) {
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

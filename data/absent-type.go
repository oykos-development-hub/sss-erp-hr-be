package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

// AbsentType struct
type AbsentType struct {
	ID                int       `db:"id,omitempty"`
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

// GetAll gets all records from the database, using Upper
func (t *AbsentType) GetAll(page *int, size *int, condition *up.Cond) ([]*AbsentType, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*AbsentType
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order by")
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *AbsentType) Get(id int) (*AbsentType, error) {
	var one AbsentType
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *AbsentType) Update(m AbsentType) error {
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
func (t *AbsentType) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *AbsentType) Insert(m AbsentType) (int, error) {
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

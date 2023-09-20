package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// JudgeNumberResolution struct
type JudgeNumberResolution struct {
	ID           int       `db:"id,omitempty"`
	Active       bool      `db:"active"`
	SerialNumber string    `db:"serial_number"`
	Year         string    `db:"year"`
	CreatedAt    time.Time `db:"created_at,omitempty"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JudgeNumberResolution) Table() string {
	return "judge_number_resolutions"
}

func (t *JudgeNumberResolution) GetAll(page *int, pageSize *int, condition *up.Cond) ([]*JudgeNumberResolution, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*JudgeNumberResolution
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

	err = res.OrderBy("updated_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *JudgeNumberResolution) Get(id int) (*JudgeNumberResolution, error) {
	var one JudgeNumberResolution
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *JudgeNumberResolution) Update(m JudgeNumberResolution) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Inactivate all resolutions excepct the one that is activated
// InactivateOtherResolutions inactivates all JudgeNumberResolutions except the one with the provided ID.
func (t *JudgeNumberResolution) InactivateOtherResolutions(id int) error {
	q := upper.SQL().Update(t.Table()).
		Where(up.Cond{"id !=": id}).
		And(up.Cond{"active": true}).
		Set("active", false)

	_, err := q.Exec()

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *JudgeNumberResolution) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *JudgeNumberResolution) Insert(m JudgeNumberResolution) (int, error) {
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

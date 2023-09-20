package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// JobPosition struct
type JobPosition struct {
	ID               int       `db:"id,omitempty"`
	Title            string    `db:"title"`
	Abbreviation     string    `db:"abbreviation"`
	SerialNumber     string    `db:"serial_number"`
	Description      *string   `db:"description"`
	Requirements     string    `db:"requirements"`
	IsJudge          *bool     `db:"is_judge"`
	IsJudgePresident *bool     `db:"is_judge_president"`
	Color            *string   `db:"color"`
	Icon             *string   `db:"icon"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JobPosition) Table() string {
	return "job_positions"
}

// GetAll gets all records from the database, using upper
func (t *JobPosition) GetAll(page *int, pageSize *int, conditions *up.AndExpr) ([]*JobPosition, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*JobPosition
	var res up.Result

	if conditions != nil {
		res = collection.Find(conditions)
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
func (t *JobPosition) Get(id int) (*JobPosition, error) {
	var one JobPosition
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *JobPosition) Update(m JobPosition) error {
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
func (t *JobPosition) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *JobPosition) Insert(m JobPosition) (int, error) {
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

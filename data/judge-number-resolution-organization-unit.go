package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// JudgeNumberResolutionOrganizationUnit struct
type JudgeNumberResolutionOrganizationUnit struct {
	ID                 int       `db:"id,omitempty"`
	Resoultion         int       `db:"resolution_id"`
	OrganizationUnitID int       `db:"organization_unit_id"`
	NumberOfJudges     int       `db:"number_of_judges"`
	NumberOfPresidents int       `db:"number_of_presidents"`
	CreatedAt          time.Time `db:"created_at,omitempty"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JudgeNumberResolutionOrganizationUnit) Table() string {
	return "judge_number_resolution_organization_units"
}

func (t *JudgeNumberResolutionOrganizationUnit) GetAll(page *int, pageSize *int, condition *up.Cond) ([]*JudgeNumberResolutionOrganizationUnit, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*JudgeNumberResolutionOrganizationUnit
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

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *JudgeNumberResolutionOrganizationUnit) Get(id int) (*JudgeNumberResolutionOrganizationUnit, error) {
	var one JudgeNumberResolutionOrganizationUnit
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *JudgeNumberResolutionOrganizationUnit) Update(m JudgeNumberResolutionOrganizationUnit) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *JudgeNumberResolutionOrganizationUnit) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *JudgeNumberResolutionOrganizationUnit) Insert(m JudgeNumberResolutionOrganizationUnit) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

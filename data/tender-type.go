package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// TenderType struct
type TenderType struct {
	ID               int       `db:"id,omitempty"`
	Title            string    `db:"title"`
	Abbreviation     *string   `db:"abbreviation"`
	IsJudge          bool      `db:"is_judge"`
	IsJudgePresident bool      `db:"is_judge_president"`
	Description      *string   `db:"description"`
	Color            *string   `db:"color"`
	Value            *string   `db:"value"`
	Icon             *string   `db:"icon"`
	CreatedAt        time.Time `db:"created_at,omitempty"`
	UpdatedAt        time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *TenderType) Table() string {
	return "tender_types"
}

// GetAll gets all records from the database, using Upper
func (t *TenderType) GetAll(condition *up.AndExpr) ([]*TenderType, error) {
	collection := Upper.Collection(t.Table())
	var all []*TenderType
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using Upper
func (t *TenderType) Get(id int) (*TenderType, error) {
	var one TenderType
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *TenderType) Update(m TenderType) error {
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
func (t *TenderType) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *TenderType) Insert(m TenderType) (int, error) {
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

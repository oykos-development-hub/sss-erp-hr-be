package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Evaluation struct
type Evaluation struct {
	ID               int        `db:"id,omitempty"`
	UserProfileID    int        `db:"user_profile_id"`
	EvaluationTypeID int        `db:"evaluation_type_id"`
	Score            string     `db:"score"`
	DateOfEvaluation *time.Time `db:"date_of_evaluation"`
	Evaluator        string     `db:"evaluator"`
	IsRelevant       *bool      `db:"is_relevant"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	FileID           *int       `db:"file_id"`
}

// Table returns the table name
func (t *Evaluation) Table() string {
	return "evaluations"
}

// GetAll gets all records from the database, using upper
func (t *Evaluation) GetAll(condition *up.Cond) ([]*Evaluation, error) {
	collection := upper.Collection(t.Table())
	var all []*Evaluation
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("updated_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Evaluation) Get(id int) (*Evaluation, error) {
	var one Evaluation
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Evaluation) Update(m Evaluation) error {
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
func (t *Evaluation) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Evaluation) Insert(m Evaluation) (int, error) {
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

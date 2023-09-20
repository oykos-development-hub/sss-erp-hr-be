package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// UserNormFulfilment struct
type UserNormFulfilment struct {
	ID                       int       `db:"id,omitempty"`
	UserProfileID            int       `db:"user_profile_id"`
	Topic                    string    `db:"topic"`
	Title                    string    `db:"title"`
	NumberOfNormDecrease     int       `db:"number_of_norm_decrease"`
	NumberOfItems            int       `db:"number_of_items"`
	NumberOfItemsSolved      int       `db:"number_of_items_solved"`
	EvaluationID             *int      `db:"evaluation_id"`
	DateOfEvaluation         time.Time `db:"date_of_evaluation"`
	DateOfEvaluationValidity time.Time `db:"date_of_evaluation_validity"`
	RelocationID             *int      `db:"relocation_id"`
	FileID                   *int      `db:"file_id"`
	CreatedAt                time.Time `db:"created_at,omitempty"`
	UpdatedAt                time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *UserNormFulfilment) Table() string {
	return "user_norm_fulfilments"
}

// GetAll gets all records from the database, using upper
func (t *UserNormFulfilment) GetAll(condition *up.Cond) ([]*UserNormFulfilment, error) {
	collection := upper.Collection(t.Table())
	var all []*UserNormFulfilment
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
func (t *UserNormFulfilment) Get(id int) (*UserNormFulfilment, error) {
	var one UserNormFulfilment
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *UserNormFulfilment) Update(m UserNormFulfilment) error {
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
func (t *UserNormFulfilment) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *UserNormFulfilment) Insert(m UserNormFulfilment) (int, error) {
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

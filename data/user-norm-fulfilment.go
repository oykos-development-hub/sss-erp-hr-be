package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
)

// UserNormFulfilment struct
type UserNormFulfilment struct {
	ID                       int        `db:"id,omitempty"`
	UserProfileID            int        `db:"user_profile_id"`
	Topic                    string     `db:"topic"`
	Title                    string     `db:"title"`
	NumberOfNormDecrease     int        `db:"number_of_norm_decrease"`
	NumberOfItems            int        `db:"number_of_items"`
	NumberOfItemsSolved      int        `db:"number_of_items_solved"`
	EvaluationID             *int       `db:"evaluation_id"`
	DateOfEvaluation         *time.Time `db:"date_of_evaluation"`
	DateOfEvaluationValidity *time.Time `db:"date_of_evaluation_validity"`
	NormStartDate            string     `db:"norm_start_date"`
	NormEndDate              string     `db:"norm_end_date"`
	RelocationID             *int       `db:"relocation_id"`
	FileID                   *int       `db:"file_id"`
	CreatedAt                time.Time  `db:"created_at,omitempty"`
	UpdatedAt                time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *UserNormFulfilment) Table() string {
	return "user_norm_fulfilments"
}

// GetAll gets all records from the database, using Upper
func (t *UserNormFulfilment) GetAll(condition *up.AndExpr) ([]*UserNormFulfilment, error) {
	collection := Upper.Collection(t.Table())
	var all []*UserNormFulfilment
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
func (t *UserNormFulfilment) Get(id int) (*UserNormFulfilment, error) {
	var one UserNormFulfilment
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *UserNormFulfilment) Update(ctx context.Context, m UserNormFulfilment) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *UserNormFulfilment) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *UserNormFulfilment) Insert(ctx context.Context, m UserNormFulfilment) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

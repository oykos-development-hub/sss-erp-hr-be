package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

// JobPosition struct
type JobPosition struct {
	ID               int       `db:"id,omitempty"`
	Title            string    `db:"title"`
	Abbreviation     string    `db:"abbreviation"`
	SerialNumber     string    `db:"serial_number"`
	Description      *string   `db:"description"`
	Requirements     string    `db:"requirements"`
	IsActive         bool      `db:"is_active"`
	IsJudge          *bool     `db:"is_judge"`
	IsJudgePresident *bool     `db:"is_judge_president"`
	Color            *string   `db:"color"`
	Icon             *string   `db:"icon"`
	CreatedAt        time.Time `db:"created_at,omitempty"`
	UpdatedAt        time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *JobPosition) Table() string {
	return "job_positions"
}

// GetAll gets all records from the database, using Upper
func (t *JobPosition) GetAll(page *int, pageSize *int, conditions *up.AndExpr) ([]*JobPosition, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*JobPosition
	var res up.Result

	if conditions != nil {
		res = collection.Find(conditions)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order by")
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *JobPosition) Get(id int) (*JobPosition, error) {
	var one JobPosition
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *JobPosition) Update(ctx context.Context, m JobPosition) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return newErrors.Wrap(err, "upper update")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *JobPosition) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return newErrors.Wrap(err, "upper delete")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *JobPosition) Insert(ctx context.Context, m JobPosition) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return 0, newErrors.Wrap(err, "context get user id")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return newErrors.Wrap(err, "upper insert")
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

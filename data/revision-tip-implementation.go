package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
	newErrors "gitlab.sudovi.me/erp/hr-ms-api/pkg/errors"
)

const StatusTipImplFinished = "Sprovedena"
const StatusTipImplNotFinished = "Nesprovedena"

// RevisionTipImplementation struct
type RevisionTipImplementation struct {
	ID                     int           `db:"id,omitempty"`
	TipID                  int           `db:"tip_id"`
	Status                 string        `db:"status"`
	NewDueDate             *int          `db:"new_due_date"`
	NewDateOfExecution     *time.Time    `db:"new_date_of_execution"`
	ReasonsForNonExecuting *string       `db:"reasons_for_non_executing"`
	RevisorID              *int          `db:"revisor_id"`
	Documents              string        `db:"documents"`
	FileIDs                pq.Int64Array `db:"file_ids"`
	CreatedAt              time.Time     `db:"created_at,omitempty"`
	UpdatedAt              time.Time     `db:"updated_at"`
}

// Table returns the table name
func (t *RevisionTipImplementation) Table() string {
	return "revision_tip_implementations"
}

// GetAll gets all records from the database, using Upper
func (t *RevisionTipImplementation) GetAll(condition *up.Cond) ([]*RevisionTipImplementation, error) {
	collection := Upper.Collection(t.Table())
	var all []*RevisionTipImplementation
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper order by")
	}

	return all, err
}

// Get gets one record from the database, by id, using Upper
func (t *RevisionTipImplementation) Get(id int) (*RevisionTipImplementation, error) {
	var one RevisionTipImplementation
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *RevisionTipImplementation) Update(ctx context.Context, m RevisionTipImplementation) error {
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
func (t *RevisionTipImplementation) Delete(ctx context.Context, id int) error {
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
func (t *RevisionTipImplementation) Insert(ctx context.Context, m RevisionTipImplementation) (int, error) {
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

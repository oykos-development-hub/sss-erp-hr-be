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

// RevisionTip struct
type RevisionTip struct {
	ID                     int        `db:"id,omitempty"`
	RevisionID             int        `db:"revision_id"`
	UserProfileID          *int       `db:"user_profile_id"`
	DateOfAccept           *time.Time `db:"date_of_accept"`
	DueDate                int        `db:"due_date"`
	NewDueDate             *int       `db:"new_due_date"`
	RevisionPriority       *string    `db:"revision_priority"`
	DateOfReject           *time.Time `db:"date_of_reject"`
	EndDate                *time.Time `db:"end_date"`
	DateOfExecution        *time.Time `db:"date_of_execution"`
	NewDateOfExecution     *time.Time `db:"new_date_of_execution"`
	Recommendation         string     `db:"recommendation"`
	Status                 *string    `db:"status"`
	Documents              *string    `db:"documents"`
	ReasonsForNonExecuting *string    `db:"reasons_for_non_executing"`
	FileID                 *int       `db:"file_id"`
	ResponsiblePerson      *string    `db:"responsible_person"`
	CreatedAt              time.Time  `db:"created_at,omitempty"`
	UpdatedAt              time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *RevisionTip) Table() string {
	return "revision_tips"
}

// GetAll gets all records from the database, using Upper
func (t *RevisionTip) GetAll(page *int, size *int, condition *up.Cond) ([]*RevisionTip, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*RevisionTip
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
func (t *RevisionTip) Get(id int) (*RevisionTip, error) {
	var one RevisionTip
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *RevisionTip) Update(ctx context.Context, m RevisionTip) error {
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
func (t *RevisionTip) Delete(ctx context.Context, id int) error {
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
func (t *RevisionTip) Insert(ctx context.Context, m RevisionTip) (int, error) {
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

package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/contextutil"
)

// Revision struct
type Revision struct {
	ID                      int           `db:"id,omitempty"`
	Title                   string        `db:"title"`
	PlanID                  int           `db:"plan_id"`
	SerialNumber            string        `db:"serial_number"`
	DateOfRevision          time.Time     `db:"date_of_revision"`
	RevisionQuartal         string        `db:"revision_quartal"`
	InternalRevisionSubject pq.Int64Array `db:"internal_revision_subject_id"`
	ExternalRevisionSubject *int          `db:"external_revision_subject"`
	Revisor                 pq.Int64Array `db:"revisor_id"`
	RevisionType            int           `db:"revision_type_id"`
	FileID                  *int          `db:"file_id"`
	TipsFileID              *int          `db:"tips_file_id"`
	CreatedAt               time.Time     `db:"created_at,omitempty"`
	UpdatedAt               time.Time     `db:"updated_at"`
}

// Table returns the table name
func (t *Revision) Table() string {
	return "revisions"
}

// GetAll gets all records from the database, using Upper
func (t *Revision) GetAll(Page *int, Size *int, condition *up.Cond) ([]*Revision, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Revision
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

	if Page != nil && Size != nil {
		res = paginateResult(res, *Page, *Size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *Revision) Get(id int) (*Revision, error) {
	var one Revision
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Revision) Update(ctx context.Context, m Revision) error {
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
func (t *Revision) Delete(ctx context.Context, id int) error {
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
func (t *Revision) Insert(ctx context.Context, m Revision) (int, error) {
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

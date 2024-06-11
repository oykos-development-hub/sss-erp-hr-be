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

// OrganizationUnit struct
type OrganizationUnit struct {
	ID             int            `db:"id,omitempty"`
	ParentID       *int           `db:"parent_id"`
	Title          string         `db:"title"`
	Pib            *string        `db:"pib"`
	Abbreviation   *string        `db:"abbreviation"`
	NumberOfJudges *int           `db:"number_of_judges"`
	OrderID        *int           `db:"order_id"`
	Color          *string        `db:"color"`
	Icon           *string        `db:"icon"`
	Address        *string        `db:"address"`
	City           *string        `db:"city"`
	Description    *string        `db:"description"`
	FolderID       *int           `db:"folder_id"`
	BankAccounts   pq.StringArray `db:"bank_accounts"`
	Code           string         `db:"code"`
	CreatedAt      time.Time      `db:"created_at,omitempty"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

// Table returns the table name
func (t *OrganizationUnit) Table() string {
	return "organization_units"
}

// GetAll gets all records from the database, using Upper
func (t *OrganizationUnit) GetAll(page *int, pageSize *int, conditions *up.AndExpr) ([]*OrganizationUnit, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*OrganizationUnit
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

	err = res.OrderBy("order_id asc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *OrganizationUnit) Get(id int) (*OrganizationUnit, error) {
	var one OrganizationUnit
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *OrganizationUnit) Update(ctx context.Context, m OrganizationUnit) error {
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
func (t *OrganizationUnit) Delete(ctx context.Context, id int) error {
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
func (t *OrganizationUnit) Insert(ctx context.Context, m OrganizationUnit) (int, error) {
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

package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// OrganizationUnit struct
type OrganizationUnit struct {
	ID             int       `db:"id,omitempty"`
	ParentID       *int      `db:"parent_id"`
	Title          string    `db:"title"`
	Pib            *string   `db:"pib"`
	Abbreviation   *string   `db:"abbreviation"`
	NumberOfJudges *int      `db:"number_of_judges"`
	OrderID        *int      `db:"order_id"`
	Color          *string   `db:"color"`
	Icon           *string   `db:"icon"`
	Address        *string   `db:"address"`
	City           *string   `db:"city"`
	Description    *string   `db:"description"`
	FolderID       *int      `db:"folder_id"`
	CreatedAt      time.Time `db:"created_at,omitempty"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *OrganizationUnit) Table() string {
	return "organization_units"
}

// GetAll gets all records from the database, using upper
func (t *OrganizationUnit) GetAll(page *int, pageSize *int, conditions *up.AndExpr) ([]*OrganizationUnit, *uint64, error) {
	collection := upper.Collection(t.Table())
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

	err = res.OrderBy("order_id desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *OrganizationUnit) Get(id int) (*OrganizationUnit, error) {
	var one OrganizationUnit
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *OrganizationUnit) Update(m OrganizationUnit) error {
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
func (t *OrganizationUnit) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *OrganizationUnit) Insert(m OrganizationUnit) (int, error) {
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

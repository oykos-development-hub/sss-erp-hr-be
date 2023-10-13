package data

import (
	"time"

	up "github.com/upper/db/v4"
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

// GetAll gets all records from the database, using upper
func (t *RevisionTip) GetAll(page *int, size *int, condition *up.Cond) ([]*RevisionTip, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*RevisionTip
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

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *RevisionTip) Get(id int) (*RevisionTip, error) {
	var one RevisionTip
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *RevisionTip) Update(m RevisionTip) error {
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
func (t *RevisionTip) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *RevisionTip) Insert(m RevisionTip) (int, error) {
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

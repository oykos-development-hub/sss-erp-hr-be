package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Foreigner struct
type Foreigner struct {
	ID                              int        `db:"id,omitempty"`
	UserProfileId                   int        `db:"user_profile_id"`
	WorkPermitNumber                string     `db:"work_permit_number"`
	WorkPermitIssuer                string     `db:"work_permit_issuer"`
	WorkPermitDateOfStart           time.Time  `db:"work_permit_date_of_start"`
	WorkPermitDateOfEnd             *time.Time `db:"work_permit_date_of_end"`
	WorkPermitIndefiniteLength      *bool      `db:"work_permit_indefinite_length"`
	ResidencePermitDateOfEnd        *time.Time `db:"residence_permit_date_of_end"`
	ResidencePermitIndefiniteLength *bool      `db:"residence_permit_indefinite_length"`
	ResidencePermitNumber           string     `db:"residence_permit_number"`
	CountryOfOrigin                 string     `db:"country_of_origin"`
	WorkPermitFileId                *int       `db:"work_permit_file_id"`
	ResidencePermitFileId           *int       `db:"residence_permit_file_id"`
	CreatedAt                       time.Time  `db:"created_at,omitempty"`
	UpdatedAt                       time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Foreigner) Table() string {
	return "foreigners"
}

// GetAll gets all records from the database, using Upper
func (t *Foreigner) GetAll(condition *up.Cond) ([]*Foreigner, error) {
	collection := Upper.Collection(t.Table())
	var all []*Foreigner
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil && err != up.ErrWarnSlowQuery {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using Upper
func (t *Foreigner) Get(id int) (*Foreigner, error) {
	var one Foreigner
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Foreigner) Update(m Foreigner) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Foreigner) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Foreigner) Insert(m Foreigner) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

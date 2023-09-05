package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type UserProfile struct {
	ID                        int        `db:"id,omitempty"`
	UserAccountId             int        `db:"user_account_id,omitempty"`
	FirstName                 string     `db:"first_name"`
	MiddleName                *string    `db:"middle_name"`
	LastName                  string     `db:"last_name"`
	BirthLastName             *string    `db:"birth_last_name"`
	FatherName                string     `db:"father_name"`
	MotherName                string     `db:"mother_name"`
	MotherBirthLastName       *string    `db:"mother_birth_last_name"`
	DateOfBirth               time.Time  `db:"date_of_birth"`
	CountryOfBirth            string     `db:"country_of_birth"`
	CityOfBirth               string     `db:"city_of_birth"`
	Nationality               string     `db:"nationality"`
	NationalMinority          *string    `db:"national_minority"`
	Citizenship               string     `db:"citizenship"`
	Address                   string     `db:"address"`
	BankAccount               *string    `db:"bank_account"`
	BankName                  *string    `db:"bank_name"`
	OfficialPersonalID        string     `db:"official_personal_id"`
	OfficialPersonalDocNumber string     `db:"official_personal_document_number"`
	OfficialPersonalDocIssuer string     `db:"official_personal_document_issuer"`
	Gender                    string     `db:"gender"`
	SingleParent              *bool      `db:"single_parent"`
	HousingDone               *bool      `db:"housing_done"`
	HousingDescription        string     `db:"housing_description"`
	MartialStatus             string     `db:"martial_status"`
	DateOfTakingOath          *time.Time `db:"date_of_taking_oath"`
	DateOfBecomingJudge       *string    `db:"date_of_becoming_judge"`
	ActiveContract            *bool      `db:"active_contract"`
	RevisorRole               *bool      `db:"revisor_role"`
	EngagementTypeID          *int       `db:"engagement_type_id,omitempty"`
	CreatedAt                 time.Time  `db:"created_at"`
	UpdatedAt                 time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *UserProfile) Table() string {
	return "user_profiles"
}

// GetAll gets all records from the database, using upper
func (t *UserProfile) GetAll(page *int, pageSize *int, condition *up.Cond) ([]*UserProfile, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*UserProfile
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

	if page != nil && pageSize != nil {
		res = paginateResult(res, *page, *pageSize)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *UserProfile) Get(id int) (*UserProfile, error) {
	var one UserProfile
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *UserProfile) Update(m UserProfile) error {
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
func (t *UserProfile) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *UserProfile) Insert(m UserProfile) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil && err != up.ErrWarnSlowQuery {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

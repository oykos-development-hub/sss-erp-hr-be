package data

import (
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/hr-ms-api/errors"
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
	Nationality               *string    `db:"nationality"`
	NationalMinority          *string    `db:"national_minority"`
	Citizenship               string     `db:"citizenship"`
	Address                   string     `db:"address"`
	BankAccount               *string    `db:"bank_account"`
	BankName                  *string    `db:"bank_name"`
	PersonalID                *string    `db:"personal_id"`
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
	IsJudge                   *bool      `db:"is_judge"`
	EngagementTypeID          *int       `db:"engagement_type_id,omitempty"`
	CreatedAt                 time.Time  `db:"created_at,omitempty"`
	UpdatedAt                 time.Time  `db:"updated_at"`
}

type Revisor struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

	err = res.OrderBy("created_at desc").All(&all)
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

func (t *UserProfile) GetBy(key string, value interface{}) ([]*UserProfile, error) {
	var all []*UserProfile
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{key: value})
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (t *UserProfile) GetRevisors() ([]*Revisor, error) {
	var all []*Revisor

	query := `select e.user_profile_id, u.first_name, u.last_name from employees_in_organization_units e, 
	systematizations s, job_positions_in_organization_units j, job_positions jp, user_profiles u
	where jp.Title = 'Revizor' and s.active = 2 and j.systematization_id = s.id and j.job_position_id = jp.id 
	and e.position_in_organization_unit_id = j.id and e.user_profile_id = u.id;`

	rows, err := upper.SQL().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Revisor
		err = rows.Scan(&item.ID, &item.FirstName, &item.LastName)
		if err != nil {
			return nil, err
		}
		all = append(all, &item)
	}

	return all, err
}

// Update updates a record in the database, using upper
func (t *UserProfile) Update(m UserProfile) error {
	m.UpdatedAt = time.Now()

	userByOfficialPersonalID, _ := t.GetBy("official_personal_id", m.OfficialPersonalID)
	if len(userByOfficialPersonalID) > 0 {
		if userByOfficialPersonalID[0].ID != m.ID {
			return errors.ErrUserJMBGExists
		}
	}

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

	userByOfficialPersonalID, _ := t.GetBy("official_personal_id", m.OfficialPersonalID)
	if len(userByOfficialPersonalID) > 0 {
		return 0, errors.ErrUserJMBGExists
	}

	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

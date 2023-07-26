package dto

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/guregu/null"
)

type JSONTime time.Time

const jsonTimeFormat = "2006-01-02"

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(jsonTimeFormat))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	nt, err := time.Parse(jsonTimeFormat, s)
	if err != nil {
		return err
	}
	*t = JSONTime(nt)
	return nil
}

type JSONTimeNullable struct {
	null.Time
}

func (t JSONTimeNullable) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Time.Time.Format(jsonTimeFormat))
	return []byte(stamp), nil
}

func (t *JSONTimeNullable) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		t.Valid = false
		return nil
	}

	nt, err := time.Parse(jsonTimeFormat, s)
	if err != nil {
		return err
	}

	t.Time.Time = nt
	t.Valid = true
	return nil
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}

	err := json.Unmarshal(data, &nt.Time)
	if err != nil {
		return err
	}

	nt.Valid = true
	return nil
}

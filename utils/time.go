package main

import (
  "time"
  "database/sql/driver"
)

type NullTime struct {
    Time  time.Time `json:"time"`
    Valid bool `json:"-"`
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
    nt.Time, nt.Valid = value.(time.Time)
    return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
    if !nt.Valid {
        return nil, nil
    }
    return nt.Time, nil
}
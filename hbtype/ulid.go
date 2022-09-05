package hbtype

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/oklog/ulid/v2"
)

type ULID struct {
	ULID ulid.ULID
}

// Creates a new ULID.
func NewULID() ULID {
	return ULID{ULID: ulid.Make()}
}

// Transforms a ULID type to database value.
func (u ULID) Value() (driver.Value, error) {
	return u.ULID.Value()
}

// Scans a database value to a ULID Go type.
func (u *ULID) Scan(value interface{}) error {
	return u.ULID.Scan(value)
}

// Marshals a ULID value into a JSON value.
func (u ULID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.ULID.String())
}

// Unmarshal a JSON value into a ULID value.
func (u *ULID) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &u.ULID)
	return err
}

// Returns a ULID value as a string.
func (u ULID) String() string {
	return u.ULID.String()
}

// Marshals a ULID into a string value.
func (u ULID) MarshalText() ([]byte, error) {
	return u.ULID.MarshalText()
}

// Unmarshals a string value into a ULID Go Value.
func (u *ULID) UnmarshalText(data []byte) error {
	return u.ULID.UnmarshalText(data)
}

// Marshals a ULID value into binary.
func (u ULID) MarshalBinary() ([]byte, error) {
	return u.ULID.MarshalBinary()
}

// Unmarshals a binary value into a ULID.
func (u *ULID) UnmarshalBinary(data []byte) error {
	return u.ULID.UnmarshalBinary(data)
}

package hbtype

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"

	"github.com/oklog/ulid/v2"
)

type NullULID struct {
	ULID  ulid.ULID
	Valid bool // Valid is true if ULID is not null.
}

var ZeroValueULID ulid.ULID = ulid.ULID{}
var jsonNull = []byte("null")

// Creates a new nullable ULID.
func NewNullULID() NullULID {
	return NullULID{ULID: ulid.Make(), Valid: true}
}

// Transforms a NullULID type to database value.
func (nu NullULID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}

	// Delegate to ULID Value function
	return nu.ULID.Value()
}

// Scans a database value to a NullULID Go type.
func (nu *NullULID) Scan(value interface{}) error {
	if value == nil {
		nu.ULID, nu.Valid = ZeroValueULID, false
		return nil
	}

	err := nu.ULID.Scan(value)
	if err != nil {
		nu.ULID = ZeroValueULID
		nu.Valid = false
		return err
	}

	nu.Valid = true
	return nil
}

// Marshals a NullULID value into a JSON value.
func (nu NullULID) MarshalJSON() ([]byte, error) {
	if nu.Valid {
		return json.Marshal(nu.ULID.String())
	}
	return json.Marshal(nil)
}

// Unmarshal a JSON value into a NullULID value.
func (nu *NullULID) UnmarshalJSON(data []byte) error {
	*nu = NullULID{}

	// Check if null ULID.
	if bytes.Equal(data, jsonNull) {
		return nil
	}

	err := json.Unmarshal(data, &nu.ULID)
	nu.Valid = err == nil
	return err
}

// Returns a NullULID value as a string.
func (nu NullULID) String() string {
	if nu.Valid {
		return nu.ULID.String()
	}

	return ""
}

// Marshals a NullULID into a string value.
func (nu NullULID) MarshalText() ([]byte, error) {
	// If NullULID is not null, return string representation, else return "null".
	if nu.Valid {
		return nu.ULID.MarshalText()
	}

	return jsonNull, nil
}

// Unmarshals a string value into a NullULID Go Value.
func (nu *NullULID) UnmarshalText(data []byte) error {
	err := nu.ULID.UnmarshalText(data)
	if err != nil {
		nu.ULID = ZeroValueULID
		nu.Valid = false
		return err
	}

	nu.Valid = true
	return nil
}

// Marshals a NullULID value into binary.
func (nu NullULID) MarshalBinary() ([]byte, error) {
	if nu.Valid {
		return nu.ULID.MarshalBinary()
	}

	// Return nil bytes for null ULID.
	return []byte(nil), nil
}

// Unmarshals a binary value into a NullULID.
func (nu *NullULID) UnmarshalBinary(data []byte) error {
	// Otherwise copy binary data into ULID field.
	err := nu.ULID.UnmarshalBinary(data)
	if err != nil {
		nu.ULID = ZeroValueULID
		nu.Valid = false
		return err
	}

	nu.Valid = true
	return nil
}

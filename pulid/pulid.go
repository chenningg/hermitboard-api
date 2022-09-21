package pulid

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/oklog/ulid/v2"
)

// Prefixed ULID with object type - globally universal ID.
type PULID string

// Separator between prefix and the ULID.
const separator = '_'

// ErrIncorrectPULIDFormat is returned when the PULID is not in the correct format.
var ErrIncorrectPULIDFormat = fmt.Errorf("pulid: incorrect format")

// Creates a new PULID.
func NewPULID(prefix string) PULID {
	return PULID(fmt.Sprintf("%s%c%s", strings.ToUpper(prefix), separator, ulid.Make().String()))
}

// ParsePrefix returns the prefix from a Prefixed-ULID.
func ParsePrefix(pulid PULID) (string, error) {
	separatorIdx := strings.IndexRune(pulid.String(), separator)

	// If separator is not found, return error
	if separatorIdx == -1 {
		return "", fmt.Errorf("pulid.ParsePrefix(): missing separator '%c': %w", separator, ErrIncorrectPULIDFormat)
	}

	return string(pulid[:separatorIdx]), nil
}

// ParseULID returns the ULID from a Prefixed-ULID.
func ParseULID(pulid PULID) (string, error) {
	separatorIdx := strings.IndexRune(pulid.String(), separator)

	// If separator is not found, return error
	if separatorIdx == -1 {
		return "", fmt.Errorf("pulid.ParseULID(): missing separator '%c': %w", separator, ErrIncorrectPULIDFormat)
	}

	return string(pulid[separatorIdx:]), nil
}

// IsAPULID checks if this is a valid PULID.
func IsAPULID(pulid PULID) error {
	prefix, err := ParsePrefix(pulid)
	if err != nil {
		return fmt.Errorf("pulid.IsAPULID(): %w, %v", ErrIncorrectPULIDFormat, err)
	}

	id, err := ParseULID(pulid)
	if err != nil {
		return fmt.Errorf("pulid.IsAPULID(): %w, %v", ErrIncorrectPULIDFormat, err)
	}

	if len(prefix) != 3 {
		return fmt.Errorf("pulid.IsAPULID(): prefix must be 3 alphabetic characters: %w", ErrIncorrectPULIDFormat)
	}

	for _, chr := range prefix {
		if !unicode.IsLetter(chr) {
			return fmt.Errorf("pulid.IsAPULID(): prefix must be 3 alphabetic characters: %w", ErrIncorrectPULIDFormat)
		}
	}

	_, err = ulid.ParseStrict(id)
	if err != nil {
		return fmt.Errorf("pulid.IsAPULID(): ulid is not valid: %v", err)
	}

	return nil
}

// Transforms a PULID type to database value string.
func (pulid PULID) Value() (driver.Value, error) {
	return pulid.String(), nil
}

// Scans a database value to a PULID Go type.
func (pulid *PULID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("pulid.Scan(): expected a string, got %t", src)
	}

	switch s := src.(type) {
	case string:
		*pulid = PULID(s)
	case []byte:
		*pulid = PULID(s)
	default:
		return fmt.Errorf("pulid.Scan(): expected a string, got %t", src)
	}

	return nil
}

// Returns a PULID value as a string.
func (pulid PULID) String() string {
	return string(pulid)
}

// Marshals a PULID into a graphql scalar string.
func (pulid PULID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(pulid.String()))
}

// Unmarshals a graphql scalar into a PULID Go type.
func (pulid *PULID) UnmarshalGQL(v interface{}) error {
	return pulid.Scan(v)
}

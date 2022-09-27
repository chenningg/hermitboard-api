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
const PrefixLength = 3
const Separator = '_'

// ErrIncorrectPULIDFormat is returned when the PULID is not in the correct format.
var ErrIncorrectPULIDFormat = fmt.Errorf("pulid: incorrect format")

// Creates a new PULID.
func NewPULID(prefix string) PULID {
	return PULID(fmt.Sprintf("%s%c%s", strings.ToUpper(prefix), Separator, ulid.Make().String()))
}

// ParsePULID returns the prefix and ULID from a PULID.
func ParsePULID(pulid PULID) (string, string, error) {
	separatorIdx := strings.IndexRune(pulid.String(), Separator)

	// If separator is not found, return error
	if separatorIdx == -1 {
		return "", "", fmt.Errorf("pulid.ParsePULID(): missing separator '%c': %w", Separator, ErrIncorrectPULIDFormat)
	}

	return string(pulid[:separatorIdx]), string(pulid[separatorIdx+1:]), nil
}

// IsAPULID checks if this is a valid PULID.
func IsAPULID(pulid PULID) error {
	prefix, id, err := ParsePULID(pulid)
	if err != nil {
		return fmt.Errorf("pulid.IsAPULID(): %w, %v", ErrIncorrectPULIDFormat, err)
	}

	if len(prefix) != PrefixLength {
		return fmt.Errorf(
			"pulid.IsAPULID(): prefix must be %d alphabetic characters: %w", PrefixLength, ErrIncorrectPULIDFormat,
		)
	}

	for _, chr := range prefix {
		if !unicode.IsLetter(chr) {
			return fmt.Errorf(
				"pulid.IsAPULID(): prefix must be %d alphabetic characters: %w", PrefixLength, ErrIncorrectPULIDFormat,
			)
		}
	}

	_, err = ulid.ParseStrict(id)
	if err != nil {
		return fmt.Errorf("pulid.IsAPULID(): ulid portion of the PULID is not valid: %v", err)
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

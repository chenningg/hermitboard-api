// Code generated by "enumer -type=AssetClass -json -text -sql"; DO NOT EDIT.

package hbtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _AssetClassName = "CASH_OR_CASH_EQUIVALENTCOMMODITYCRYPTOCURRENCYEQUITYFIXED_INCOMEFUTUREREAL_ESTATE"

var _AssetClassIndex = [...]uint8{0, 23, 32, 46, 52, 64, 70, 81}

const _AssetClassLowerName = "cash_or_cash_equivalentcommoditycryptocurrencyequityfixed_incomefuturereal_estate"

func (i AssetClass) String() string {
	i -= 1
	if i < 0 || i >= AssetClass(len(_AssetClassIndex)-1) {
		return fmt.Sprintf("AssetClass(%d)", i+1)
	}
	return _AssetClassName[_AssetClassIndex[i]:_AssetClassIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AssetClassNoOp() {
	var x [1]struct{}
	_ = x[CASH_OR_CASH_EQUIVALENT-(1)]
	_ = x[COMMODITY-(2)]
	_ = x[CRYPTOCURRENCY-(3)]
	_ = x[EQUITY-(4)]
	_ = x[FIXED_INCOME-(5)]
	_ = x[FUTURE-(6)]
	_ = x[REAL_ESTATE-(7)]
}

var _AssetClassValues = []AssetClass{CASH_OR_CASH_EQUIVALENT, COMMODITY, CRYPTOCURRENCY, EQUITY, FIXED_INCOME, FUTURE, REAL_ESTATE}

var _AssetClassNameToValueMap = map[string]AssetClass{
	_AssetClassName[0:23]:       CASH_OR_CASH_EQUIVALENT,
	_AssetClassLowerName[0:23]:  CASH_OR_CASH_EQUIVALENT,
	_AssetClassName[23:32]:      COMMODITY,
	_AssetClassLowerName[23:32]: COMMODITY,
	_AssetClassName[32:46]:      CRYPTOCURRENCY,
	_AssetClassLowerName[32:46]: CRYPTOCURRENCY,
	_AssetClassName[46:52]:      EQUITY,
	_AssetClassLowerName[46:52]: EQUITY,
	_AssetClassName[52:64]:      FIXED_INCOME,
	_AssetClassLowerName[52:64]: FIXED_INCOME,
	_AssetClassName[64:70]:      FUTURE,
	_AssetClassLowerName[64:70]: FUTURE,
	_AssetClassName[70:81]:      REAL_ESTATE,
	_AssetClassLowerName[70:81]: REAL_ESTATE,
}

var _AssetClassNames = []string{
	_AssetClassName[0:23],
	_AssetClassName[23:32],
	_AssetClassName[32:46],
	_AssetClassName[46:52],
	_AssetClassName[52:64],
	_AssetClassName[64:70],
	_AssetClassName[70:81],
}

// AssetClassString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AssetClassString(s string) (AssetClass, error) {
	if val, ok := _AssetClassNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AssetClassNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AssetClass values", s)
}

// AssetClassValues returns all values of the enum
func AssetClassValues() []AssetClass {
	return _AssetClassValues
}

// AssetClassStrings returns a slice of all String values of the enum
func AssetClassStrings() []string {
	strs := make([]string, len(_AssetClassNames))
	copy(strs, _AssetClassNames)
	return strs
}

// IsAAssetClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AssetClass) IsAAssetClass() bool {
	for _, v := range _AssetClassValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AssetClass
func (i AssetClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetClass
func (i *AssetClass) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AssetClass should be a string, got %s", data)
	}

	var err error
	*i, err = AssetClassString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for AssetClass
func (i AssetClass) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AssetClass
func (i *AssetClass) UnmarshalText(text []byte) error {
	var err error
	*i, err = AssetClassString(string(text))
	return err
}

func (i AssetClass) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *AssetClass) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AssetClass: %[1]T(%[1]v)", value)
	}

	val, err := AssetClassString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

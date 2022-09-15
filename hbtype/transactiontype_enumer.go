// Code generated by "enumer -type=TransactionType -json -text -sql"; DO NOT EDIT.

package hbtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _TransactionTypeName = "BUYSELLSTAKEDIVIDEND_INCOMERENT_PAYMENTRENT_INCOMESTOCK_DIVIDEND"

var _TransactionTypeIndex = [...]uint8{0, 3, 7, 12, 27, 39, 50, 64}

const _TransactionTypeLowerName = "buysellstakedividend_incomerent_paymentrent_incomestock_dividend"

func (i TransactionType) String() string {
	i -= 1
	if i < 0 || i >= TransactionType(len(_TransactionTypeIndex)-1) {
		return fmt.Sprintf("TransactionType(%d)", i+1)
	}
	return _TransactionTypeName[_TransactionTypeIndex[i]:_TransactionTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TransactionTypeNoOp() {
	var x [1]struct{}
	_ = x[BUY-(1)]
	_ = x[SELL-(2)]
	_ = x[STAKE-(3)]
	_ = x[DIVIDEND_INCOME-(4)]
	_ = x[RENT_PAYMENT-(5)]
	_ = x[RENT_INCOME-(6)]
	_ = x[STOCK_DIVIDEND-(7)]
}

var _TransactionTypeValues = []TransactionType{BUY, SELL, STAKE, DIVIDEND_INCOME, RENT_PAYMENT, RENT_INCOME, STOCK_DIVIDEND}

var _TransactionTypeNameToValueMap = map[string]TransactionType{
	_TransactionTypeName[0:3]:        BUY,
	_TransactionTypeLowerName[0:3]:   BUY,
	_TransactionTypeName[3:7]:        SELL,
	_TransactionTypeLowerName[3:7]:   SELL,
	_TransactionTypeName[7:12]:       STAKE,
	_TransactionTypeLowerName[7:12]:  STAKE,
	_TransactionTypeName[12:27]:      DIVIDEND_INCOME,
	_TransactionTypeLowerName[12:27]: DIVIDEND_INCOME,
	_TransactionTypeName[27:39]:      RENT_PAYMENT,
	_TransactionTypeLowerName[27:39]: RENT_PAYMENT,
	_TransactionTypeName[39:50]:      RENT_INCOME,
	_TransactionTypeLowerName[39:50]: RENT_INCOME,
	_TransactionTypeName[50:64]:      STOCK_DIVIDEND,
	_TransactionTypeLowerName[50:64]: STOCK_DIVIDEND,
}

var _TransactionTypeNames = []string{
	_TransactionTypeName[0:3],
	_TransactionTypeName[3:7],
	_TransactionTypeName[7:12],
	_TransactionTypeName[12:27],
	_TransactionTypeName[27:39],
	_TransactionTypeName[39:50],
	_TransactionTypeName[50:64],
}

// TransactionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TransactionTypeString(s string) (TransactionType, error) {
	if val, ok := _TransactionTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TransactionTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TransactionType values", s)
}

// TransactionTypeValues returns all values of the enum
func TransactionTypeValues() []TransactionType {
	return _TransactionTypeValues
}

// TransactionTypeStrings returns a slice of all String values of the enum
func TransactionTypeStrings() []string {
	strs := make([]string, len(_TransactionTypeNames))
	copy(strs, _TransactionTypeNames)
	return strs
}

// IsATransactionType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TransactionType) IsATransactionType() bool {
	for _, v := range _TransactionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TransactionType
func (i TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionType
func (i *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TransactionType should be a string, got %s", data)
	}

	var err error
	*i, err = TransactionTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for TransactionType
func (i TransactionType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for TransactionType
func (i *TransactionType) UnmarshalText(text []byte) error {
	var err error
	*i, err = TransactionTypeString(string(text))
	return err
}

func (i TransactionType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *TransactionType) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of TransactionType: %[1]T(%[1]v)", value)
	}

	val, err := TransactionTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

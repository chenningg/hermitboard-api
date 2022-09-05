// Code generated by "enumer -type=TransactionType -json -text -sql"; DO NOT EDIT.

package hbtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _TransactionTypeName = "BuySellStakeDividendIncomeRentPaymentRentIncomeStockDividend"

var _TransactionTypeIndex = [...]uint8{0, 3, 7, 12, 26, 37, 47, 60}

const _TransactionTypeLowerName = "buysellstakedividendincomerentpaymentrentincomestockdividend"

func (i TransactionType) String() string {
	if i < 0 || i >= TransactionType(len(_TransactionTypeIndex)-1) {
		return fmt.Sprintf("TransactionType(%d)", i)
	}
	return _TransactionTypeName[_TransactionTypeIndex[i]:_TransactionTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TransactionTypeNoOp() {
	var x [1]struct{}
	_ = x[Buy-(0)]
	_ = x[Sell-(1)]
	_ = x[Stake-(2)]
	_ = x[DividendIncome-(3)]
	_ = x[RentPayment-(4)]
	_ = x[RentIncome-(5)]
	_ = x[StockDividend-(6)]
}

var _TransactionTypeValues = []TransactionType{Buy, Sell, Stake, DividendIncome, RentPayment, RentIncome, StockDividend}

var _TransactionTypeNameToValueMap = map[string]TransactionType{
	_TransactionTypeName[0:3]:        Buy,
	_TransactionTypeLowerName[0:3]:   Buy,
	_TransactionTypeName[3:7]:        Sell,
	_TransactionTypeLowerName[3:7]:   Sell,
	_TransactionTypeName[7:12]:       Stake,
	_TransactionTypeLowerName[7:12]:  Stake,
	_TransactionTypeName[12:26]:      DividendIncome,
	_TransactionTypeLowerName[12:26]: DividendIncome,
	_TransactionTypeName[26:37]:      RentPayment,
	_TransactionTypeLowerName[26:37]: RentPayment,
	_TransactionTypeName[37:47]:      RentIncome,
	_TransactionTypeLowerName[37:47]: RentIncome,
	_TransactionTypeName[47:60]:      StockDividend,
	_TransactionTypeLowerName[47:60]: StockDividend,
}

var _TransactionTypeNames = []string{
	_TransactionTypeName[0:3],
	_TransactionTypeName[3:7],
	_TransactionTypeName[7:12],
	_TransactionTypeName[12:26],
	_TransactionTypeName[26:37],
	_TransactionTypeName[37:47],
	_TransactionTypeName[47:60],
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

package authrole

import (
    "encoding/json"
    "fmt"
)

func (e Value) MarshalBinary() (data []byte, err error) {
    return json.Marshal(e.String())
}

func (e *Value) UnmarshalBinary(data []byte) error {
    err := json.Unmarshal(data, e)
    if err != nil {
        return fmt.Errorf("failed to unmarshal")
    }

    if err := ValueValidator(*e); err != nil {
        return fmt.Errorf("%s is not a valid Value", string(data))
    }
    return nil
}

package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json_value := fmt.Sprintf("%d mins", r)
	quoted_json_value := strconv.Quote(json_value)
	return []byte(quoted_json_value), nil
}

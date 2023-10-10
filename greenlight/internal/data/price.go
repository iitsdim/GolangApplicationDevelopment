package data

import (
	"fmt"
	"strconv"
)

type Price int32

func (p Price) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d $", p)
	quotedValue := strconv.Quote(jsonValue)

	return []byte(quotedValue), nil
}

package telegram

import (
	"fmt"
	"strings"

	"github.com/TBXark/jsoncompressor"
)

// query prefix must be unique and has suffix ":" to separate the data
// update.CallbackQuery.Data format: $route:json($data)

func UnmarshalData[T any](data string) (string, *T, error) {
	cmp := strings.SplitN(data, ":", 2)
	if len(cmp) != 2 {
		return "", nil, fmt.Errorf("invalid data format")
	}
	var v T
	err := jsoncompressor.Unmarshal([]byte(cmp[1]), &v)
	if err != nil {
		return cmp[0], nil, err
	}
	return cmp[0], &v, nil
}

func MarshalData[T any](route string, data T) string {
	b, _ := jsoncompressor.Marshal(data)
	return fmt.Sprintf("%s:%s", route, string(b))
}

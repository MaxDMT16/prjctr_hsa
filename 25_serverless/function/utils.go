package function

import (
	"encoding/json"
	"fmt"
)

func MapToStruct[T any](m map[string]interface{}) (T, error) {
	var s T

	mapJSON, err := json.Marshal(m)
	if err != nil {
		return s, fmt.Errorf("marshal input map: %w", err)
	}

	if err := json.Unmarshal(mapJSON, &s); err != nil {
		return s, fmt.Errorf("unmarshal to %T struct: %w", s, err)
	}

	return s, nil
}

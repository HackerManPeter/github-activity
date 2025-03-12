package internal

import "encoding/json"

func MapToStruct[T any](m map[string]any, s *T) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, s)
	if err != nil {
		return err
	}

	return nil
}

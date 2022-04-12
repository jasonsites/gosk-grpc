package utils

import (
	"encoding/json"
	"fmt"
)

// ToJSON marshals input to JSON
func ToJSON(i interface{}) string {
	s, err := json.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("error marshalling json %v", err))
	}
	return string(s)
}

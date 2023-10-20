package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseJSON(jsonStr string) error {
	var data interface{}

	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&data); err != nil {
		offset := int(decoder.InputOffset())
		if offset <= 0 {
			return fmt.Errorf("JSON is invalid: %s", err)
		}
		character := string(jsonStr[offset-1])
		line, col := findErrorPosition(jsonStr, offset)
		return fmt.Errorf("error: unexpected character '%s' at line %d, column %d", character, line, col)
	}

	if decoder.More() {
		return fmt.Errorf("JSON is invalid: multiple top-level entities or extra data")
	}

	if strings.Count(jsonStr, "{") < strings.Count(jsonStr, "}") {
		return fmt.Errorf("JSON is invalid: unmatched closing brace '}'")
	}

	return nil
}

func findErrorPosition(jsonStr string, offset int) (int, int) {
	lines := strings.Split(jsonStr[:offset], "\n")

	line := len(lines)
	col := len(lines[line-1]) + 1

	return line, col
}

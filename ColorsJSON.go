package chalk

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CachedFormatters for json
var (
	// Array
	jsonArrayFormat = Blue().Bold()
	jsonArrayStart  = jsonArrayFormat.Sprint("[")
	jsonArrayEnd    = jsonArrayFormat.Sprint("]")
	jsonArrayComma  = jsonArrayFormat.Sprint(",")

	// Object
	jsonObjectFormat = Red().Bold()
	jsonObjectStart  = jsonObjectFormat.Sprint("{")
	jsonObjectEnd    = jsonObjectFormat.Sprint("}")
	jsonObjectColon  = jsonObjectFormat.Sprint(":")
	jsonObjectComma  = jsonObjectFormat.Sprint(",")

	// Key
	jsonKeyFormat     = Green().Bold().Cache()
	jsonKeyQuote      = jsonKeyFormat.Sprint(`"`)
	jsonKeyTextFormat = Green().Italic().Cache()

	// Value
	jsonValueFormat     = Yellow().Bold().Cache()
	jsonValueQuote      = jsonValueFormat.Sprint(`"`)
	jsonValueTextFormat = Yellow().Italic().Cache()
)

// JSON repoerts string with colorized syntax of given json string
// This is experimental, but works on basic cases
func JSON(v interface{}) string {
	if !coloringEnabled {
		b, err := json.Marshal(v)
		if err != nil {
			return err.Error()
		}
		return string(b)
	}

	// Marshal / Unmarshal to get interface values instead of real types
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err.Error()
	}
	return jsonFormat(v, "", "")
}

// JSONIndent repoerts string with colorized syntax of given json string
// This is experimental, but works on basic cases
func JSONIndent(v interface{}, prefix, indent string) string {
	if !coloringEnabled {
		b, err := json.MarshalIndent(v, prefix, indent)
		if err != nil {
			return err.Error()
		}
		return string(b)
	}

	// Marshal / Unmarshal to get interface values instead of real types
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err.Error()
	}
	return jsonFormat(v, prefix, indent)
}

// jsonFormat reports given value as color formatted
func jsonFormat(v interface{}, indent, indentStep string) string {
	// New line token
	newLine, colon := "\n", jsonObjectColon+" "
	if indent == "" && indentStep == "" {
		newLine = ""
		colon = jsonObjectColon
	}

	// Select type
	switch v.(type) {
	case []interface{}: // Array []
		list := v.([]interface{})
		values := make([]string, len(list))
		for i, value := range list {
			values[i] = jsonFormat(value, indent+indentStep, indentStep)
		}
		return indent + fmt.Sprintf("%s%s%s%s%s", jsonArrayStart, newLine, strings.Join(values, jsonArrayComma+newLine), newLine, indent+jsonArrayEnd)

	case map[string]interface{}: // Object {}
		obj := v.(map[string]interface{})
		fields, i := make([]string, len(obj)), 0
		for key, value := range obj {
			fields[i] = indent + indentStep + fmt.Sprintf(`%s%s%s`, jsonKeyQuote, jsonKeyTextFormat.Sprint(key), jsonKeyQuote)
			fields[i] += colon + jsonFormat(value, "", "")
			i++
		}
		return indent + fmt.Sprintf("%s%s%s%s%s", jsonObjectStart, newLine, strings.Join(fields, jsonObjectComma+newLine), newLine, indent+jsonObjectEnd)

	case string: // String "abc"
		return indent + fmt.Sprintf(`%s%s%s`, jsonValueQuote, jsonValueTextFormat.Sprint(v), jsonValueQuote)

	default: // Other values
		return indent + jsonValueTextFormat.Sprint(v)
	}
}

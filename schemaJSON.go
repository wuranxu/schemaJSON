// The package can translate your json-schema data to json data, if you have set default value, it will be translated too
// It depends on simpleson: https://github.com/simplejson/simplejson
package schemaJSON

import (
	"encoding/json"
	"errors"
	"fmt"
)

// returns the current implementation version
func Version() string {
	return "0.0.0"
}

const (
	NotValidJSONSchema = "valid JSON schema data"
	NotValidMap        = "is not a map"
)

type schema struct {
	data string
}

func NewSchema(data string) *schema {
	return &schema{data}
}

func (s *schema) Generate() (result interface{}, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%+v", e)
		}
	}()
	js, e := NewJSON(s.data)
	if e != nil {
		return nil, e
	}
	data, e := js.Map()
	if e != nil {
		return nil, e
	}
	return s.SchemaToJSON(data)
}

func (s *schema) GenerateJSON() (string, error) {
	data, err := s.Generate()
	if err != nil {
		return "", err
	}
	return s.toJSON(data)
}

func (s *schema) toJSON(data interface{}) (string, error) {
	bt, err := json.Marshal(data)
	return string(bt), err
}

// json-schema to json
func (s *schema) SchemaToJSON(data map[string]interface{}) (interface{}, error) {
	var result interface{} // json result
	if tp, ok := data["type"]; ok {
		switch tp {
		case "object":
			// when son is object
			result = make(map[string]interface{})
			properties, ok := data["properties"]
			if !ok {
				return nil, fmt.Errorf("map %+v doesn't contain properties", data)
			}
			if now, ok := properties.(map[string]interface{}); ok {
				for k, v := range now {
					or := result.(map[string]interface{})
					if m, ok := v.(map[string]interface{}); ok {
						or[k], _ = s.SchemaToJSON(m)
					} else {
						return nil, fmt.Errorf("%+v %s", v, NotValidMap)
					}
				}
			} else {
				return nil, fmt.Errorf("%+v %s", properties, NotValidMap)
			}
		case "array":
			result = make([]interface{}, 0)
			items, ok := data["items"]
			if !ok {
				return nil, fmt.Errorf("array %+v doesn't contain items", data)
			}
			if items, ok := items.(map[string]interface{}); ok {
				d, _ := s.SchemaToJSON(items)
				result = append(result.([]interface{}), d)
			} else {
				return nil, fmt.Errorf("%+v %s", items, NotValidMap)
			}
		case "string":
			result = ""
		case "number", "integer":
			result = 0
		case "boolean":
			result = false
		default:
			result = nil
		}
		if val, ok := data["default"]; ok {
			result = val
		}
	} else {
		return nil, errors.New(NotValidJSONSchema)
	}

	return result, nil
}

package schemaJSON

import (
	jsob "encoding/json"
	"errors"
	sj "github.com/bitly/go-simplejson"
	"strconv"
)

const (
	NotValid = "data is not valid"
)

func NewJSON(data interface{}) (*JSONData, error) {
	j := new(sj.Json)
	switch data.(type) {
	case []byte:
		err := j.UnmarshalJSON(data.([]byte))
		if err != nil {
			return nil, err
		}
		return &JSONData{data.([]byte)}, nil
	case string:
		err := j.UnmarshalJSON([]byte(data.(string)))
		if err != nil {
			return nil, err
		}
		return &JSONData{[]byte(data.(string))}, nil
	default:
		return nil, errors.New(NotValid)
	}
}

type JSONData struct {
	data []byte
}

// Get json data by path, return a simpleJson obj if you like
func (j *JSONData) get(path ...interface{}) (*sj.Json, error) {
	json, err := sj.NewJson(j.data)
	if err != nil {
		return nil, err
	}
	for _, p := range path {
		switch p.(type) {
		case string:
			dt, err := strconv.Atoi(p.(string))
			if err == nil {
				json = json.GetIndex(dt)
				continue
			}
			json = json.Get(p.(string))
		case int:
			json = json.GetIndex(p.(int))
		case []string:
			data := p.([]string)
			for _, p2 := range data {
				dt, err := strconv.Atoi(p2)
				if err == nil {
					json = json.GetIndex(dt)
					continue
				}
				json = json.Get(p2)
			}
		case []int:
			data := p.([]int)
			for _, p2 := range data {
				json = json.GetIndex(p2)
			}
		case []interface{}:
			data := p.([]interface{})
			for _, p2 := range data {
				switch p2.(type) {
				case string:
					dt, err := strconv.Atoi(p2.(string))
					if err == nil {
						json = json.GetIndex(dt)
						continue
					}
					json = json.Get(p2.(string))
				case int:
					json = json.GetIndex(p2.(int))
				}
			}
		}
	}
	return json, err
}

// Get JSONData
func (j *JSONData) GetJSON(path ...interface{}) (*JSONData, error) {
	json, err := sj.NewJson(j.data)
	if err != nil {
		return nil, err
	}
	for _, p := range path {
		switch p.(type) {
		case string:
			json = json.Get(p.(string))
		case int:
			json = json.GetIndex(p.(int))
		}
	}
	dt, err := json.Encode()
	if err != nil {
		return j, err
	}
	return &JSONData{dt}, err
}

// get string by path, return string and error
func (j *JSONData) String(path ...interface{}) (string, error) {
	json, err := j.get(path...)
	return json.MustString(), err
}

// return map and error
func (j *JSONData) Map(path ...interface{}) (map[string]interface{}, error) {
	json, err := j.get(path...)
	return json.MustMap(), err
}

// return array and error
func (j *JSONData) Array(path ...interface{}) ([]interface{}, error) {
	json, err := j.get(path...)
	return json.MustArray(), err
}

// return bool and error
func (j *JSONData) Bool(path ...interface{}) (bool, error) {
	json, err := j.get(path...)
	return json.MustBool(), err
}

// return int and error
func (j *JSONData) Int(path ...interface{}) (int, error) {
	json, err := j.get(path...)
	return json.MustInt(), err
}

// return int64 and error
func (j *JSONData) Int64(path ...interface{}) (int64, error) {
	json, err := j.get(path...)
	return json.MustInt64(), err
}

// return float64 and error
func (j *JSONData) Float64(path ...interface{}) (float64, error) {
	json, err := j.get(path...)
	return json.MustFloat64(), err
}

// return Uint64 and error
func (j *JSONData) Uint64(path ...interface{}) (uint64, error) {
	json, err := j.get(path...)
	return json.MustUint64(), err
}

// return []string and error
func (j *JSONData) StrArray(path ...interface{}) ([]string, error) {
	json, err := j.get(path...)
	return json.MustStringArray(), err
}

// return interface{} and error
func (j *JSONData) Interface(path ...interface{}) (interface{}, error) {
	json, err := j.get(path...)
	return json.Interface(), err
}

// set key
func (j *JSONData) Set(key string, value interface{}) (string, error) {
	body, err := sj.NewJson(j.data)
	if err != nil {
		return "", err
	}
	body.Set(key, value)
	b, err := jsob.Marshal(body)
	return string(b), err
}

// get origin data
func (j *JSONData) Data() []byte {
	return j.data
}

// set key-value by path
func (j *JSONData) SetPath(v interface{}, path ...string) error {
	json, err := sj.NewJson(j.data)
	if err != nil {
		return err
	}
	json.SetPath(path, v)
	bt, err := json.MarshalJSON()
	if err != nil {
		return err
	}
	j.data = bt
	return nil
}

// translate []byte to string
func (j *JSONData) FormatString() string {
	return string(j.data)
}

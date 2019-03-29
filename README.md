# schemaJSON
translate your json-schema data to json data, if you have set default value, it will be translated too

将json-schema数据转为默认json数据。


# Get Started

### Install Package

  ```
  go get github.com/wuranxu/schemaJSON
 
  ```


### Example

```golang
package main

import (
	"fmt"
	"github.com/wuranxu"
)

func main() {
	data := `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "default": "wuranxu",
      "description": "f"
    },
    "age": {
      "type": "number",
      "default": 10,
      "description": "f"
    },
    "class": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "default": "math"
          },
          "score": {
            "type": "number",
            "default": 120
          }
        },
        "required": [
          "name",
          "score"
        ],
        "description": "f"
      },
      "description": "f"
    }
  },
  "description": "few"
}`
	s := schemaJSON.NewSchema(data)
	data, err := s.GenerateJSON() // return string
	fmt.Println(data, err)
	data2, err := s.Generate() // return interface{}
	fmt.Println(data2, err)
}

// {"age":10,"class":[{"name":"math","score":120}],"name":"wuranxu"} <nil>
// map[age:10 class:[map[name:math score:120]] name:wuranxu] <nil>

```

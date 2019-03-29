package schemaJSON

import (
	"fmt"
	"testing"
)

func TestSchema_GenerateJSON(t *testing.T) {
	data := `
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "classes": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "haha": {
            "type": "string",
            "default": "这事一条数据",
            "description": "fwef"
          },
          "xixi": {
            "type": "array",
            "items": {
              "type": "string",
              "description": "fwe"
            },
            "description": "fewf"
          }
        },
        "description": "fewf"
      },
      "description": "fwef"
    },
    "match": {
      "type": "object",
      "properties": {
        "day": {
          "type": "object",
          "properties": {
            "month": {
              "type": "number",
              "default": 5
            }
          }
        }
      },
      "description": "fwe"
    },
    "name": {
      "type": "string",
      "default": "wuranxu",
      "description": "fewfwe"
    }
  },
  "description": "fwef"
}
	`
	j := NewSchema(data)
	fmt.Println(j.GenerateJSON())
}

func TestSchema_SchemaToJSON(t *testing.T) {
	data := `{"type":"object","properties":{"token":{"description":"用户标识，必填","type":"string","default":""},"action":{"description":"方法名，必填","type":"string","default":"user.ride.check"}}}`
	j := NewSchema(data)
	fmt.Println(j.GenerateJSON())
}

func TestSchema_Generate(t *testing.T) {
	data := `{"type":"object","properties":{"token":{"description":"用户标识，必填","type":"string","default":""},"action":{"description":"方法名，必填","type":"string","default":"user.ride.check"}}}`
	j := NewSchema(data)
	fmt.Println(j.Generate())
}
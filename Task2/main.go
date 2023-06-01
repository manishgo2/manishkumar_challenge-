package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

// JSONKeyValue represents a key-value pair in the JSON object
type JSONKeyValue struct {
	Key   string      `json:"key"`   // Key of the JSON object
	Value interface{} `json:"value"` // Value associated with the key
}

func main() {
	input := `{
		"number_1": {
			"N": "1.50"
		},
		"string_1": {
			"S": "784498"
		},
		"string_2": {
			"S": "2014-07-16T20:55:46Z"
		},
		"map_1": {
			"M": {
				"bool_1": {
					"BOOL": "truthy"
				},
				"null_1": {
					"NULL": "true"
				},
				"list_1": {
					"L": [
						{
							"S": ""
						},
						{
							"N": "011"
						},
						{
							"N": "5215s"
						},
						{
							"BOOL": "f"
						},
						{
							"NULL": "0"
						}
					]
				}
			}
		},
		"list_2": {
			"L": ["noop"]
		},
		"list_3": {
			"L": ["noop"]
		},
		"": {
			"S": "noop"
		}
	}`

	output, err := ProcessJSON(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

// ProcessJSON processes the input JSON string and returns the modified JSON string
func ProcessJSON(input string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}

	var dataModified []JSONKeyValue
	for key, value := range data {
		if key == "" {
			continue
		}
		modifyVal, err := ModifyValue(value) // Modify the value based on its type
		if err != nil {
			continue // Skip fields with unsupported data types
		}
		dataModified = append(dataModified, JSONKeyValue{
			Key:   key,
			Value: modifyVal,
		}) // Append the modified key-value pair to the dataModified slice
	}

	sort.Slice(dataModified, func(i, j int) bool {
		return dataModified[i].Key < dataModified[j].Key // Sort the dataModified slice by key in ascending order
	})

	jsonOutput, err := json.MarshalIndent(dataModified, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonOutput), nil // Return the modified JSON string
}

// ModifyValue modifies the value based on its type
func ModifyValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		return ModifyMap(v) // If the value is a map, recursively call ModifyMap to modify its contents
	case string:
		return ModifyString(v) // If the value is a string, call ModifyString to modify it
	case []interface{}:
		return ModifyList(v) // If the value is a list, call ModifyList to modify its elements
	default:
		return nil, fmt.Errorf("unsupported data type") // If the value has an unsupported data type, return an error
	}
}

// ModifyMap modifies the map value recursively
func ModifyMap(data map[string]interface{}) (map[string]interface{}, error) {
	modMap := make(map[string]interface{})
	for key, value := range data {
		modifyVal, err := ModifyValue(value) // Modify the value based on its type
		if err != nil {
			continue // Skip fields with unsupported data types
		}
		modMap[key] = modifyVal // Update the map with the modified value
	}

	if len(modMap) == 0 {
		return nil, fmt.Errorf("empty map") // If the modified map is empty, return an error indicating an empty map
	}
	return modMap, nil // Return the modified map
}

// ModifyString modifies a string value
func ModifyString(value string) (interface{}, error) {
	value = strings.TrimSpace(value) // Remove leading and trailing white spaces from the string

	t, err := time.Parse(time.RFC3339, value) // Attempt to parse the string as a time in RFC3339 format
	if err == nil {
		return t.Unix(), nil // If the parsing is successful, return the Unix timestamp
	}

	if value == "" {
		return nil, fmt.Errorf("empty string") // If the string is empty, return an error indicating an empty string
	}
	return value, nil // If no modifications are needed, return the original string value
}

// ModifyList modifies a list value
func ModifyList(data []interface{}) ([]interface{}, error) {
	modifiedList := make([]interface{}, 0, len(data)) // Create a new slice to store the modified list values
	for _, value := range data {
		modifyVal, err := ModifyValue(value) // Recursively modify each element in the list
		if err != nil {
			// Skip fields with unsupported data types
			continue
		}
		modifiedList = append(modifiedList, modifyVal) // Append the modified value to the modified list slice
	}

	if len(modifiedList) == 0 {
		return nil, fmt.Errorf("empty list") // If the modified list is empty, return an error indicating an empty list
	}
	return modifiedList, nil // Return the modified list
}

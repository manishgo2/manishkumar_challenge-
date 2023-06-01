
The program will process the input JSON and print the modified JSON output.

## Code Explanation

The `main.go` file contains the following functions:

- `main()`: The main function of the program. It initializes the input JSON string, calls the `ProcessJSON` function, and prints the output.

- `ProcessJSON(input string) (string, error)`: This function processes the input JSON string and returns the modified JSON string. It performs the following steps:
- Unmarshals the input JSON string into a map structure.
- Iterates over the key-value pairs in the map and modifies the values based on their types.
- Supports modifications for string, map, and list values.
- Sorts the modified key-value pairs by key in ascending order.
- Marshals the modified map into a JSON string with proper indentation.

- `ModifyValue(value interface{}) (interface{}, error)`: This function modifies the value based on its type. It supports modification for map, string, and list values.

- `ModifyMap(data map[string]interface{}) (map[string]interface{}, error)`: This function recursively modifies the map value by iterating over its key-value pairs and calling `ModifyValue` on each value.

- `ModifyString(value string) (interface{}, error)`: This function modifies a string value by trimming leading and trailing white spaces. If the string can be parsed as a valid RFC3339 formatted time, it converts it to a Unix timestamp.

- `ModifyList(data []interface{}) ([]interface{}, error)`: This function modifies a list value by iterating over its elements and calling `ModifyValue` on each element.

## Example Input and Output

The program is designed to handle JSON objects like the following example:

**Input JSON:**

```json
{
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
}

**Output JSON:**

```json

[
    {
        "key": "list_2",
        "value": ["noop"]
    },
    {
        "key": "list_3",
        "value": ["noop"]
    },
    {
        "key": "map_1",
        "value": {
            "bool_1": {
                "BOOL": "truthy"
            },
            "list_1": {
                "L": [
                    {
                        "N": "11"
                    },
                    {
                        "N": "5215s"
                    },
                    {
                        "BOOL": false
                    }
                ]
            },
            "null_1": {
                "NULL": true
            }
        }
    },
    ...
]

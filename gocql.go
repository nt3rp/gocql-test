package main

import (
	"fmt"
)

type Arguments struct {
	Statement string
	Values    []interface{}
}

type ArgumentMap map[string]Arguments
type ArgumentMap2D map[string]ArgumentMap
type ArgumentMap3D map[string]ArgumentMap2D

var (
	undefinedInt    int
	undefinedString string
	undefinedList   []string
)

var arguments = ArgumentMap3D{
	"simple": ArgumentMap2D{
		"all": ArgumentMap{
			"initial":  Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{10, "text", 10}},
			"complete": Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{10, "text", 10}},
			"partial":  Arguments{"INSERT INTO simple (k, s)    VALUES (?, ?   )", []interface{}{10, "text"}},
		},
		"some": ArgumentMap{
			"initial":  Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{11, "text", 11}},
			"complete": Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{11, "text", nil}},
			"partial":  Arguments{"INSERT INTO simple (k, s)    VALUES (?, ?   )", []interface{}{11, nil}},
		},
		"undefined": ArgumentMap{
			"initial":  Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{12, "text", 12}},
			"complete": Arguments{"INSERT INTO simple (k, s, i) VALUES (?, ?, ?)", []interface{}{12, "text", undefinedInt}},
			"partial":  Arguments{"INSERT INTO simple (k, s)    VALUES (?, ?   )", []interface{}{12, undefinedString}},
		},
	},
	"complex": ArgumentMap2D{
		"all": ArgumentMap{
			"initial":  Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{13, "text", []string{"a", "b"}}},
			"complete": Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{13, "text", []string{"a", "b"}}},
			"partial":  Arguments{"INSERT INTO complex (k, l)    VALUES (?, ?   )", []interface{}{13, []string{"a", "b"}}},
		},
		"some": ArgumentMap{
			"initial":  Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{14, "text", []string{"a", "b"}}},
			"complete": Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{14, "text", nil}},
			"partial":  Arguments{"INSERT INTO complex (k, l)    VALUES (?, ?   )", []interface{}{14, nil}},
		},
		"undefined": ArgumentMap{
			"initial":  Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{15, "text", []string{"a", "b"}}},
			"complete": Arguments{"INSERT INTO complex (k, s, l) VALUES (?, ?, ?)", []interface{}{15, "text", undefinedList}},
			"partial":  Arguments{"INSERT INTO complex (k, l)    VALUES (?, ?   )", []interface{}{15, undefinedList}},
		},
	},
}

func gocql_query(arguments Arguments) {
	fmt.Printf("%+v\n", arguments)
	err := session.Query(arguments.Statement, arguments.Values...).Exec()

	if err != nil {
		fmt.Println(err)
	}
}

func gocql_case(tableType string, columnType string, updateType string) {
	var (
		initialArgs Arguments
		actualArgs  Arguments
	)

	initialArgs = arguments[tableType]["all"]["initial"]
	gocql_query(initialArgs)

	actualArgs = arguments[tableType][columnType][updateType]
	gocql_query(actualArgs)
}

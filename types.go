package main

import (
	"fmt"
	"log"
)

const (
	tpl = `type %s struct {
%s}
`
)

var (
	//key - pg type, value - go type
	nonNullableTypes map[string]string = map[string]string{
		"int2":        "int",
		"int4":        "int",
		"int8":        "int",
		"float4":      "float64",
		"float8":      "float64",
		"numeric":     "float64",
		"money":       "float64",
		"bpchar":      "string",
		"varchar":     "string",
		"text":        "string",
		"bytea":       "[]byte",
		"uuid":        "uuid.UUID",
		"timestamp":   "time.Time",
		"timestamptz": "time.Time",
		"time":        "time.Time",
		"timetz":      "time.Time",
		"date":        "time.Time",
		"interval":    "time.Time",
		"bool":        "bool",
		"bit":         "uint32",
		"varbit":      "uint32",
		"json":        "struct{...}",
		"jsonb":       "string",
		"xml":         "struct{...}",
	}

	nullableTypes map[string]string = map[string]string{
		"int2":        "sql.NullInt64",
		"int4":        "sql.NullInt64",
		"int8":        "sql.NullInt64",
		"float4":      "sql.NullFloat64",
		"float8":      "sql.NullFloat64",
		"numeric":     "sql.NullFloat64",
		"money":       "sql.NullFloat64",
		"bpchar":      "sql.NullString",
		"varchar":     "sql.NullString",
		"text":        "sql.NullString",
		"bytea":       "*[]byte",
		"uuid":        "*uuid.UUID",
		"timestamp":   "pq.NullTime",
		"timestamptz": "pq.NullTime",
		"time":        "pq.NullTime",
		"timetz":      "pq.NullTime",
		"date":        "pq.NullTime",
		"interval":    "pq.NullTime",
		"bool":        "sql.NullBool",
		"bit":         "*uint32",
		"varbit":      "*uint32",
		"json":        "*struct{...}",
		"jsonb":       "sql.NullString",
		"xml":         "*struct{...}",
	}
)

func getStruct(tab string, cols []*column) string {
	var body string
	for _, c := range cols {
		body += fmt.Sprintf("\t%s %s `db:\"%s\"` // sqltype: %s\n",
			snake2Camel(c.Name), convertType(c), c.Name, c.Type)
	}

	return fmt.Sprintf(tpl, getTableName(tab), body)
}

func getMap(n bool) map[string]string {
	if n {
		return nullableTypes
	}
	return nonNullableTypes
}

func convertType(c *column) string {
	t, ok := getMap(c.IsNull)[c.Type]
	if !ok {
		log.Fatalln("unknown type: ", t)
	}

	if c.IsArray {
		return "[]" + t
	}
	return t
}

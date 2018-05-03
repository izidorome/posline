package posline

import (
	"fmt"
	"reflect"
	"strings"
)

// Marshal parsers all structs and transform into one string with all lines
func Marshal(v interface{}) (string, error) {
	var lines strings.Builder

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Struct:
		l, err := marshalStruct(rv)

		if err != nil {
			fmt.Println("err", err)
			return "", err
		}

		lines.WriteString(l)
		break
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			l, err := marshalStruct(rv.Index(i))

			if err != nil {
				fmt.Println("err:", err)
			}

			lines.WriteString(l)

			if i != (rv.Len() - 1) {
				lines.WriteString("\n")
			}
		}
	}

	return lines.String(), nil
}

func marshalStruct(rv reflect.Value) (string, error) {
	var c tagCollection

	t := rv.Type()

	c, err := parseTags(t)

	if err != nil {
		return "", err
	}

	content, err := parseValue(rv, c)

	return content, nil
}

package posline

import (
	"errors"
	"reflect"
	"strings"
)

const tagname = "posline"

var (
	// ErrInvalidSize is raised when the size tag are not int
	ErrInvalidSize = errors.New("posline: tag size should be an integer")
)

type tagCollection struct {
	Name string
	Tags []tag
}

type tag struct {
	Name         string
	Size         int
	LeftPad      bool
	ZeroFill     bool
	NoFloatPoint bool
}

// Marshal parsers all structs and transform into one string with all lines
func Marshal(v interface{}) (string, error) {
	var lines strings.Builder

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Struct:
		l, err := marshalStruct(rv)

		if err != nil {
			return "", err
		}

		lines.WriteString(l)
		break
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			l, err := marshalStruct(rv.Index(i))

			if err != nil {
				return "", err
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

	content := parseValue(rv, c)

	return content, nil
}

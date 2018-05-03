package posline

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/noverde/posline/pad"
)

func parseTags(t reflect.Type) (tagCollection, error) {
	var tags []tag

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		ftag := field.Tag.Get(tagname)

		if ftag == "" {
			continue
		}

		opts := strings.Split(ftag, ",")

		size, err := strconv.Atoi(opts[0])

		if err != nil {
			return tagCollection{}, ErrInvalidSize
		}

		zerofill := false
		leftpad := false
		nofp := false

		modifiers := opts[1:]

		for _, m := range modifiers {
			if m == "zerofill" {
				zerofill = true
			}

			if m == "leftpad" {
				leftpad = true
			}

			if m == "nofp" {
				nofp = true
			}
		}

		t := tag{
			Name:         field.Name,
			Size:         size,
			LeftPad:      leftpad,
			ZeroFill:     zerofill,
			NoFloatPoint: nofp,
		}

		tags = append(tags, t)
	}

	line := tagCollection{
		Name: t.Name(),
		Tags: tags,
	}

	return line, nil
}

func parseValue(rv reflect.Value, line tagCollection) string {
	var content strings.Builder
	t := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := t.Field(i)
		value := rv.Field(i)

		tg := tags(line)[field.Name]

		if tg == (tag{}) {
			continue
		}

		fieldContent := convert(value, tg)

		var sep string
		if tg.ZeroFill {
			sep = "0"
		} else {
			sep = " "
		}

		var fline string
		if tg.LeftPad {
			fline = pad.Left(fieldContent, tg.Size, sep)
		} else {
			fline = pad.Right(fieldContent, tg.Size, sep)
		}

		content.WriteString(fline)
	}

	return content.String()
}

func convert(v reflect.Value, t tag) string {
	var content string

	switch v.Kind() {
	case reflect.String:
		content = v.Interface().(string)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		content = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		content = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		value := fmt.Sprintf("%.2f", v.Float())

		if t.NoFloatPoint {
			content = strings.Replace(value, ".", "", 1)
		} else {
			content = value
		}
	case reflect.Bool:
		value := v.Interface().(bool)

		if value {
			content = "1"
		} else {
			content = "0"
		}
		break
	}

	return content
}

func tags(l tagCollection) map[string]tag {
	tags := make(map[string]tag)

	for _, t := range l.Tags {
		tags[t.Name] = t
	}

	return tags
}

package output

import (
	"fmt"
	"reflect"
	"strings"
)

func maxLabelSize(entries []interface{}) int {
	var labels []string
	for _, entry := range entries {
		iface := reflect.ValueOf(entry)
		for i := 0; i < iface.NumField(); i++ {
			labels = append(labels, iface.Type().Field(i).Name)
		}
	}

	var max int
	for _, label := range labels {
		if max < len(label) {
			max = len(label)
		}
	}

	return max
}

func maxPropertySize(properties []string) int {
	var labels []string
	for i := 0; i < len(properties); i++ {
		labels = append(labels, properties[i])
	}

	var max int
	for _, label := range labels {
		if max < len(label) {
			max = len(label)
		}
	}

	return max
}

func padLabel(label string, maxSize int) string {
	l := label
	for i := 0; i < maxSize - len(label); i++ {
		l += " "
	}

	return l
}

func labelExist(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func buildListEntry(builder *strings.Builder, field *reflect.Value, fieldName string, maxLabelSize int) {
	builder.WriteString(padLabel(fieldName, maxLabelSize))
	builder.WriteString(" : ")

	switch field.Kind() {
	case reflect.Slice:
		for f := 0; f < field.Len(); f++ {
			if f != field.Len() - 1 {
				builder.WriteString(fmt.Sprintf("%v, ", field.Index(f)))
			} else {
				builder.WriteString(fmt.Sprintf("%v", field.Index(f)))
			}
		}
	default:
		builder.WriteString(fmt.Sprintf("%v", field))
	}
}

func FormatList(entries []interface{}, properties []string) string {
	builder := strings.Builder{}

	for index, entry := range entries {
		iface := reflect.Indirect(reflect.ValueOf(entry))

		if properties != nil {
			maxLabelSize := maxPropertySize(properties)
			for i := 0; i < len(properties); i++ {
				field := iface.Field(i)
				fieldName := iface.Type().Field(i).Name

				buildListEntry(&builder, &field, fieldName, maxLabelSize)

				if i < len(properties) - 1 {
					builder.WriteString("\n")
				}
			}

			if index < len(entries) - 1 {
				builder.WriteString("\n\n")
			}
		} else {
			maxLabelSize := maxLabelSize(entries)
			for i := 0; i < iface.NumField(); i++ {
				field := iface.Field(i)
				fieldName := iface.Type().Field(i).Name

				buildListEntry(&builder, &field, fieldName, maxLabelSize)

				if i < iface.NumField() - 1 {
					builder.WriteString("\n")
				}
			}

			if index < len(entries) - 1 {
				builder.WriteString("\n\n")
			}
		}
	}

	return builder.String()
}

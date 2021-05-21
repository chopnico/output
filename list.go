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

func validProperties(p []string, v *reflect.Value) []string {
	var validProperties []string

	for i := 0; i < len(p); i++ {
		fieldName, valid := v.Type().FieldByName(p[i])

		if valid {
			validProperties = append(validProperties, fieldName.Name)
		}
	}
	return validProperties
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
				builder.WriteString(fmt.Sprintf("%s, ", field.Index(f).String()))
			} else {
				builder.WriteString(fmt.Sprintf("%s", field.Index(f).String()))
			}
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		builder.WriteString(fmt.Sprintf("%d", field.Int()))
	case reflect.Bool:
		builder.WriteString(fmt.Sprintf("%t", field.Bool()))
	case reflect.Ptr:
		builder.WriteString(fmt.Sprintf("%s", field.Elem()))
	default:
		if field.Interface() == nil || field == nil {
			builder.WriteString("")
		} else {
			builder.WriteString(fmt.Sprintf("%v (%s)", field.Interface(), field.Kind()))
		}
	}
}

func FormatList(entries []interface{}, properties []string) string {
	builder := strings.Builder{}

	for index, entry := range entries {
		value := reflect.Indirect(reflect.ValueOf(entry))

		if properties != nil {
			validProperties := validProperties(properties, &value)
			maxPropertySize := maxPropertySize(validProperties)

			for i := 0; i < len(validProperties); i++ {
				field := value.FieldByName(validProperties[i])

				buildListEntry(&builder, &field, validProperties[i], maxPropertySize)

				if i < len(validProperties) - 1 {
					builder.WriteString("\n")
				}
			}

			if index < len(entries) - 1 {
				builder.WriteString("\n\n")
			}

		} else {
			maxLabelSize := maxLabelSize(entries)

			for i := 0; i < value.NumField(); i++ {
				field := value.Field(i)
				fieldName := value.Type().Field(i).Name

				buildListEntry(&builder, &field, fieldName, maxLabelSize)

				if i < value.NumField() - 1 {
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

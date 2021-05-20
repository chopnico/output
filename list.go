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

func padLabel(label string, maxSize int, alignment string) string {
	l := label
	for i := 0; i < maxSize - len(label); i++ {
		l += " "
	}

	return l
}

func FormatList(entries []interface{}, seperator, alignment string) string {
	maxLabelSize := maxLabelSize(entries)
	builder := strings.Builder{}

	for index, entry := range entries {
		iface := reflect.ValueOf(entry)

		for i := 0; i < iface.NumField(); i++ {
			builder.WriteString(padLabel(iface.Type().Field(i).Name, maxLabelSize, alignment))
			builder.WriteString(" " + seperator + " ")

			switch iface.Field(i).Kind() {
			case reflect.Slice:
				for f := 0; f < iface.Field(i).Len(); f++ {
					if f != iface.Field(i).Len() - 1{
						builder.WriteString(fmt.Sprintf("%v, ", iface.Field(i).Index(f)))
					} else {
						builder.WriteString(fmt.Sprintf("%v", iface.Field(i).Index(f)))
					}
				}
			default:
				builder.WriteString(fmt.Sprintf("%v", iface.Field(i)))
			}

			if i < iface.NumField() - 1 {
				builder.WriteString("\n")
			}
		}

		if index < len(entries) - 1 {
			builder.WriteString("\n\n")
		}
	}

	return builder.String()
}

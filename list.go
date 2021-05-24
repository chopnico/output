package output

import (
	"fmt"
	"reflect"
	"strings"
)

func maxLabelSize(v *reflect.Value) int {
	var m int
	for i := 0; i < v.NumField(); i++ {
		l := v.Type().Field(i).Name
		if len(l) > m {
			m = len(l)
		}
	}
	return m
}

func maxPropertyLabelSize(p []string) int {
	var m int
	for i := 0; i < len(p); i++ {
		l := p[i]
		if len(l) > m {
			m = len(l)
		}
	}
	return m
}

func validProperties(p []string, v *reflect.Value) []string {
	var vp []string

	for i := 0; i < len(p); i++ {
		f, valid := v.Type().FieldByName(p[i])
		if valid {
			vp = append(vp, f.Name)
		}
	}
	return vp
}

func fieldValueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s", v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.2f", v.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.Array:
		var l string
		for i := 0; i < v.Len(); i++ {
			l += fmt.Sprintf("%s ", fieldValueToString(v.Index(i)))
		}
		return l
	case reflect.Slice:
		var l string
		for i := 0; i < v.Len(); i++ {
			l += fmt.Sprintf("%s ", fieldValueToString(v.Index(i)))
		}
		return l
	case reflect.Ptr:
		return fmt.Sprintf("%s", fieldValueToString(v.Elem()))
	case reflect.Struct:
		m := maxLabelSize(&v)
		var l string
		for i := 0; i < v.NumField(); i++ {
			fieldName := padLabel(v.Type().Field(i).Name, m)
			if i == 0 {
				l += "\n  - " + fieldName + " : "
				l += fmt.Sprintf("%s", fieldValueToString(v.Field(i)))
			} else {
				l += "\n    " + fieldName + " : "
				l += fmt.Sprintf("%s", fieldValueToString(v.Field(i)))
			}
		}
		return l
	default:
		return fmt.Sprintf("%s", v.Kind())
	}
}

func padLabel(l string, m int) string {
	t := l
	for i := 0; i < m - len(l); i++ {
		t += " "
	}
	return t
}

func buildList(v *reflect.Value, b *strings.Builder, p []string) {
	if p != nil {
		pm := maxPropertyLabelSize(p)
		vp := validProperties(p, v)
		for i := 0; i < len(vp); i++ {
			fieldName := padLabel(vp[i], pm)

			b.WriteString(fieldName + " : ")
			b.WriteString(fieldValueToString(v.FieldByName(vp[i])))

			if i != v.Type().NumField() {
				b.WriteString("\n")
			}
		}
	} else {
		m := maxLabelSize(v)
		for i := 0; i < v.Type().NumField(); i++ {
			fieldName := padLabel(v.Type().Field(i).Name, m)

			b.WriteString(fieldName + " : ")
			b.WriteString(fieldValueToString(v.Field(i)))

			if i != v.Type().NumField() {
				b.WriteString("\n")
			}
		}
	}
}

func FormatList(t *[]interface{}, p []string) string {
	builder := strings.Builder{}

	for i, x := range *t {
		value := reflect.Indirect(reflect.ValueOf(x))
		fmt.Printf("%s\n", value.Kind())

		buildList(&value, &builder, p)

		if i != len(*t) - 1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

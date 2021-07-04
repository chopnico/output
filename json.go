package output

import (
	"encoding/json"
	"reflect"
	"strings"
)

// exported function to print an item as json
func FormatItemAsJson(t interface{}) string {
	b := strings.Builder{}

	j, _ := json.Marshal(t)
	b.WriteString(string(j) + "\n")

	return b.String()
}

// exported function to print an item as pretty json
func FormatItemAsPrettyJson(t interface{}) string {
	b := strings.Builder{}

	j, _ := json.MarshalIndent(t, "", "  ")
	b.WriteString(string(j) + "\n")

	return b.String()
}

// exported function to print items as pretty json
func FormatItemsAsPrettyJson(t interface{}) string {
	a := reflect.ValueOf(t)
	b := strings.Builder{}

	for i := 0; i < a.Len(); i++ {
		j, _ := json.MarshalIndent(a.Index(i).Interface(), "", "  ")

		b.WriteString(string(j) + "\n")

		if i != a.Len()-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// exported function to print items as json
func FormatItemsAsJson(t interface{}) string {
	a := reflect.ValueOf(t)
	b := strings.Builder{}

	for i := 0; i < a.Len(); i++ {
		j, _ := json.Marshal(a.Index(i).Interface())

		b.WriteString(string(j) + "\n")

		if i != a.Len()-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

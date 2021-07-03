package output

import (
	"encoding/json"
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
func FormatItemsAsPrettyJson(t []interface{}) string {
	b := strings.Builder{}

	for i := 0; i < len(t); i++ {
		j, _ := json.MarshalIndent(t[i], "", "  ")

		b.WriteString(string(j) + "\n")

		if i != len(t)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// exported function to print items as json
func FormatItemsAsJson(t []interface{}) string {
	b := strings.Builder{}

	for i := 0; i < len(t); i++ {
		j, _ := json.Marshal(t[i])

		b.WriteString(string(j) + "\n")

		if i != len(t)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

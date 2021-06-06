package output

import (
	"strings"
	"encoding/json"
)

func FormatPrettyJson(t *[]interface{}) string {
	b := strings.Builder{}

	for i := 0; i < len((*t)); i++ {
		j, _ := json.MarshalIndent((*t)[i], "", "  ")

		b.WriteString(string(j) + "\n")

		if i != len((*t)) - 1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func FormatJson(t *[]interface{}) string {
	b := strings.Builder{}

	for i := 0; i < len((*t)); i++ {
		j, _ := json.Marshal((*t)[i])

		b.WriteString(string(j) + "\n")

		if i != len((*t)) - 1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

package output

import (
	"strings"
	"encoding/json"

	"github.com/chopnico/structs"
)

func FormatJson(t *[]interface{}) string {
	b := strings.Builder{}

	for i := 0; i < len((*t)); i++ {
		s := structs.New((*t)[i])
		j, _ := json.Marshal(s.Map())

		b.WriteString(string(j) + "\n")

		if i != len((*t)) - 1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

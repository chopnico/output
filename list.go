package output

import (
	"fmt"
	"strings"

	"github.com/chopnico/structs"
)

func label(l string, m int) string {
	s := l
	for i := 0; i < m; i++ {
		if len(s) < m {
			s += " "
		}
	}
	return s
}

func maxLabelSize(l []string) int {
	var max int
	for i := range l {
		if max < len(l[i]) {
			max = len(l[i])
		}
	}
	return max
}

func list(t interface{}, p []string, b *strings.Builder) {
	if p == nil || len(p) == 0 {
		if structs.IsStruct(t) {
			m := maxLabelSize(structs.Names(t))
			s := structs.New(t)
			n := s.Names()

			for i := range n {
				l := label(n[i], m)
				f, ok := s.FieldOk(n[i])
				if ok {
					b.WriteString(fmt.Sprintf("%s : %v", l, f.Value()))
					if i != len(n) {
						b.WriteString("\n")
					}
				}
			}
		}
	} else {
		if structs.IsStruct(t) {
			m := maxLabelSize(p)
			s := structs.New(t)

			for i := range p {
				l := label(p[i], m)
				f, ok := s.FieldOk(p[i])
				if ok {
					b.WriteString(fmt.Sprintf("%s : %v", l, f.Value()))
					if i != len(p) {
						b.WriteString("\n")
					}
				}
			}
		}
	}
}

func FormatList(t *[]interface{}, p []string) string {
	b := strings.Builder{}

	for i := 0; i < len((*t)); i++ {
		list((*t)[i], p, &b)

		if i != len((*t)) - 1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

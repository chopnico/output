package output

import (
	"fmt"
	"strings"

	"github.com/chopnico/structs"
)

// create the label based on the max size
// essentially, a glorified padder...
func label(l string, m int) string {
	s := l
	for i := 0; i < m; i++ {
		if len(s) < m {
			s += " "
		}
	}
	return s
}

// get the max label size
// this will determine which field has the most characters
func maxLabelSize(t interface{}, p []string) int {
	var max int
	s := structs.New(t)
	for i := range p {
		if _, ok := s.FieldOk(p[i]); ok {
			if max < len(p[i]) {
				max = len(p[i])
			}
		}
	}
	return max
}

// the function to actually build the list
// this uses a forked version of faiths' struct package
// github.com/chopnico/structs
func list(t interface{}, p []string, b *strings.Builder) {
	if p == nil {
		if structs.IsStruct(t) {
			s := structs.New(t)
			n := s.Names()
			m := maxLabelSize(t, n)

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
			m := maxLabelSize(t, p)
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

// exported function to print struct items as a list
func FormatItemsAsList(t []interface{}, p []string) string {
	b := strings.Builder{}

	for i := 0; i < len(t); i++ {
		list(t[i], p, &b)

		if i != len(t)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// exported fuction to print a single struct item as a list
func FormatItemAsList(t interface{}, p []string) string {
	b := strings.Builder{}

	list(t, p, &b)

	return b.String()
}

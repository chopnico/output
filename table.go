package output

import(
	"strings"

	"github.com/olekukonko/tablewriter"
)

func FormatTable(data [][]string, header []string) string {
	b := &strings.Builder{}
	t := tablewriter.NewWriter(b)
	t.SetHeader(header)
	t.SetAutoWrapText(false)
	t.SetAutoFormatHeaders(true)
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetRowSeparator("")
	t.SetHeaderLine(false)
	t.SetBorder(false)
	t.SetTablePadding("\t")
	t.SetNoWhiteSpace(true)
	t.AppendBulk(data)
	t.Render()

	return b.String()
}

// Copyright 2021 chopnico All rights reserved.
// Use of this source code is governed by a GNU GPLv3
// license that can be found in the LICENSE file.

package output

import (
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

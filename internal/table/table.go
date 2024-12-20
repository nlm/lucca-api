package table

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func New(headers ...any) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	return table.New(headers...).WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(2)
}

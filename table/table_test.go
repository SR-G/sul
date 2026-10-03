package table

import "testing"

type tableTestRow struct{ Name, Qty string }

func (r tableTestRow) TableCells() []string { return []string{r.Name, r.Qty} }

func TestTable(t *testing.T) {
	tbl := NewTable[tableTestRow]("Name", "Qty").
		Justify(1, JustifyRight).
		AddRows(tableTestRow{"apple", "3"}, tableTestRow{"kiwi", "120"})
	expected := "Name  | Qty\n------+----\napple |   3\nkiwi  | 120\n"
	if got := tbl.String(); got != expected {
		t.Errorf("got:\n%q\nexpected:\n%q", got, expected)
	}
}

func TestTableBorders(t *testing.T) {
	tbl := NewTable[tableTestRow]("Name", "Qty").
		Borders(true).
		Justify(1, JustifyRight).
		AddRows(tableTestRow{"apple", "3"}, tableTestRow{"kiwi", "120"})
	expected := "| Name  | Qty |\n+-------+-----+\n| apple |   3 |\n| kiwi  | 120 |\n"
	if got := tbl.String(); got != expected {
		t.Errorf("got:\n%q\nexpected:\n%q", got, expected)
	}
}

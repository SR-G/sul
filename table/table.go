package table

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

// cellWidth ignores ANSI escape sequences so colored cells stay aligned.
func cellWidth(s string) int {
	return utf8.RuneCountInString(ansiEscape.ReplaceAllString(s, ""))
}

var cellSanitizer = strings.NewReplacer("\r\n", " ", "\n", " ", "\r", " ", "\t", " ")

type Justification int

const (
	JustifyLeft Justification = iota
	JustifyRight
	JustifyCenter
)

// TableRow is implemented by any struct that can be rendered as a table row.
type TableRow interface {
	// TableCells returns one string per column, in header order.
	TableCells() []string
}

// Table renders rows of type T as a text table. Columns are left-aligned unless overridden.
type Table[T TableRow] struct {
	headers        []string
	justifications map[int]Justification
	rows           []T
	borders        bool
}

func NewTable[T TableRow](headers ...string) *Table[T] {
	return &Table[T]{headers: headers, justifications: map[int]Justification{}}
}

// Justify overrides the alignment of the given column index (0-based).
func (t *Table[T]) Justify(column int, j Justification) *Table[T] {
	t.justifications[column] = j
	return t
}

// JustifyHeader overrides alignment by header name.
func (t *Table[T]) JustifyHeader(header string, j Justification) *Table[T] {
	for i, h := range t.headers {
		if h == header {
			t.justifications[i] = j
		}
	}
	return t
}

// Borders enables left and right borders on every line.
func (t *Table[T]) Borders(enabled bool) *Table[T] {
	t.borders = enabled
	return t
}

func (t *Table[T]) AddRows(rows ...T) *Table[T] {
	t.rows = append(t.rows, rows...)
	return t
}

func (t *Table[T]) String() string {
	cells := make([][]string, 0, len(t.rows))
	widths := make([]int, len(t.headers))
	for i, h := range t.headers {
		widths[i] = cellWidth(h)
	}
	for _, r := range t.rows {
		c := r.TableCells()
		row := make([]string, len(t.headers))
		for i := range row {
			if i < len(c) {
				row[i] = cellSanitizer.Replace(c[i])
			}
			if w := cellWidth(row[i]); w > widths[i] {
				widths[i] = w
			}
		}
		cells = append(cells, row)
	}

	var sb strings.Builder
	t.writeLine(&sb, t.headers, widths)
	seps := make([]string, len(widths))
	for i, w := range widths {
		seps[i] = strings.Repeat("-", w)
	}
	if t.borders {
		fmt.Fprintf(&sb, "+-%s-+\n", strings.Join(seps, "-+-"))
	} else {
		fmt.Fprintf(&sb, "%s\n", strings.Join(seps, "-+-"))
	}
	for _, row := range cells {
		t.writeLine(&sb, row, widths)
	}
	return sb.String()
}

// Print writes the table to w.
func (t *Table[T]) Print(w io.Writer) {
	fmt.Fprint(w, t.String())
}

func (t *Table[T]) writeLine(sb *strings.Builder, row []string, widths []int) {
	parts := make([]string, len(row))
	for i, s := range row {
		parts[i] = pad(s, widths[i], t.justifications[i])
	}
	if t.borders {
		fmt.Fprintf(sb, "| %s |\n", strings.Join(parts, " | "))
		return
	}
	fmt.Fprintf(sb, "%s\n", strings.TrimRight(strings.Join(parts, " | "), " "))
}

func pad(s string, width int, j Justification) string {
	gap := width - cellWidth(s)
	if gap <= 0 {
		return s
	}
	switch j {
	case JustifyRight:
		return strings.Repeat(" ", gap) + s
	case JustifyCenter:
		left := gap / 2
		return strings.Repeat(" ", left) + s + strings.Repeat(" ", gap-left)
	default:
		return s + strings.Repeat(" ", gap)
	}
}

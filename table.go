package term

import (
	"io"
	"sort"
	"strings"
)

// Table holds and prints tabular data. Methods mutating the table operate
// in-place, but return the Table itself for convenience.
type Table interface {
	// SetPadRight sets the number of spaces to use for right padding. Defaults
	// to 1.
	SetPadRight(int) Table
	// Sort sorts the table.
	Sort(SortFunc) Table
	// WriterTo writes the table to the given io.Writer.
	io.WriterTo
}

// SortFunc returns if row a should come before row b.
type SortFunc func(a, b []string) bool

// NewTable returns a new table with the given rows.
func NewTable(rows [][]string) Table {
	return &table{
		rows:     rows,
		padRight: 1,
	}
}

// table implements the Table interface.
type table struct {
	rows     [][]string
	padRight int
}

// Sort is part of the table interface.
func (t *table) Sort(fn SortFunc) Table {
	s := &sortableRows{rows: t.rows, fn: fn}
	sort.Sort(s)
	return t
}

// sortableRows implements the sort.Interface.
type sortableRows struct {
	rows [][]string
	fn   SortFunc
}

// Less is part of the sort.Interface.
func (r *sortableRows) Less(i, j int) bool {
	return r.fn(r.rows[i], r.rows[j])
}

// Swap is part of the sort.Interface.
func (r *sortableRows) Swap(i, j int) {
	r.rows[i], r.rows[j] = r.rows[j], r.rows[i]
}

// Len is part of the sort.Interface.
func (r *sortableRows) Len() int {
	return len(r.rows)
}

// WriteTo is part of the Table interface.
func (t *table) WriteTo(w io.Writer) (int64, error) {
	widths := t.widths()
	for _, r := range t.rows {
		for i, width := range widths {
			var c string
			if i < len(r) {
				c = r[i]
			}
			wantWidth := width + t.padRight
			if l := len(c); l < wantWidth {
				c = c + strings.Repeat(" ", wantWidth-l)
			}
			io.WriteString(w, c)
		}
		io.WriteString(w, "\n")
	}
	return 0, nil
}

// SetPadRight is part of the Table interface.
func (t *table) SetPadRight(w int) Table {
	t.padRight = w
	return t
}

// widths returns the maximum number of characters for each column.
func (t *table) widths() []int {
	var widths []int
	for _, r := range t.rows {
		for i, c := range r {
			if i >= len(widths) {
				widths = append(widths, 0)
			}
			if l := len(c); l > widths[i] {
				widths[i] = l
			}
		}
	}
	return widths
}

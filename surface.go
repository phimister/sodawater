package soda

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Cell struct {
	char  rune
	empty bool
	style *lipgloss.Style
}

type Surface struct {
	width  int
	height int
	cells  []Cell
}

func NewSurface(w int, h int) *Surface {
	return &Surface{
		width: w,
		height: h,
		cells: make([]Cell, w*h),
	}
}

func (s *Surface) Set(x int, y int, cell Cell) {
	if x < 0 || y < 0 || x >= s.width || y >= s.height {
		return
	}

	s.cells[(y * s.width) + x] = cell
}

func (s *Surface) Get(x int, y int) Cell {
	if x < 0 || y < 0 || x >= s.width || y >= s.height {
		return Cell{empty: true}
	}

	return s.cells[(y * s.width) + x]
}

func (s *Surface) Blit(x int, y int, src *Surface) {
	// Source surface clipping
	srcX0, srcY0 := 0, 0
	srcX1, srcY1 := src.width, src.height

	if x < 0 { srcX0 = -x }
	if y < 0 { srcY0 = -y }
	if x + src.width > s.width { srcX1 = s.width - x }
	if y + src.height > s.height { srcY1 = s.height - y }

	if srcX0 >= srcX1 || srcY0 >= srcY1 {
		return // Fully OOB
	}

	// Block transfer
	srcIdx := (srcY0 * src.width) + srcX0
	dstIdx := ((y + srcY0) * s.width) + (x + srcX0)

	for range srcY1 - srcY0 {
		si := srcIdx
		di := dstIdx

		for range srcX1 - srcX0 {
			cell := src.cells[si]

			if !cell.empty {
				s.cells[di] = cell
			}

			si++
			di++
		}

		srcIdx += src.width
		dstIdx += s.width
	}
}

func (s *Surface) Render() string {
	cells := s.cells
	w := s.width
	h := s.height

	var out strings.Builder
	out.Grow((w + 1) * h)

	var run strings.Builder
	var currentStyle *lipgloss.Style

	flush := func() {
		if run.Len() == 0 {
			return
		}

		if currentStyle != nil {
			out.WriteString(currentStyle.Render(run.String()))
		} else {
			out.WriteString(run.String())
		}

		run.Reset()
	}

	for y := range h {
		row := y * w

		for x := range w {
			c := cells[row + x]

			ch := ' '
			if !c.empty {
				ch = c.char
			}

			if run.Len() == 0 {
				currentStyle = c.style
			}

			if c.style != currentStyle {
				flush()
				currentStyle = c.style
			}

			run.WriteRune(ch)
		}

		flush()

		if y < h - 1 {
			out.WriteByte('\n')
		}
	}

	return out.String()
}

package soda

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Cell struct {
	char  rune
	empty bool
	style lipgloss.Style
	styled bool
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

	var sb strings.Builder
	sb.Grow(( w + 1 ) * h) // width + 1 for newlines on each row

	for y := range h {
		row := y * w

		for x := range w {
			c := cells[row+x]

			if c.empty {
				sb.WriteRune(' ')
				continue
			}

			if c.styled { // FIXME: Batch styling for better performance
				sb.WriteString(c.style.Render(string(c.char)))
			} else {
				sb.WriteRune(c.char)
			}
		}

		if y < h - 1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

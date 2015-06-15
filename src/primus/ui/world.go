package ui

import (
	ui "github.com/gizak/termui"
)

type WorldView struct {
	ui.Block

	CenterX int
	CenterY int

	OutOfBounds ui.Point

	WorldWidth  int
	WorldHeight int
	World       [][]ui.Point
}

func NewWorldView(b ui.Block, w, h int) *WorldView {
	v := &WorldView{
		Block:       b,
		WorldWidth:  w,
		WorldHeight: h,
	}

	v.World = make([][]ui.Point, w)
	for x := range v.World {
		v.World[x] = make([]ui.Point, h)
	}

	return v
}

func (v *WorldView) Buffer() []ui.Point {
	ps := v.Block.Buffer()

	x, y, w, h := v.Block.InnerBounds()

	cdx := x + w/2
	cdy := y + h/2

	for dx := x; dx <= w; dx++ {
		wx := dx + v.CenterX - cdx
		for dy := y; dy <= h; dy++ {
			wy := dy + v.CenterY - cdy
			if wy < 0 || wy >= v.WorldHeight || wx < 0 || wx >= v.WorldWidth {
				p := v.OutOfBounds
				if p.Ch == rune(0) {
					continue
				}
				p.X = dx
				p.Y = dy
				ps = append(ps, p)
				continue
			}
			p := v.World[wx][wy]
			if p.Ch == rune(0) {
				continue
			}
			p.X = dx
			p.Y = dy
			ps = append(ps, p)
		}
	}

	return ps
}

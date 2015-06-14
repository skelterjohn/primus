package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui"
)

var lg = makeLog()

func makeLog() *log.Logger {
	fout, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	return log.New(fout, "", 0)
}

type ViewBlock struct {
	ui.Block

	VerticalOffset   int
	HorizontalOffset int

	WorldWidth  int
	WorldHeight int
	World       [][]ui.Point
}

func NewViewBlock(b ui.Block, w, h int) *ViewBlock {
	v := &ViewBlock{
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

func (v *ViewBlock) Buffer() []ui.Point {
	ps := v.Block.Buffer()

	x, y, w, h := v.Block.InnerBounds()

	for dx := x; dx <= w; dx++ {
		wx := dx + v.HorizontalOffset
		if wx < 0 || wx > v.WorldWidth {
			continue
		}
		for dy := y; dy <= h; dy++ {
			wy := dy + v.VerticalOffset
			if wy < 0 || wy > v.WorldHeight {
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

func main() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	b := ui.Block{
		HasBorder: true,
		IsDisplay: true,
	}
	v := NewViewBlock(b, 20, 20)
	v.HorizontalOffset = 0
	v.VerticalOffset = 0
	v.Width = 10
	v.Height = 10

	v.World[1][1] = ui.Point{
		Ch: 'x',
		Fg: ui.ColorRed,
	}
	v.World[2][2] = ui.Point{
		Ch: 'y',
		Fg: ui.ColorRed,
	}
	v.World[3][5] = ui.Point{
		Ch: 'z',
		Fg: ui.ColorRed,
	}
	v.World[7][5] = ui.Point{
		Ch: '1',
		Fg: ui.ColorRed,
	}
	v.World[8][5] = ui.Point{
		Ch: '2',
		Fg: ui.ColorRed,
	}
	v.World[9][5] = ui.Point{
		Ch: '3',
		Fg: ui.ColorRed,
	}

	ui.Render(v)
	<-ui.EventCh()
	v.HorizontalOffset += 1
	ui.Render(v)
	<-ui.EventCh()
}
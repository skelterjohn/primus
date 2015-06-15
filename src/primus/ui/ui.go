package ui

import (
	"log"

	ui "github.com/gizak/termui"
)

func Init() {
	err := ui.Init()
	if err != nil {
		log.Fatalf("initializing termui: %q", err)
	}
}

func Close() {
	ui.Close()
}

type T struct {
	View *WorldView
}

func NewT() *T {
	b := ui.Block{
		HasBorder: true,
		IsDisplay: true,
		Width:     30,
		Height:    30,
	}
	w := NewWorldView(b, 20, 20)
	w.CenterX = 5
	w.CenterY = 5
	w.OutOfBounds = ui.Point{
		Ch: '#',
	}
	return &T{
		View: w,
	}
}

func (t *T) Render() {
	ui.Render(t.View)
}

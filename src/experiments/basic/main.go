package main

import (
	"log"
	"os"

	"github.com/gizak/termui"

	"primus/ui"
)

var lg = makeLog()

func makeLog() *log.Logger {
	fout, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	return log.New(fout, "", 0)
}

func main() {

	ui.Init()
	defer ui.Close()

	t := ui.NewT()

	t.View.World[1][1] = termui.Point{
		Ch: 'x',
		Fg: termui.ColorRed,
	}
	t.View.World[2][2] = termui.Point{
		Ch: 'y',
		Fg: termui.ColorRed,
	}
	t.View.World[3][5] = termui.Point{
		Ch: 'z',
		Fg: termui.ColorRed,
	}
	t.View.World[7][5] = termui.Point{
		Ch: '1',
		Fg: termui.ColorRed,
	}
	t.View.World[8][5] = termui.Point{
		Ch: '2',
		Fg: termui.ColorRed,
	}
	t.View.World[9][5] = termui.Point{
		Ch: '3',
		Fg: termui.ColorRed,
	}

	t.Render()
	<-termui.EventCh()
	for i := 0; i < 5; i++ {
		t.View.CenterX += 1
		t.Render()
		<-termui.EventCh()
	}
}

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/style"
	"github.com/wizenerd/color"
	"github.com/wizenerd/grid"
	"github.com/wizenerd/ui/examples"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	a := New(gridExamples()...)
	vecty.RenderBody(a)
}

type App struct {
	vecty.Core
	grids []*examples.Example
}

func New(e ...*examples.Example) *App {
	return &App{
		grids: e,
	}
}

func (a *App) Render() *vecty.HTML {
	var v []vecty.MarkupOrComponentOrHTML
	for _, e := range a.grids {
		v = append(v, e)
	}
	return elem.Body(v...)
}

func gridExamples() []*examples.Example {
	return []*examples.Example{
		&examples.Example{
			View: grid1(),
			Code: grid1Txt,
		},
		&examples.Example{
			View: grid2(),
		},
		&examples.Example{
			View: grid3(),
		},
		&examples.Example{
			View: grid4(),
		},
	}
}

var cellColor = color.Class(color.Blue, color.A400)

func cellChild(m ...vecty.MarkupOrComponentOrHTML) vecty.MarkupOrComponentOrHTML {
	var v []vecty.MarkupOrComponentOrHTML
	v = append(v, m...)
	v = append(v, style.Height(style.Px(50)))
	v = append(v, style.Color("white"))
	v = append(v, vecty.Style("background-color", "#BDBDBD"))
	v = append(v, vecty.Style("box-sizing", "border-box"))
	v = append(v, vecty.Style("padding-leftr", string(style.Px(8))))
	v = append(v, vecty.Style("padding-top", string(style.Px(4))))
	return vecty.List(v)
}

func cellChild2(m ...vecty.MarkupOrComponentOrHTML) vecty.MarkupOrComponentOrHTML {
	var v []vecty.MarkupOrComponentOrHTML
	v = append(v, m...)
	v = append(v, style.Height(style.Px(200)))
	v = append(v, style.Color("white"))
	v = append(v, vecty.Style("background-color", "#BDBDBD"))
	v = append(v, vecty.Style("box-sizing", "border-box"))
	v = append(v, vecty.Style("padding-leftr", string(style.Px(8))))
	v = append(v, vecty.Style("padding-top", string(style.Px(4))))
	return vecty.List(v)
}

func grid1() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
		},
	}
}

var grid1Txt = `
var cellColor = color.Class(color.Blue, color.A400)

func cellChild(m ...vecty.MarkupOrComponentOrHTML) vecty.MarkupOrComponentOrHTML {
	var v []vecty.MarkupOrComponentOrHTML
	v = append(v, m...)
	v = append(v, style.Height(style.Px(50)))
	v = append(v, style.Color("white"))
	v = append(v, vecty.Style("background-color", "#BDBDBD"))
	v = append(v, vecty.Style("box-sizing", "border-box"))
	v = append(v, vecty.Style("padding-leftr", string(style.Px(8))))
	v = append(v, vecty.Style("padding-top", string(style.Px(4))))
	return vecty.List(v)
}

func grid1() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
		},
	}
}
`

func grid2() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:     grid.Default,
				Size:     4,
				Children: cellChild2(vecty.Text("4")),
			},
			{
				Mode:     grid.Default,
				Size:     4,
				Children: cellChild2(vecty.Text("4")),
			},
			{
				Mode:     grid.Default,
				Size:     4,
				Children: cellChild2(vecty.Text("4")),
			},
		},
	}
}

func grid3() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:     grid.Default,
				Size:     6,
				Children: cellChild2(vecty.Text("6")),
			},
			{
				Mode:     grid.Default,
				Size:     4,
				Children: cellChild2(vecty.Text("4")),
			},
			{
				Mode:     grid.Default,
				Size:     2,
				Children: cellChild2(vecty.Text("2")),
			},
		},
	}
}

func grid4() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:       grid.Default | grid.Tablet,
				Size:       6,
				TabletSize: 8,
				Children:   cellChild2(vecty.Text("6 (8 tablet)")),
			},
			{
				Mode:       grid.Default | grid.Tablet,
				Size:       4,
				TabletSize: 6,
				Children:   cellChild2(vecty.Text("4 (6 tablet)")),
			},
			{
				Mode:      grid.Default | grid.Phone,
				Size:      2,
				PhoneSIze: 4,
				Children:  cellChild2(vecty.Text("2 (4 phone)")),
			},
		},
	}
}

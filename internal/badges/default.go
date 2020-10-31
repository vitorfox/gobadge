package badges

import (
	"fmt"
	"math"

	"github.com/vitorfox/gobadge/package/svg"
)

type Default struct {
	Name string
}

func NewDefault(name string) *Default {
	return &Default{
		Name: name,
	}
}

func (d *Default) Build(params map[string]string) *svg.Node {

	value := params["value"]
	backgroundColor := params["background-value-color"]
	textColor := params["text-value-color"]

	nameSize := int64(math.Round(7.5 * float64(len(d.Name))))
	nameSizeS := fmt.Sprintf("%d", nameSize)

	valueSize := int64(math.Round(7.5*float64(len(value)))) + 10
	valueSizeS := fmt.Sprintf("%d", valueSize)

	totalSizeS := fmt.Sprintf("%d", nameSize+valueSize)

	nameXPosition := fmt.Sprintf("%d", nameSize/2)
	valueXPosition := fmt.Sprintf("%d", (nameSize+valueSize)-(valueSize/2))

	pathP1 := fmt.Sprintf("M0 0 h%s v20 H0 z", nameSizeS)
	pathP2 := fmt.Sprintf("M%s 0 h%s v20 H%s z", nameSizeS, valueSizeS, nameSizeS)
	pathP3 := fmt.Sprintf("M0 0 h%s v20 H0 z", totalSizeS)

	s := svg.New(svg.String("width", totalSizeS), svg.String("height", "20"))
	lg := s.Add("linearGradient",
		svg.String("id", "b"),
		svg.String("x2", "0"),
		svg.String("y2", "100%"),
	)

	lg.Add("stop",
		svg.String("offset", "0%"),
		svg.String("stop-color", "#bbb"),
		svg.String("stop-opacity", ".1"),
	)

	lg.Add("stop",
		svg.String("offset", "1%"),
		svg.String("stop-opacity", ".1"),
	)

	s.Add("mask", svg.String("id", "a")).
		Add("rect",
			svg.String("width", totalSizeS),
			svg.String("height", "20"),
			svg.String("rx", "3"),
			svg.String("fill", "#fff"),
		)

	g := s.Add("g", svg.String("mask", "url(#a)"))
	g.Add("path", svg.String("fill", "#555"), svg.String("d", pathP1))
	g.Add("path", svg.String("fill", backgroundColor), svg.String("d", pathP2))
	g.Add("path", svg.String("fill", "url(#b)"), svg.String("d", pathP3))

	g = s.Add("g", svg.String("fill", "#fff"), svg.String("text-anchor", "middle")).
		Add("g", svg.String("font-family", "DejaVu Sans,Verdana,Geneva,sans-serif"), svg.String("font-size", "11"))

	g.Add("text", svg.String("x", nameXPosition), svg.String("y", "15"), svg.String("fill", "#010101"), svg.String("fill-opacity", ".3")).Add("", svg.Simple(d.Name))
	g.Add("text", svg.String("x", nameXPosition), svg.String("y", "14")).Add("", svg.Simple(d.Name))
	g.Add("text", svg.String("x", valueXPosition), svg.String("y", "15"), svg.String("fill", "#010101"), svg.String("fill-opacity", ".3")).Add("", svg.Simple(value))
	g.Add("text", svg.String("x", valueXPosition), svg.String("y", "14"), svg.String("fill", textColor)).Add("", svg.Simple(value))

	return s
}

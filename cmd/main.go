package main

import (
	"os"

	"github.com/vitorfox/gobadge/package/svg"
)

func main() {

	s := svg.New(svg.String("width", "120"), svg.String("height", "20"))
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
			svg.String("width", "120"),
			svg.String("height", "20"),
			svg.String("rx", "3"),
			svg.String("fill", "#fff"),
		)

	g := s.Add("g", svg.String("mask", "url(#a)"))
	g.Add("path", svg.String("fill", "#555"), svg.String("d", "M0 0 h62 v20 H0 z"))
	g.Add("path", svg.String("fill", "#9f9f9f"), svg.String("d", "M62 0 h58 v20 H62 z"))
	g.Add("path", svg.String("fill", "url(#b)"), svg.String("d", "M0 0 h120 v20 H0 z"))

	g = s.Add("g", svg.String("fill", "#fff"), svg.String("text-anchor", "middle")).
		Add("g", svg.String("font-family", "DejaVu Sans,Verdana,Geneva,sans-serif"), svg.String("font-size", "11"))

	g.Add("text", svg.String("x", "31"), svg.String("y", "15"), svg.String("fill", "#010101"), svg.String("fill-opacity", ".3")).Add("", svg.Simple("coverage"))
	g.Add("text", svg.String("x", "31"), svg.String("y", "14")).Add("", svg.Simple("coverage"))
	g.Add("text", svg.String("x", "91"), svg.String("y", "15"), svg.String("fill", "#010101"), svg.String("fill-opacity", ".3")).Add("", svg.Simple("unknown"))
	g.Add("text", svg.String("x", "91"), svg.String("y", "14")).Add("", svg.Simple("unknown"))

	s.Build(os.Stdout)

}

package main

import (
	"flag"
	"log"

	"github.com/fullstack-lang/gongsvg/go/fullstack"
	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/static"
)

var ()

const RectangleWidth = 160
const RectangleHeigth = 50
const numberOfRectanglePerLine = 6

func main() {

	log.SetPrefix("gongsvg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup GORM
	r := static.ServeStaticFiles(false)
	stage := fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsvg/go/models")

	svg := (&models.Layer{Name: "SVG2"}).Stage(stage)
	svg.Display = true

	// create an array of rectangle
	for idx, color := range models.Colors {
		column := float64(idx % numberOfRectanglePerLine)
		line := float64(int(
			float64(idx) /
				float64(numberOfRectanglePerLine)))

		appendRect(stage, svg, color,
			RectangleWidth*column,
			RectangleHeigth*line)
	}

	stage.Commit()

	log.Printf("Server ready serve on localhost:8082")
	r.Run(":8082")
}

func appendRect(stage *models.StageStruct, svg *models.Layer, color models.ColorType, x, y float64) {
	rect := (&models.Rect{Name: color.ToString()}).Stage(stage)
	rect.Color = color.ToString()
	rect.Height = RectangleHeigth
	rect.Width = RectangleWidth
	rect.X = x
	rect.Y = y
	rect.FillOpacity = 50.0

	text := (&models.Text{Name: color.ToString()}).Stage(stage)
	text.X = x + 2
	text.Y = y + RectangleHeigth/2
	text.Content = color.ToString()
	text.FillOpacity = 100.0

	svg.Texts = append(svg.Texts, text)
	svg.Rects = append(svg.Rects, rect)
}

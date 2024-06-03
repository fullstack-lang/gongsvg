package main

import (
	"flag"
	"log"
	"strconv"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"
	gongsvg_static "github.com/fullstack-lang/gongsvg/go/static"
)

var (
	port = flag.Int("port", 8082, "port server")

	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
)

const RectangleWidth = 160
const RectangleHeigth = 50
const numberOfRectanglePerLine = 6

func main() {

	log.SetPrefix("gongsvg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongsvg_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongsvg_stack.NewStack(r, "gongsvg", "", "", "", false, true)
	stack.Probe.Refresh()

	colorLayer := (&gongsvg_models.Layer{Name: "color"}).Stage(stack.Stage)
	colorLayer.Display = true

	textLayer := (&gongsvg_models.Layer{Name: "text"}).Stage(stack.Stage)
	textLayer.Display = true

	// create an array of rectangle
	for idx, color := range gongsvg_models.Colors {
		column := float64(idx % numberOfRectanglePerLine)
		line := float64(int(
			float64(idx) /
				float64(numberOfRectanglePerLine)))

		appendRect(stack.Stage,
			colorLayer,
			textLayer,
			color,
			RectangleWidth*column,
			RectangleHeigth*line)
	}

	svg := new(gongsvg_models.SVG).Stage(stack.Stage)

	svg.Name = "SVG"
	svg.Layers = append(svg.Layers, colorLayer)
	svg.Layers = append(svg.Layers, textLayer)

	stack.Stage.Commit()

	log.Printf("Server ready serve on localhost" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func appendRect(
	stage *gongsvg_models.StageStruct,
	colorLayer *gongsvg_models.Layer,
	textLayer *gongsvg_models.Layer,

	color gongsvg_models.ColorType, x, y float64) {
	rect := (&gongsvg_models.Rect{Name: color.ToString()}).Stage(stage)
	rect.Color = color.ToString()
	rect.Height = RectangleHeigth
	rect.Width = RectangleWidth
	rect.X = x
	rect.Y = y
	rect.FillOpacity = 50.0

	text := (&gongsvg_models.Text{Name: color.ToString()}).Stage(stage)
	text.X = x + 2
	text.Y = y + RectangleHeigth/2
	text.Content = color.ToString()
	text.FillOpacity = 100.0

	textLayer.Texts = append(textLayer.Texts, text)
	colorLayer.Rects = append(colorLayer.Rects, rect)
}

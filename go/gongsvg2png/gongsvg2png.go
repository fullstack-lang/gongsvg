package gongsvg2png

import (
	"fmt"
	"strings"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	"image"
	"image/color"
	"log"
	"math"

	"golang.org/x/image/colornames"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func Gongsvg2png(stage *gongsvg_models.StageStruct, path string) {

	gongsvg_models.Stage.Checkout()
	for svg := range stage.SVGs {

		// Initialize the graphic context on an RGBA image
		dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
		gc := draw2dimg.NewGraphicContext(dest)

		for _, line := range svg.Lines {
			line2gc(gc, line)
		}

		for _, rect := range svg.Rects {
			rect2gc(gc, rect)
		}

		pathPng := fmt.Sprintf("./%s.png", svg.Name)

		// Save to png
		err := draw2dimg.SaveToPngFile(pathPng, dest)
		if err != nil {
			log.Printf("Saving %q failed: %v", pathPng, err)
			return
		}
	}

}

func rect2gc(gc draw2d.GraphicContext, rect *gongsvg_models.Rect) {

	// set rect properties
	gc.SetLineCap(draw2d.RoundCap)
	gc.SetLineWidth(rect.StrokeWidth)
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	fillOpacityUint := uint8(rect.FillOpacity * 255.0)

	var rectFillColor color.Color
	if strings.HasPrefix(rect.Color, "rgb") {
		var red, green, blue int
		_, err := fmt.Sscanf(rect.Color, "rgb(%d,%d,%d)", &red, &green, &blue)

		if err != nil {
			panic(err)
		}
		rectFillColor = color.RGBA{uint8(red), uint8(green), uint8(blue), fillOpacityUint}

	} else {
		rectFillColor = colornames.Map[rect.Color]
		r, g, b, _ := rectFillColor.RGBA()
		rectFillColor = color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(fillOpacityUint)}
	}
	gc.SetFillColor(rectFillColor)

	gc.BeginPath()
	gc.MoveTo(rect.X, rect.Y)
	gc.LineTo(rect.X+rect.Width, rect.Y)
	gc.LineTo(rect.X+rect.Width, rect.Y+rect.Height)
	gc.LineTo(rect.X, rect.Y+rect.Height)
	gc.Close()
	gc.FillStroke()
	path := gc.GetPath()
	gc.Fill(&path)
	gc.Stroke()

}

func line2gc(gc draw2d.GraphicContext, line *gongsvg_models.Line) {

	// set line properties
	gc.SetLineCap(draw2d.RoundCap)
	gc.SetLineWidth(5)
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

	gc.MoveTo(line.X1, line.Y1)
	gc.LineTo(line.X2, line.Y2)
	gc.Stroke()
}

// Draw the droid on a certain position.
func Draw(gc draw2d.GraphicContext, x, y float64) {
	// set the fill and stroke color of the droid
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

	// set line properties
	gc.SetLineCap(draw2d.RoundCap)
	gc.SetLineWidth(5)

	// head
	gc.MoveTo(x+30, y+70)
	gc.ArcTo(x+80, y+70, 50, 50, 180*(math.Pi/180), 180*(math.Pi/180))
	gc.Close()
	gc.FillStroke()
	gc.MoveTo(x+60, y+25)
	gc.LineTo(x+50, y+10)
	gc.MoveTo(x+100, y+25)
	gc.LineTo(x+110, y+10)
	gc.Stroke()

	// left eye
	draw2dkit.Circle(gc, x+60, y+45, 5)
	gc.FillStroke()

	// right eye
	draw2dkit.Circle(gc, x+100, y+45, 5)
	gc.FillStroke()

	// body
	draw2dkit.RoundedRectangle(gc, x+30, y+75, x+30+100, y+75+90, 10, 10)
	gc.FillStroke()
	draw2dkit.Rectangle(gc, x+30, y+75, x+30+100, y+75+80)
	gc.FillStroke()

	// left arm
	draw2dkit.RoundedRectangle(gc, x+5, y+80, x+5+20, y+80+70, 10, 10)
	gc.FillStroke()

	// right arm
	draw2dkit.RoundedRectangle(gc, x+135, y+80, x+135+20, y+80+70, 10, 10)
	gc.FillStroke()

	// left leg
	draw2dkit.RoundedRectangle(gc, x+50, y+150, x+50+20, y+150+50, 10, 10)
	gc.FillStroke()

	// right leg
	draw2dkit.RoundedRectangle(gc, x+90, y+150, x+90+20, y+150+50, 10, 10)
	gc.FillStroke()
}

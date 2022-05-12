package gongsvg2png

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	"image"
	"image/color"
	"log"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func Gongsvg2png(stage *gongsvg_models.StageStruct, path string) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Draw(gc, 65, 0)
	gongsvg_models.Stage.Checkout()
	for line := range stage.Lines {
		line2gc(gc, line, 20, 20)
	}

	pathPng := "./test.png"

	// Save to png
	err := draw2dimg.SaveToPngFile(pathPng, dest)
	if err != nil {
		log.Printf("Saving %q failed: %v", pathPng, err)
		return
	}

}

func line2gc(gc draw2d.GraphicContext, line *gongsvg_models.Line, x, y float64) {

	// set line properties
	gc.SetLineCap(draw2d.RoundCap)
	gc.SetLineWidth(5)
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

	gc.MoveTo(x+line.X1, y+line.Y1)
	gc.LineTo(x+line.X2, y+line.Y2)
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

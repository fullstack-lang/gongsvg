package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_rectangle models.StageStruct
var ___dummy__Time_rectangle time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_rectangle map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["rectangle"] = rectangleInjection
// }

// rectangleInjection will stage objects of database "rectangle"
func rectangleInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle
	__Circle__000000_Test := (&models.Circle{Name: `Test`}).Stage(stage)

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_Circle_Layer := (&models.Layer{Name: `Circle Layer`}).Stage(stage)
	__Layer__000001_Rectangle_layer := (&models.Layer{Name: `Rectangle layer`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Path

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Sample := (&models.Rect{Name: `Sample`}).Stage(stage)

	// Declarations of staged instances of SVG
	__SVG__000000_SVG := (&models.SVG{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of Text
	__Text__000000_Essai := (&models.Text{Name: `Essai`}).Stage(stage)

	// Setup of values

	// Circle values setup
	__Circle__000000_Test.Name = `Test`
	__Circle__000000_Test.CX = 400.000000
	__Circle__000000_Test.CY = 300.000000
	__Circle__000000_Test.Radius = 100.000000
	__Circle__000000_Test.Color = `lavender`
	__Circle__000000_Test.FillOpacity = 50.000000
	__Circle__000000_Test.Stroke = ``
	__Circle__000000_Test.StrokeWidth = 0.000000
	__Circle__000000_Test.StrokeDashArray = ``
	__Circle__000000_Test.Transform = ``

	// Layer values setup
	__Layer__000000_Circle_Layer.Display = false
	__Layer__000000_Circle_Layer.Name = `Circle Layer`

	// Layer values setup
	__Layer__000001_Rectangle_layer.Display = false
	__Layer__000001_Rectangle_layer.Name = `Rectangle layer`

	// Rect values setup
	__Rect__000000_Sample.Name = `Sample`
	__Rect__000000_Sample.X = 104.000000
	__Rect__000000_Sample.Y = 207.000000
	__Rect__000000_Sample.Width = 614.000000
	__Rect__000000_Sample.Height = 100.000000
	__Rect__000000_Sample.RX = 3.000000
	__Rect__000000_Sample.Color = `lightcyan`
	__Rect__000000_Sample.FillOpacity = 100.000000
	__Rect__000000_Sample.Stroke = `darkcyan`
	__Rect__000000_Sample.StrokeWidth = 2.000000
	__Rect__000000_Sample.StrokeDashArray = ``
	__Rect__000000_Sample.Transform = ``
	__Rect__000000_Sample.Selected = false

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`

	// Text values setup
	__Text__000000_Essai.Name = `Essai`
	__Text__000000_Essai.X = 50.000000
	__Text__000000_Essai.Y = 300.000000
	__Text__000000_Essai.Content = `Hello`
	__Text__000000_Essai.Color = ``
	__Text__000000_Essai.FillOpacity = 0.000000
	__Text__000000_Essai.Stroke = `black`
	__Text__000000_Essai.StrokeWidth = 2.000000
	__Text__000000_Essai.StrokeDashArray = ``
	__Text__000000_Essai.Transform = ``

	// Setup of pointers
	__Layer__000000_Circle_Layer.Circles = append(__Layer__000000_Circle_Layer.Circles, __Circle__000000_Test)
	__Layer__000001_Rectangle_layer.Rects = append(__Layer__000001_Rectangle_layer.Rects, __Rect__000000_Sample)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Circle_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000001_Rectangle_layer)
}



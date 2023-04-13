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
	__Layer__000000_Bottom_Rectangle_Layer := (&models.Layer{Name: `Bottom Rectangle Layer`}).Stage(stage)
	__Layer__000001_Circle_Layer := (&models.Layer{Name: `Circle Layer`}).Stage(stage)
	__Layer__000002_Line_layer := (&models.Layer{Name: `Line layer`}).Stage(stage)
	__Layer__000003_Top_Rectangle_layer := (&models.Layer{Name: `Top Rectangle layer`}).Stage(stage)

	// Declarations of staged instances of Line
	__Line__000000_Line_connecting_rect_Bottom_to_Top := (&models.Line{Name: `Line connecting rect Bottom to Top`}).Stage(stage)

	// Declarations of staged instances of Path

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Bottom := (&models.Rect{Name: `Bottom`}).Stage(stage)
	__Rect__000001_Top := (&models.Rect{Name: `Top`}).Stage(stage)

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
	__Layer__000000_Bottom_Rectangle_Layer.Display = false
	__Layer__000000_Bottom_Rectangle_Layer.Name = `Bottom Rectangle Layer`

	// Layer values setup
	__Layer__000001_Circle_Layer.Display = false
	__Layer__000001_Circle_Layer.Name = `Circle Layer`

	// Layer values setup
	__Layer__000002_Line_layer.Display = false
	__Layer__000002_Line_layer.Name = `Line layer`

	// Layer values setup
	__Layer__000003_Top_Rectangle_layer.Display = false
	__Layer__000003_Top_Rectangle_layer.Name = `Top Rectangle layer`

	// Line values setup
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Name = `Line connecting rect Bottom to Top`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.X1 = 224.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Y1 = 697.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.X2 = 158.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Y2 = 626.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Color = `olivedrab`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.FillOpacity = 0.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Stroke = `olivedrab`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeWidth = 3.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeDashArray = ``
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Transform = ``
	__Line__000000_Line_connecting_rect_Bottom_to_Top.MouseClickX = 195.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.MouseClickY = 718.000000

	// Rect values setup
	__Rect__000000_Bottom.Name = `Bottom`
	__Rect__000000_Bottom.X = 492.000000
	__Rect__000000_Bottom.Y = 92.000000
	__Rect__000000_Bottom.Width = 234.000000
	__Rect__000000_Bottom.Height = 300.000000
	__Rect__000000_Bottom.RX = 5.000000
	__Rect__000000_Bottom.Color = `bisque`
	__Rect__000000_Bottom.FillOpacity = 50.000000
	__Rect__000000_Bottom.Stroke = `lightcoral`
	__Rect__000000_Bottom.StrokeWidth = 3.000000
	__Rect__000000_Bottom.StrokeDashArray = ``
	__Rect__000000_Bottom.Transform = ``
	__Rect__000000_Bottom.IsSelectable = true
	__Rect__000000_Bottom.IsSelected = true
	__Rect__000000_Bottom.CanHaveHorizontalHandles = true
	__Rect__000000_Bottom.HasHorizontalHandles = true
	__Rect__000000_Bottom.CanMoveHorizontaly = false
	__Rect__000000_Bottom.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Top.Name = `Top`
	__Rect__000001_Top.X = 102.000000
	__Rect__000001_Top.Y = 326.000000
	__Rect__000001_Top.Width = 234.000000
	__Rect__000001_Top.Height = 300.000000
	__Rect__000001_Top.RX = 3.000000
	__Rect__000001_Top.Color = `lightcyan`
	__Rect__000001_Top.FillOpacity = 100.000000
	__Rect__000001_Top.Stroke = `darkcyan`
	__Rect__000001_Top.StrokeWidth = 2.000000
	__Rect__000001_Top.StrokeDashArray = ``
	__Rect__000001_Top.Transform = ``
	__Rect__000001_Top.IsSelectable = true
	__Rect__000001_Top.IsSelected = false
	__Rect__000001_Top.CanHaveHorizontalHandles = true
	__Rect__000001_Top.HasHorizontalHandles = false
	__Rect__000001_Top.CanMoveHorizontaly = true
	__Rect__000001_Top.CanMoveVerticaly = false

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.DrawingState = models.NOT_DRAWING_LINE

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
	__Layer__000000_Bottom_Rectangle_Layer.Rects = append(__Layer__000000_Bottom_Rectangle_Layer.Rects, __Rect__000000_Bottom)
	__Layer__000000_Bottom_Rectangle_Layer.Rects = append(__Layer__000000_Bottom_Rectangle_Layer.Rects, __Rect__000001_Top)
	__Layer__000001_Circle_Layer.Circles = append(__Layer__000001_Circle_Layer.Circles, __Circle__000000_Test)
	__Layer__000002_Line_layer.Lines = append(__Layer__000002_Line_layer.Lines, __Line__000000_Line_connecting_rect_Bottom_to_Top)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Bottom_Rectangle_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000003_Top_Rectangle_layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000001_Circle_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000002_Line_layer)
	__SVG__000000_SVG.StartRect = __Rect__000001_Top
	__SVG__000000_SVG.EndRect = __Rect__000000_Bottom
}



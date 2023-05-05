package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var ___dummy__Time_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of AnchoredText

	// Declarations of staged instances of Animate
	__Animate__000000_A1 := (&models.Animate{Name: `A1`}).Stage(stage)
	__Animate__000001_Animate_oppacity := (&models.Animate{Name: `Animate oppacity`}).Stage(stage)
	__Animate__000002_C1a_animation := (&models.Animate{Name: `C1a animation`}).Stage(stage)
	__Animate__000003_Move_text := (&models.Animate{Name: `Move text`}).Stage(stage)

	// Declarations of staged instances of Circle
	__Circle__000000_C1a := (&models.Circle{Name: `C1a`}).Stage(stage)

	// Declarations of staged instances of Ellipse
	__Ellipse__000000_Ellipse_Example_SVG := (&models.Ellipse{Name: `Ellipse Example SVG`}).Stage(stage)

	// Declarations of staged instances of Layer
	__Layer__000000_SVG2 := (&models.Layer{Name: `SVG2`}).Stage(stage)

	// Declarations of staged instances of Line
	__Line__000000_Line := (&models.Line{Name: `Line`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Path
	__Path__000000_Path_example := (&models.Path{Name: `Path example`}).Stage(stage)

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone
	__Polygone__000000_Polygone_example_SVG := (&models.Polygone{Name: `Polygone example SVG`}).Stage(stage)

	// Declarations of staged instances of Polyline
	__Polyline__000000_Polyline_example_SVG := (&models.Polyline{Name: `Polyline example SVG`}).Stage(stage)

	// Declarations of staged instances of Rect
	__Rect__000000_R1 := (&models.Rect{Name: `R1`}).Stage(stage)
	__Rect__000001_R2 := (&models.Rect{Name: `R2`}).Stage(stage)
	__Rect__000002_R3 := (&models.Rect{Name: `R3`}).Stage(stage)
	__Rect__000003_R4_rounded := (&models.Rect{Name: `R4 rounded`}).Stage(stage)
	__Rect__000004_Test_Rect := (&models.Rect{Name: `Test Rect`}).Stage(stage)

	// Declarations of staged instances of SVG
	__SVG__000000_SVG := (&models.SVG{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of Text
	__Text__000000_Bonjorno := (&models.Text{Name: `Bonjorno`}).Stage(stage)
	__Text__000001_Hello := (&models.Text{Name: `Hello`}).Stage(stage)

	// Setup of values

	// Animate values setup
	__Animate__000000_A1.Name = `A1`
	__Animate__000000_A1.AttributeName = `x`
	__Animate__000000_A1.Values = `0;100;0`
	__Animate__000000_A1.Dur = `4s`
	__Animate__000000_A1.RepeatCount = `indefinite`

	// Animate values setup
	__Animate__000001_Animate_oppacity.Name = `Animate oppacity`
	__Animate__000001_Animate_oppacity.AttributeName = `fill-opacity`
	__Animate__000001_Animate_oppacity.Values = `0.2;0.8;0.2`
	__Animate__000001_Animate_oppacity.Dur = `5s`
	__Animate__000001_Animate_oppacity.RepeatCount = `indefinite`

	// Animate values setup
	__Animate__000002_C1a_animation.Name = `C1a animation`
	__Animate__000002_C1a_animation.AttributeName = `r`
	__Animate__000002_C1a_animation.Values = `50;150;50`
	__Animate__000002_C1a_animation.Dur = `4s`
	__Animate__000002_C1a_animation.RepeatCount = `indefinite`

	// Animate values setup
	__Animate__000003_Move_text.Name = `Move text`
	__Animate__000003_Move_text.AttributeName = `x`
	__Animate__000003_Move_text.Values = `30;100;30`
	__Animate__000003_Move_text.Dur = `4s`
	__Animate__000003_Move_text.RepeatCount = `indefinite`

	// Circle values setup
	__Circle__000000_C1a.Name = `C1a`
	__Circle__000000_C1a.CX = 350.000000
	__Circle__000000_C1a.CY = 300.000000
	__Circle__000000_C1a.Radius = 50.000000
	__Circle__000000_C1a.Color = `greenlight`
	__Circle__000000_C1a.FillOpacity = 0.800000
	__Circle__000000_C1a.Stroke = `blue`
	__Circle__000000_C1a.StrokeWidth = 1.000000
	__Circle__000000_C1a.StrokeDashArray = ``
	__Circle__000000_C1a.StrokeDashArrayWhenSelected = ``
	__Circle__000000_C1a.Transform = ``

	// Ellipse values setup
	__Ellipse__000000_Ellipse_Example_SVG.Name = `Ellipse Example SVG`
	__Ellipse__000000_Ellipse_Example_SVG.CX = 150.000000
	__Ellipse__000000_Ellipse_Example_SVG.CY = 150.000000
	__Ellipse__000000_Ellipse_Example_SVG.RX = 200.000000
	__Ellipse__000000_Ellipse_Example_SVG.RY = 50.000000
	__Ellipse__000000_Ellipse_Example_SVG.Color = `red`
	__Ellipse__000000_Ellipse_Example_SVG.FillOpacity = 0.300000
	__Ellipse__000000_Ellipse_Example_SVG.Stroke = `black`
	__Ellipse__000000_Ellipse_Example_SVG.StrokeWidth = 2.000000
	__Ellipse__000000_Ellipse_Example_SVG.StrokeDashArray = ``
	__Ellipse__000000_Ellipse_Example_SVG.StrokeDashArrayWhenSelected = ``
	__Ellipse__000000_Ellipse_Example_SVG.Transform = ``

	// Layer values setup
	__Layer__000000_SVG2.Display = true
	__Layer__000000_SVG2.Name = `SVG2`

	// Line values setup
	__Line__000000_Line.Name = `Line`
	__Line__000000_Line.X1 = 30.000000
	__Line__000000_Line.Y1 = 100.000000
	__Line__000000_Line.X2 = 80.000000
	__Line__000000_Line.Y2 = 300.000000
	__Line__000000_Line.Color = `red`
	__Line__000000_Line.FillOpacity = 1.000000
	__Line__000000_Line.Stroke = `red`
	__Line__000000_Line.StrokeWidth = 1.000000
	__Line__000000_Line.StrokeDashArray = ``
	__Line__000000_Line.StrokeDashArrayWhenSelected = ``
	__Line__000000_Line.Transform = ``
	__Line__000000_Line.MouseClickX = 0.000000
	__Line__000000_Line.MouseClickY = 0.000000

	// Path values setup
	__Path__000000_Path_example.Name = `Path example`
	__Path__000000_Path_example.Definition = `M 10,30            A 20,20 0,0,1 50,30            A 20,20 0,0,1 90,30            Q 90,60 50,90            Q 10,60 10,30 z`
	__Path__000000_Path_example.Color = `black`
	__Path__000000_Path_example.FillOpacity = 0.500000
	__Path__000000_Path_example.Stroke = `blue`
	__Path__000000_Path_example.StrokeWidth = 2.000000
	__Path__000000_Path_example.StrokeDashArray = `4 4`
	__Path__000000_Path_example.StrokeDashArrayWhenSelected = ``
	__Path__000000_Path_example.Transform = ``

	// Polygone values setup
	__Polygone__000000_Polygone_example_SVG.Name = `Polygone example SVG`
	__Polygone__000000_Polygone_example_SVG.Points = `100,100 150,25 150,75 200,0`
	__Polygone__000000_Polygone_example_SVG.Color = `green`
	__Polygone__000000_Polygone_example_SVG.FillOpacity = 0.800000
	__Polygone__000000_Polygone_example_SVG.Stroke = ``
	__Polygone__000000_Polygone_example_SVG.StrokeWidth = 0.000000
	__Polygone__000000_Polygone_example_SVG.StrokeDashArray = ``
	__Polygone__000000_Polygone_example_SVG.StrokeDashArrayWhenSelected = ``
	__Polygone__000000_Polygone_example_SVG.Transform = ``

	// Polyline values setup
	__Polyline__000000_Polyline_example_SVG.Name = `Polyline example SVG`
	__Polyline__000000_Polyline_example_SVG.Points = `350,100 150,25 150,75 200,0`
	__Polyline__000000_Polyline_example_SVG.Color = `yellow`
	__Polyline__000000_Polyline_example_SVG.FillOpacity = 0.700000
	__Polyline__000000_Polyline_example_SVG.Stroke = ``
	__Polyline__000000_Polyline_example_SVG.StrokeWidth = 0.000000
	__Polyline__000000_Polyline_example_SVG.StrokeDashArray = ``
	__Polyline__000000_Polyline_example_SVG.StrokeDashArrayWhenSelected = ``
	__Polyline__000000_Polyline_example_SVG.Transform = ``

	// Rect values setup
	__Rect__000000_R1.Name = `R1`
	__Rect__000000_R1.X = 200.000000
	__Rect__000000_R1.Y = 0.100000
	__Rect__000000_R1.Width = 50.000000
	__Rect__000000_R1.Height = 120.000000
	__Rect__000000_R1.RX = 0.000000
	__Rect__000000_R1.Color = `rgb(255, 0, 0)`
	__Rect__000000_R1.FillOpacity = 0.500000
	__Rect__000000_R1.Stroke = `blue`
	__Rect__000000_R1.StrokeWidth = 1.000000
	__Rect__000000_R1.StrokeDashArray = `4 4`
	__Rect__000000_R1.StrokeDashArrayWhenSelected = ``
	__Rect__000000_R1.Transform = ``
	__Rect__000000_R1.IsSelectable = false
	__Rect__000000_R1.IsSelected = false
	__Rect__000000_R1.CanHaveLeftHandle = false
	__Rect__000000_R1.HasLeftHandle = false
	__Rect__000000_R1.CanHaveRightHandle = false
	__Rect__000000_R1.HasRightHandle = false
	__Rect__000000_R1.CanHaveTopHandle = false
	__Rect__000000_R1.HasTopHandle = false
	__Rect__000000_R1.CanHaveBottomHandle = false
	__Rect__000000_R1.HasBottomHandle = false
	__Rect__000000_R1.CanMoveHorizontaly = false
	__Rect__000000_R1.CanMoveVerticaly = false

	// Rect values setup
	__Rect__000001_R2.Name = `R2`
	__Rect__000001_R2.X = 10.000000
	__Rect__000001_R2.Y = 100.000000
	__Rect__000001_R2.Width = 150.000000
	__Rect__000001_R2.Height = 200.000000
	__Rect__000001_R2.RX = 0.000000
	__Rect__000001_R2.Color = `greenyellow`
	__Rect__000001_R2.FillOpacity = 0.900000
	__Rect__000001_R2.Stroke = `yellow`
	__Rect__000001_R2.StrokeWidth = 3.000000
	__Rect__000001_R2.StrokeDashArray = `4 1`
	__Rect__000001_R2.StrokeDashArrayWhenSelected = ``
	__Rect__000001_R2.Transform = ``
	__Rect__000001_R2.IsSelectable = false
	__Rect__000001_R2.IsSelected = false
	__Rect__000001_R2.CanHaveLeftHandle = false
	__Rect__000001_R2.HasLeftHandle = false
	__Rect__000001_R2.CanHaveRightHandle = false
	__Rect__000001_R2.HasRightHandle = false
	__Rect__000001_R2.CanHaveTopHandle = false
	__Rect__000001_R2.HasTopHandle = false
	__Rect__000001_R2.CanHaveBottomHandle = false
	__Rect__000001_R2.HasBottomHandle = false
	__Rect__000001_R2.CanMoveHorizontaly = false
	__Rect__000001_R2.CanMoveVerticaly = false

	// Rect values setup
	__Rect__000002_R3.Name = `R3`
	__Rect__000002_R3.X = 10.000000
	__Rect__000002_R3.Y = 40.000000
	__Rect__000002_R3.Width = 50.000000
	__Rect__000002_R3.Height = 50.000000
	__Rect__000002_R3.RX = 0.000000
	__Rect__000002_R3.Color = `rgb(124, 0, 0)`
	__Rect__000002_R3.FillOpacity = 0.200000
	__Rect__000002_R3.Stroke = `pink`
	__Rect__000002_R3.StrokeWidth = 6.000000
	__Rect__000002_R3.StrokeDashArray = ``
	__Rect__000002_R3.StrokeDashArrayWhenSelected = ``
	__Rect__000002_R3.Transform = ``
	__Rect__000002_R3.IsSelectable = false
	__Rect__000002_R3.IsSelected = false
	__Rect__000002_R3.CanHaveLeftHandle = false
	__Rect__000002_R3.HasLeftHandle = false
	__Rect__000002_R3.CanHaveRightHandle = false
	__Rect__000002_R3.HasRightHandle = false
	__Rect__000002_R3.CanHaveTopHandle = false
	__Rect__000002_R3.HasTopHandle = false
	__Rect__000002_R3.CanHaveBottomHandle = false
	__Rect__000002_R3.HasBottomHandle = false
	__Rect__000002_R3.CanMoveHorizontaly = false
	__Rect__000002_R3.CanMoveVerticaly = false

	// Rect values setup
	__Rect__000003_R4_rounded.Name = `R4 rounded`
	__Rect__000003_R4_rounded.X = 300.000000
	__Rect__000003_R4_rounded.Y = 300.000000
	__Rect__000003_R4_rounded.Width = 300.000000
	__Rect__000003_R4_rounded.Height = 400.000000
	__Rect__000003_R4_rounded.RX = 50.000000
	__Rect__000003_R4_rounded.Color = `red`
	__Rect__000003_R4_rounded.FillOpacity = 0.500000
	__Rect__000003_R4_rounded.Stroke = ``
	__Rect__000003_R4_rounded.StrokeWidth = 0.000000
	__Rect__000003_R4_rounded.StrokeDashArray = ``
	__Rect__000003_R4_rounded.StrokeDashArrayWhenSelected = ``
	__Rect__000003_R4_rounded.Transform = ``
	__Rect__000003_R4_rounded.IsSelectable = false
	__Rect__000003_R4_rounded.IsSelected = false
	__Rect__000003_R4_rounded.CanHaveLeftHandle = false
	__Rect__000003_R4_rounded.HasLeftHandle = false
	__Rect__000003_R4_rounded.CanHaveRightHandle = false
	__Rect__000003_R4_rounded.HasRightHandle = false
	__Rect__000003_R4_rounded.CanHaveTopHandle = false
	__Rect__000003_R4_rounded.HasTopHandle = false
	__Rect__000003_R4_rounded.CanHaveBottomHandle = false
	__Rect__000003_R4_rounded.HasBottomHandle = false
	__Rect__000003_R4_rounded.CanMoveHorizontaly = false
	__Rect__000003_R4_rounded.CanMoveVerticaly = false

	// Rect values setup
	__Rect__000004_Test_Rect.Name = `Test Rect`
	__Rect__000004_Test_Rect.X = 800.000000
	__Rect__000004_Test_Rect.Y = 800.000000
	__Rect__000004_Test_Rect.Width = 400.000000
	__Rect__000004_Test_Rect.Height = 400.000000
	__Rect__000004_Test_Rect.RX = 0.000000
	__Rect__000004_Test_Rect.Color = `blue`
	__Rect__000004_Test_Rect.FillOpacity = 0.300000
	__Rect__000004_Test_Rect.Stroke = ``
	__Rect__000004_Test_Rect.StrokeWidth = 0.000000
	__Rect__000004_Test_Rect.StrokeDashArray = ``
	__Rect__000004_Test_Rect.StrokeDashArrayWhenSelected = ``
	__Rect__000004_Test_Rect.Transform = ``
	__Rect__000004_Test_Rect.IsSelectable = false
	__Rect__000004_Test_Rect.IsSelected = false
	__Rect__000004_Test_Rect.CanHaveLeftHandle = false
	__Rect__000004_Test_Rect.HasLeftHandle = false
	__Rect__000004_Test_Rect.CanHaveRightHandle = false
	__Rect__000004_Test_Rect.HasRightHandle = false
	__Rect__000004_Test_Rect.CanHaveTopHandle = false
	__Rect__000004_Test_Rect.HasTopHandle = false
	__Rect__000004_Test_Rect.CanHaveBottomHandle = false
	__Rect__000004_Test_Rect.HasBottomHandle = false
	__Rect__000004_Test_Rect.CanMoveHorizontaly = false
	__Rect__000004_Test_Rect.CanMoveVerticaly = false

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`

	// Text values setup
	__Text__000000_Bonjorno.Name = `Bonjorno`
	__Text__000000_Bonjorno.X = 40.000000
	__Text__000000_Bonjorno.Y = 50.000000
	__Text__000000_Bonjorno.Content = `Bonjorno Mondo`
	__Text__000000_Bonjorno.Color = `green`
	__Text__000000_Bonjorno.FillOpacity = 1.000000
	__Text__000000_Bonjorno.Stroke = `black`
	__Text__000000_Bonjorno.StrokeWidth = 0.000000
	__Text__000000_Bonjorno.StrokeDashArray = ``
	__Text__000000_Bonjorno.StrokeDashArrayWhenSelected = ``
	__Text__000000_Bonjorno.Transform = ``

	// Text values setup
	__Text__000001_Hello.Name = `Hello`
	__Text__000001_Hello.X = 30.000000
	__Text__000001_Hello.Y = 30.000000
	__Text__000001_Hello.Content = `Hello World`
	__Text__000001_Hello.Color = `green`
	__Text__000001_Hello.FillOpacity = 1.000000
	__Text__000001_Hello.Stroke = `red`
	__Text__000001_Hello.StrokeWidth = 3.000000
	__Text__000001_Hello.StrokeDashArray = ``
	__Text__000001_Hello.StrokeDashArrayWhenSelected = ``
	__Text__000001_Hello.Transform = ``

	// Setup of pointers
	__Circle__000000_C1a.Animations = append(__Circle__000000_C1a.Animations, __Animate__000002_C1a_animation)
	__Layer__000000_SVG2.Rects = append(__Layer__000000_SVG2.Rects, __Rect__000002_R3)
	__Layer__000000_SVG2.Rects = append(__Layer__000000_SVG2.Rects, __Rect__000003_R4_rounded)
	__Layer__000000_SVG2.Rects = append(__Layer__000000_SVG2.Rects, __Rect__000004_Test_Rect)
	__Layer__000000_SVG2.Rects = append(__Layer__000000_SVG2.Rects, __Rect__000000_R1)
	__Layer__000000_SVG2.Rects = append(__Layer__000000_SVG2.Rects, __Rect__000001_R2)
	__Layer__000000_SVG2.Texts = append(__Layer__000000_SVG2.Texts, __Text__000001_Hello)
	__Layer__000000_SVG2.Texts = append(__Layer__000000_SVG2.Texts, __Text__000000_Bonjorno)
	__Layer__000000_SVG2.Circles = append(__Layer__000000_SVG2.Circles, __Circle__000000_C1a)
	__Layer__000000_SVG2.Lines = append(__Layer__000000_SVG2.Lines, __Line__000000_Line)
	__Layer__000000_SVG2.Ellipses = append(__Layer__000000_SVG2.Ellipses, __Ellipse__000000_Ellipse_Example_SVG)
	__Layer__000000_SVG2.Polylines = append(__Layer__000000_SVG2.Polylines, __Polyline__000000_Polyline_example_SVG)
	__Layer__000000_SVG2.Polygones = append(__Layer__000000_SVG2.Polygones, __Polygone__000000_Polygone_example_SVG)
	__Layer__000000_SVG2.Paths = append(__Layer__000000_SVG2.Paths, __Path__000000_Path_example)
	__Rect__000003_R4_rounded.Animations = append(__Rect__000003_R4_rounded.Animations, __Animate__000000_A1)
	__Rect__000003_R4_rounded.Animations = append(__Rect__000003_R4_rounded.Animations, __Animate__000001_Animate_oppacity)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_SVG2)
	__Text__000001_Hello.Animates = append(__Text__000001_Hello.Animates, __Animate__000003_Move_text)
}



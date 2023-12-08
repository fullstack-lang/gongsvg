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
	__Layer__000000_Default_Layer := (&models.Layer{Name: `Default Layer`}).Stage(stage)
	__Layer__000001_Test_Layer := (&models.Layer{Name: `Test Layer`}).Stage(stage)

	// Declarations of staged instances of Line
	__Line__000000_Line := (&models.Line{Name: `Line`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path
	__Path__000000_Path_example := (&models.Path{Name: `Path example`}).Stage(stage)
	__Path__000001_another_play_shape := (&models.Path{Name: `another play shape`}).Stage(stage)
	__Path__000002_celebration := (&models.Path{Name: `celebration`}).Stage(stage)
	__Path__000003_path := (&models.Path{Name: `path`}).Stage(stage)
	__Path__000004_play_shape := (&models.Path{Name: `play_shape`}).Stage(stage)

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

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

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
	__Layer__000000_Default_Layer.Display = true
	__Layer__000000_Default_Layer.Name = `Default Layer`

	// Layer values setup
	__Layer__000001_Test_Layer.Display = true
	__Layer__000001_Test_Layer.Name = `Test Layer`

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

	// Path values setup
	__Path__000001_another_play_shape.Name = `another play shape`
	__Path__000001_another_play_shape.Definition = `m576-540-56-56 104-104-104-104 56-56 104 104 104-104 56 56-104 104 104 104-56 56-104-104-104 104ZM80-520l200-360 200 360H80Zm200 400q-66 0-113-47t-47-113q0-67 47-113.5T280-440q66 0 113 47t47 113q0 66-47 113t-113 47Zm0-80q33 0 56.5-23.5T360-280q0-33-23.5-56.5T280-360q-33 0-56.5 23.5T200-280q0 33 23.5 56.5T280-200Zm-64-400h128l-64-115-64 115Zm304 480v-320h320v320H520Zm80-80h160v-160H600v160ZM280-658Zm0 378Zm400 0Z`
	__Path__000001_another_play_shape.Color = `black`
	__Path__000001_another_play_shape.FillOpacity = 0.500000
	__Path__000001_another_play_shape.Stroke = `black`
	__Path__000001_another_play_shape.StrokeWidth = 1.000000
	__Path__000001_another_play_shape.StrokeDashArray = ``
	__Path__000001_another_play_shape.StrokeDashArrayWhenSelected = ``
	__Path__000001_another_play_shape.Transform = `scale(0.5 0.5) translate(200 960)`

	// Path values setup
	__Path__000002_celebration.Name = `celebration`
	__Path__000002_celebration.Definition = `m80-80 200-560 360 360L80-80Zm132-132 282-100-182-182-100 282Zm370-246-42-42 224-224q32-32 77-32t77 32l24 24-42 42-24-24q-14-14-35-14t-35 14L582-458ZM422-618l-42-42 24-24q14-14 14-34t-14-34l-26-26 42-42 26 26q32 32 32 76t-32 76l-24 24Zm80 80-42-42 144-144q14-14 14-35t-14-35l-64-64 42-42 64 64q32 32 32 77t-32 77L502-538Zm160 160-42-42 64-64q32-32 77-32t77 32l64 64-42 42-64-64q-14-14-35-14t-35 14l-64 64ZM212-212Z`
	__Path__000002_celebration.Color = `black`
	__Path__000002_celebration.FillOpacity = 1.000000
	__Path__000002_celebration.Stroke = `black`
	__Path__000002_celebration.StrokeWidth = 2.000000
	__Path__000002_celebration.StrokeDashArray = ``
	__Path__000002_celebration.StrokeDashArrayWhenSelected = ``
	__Path__000002_celebration.Transform = ``

	// Path values setup
	__Path__000003_path.Name = `path`
	__Path__000003_path.Definition = `M 200 200 C 20 20, 40 70, 250 210`
	__Path__000003_path.Color = `transparant`
	__Path__000003_path.FillOpacity = 0.000000
	__Path__000003_path.Stroke = `blue`
	__Path__000003_path.StrokeWidth = 2.000000
	__Path__000003_path.StrokeDashArray = ``
	__Path__000003_path.StrokeDashArrayWhenSelected = ``
	__Path__000003_path.Transform = ``

	// Path values setup
	__Path__000004_play_shape.Name = `play_shape`
	__Path__000004_play_shape.Definition = `m 763,-197 -43,-43 q 47,-47 73.5,-108 26.5,-61 26.5,-132 0,-71 -26.5,-132 Q 767,-673 720,-720 l 43,-43 q 54,54 85.5,127 31.5,73 31.5,156 0,83 -31.5,156 -31.5,73 -85.5,127 z`
	__Path__000004_play_shape.Color = `blue`
	__Path__000004_play_shape.FillOpacity = 1.000000
	__Path__000004_play_shape.Stroke = `red`
	__Path__000004_play_shape.StrokeWidth = 1.000000
	__Path__000004_play_shape.StrokeDashArray = ``
	__Path__000004_play_shape.StrokeDashArrayWhenSelected = ``
	__Path__000004_play_shape.Transform = `scale(0.5 0.5) translate(200 960)`

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
	__Rect__000004_Test_Rect.X = 100.000000
	__Rect__000004_Test_Rect.Y = 400.000000
	__Rect__000004_Test_Rect.Width = 1200.000000
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
	__SVG__000000_SVG.IsEditable = true

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
	__Layer__000000_Default_Layer.Rects = append(__Layer__000000_Default_Layer.Rects, __Rect__000002_R3)
	__Layer__000000_Default_Layer.Rects = append(__Layer__000000_Default_Layer.Rects, __Rect__000003_R4_rounded)
	__Layer__000000_Default_Layer.Rects = append(__Layer__000000_Default_Layer.Rects, __Rect__000004_Test_Rect)
	__Layer__000000_Default_Layer.Rects = append(__Layer__000000_Default_Layer.Rects, __Rect__000000_R1)
	__Layer__000000_Default_Layer.Rects = append(__Layer__000000_Default_Layer.Rects, __Rect__000001_R2)
	__Layer__000000_Default_Layer.Texts = append(__Layer__000000_Default_Layer.Texts, __Text__000001_Hello)
	__Layer__000000_Default_Layer.Texts = append(__Layer__000000_Default_Layer.Texts, __Text__000000_Bonjorno)
	__Layer__000000_Default_Layer.Circles = append(__Layer__000000_Default_Layer.Circles, __Circle__000000_C1a)
	__Layer__000000_Default_Layer.Lines = append(__Layer__000000_Default_Layer.Lines, __Line__000000_Line)
	__Layer__000000_Default_Layer.Ellipses = append(__Layer__000000_Default_Layer.Ellipses, __Ellipse__000000_Ellipse_Example_SVG)
	__Layer__000001_Test_Layer.Paths = append(__Layer__000001_Test_Layer.Paths, __Path__000000_Path_example)
	__Layer__000001_Test_Layer.Paths = append(__Layer__000001_Test_Layer.Paths, __Path__000003_path)
	__Layer__000001_Test_Layer.Paths = append(__Layer__000001_Test_Layer.Paths, __Path__000002_celebration)
	__Layer__000001_Test_Layer.Paths = append(__Layer__000001_Test_Layer.Paths, __Path__000004_play_shape)
	__Layer__000001_Test_Layer.Paths = append(__Layer__000001_Test_Layer.Paths, __Path__000001_another_play_shape)
	__Rect__000003_R4_rounded.Animations = append(__Rect__000003_R4_rounded.Animations, __Animate__000000_A1)
	__Rect__000003_R4_rounded.Animations = append(__Rect__000003_R4_rounded.Animations, __Animate__000001_Animate_oppacity)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Default_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000001_Test_Layer)
	__Text__000001_Hello.Animates = append(__Text__000001_Hello.Animates, __Animate__000003_Move_text)
}



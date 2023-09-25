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
	__Animate__000000_ := (&models.Animate{Name: ``}).Stage(stage)

	// Declarations of staged instances of Circle
	__Circle__000000_Test := (&models.Circle{Name: `Test`}).Stage(stage)

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_Bottom_Rectangle_Layer := (&models.Layer{Name: `Bottom Rectangle Layer`}).Stage(stage)
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom := (&models.Layer{Name: `Layer RectLinkLink Medium to Top-Bottom`}).Stage(stage)
	__Layer__000002_Link_layer_vertical_to_horizontal := (&models.Layer{Name: `Link layer vertical to horizontal`}).Stage(stage)
	__Layer__000003_Middle_Rect_Layer := (&models.Layer{Name: `Middle Rect Layer`}).Stage(stage)
	__Layer__000004_Top_Rectangle_layer := (&models.Layer{Name: `Top Rectangle layer`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal := (&models.Link{Name: `Arrow - Top to Bottom vertical to horizontal`}).Stage(stage)

	// Declarations of staged instances of LinkAnchoredText
	__LinkAnchoredText__000000_Liine_1_Line_2 := (&models.LinkAnchoredText{Name: `Liine 1
Line 2`}).Stage(stage)
	__LinkAnchoredText__000001_Start_Anchored_1 := (&models.LinkAnchoredText{Name: `Start Anchored 1`}).Stage(stage)

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Bottom := (&models.Rect{Name: `Bottom`}).Stage(stage)
	__Rect__000001_Middle_Rect := (&models.Rect{Name: `Middle Rect`}).Stage(stage)
	__Rect__000002_Top := (&models.Rect{Name: `Top`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredRect
	__RectAnchoredRect__000000_Rect_within_top := (&models.RectAnchoredRect{Name: `Rect within top`}).Stage(stage)
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width := (&models.RectAnchoredRect{Name: `Top on Bottom with same width`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredText
	__RectAnchoredText__000000_Bottom_Text := (&models.RectAnchoredText{Name: `Bottom Text`}).Stage(stage)
	__RectAnchoredText__000001_Top_Left := (&models.RectAnchoredText{Name: `Top Left`}).Stage(stage)
	__RectAnchoredText__000002_Top_anchored_bottom_middle := (&models.RectAnchoredText{Name: `Top anchored bottom middle`}).Stage(stage)
	__RectAnchoredText__000003_Top_anchored_top_middle := (&models.RectAnchoredText{Name: `Top anchored top middle`}).Stage(stage)

	// Declarations of staged instances of RectLinkLink
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link := (&models.RectLinkLink{Name: `Test Middle to Top-Bottom Link`}).Stage(stage)

	// Declarations of staged instances of SVG
	__SVG__000000_SVG := (&models.SVG{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of Text
	__Text__000000_Essai := (&models.Text{Name: `Essai`}).Stage(stage)

	// Setup of values

	// Animate values setup
	__Animate__000000_.Name = ``
	__Animate__000000_.AttributeName = ``
	__Animate__000000_.Values = ``
	__Animate__000000_.Dur = ``
	__Animate__000000_.RepeatCount = ``

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
	__Circle__000000_Test.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Test.Transform = ``

	// Layer values setup
	__Layer__000000_Bottom_Rectangle_Layer.Display = false
	__Layer__000000_Bottom_Rectangle_Layer.Name = `Bottom Rectangle Layer`

	// Layer values setup
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Display = false
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Name = `Layer RectLinkLink Medium to Top-Bottom`

	// Layer values setup
	__Layer__000002_Link_layer_vertical_to_horizontal.Display = false
	__Layer__000002_Link_layer_vertical_to_horizontal.Name = `Link layer vertical to horizontal`

	// Layer values setup
	__Layer__000003_Middle_Rect_Layer.Display = false
	__Layer__000003_Middle_Rect_Layer.Name = `Middle Rect Layer`

	// Layer values setup
	__Layer__000004_Top_Rectangle_layer.Display = false
	__Layer__000004_Top_Rectangle_layer.Name = `Top Rectangle layer`

	// Link values setup
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Name = `Arrow - Top to Bottom vertical to horizontal`
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartAnchorType = models.ANCHOR_CENTER
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndAnchorType = models.ANCHOR_CENTER
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartRatio = 0.696262
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndRatio = 0.420455
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.CornerOffsetRatio = 1.735577
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.CornerRadius = 8.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.HasEndArrow = true
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndArrowSize = 10.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Color = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.FillOpacity = 0.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Stroke = `green`
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeWidth = 4.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeDashArray = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeDashArrayWhenSelected = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Transform = ``

	// LinkAnchoredText values setup
	__LinkAnchoredText__000000_Liine_1_Line_2.Name = `Liine 1
Line 2`
	__LinkAnchoredText__000000_Liine_1_Line_2.Content = `Liine 1
Line 2`
	__LinkAnchoredText__000000_Liine_1_Line_2.X_Offset = -65.000000
	__LinkAnchoredText__000000_Liine_1_Line_2.Y_Offset = -37.000000
	__LinkAnchoredText__000000_Liine_1_Line_2.FontWeight = `normal`
	__LinkAnchoredText__000000_Liine_1_Line_2.Color = `black`
	__LinkAnchoredText__000000_Liine_1_Line_2.FillOpacity = 100.000000
	__LinkAnchoredText__000000_Liine_1_Line_2.Stroke = `black`
	__LinkAnchoredText__000000_Liine_1_Line_2.StrokeWidth = 1.000000
	__LinkAnchoredText__000000_Liine_1_Line_2.StrokeDashArray = ``
	__LinkAnchoredText__000000_Liine_1_Line_2.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000000_Liine_1_Line_2.Transform = ``

	// LinkAnchoredText values setup
	__LinkAnchoredText__000001_Start_Anchored_1.Name = `Start Anchored 1`
	__LinkAnchoredText__000001_Start_Anchored_1.Content = `Start Anchored 1
Second line
Third Line`
	__LinkAnchoredText__000001_Start_Anchored_1.X_Offset = 5.999969
	__LinkAnchoredText__000001_Start_Anchored_1.Y_Offset = -72.987488
	__LinkAnchoredText__000001_Start_Anchored_1.FontWeight = `light`
	__LinkAnchoredText__000001_Start_Anchored_1.Color = `cyan`
	__LinkAnchoredText__000001_Start_Anchored_1.FillOpacity = 100.000000
	__LinkAnchoredText__000001_Start_Anchored_1.Stroke = `cyan`
	__LinkAnchoredText__000001_Start_Anchored_1.StrokeWidth = 1.000000
	__LinkAnchoredText__000001_Start_Anchored_1.StrokeDashArray = ``
	__LinkAnchoredText__000001_Start_Anchored_1.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000001_Start_Anchored_1.Transform = ``

	// Rect values setup
	__Rect__000000_Bottom.Name = `Bottom`
	__Rect__000000_Bottom.X = 549.000000
	__Rect__000000_Bottom.Y = 122.000000
	__Rect__000000_Bottom.Width = 632.000000
	__Rect__000000_Bottom.Height = 122.000000
	__Rect__000000_Bottom.RX = 5.000000
	__Rect__000000_Bottom.Color = `bisque`
	__Rect__000000_Bottom.FillOpacity = 50.000000
	__Rect__000000_Bottom.Stroke = `lightcoral`
	__Rect__000000_Bottom.StrokeWidth = 3.000000
	__Rect__000000_Bottom.StrokeDashArray = ``
	__Rect__000000_Bottom.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000000_Bottom.Transform = ``
	__Rect__000000_Bottom.IsSelectable = true
	__Rect__000000_Bottom.IsSelected = false
	__Rect__000000_Bottom.CanHaveLeftHandle = true
	__Rect__000000_Bottom.HasLeftHandle = false
	__Rect__000000_Bottom.CanHaveRightHandle = true
	__Rect__000000_Bottom.HasRightHandle = false
	__Rect__000000_Bottom.CanHaveTopHandle = true
	__Rect__000000_Bottom.HasTopHandle = false
	__Rect__000000_Bottom.CanHaveBottomHandle = true
	__Rect__000000_Bottom.HasBottomHandle = false
	__Rect__000000_Bottom.CanMoveHorizontaly = true
	__Rect__000000_Bottom.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Middle_Rect.Name = `Middle Rect`
	__Rect__000001_Middle_Rect.X = 486.000000
	__Rect__000001_Middle_Rect.Y = 387.000000
	__Rect__000001_Middle_Rect.Width = 253.000000
	__Rect__000001_Middle_Rect.Height = 132.000000
	__Rect__000001_Middle_Rect.RX = 3.000000
	__Rect__000001_Middle_Rect.Color = `lavender`
	__Rect__000001_Middle_Rect.FillOpacity = 50.000000
	__Rect__000001_Middle_Rect.Stroke = `turquoise`
	__Rect__000001_Middle_Rect.StrokeWidth = 1.000000
	__Rect__000001_Middle_Rect.StrokeDashArray = ``
	__Rect__000001_Middle_Rect.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000001_Middle_Rect.Transform = ``
	__Rect__000001_Middle_Rect.IsSelectable = true
	__Rect__000001_Middle_Rect.IsSelected = false
	__Rect__000001_Middle_Rect.CanHaveLeftHandle = true
	__Rect__000001_Middle_Rect.HasLeftHandle = false
	__Rect__000001_Middle_Rect.CanHaveRightHandle = true
	__Rect__000001_Middle_Rect.HasRightHandle = false
	__Rect__000001_Middle_Rect.CanHaveTopHandle = true
	__Rect__000001_Middle_Rect.HasTopHandle = false
	__Rect__000001_Middle_Rect.CanHaveBottomHandle = true
	__Rect__000001_Middle_Rect.HasBottomHandle = false
	__Rect__000001_Middle_Rect.CanMoveHorizontaly = true
	__Rect__000001_Middle_Rect.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000002_Top.Name = `Top`
	__Rect__000002_Top.X = 70.999969
	__Rect__000002_Top.Y = 83.000000
	__Rect__000002_Top.Width = 208.000000
	__Rect__000002_Top.Height = 214.000000
	__Rect__000002_Top.RX = 3.000000
	__Rect__000002_Top.Color = `lightcyan`
	__Rect__000002_Top.FillOpacity = 100.000000
	__Rect__000002_Top.Stroke = `darkcyan`
	__Rect__000002_Top.StrokeWidth = 2.000000
	__Rect__000002_Top.StrokeDashArray = ``
	__Rect__000002_Top.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000002_Top.Transform = ``
	__Rect__000002_Top.IsSelectable = true
	__Rect__000002_Top.IsSelected = false
	__Rect__000002_Top.CanHaveLeftHandle = true
	__Rect__000002_Top.HasLeftHandle = false
	__Rect__000002_Top.CanHaveRightHandle = true
	__Rect__000002_Top.HasRightHandle = false
	__Rect__000002_Top.CanHaveTopHandle = false
	__Rect__000002_Top.HasTopHandle = false
	__Rect__000002_Top.CanHaveBottomHandle = false
	__Rect__000002_Top.HasBottomHandle = false
	__Rect__000002_Top.CanMoveHorizontaly = true
	__Rect__000002_Top.CanMoveVerticaly = true

	// RectAnchoredRect values setup
	__RectAnchoredRect__000000_Rect_within_top.Name = `Rect within top`
	__RectAnchoredRect__000000_Rect_within_top.X = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Y = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Width = 100.000000
	__RectAnchoredRect__000000_Rect_within_top.Height = 30.000000
	__RectAnchoredRect__000000_Rect_within_top.RX = 3.000000
	__RectAnchoredRect__000000_Rect_within_top.X_Offset = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Y_Offset = 40.000000
	__RectAnchoredRect__000000_Rect_within_top.RectAnchorType = models.RECT_ANCHOR_TOP
	__RectAnchoredRect__000000_Rect_within_top.WidthFollowRect = false
	__RectAnchoredRect__000000_Rect_within_top.HeightFollowRect = false
	__RectAnchoredRect__000000_Rect_within_top.Color = `lightgrey`
	__RectAnchoredRect__000000_Rect_within_top.FillOpacity = 40.000000
	__RectAnchoredRect__000000_Rect_within_top.Stroke = `bisque`
	__RectAnchoredRect__000000_Rect_within_top.StrokeWidth = 1.000000
	__RectAnchoredRect__000000_Rect_within_top.StrokeDashArray = ``
	__RectAnchoredRect__000000_Rect_within_top.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000000_Rect_within_top.Transform = ``

	// RectAnchoredRect values setup
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Name = `Top on Bottom with same width`
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.X = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Y = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Width = 100.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Height = 50.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.RX = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.X_Offset = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Y_Offset = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.RectAnchorType = models.RECT_ANCHOR_TOP_LEFT
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.WidthFollowRect = true
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.HeightFollowRect = false
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Color = `lightblue`
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.FillOpacity = 100.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Stroke = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeWidth = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeDashArray = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Transform = ``

	// RectAnchoredText values setup
	__RectAnchoredText__000000_Bottom_Text.Name = `Bottom Text`
	__RectAnchoredText__000000_Bottom_Text.Content = `This is an example of a note that
could be displayed on a diagram.

It could explain one aspect of the model
for instance, describing relations between structs

The text of a UML note refers a comment with the GONGNOTE keyword which is
a special case of go Note convention. See example
for details in the go code of the models.
`
	__RectAnchoredText__000000_Bottom_Text.FontWeight = `normal`
	__RectAnchoredText__000000_Bottom_Text.FontSize = 0
	__RectAnchoredText__000000_Bottom_Text.X_Offset = 0.000000
	__RectAnchoredText__000000_Bottom_Text.Y_Offset = 20.000000
	__RectAnchoredText__000000_Bottom_Text.RectAnchorType = models.RECT_ANCHOR_TOP
	__RectAnchoredText__000000_Bottom_Text.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__000000_Bottom_Text.Color = `black`
	__RectAnchoredText__000000_Bottom_Text.FillOpacity = 100.000000
	__RectAnchoredText__000000_Bottom_Text.Stroke = `black`
	__RectAnchoredText__000000_Bottom_Text.StrokeWidth = 1.000000
	__RectAnchoredText__000000_Bottom_Text.StrokeDashArray = ``
	__RectAnchoredText__000000_Bottom_Text.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000000_Bottom_Text.Transform = ``

	// RectAnchoredText values setup
	__RectAnchoredText__000001_Top_Left.Name = `Top Left`
	__RectAnchoredText__000001_Top_Left.Content = `Top Left`
	__RectAnchoredText__000001_Top_Left.FontWeight = ``
	__RectAnchoredText__000001_Top_Left.FontSize = 0
	__RectAnchoredText__000001_Top_Left.X_Offset = 0.000000
	__RectAnchoredText__000001_Top_Left.Y_Offset = 0.000000
	__RectAnchoredText__000001_Top_Left.RectAnchorType = models.RECT_ANCHOR_TOP_LEFT
	__RectAnchoredText__000001_Top_Left.Color = ``
	__RectAnchoredText__000001_Top_Left.FillOpacity = 0.000000
	__RectAnchoredText__000001_Top_Left.Stroke = `black`
	__RectAnchoredText__000001_Top_Left.StrokeWidth = 0.000000
	__RectAnchoredText__000001_Top_Left.StrokeDashArray = ``
	__RectAnchoredText__000001_Top_Left.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000001_Top_Left.Transform = ``

	// RectAnchoredText values setup
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Name = `Top anchored bottom middle`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Content = `Top anchored bottom middle`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FontWeight = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FontSize = 0
	__RectAnchoredText__000002_Top_anchored_bottom_middle.X_Offset = 0.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Y_Offset = 20.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.RectAnchorType = models.RECT_ANCHOR_BOTTOM
	__RectAnchoredText__000002_Top_anchored_bottom_middle.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Color = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FillOpacity = 100.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Stroke = `blue`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeWidth = 2.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeDashArray = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Transform = ``

	// RectAnchoredText values setup
	__RectAnchoredText__000003_Top_anchored_top_middle.Name = `Top anchored top middle`
	__RectAnchoredText__000003_Top_anchored_top_middle.Content = `Top anchored
top middle
line 3`
	__RectAnchoredText__000003_Top_anchored_top_middle.FontWeight = `bold`
	__RectAnchoredText__000003_Top_anchored_top_middle.FontSize = 0
	__RectAnchoredText__000003_Top_anchored_top_middle.X_Offset = 0.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.Y_Offset = 20.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.RectAnchorType = models.RECT_ANCHOR_TOP
	__RectAnchoredText__000003_Top_anchored_top_middle.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__000003_Top_anchored_top_middle.Color = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.FillOpacity = 100.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.Stroke = `black`
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeWidth = 1.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeDashArray = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.Transform = ``

	// RectLinkLink values setup
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Name = `Test Middle to Top-Bottom Link`
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.TargetAnchorPosition = 0.600000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Color = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.FillOpacity = 0.000000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Stroke = `lightgreen`
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeWidth = 4.000000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeDashArray = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeDashArrayWhenSelected = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Transform = ``

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.DrawingState = models.NOT_DRAWING_LINE
	__SVG__000000_SVG.IsEditable = true

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
	__Text__000000_Essai.StrokeDashArrayWhenSelected = ``
	__Text__000000_Essai.Transform = ``

	// Setup of pointers
	__Layer__000000_Bottom_Rectangle_Layer.Rects = append(__Layer__000000_Bottom_Rectangle_Layer.Rects, __Rect__000000_Bottom)
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.RectLinkLinks = append(__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.RectLinkLinks, __RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link)
	__Layer__000002_Link_layer_vertical_to_horizontal.Links = append(__Layer__000002_Link_layer_vertical_to_horizontal.Links, __Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal)
	__Layer__000003_Middle_Rect_Layer.Rects = append(__Layer__000003_Middle_Rect_Layer.Rects, __Rect__000001_Middle_Rect)
	__Layer__000004_Top_Rectangle_layer.Rects = append(__Layer__000004_Top_Rectangle_layer.Rects, __Rect__000002_Top)
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Start = __Rect__000002_Top
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.End = __Rect__000000_Bottom
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd = append(__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd, __LinkAnchoredText__000000_Liine_1_Line_2)
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowStart = append(__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowStart, __LinkAnchoredText__000001_Start_Anchored_1)
	__Rect__000000_Bottom.RectAnchoredTexts = append(__Rect__000000_Bottom.RectAnchoredTexts, __RectAnchoredText__000000_Bottom_Text)
	__Rect__000000_Bottom.RectAnchoredRects = append(__Rect__000000_Bottom.RectAnchoredRects, __RectAnchoredRect__000001_Top_on_Bottom_with_same_width)
	__Rect__000002_Top.RectAnchoredTexts = append(__Rect__000002_Top.RectAnchoredTexts, __RectAnchoredText__000003_Top_anchored_top_middle)
	__Rect__000002_Top.RectAnchoredTexts = append(__Rect__000002_Top.RectAnchoredTexts, __RectAnchoredText__000001_Top_Left)
	__Rect__000002_Top.RectAnchoredTexts = append(__Rect__000002_Top.RectAnchoredTexts, __RectAnchoredText__000002_Top_anchored_bottom_middle)
	__Rect__000002_Top.RectAnchoredRects = append(__Rect__000002_Top.RectAnchoredRects, __RectAnchoredRect__000000_Rect_within_top)
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Start = __Rect__000001_Middle_Rect
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.End = __Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Bottom_Rectangle_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000004_Top_Rectangle_layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000003_Middle_Rect_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000002_Link_layer_vertical_to_horizontal)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom)
	__SVG__000000_SVG.StartRect = __Rect__000000_Bottom
	__SVG__000000_SVG.EndRect = __Rect__000002_Top
}



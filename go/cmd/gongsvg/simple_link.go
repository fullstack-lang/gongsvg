package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_simple_link models.StageStruct
var ___dummy__Time_simple_link time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_simple_link map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["simple_link"] = simple_linkInjection
// }

// simple_linkInjection will stage objects of database "simple_link"
func simple_linkInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of AnchoredText

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_Links := (&models.Layer{Name: `Links`}).Stage(stage)
	__Layer__000001_Rects := (&models.Layer{Name: `Rects`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link
	__Link__000000_Start_to_End := (&models.Link{Name: `Start to End`}).Stage(stage)

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_End := (&models.Rect{Name: `End`}).Stage(stage)
	__Rect__000001_Start := (&models.Rect{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_test := (&models.SVG{Name: `test`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_Links.Display = false
	__Layer__000000_Links.Name = `Links`

	// Layer values setup
	__Layer__000001_Rects.Display = false
	__Layer__000001_Rects.Name = `Rects`

	// Link values setup
	__Link__000000_Start_to_End.Name = `Start to End`
	__Link__000000_Start_to_End.Type = models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	__Link__000000_Start_to_End.StartAnchorType = models.ANCHOR_CENTER
	__Link__000000_Start_to_End.EndAnchorType = models.ANCHOR_CENTER
	__Link__000000_Start_to_End.StartRatio = 0.000000
	__Link__000000_Start_to_End.EndRatio = 0.000000
	__Link__000000_Start_to_End.CornerOffsetRatio = 0.000000
	__Link__000000_Start_to_End.CornerRadius = 0.000000
	__Link__000000_Start_to_End.HasEndArrow = true
	__Link__000000_Start_to_End.EndArrowSize = 0.000000
	__Link__000000_Start_to_End.Color = ``
	__Link__000000_Start_to_End.FillOpacity = 0.000000
	__Link__000000_Start_to_End.Stroke = `lightblue`
	__Link__000000_Start_to_End.StrokeWidth = 3.000000
	__Link__000000_Start_to_End.StrokeDashArray = ``
	__Link__000000_Start_to_End.StrokeDashArrayWhenSelected = `5 5`
	__Link__000000_Start_to_End.Transform = ``

	// Rect values setup
	__Rect__000000_End.Name = `End`
	__Rect__000000_End.X = 506.000000
	__Rect__000000_End.Y = 159.000000
	__Rect__000000_End.Width = 248.000000
	__Rect__000000_End.Height = 167.000000
	__Rect__000000_End.RX = 0.000000
	__Rect__000000_End.Color = `lightcyan`
	__Rect__000000_End.FillOpacity = 50.000000
	__Rect__000000_End.Stroke = ``
	__Rect__000000_End.StrokeWidth = 0.000000
	__Rect__000000_End.StrokeDashArray = ``
	__Rect__000000_End.StrokeDashArrayWhenSelected = ``
	__Rect__000000_End.Transform = ``
	__Rect__000000_End.IsSelectable = true
	__Rect__000000_End.IsSelected = true
	__Rect__000000_End.CanHaveLeftHandle = true
	__Rect__000000_End.HasLeftHandle = true
	__Rect__000000_End.CanHaveRightHandle = true
	__Rect__000000_End.HasRightHandle = true
	__Rect__000000_End.CanHaveTopHandle = true
	__Rect__000000_End.HasTopHandle = true
	__Rect__000000_End.CanHaveBottomHandle = true
	__Rect__000000_End.HasBottomHandle = true
	__Rect__000000_End.CanMoveHorizontaly = true
	__Rect__000000_End.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Start.Name = `Start`
	__Rect__000001_Start.X = 172.000000
	__Rect__000001_Start.Y = 116.000000
	__Rect__000001_Start.Width = 169.000000
	__Rect__000001_Start.Height = 200.000000
	__Rect__000001_Start.RX = 0.000000
	__Rect__000001_Start.Color = `bisque`
	__Rect__000001_Start.FillOpacity = 50.000000
	__Rect__000001_Start.Stroke = ``
	__Rect__000001_Start.StrokeWidth = 0.000000
	__Rect__000001_Start.StrokeDashArray = ``
	__Rect__000001_Start.StrokeDashArrayWhenSelected = ``
	__Rect__000001_Start.Transform = ``
	__Rect__000001_Start.IsSelectable = true
	__Rect__000001_Start.IsSelected = true
	__Rect__000001_Start.CanHaveLeftHandle = true
	__Rect__000001_Start.HasLeftHandle = true
	__Rect__000001_Start.CanHaveRightHandle = true
	__Rect__000001_Start.HasRightHandle = true
	__Rect__000001_Start.CanHaveTopHandle = true
	__Rect__000001_Start.HasTopHandle = true
	__Rect__000001_Start.CanHaveBottomHandle = true
	__Rect__000001_Start.HasBottomHandle = true
	__Rect__000001_Start.CanMoveHorizontaly = true
	__Rect__000001_Start.CanMoveVerticaly = true

	// SVG values setup
	__SVG__000000_test.Name = `test`
	__SVG__000000_test.IsEditable = false

	// Setup of pointers
	__Layer__000000_Links.Links = append(__Layer__000000_Links.Links, __Link__000000_Start_to_End)
	__Layer__000001_Rects.Rects = append(__Layer__000001_Rects.Rects, __Rect__000001_Start)
	__Layer__000001_Rects.Rects = append(__Layer__000001_Rects.Rects, __Rect__000000_End)
	__Link__000000_Start_to_End.Start = __Rect__000001_Start
	__Link__000000_Start_to_End.End = __Rect__000000_End
	__SVG__000000_test.Layers = append(__SVG__000000_test.Layers, __Layer__000001_Rects)
	__SVG__000000_test.Layers = append(__SVG__000000_test.Layers, __Layer__000000_Links)
}


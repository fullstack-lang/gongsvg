package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_new models.StageStruct
var ___dummy__Time_new time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_new map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["new"] = newInjection
// }

// newInjection will stage objects of database "new"
func newInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_Paths := (&models.Layer{Name: `Paths`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path
	__Path__000000_Tree_lines := (&models.Path{Name: `Tree lines`}).Stage(stage)

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Center_100_100 := (&models.Rect{Name: `Center 100 100`}).Stage(stage)
	__Rect__000001_Center_300_300 := (&models.Rect{Name: `Center 300 300`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredPath
	__RectAnchoredPath__000000_Ref_anchored_path := (&models.RectAnchoredPath{Name: `Ref anchored path`}).Stage(stage)
	__RectAnchoredPath__000001_new_rect_anchored_path := (&models.RectAnchoredPath{Name: `new rect anchored path`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_SVG := (&models.SVG{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_Paths.Display = true
	__Layer__000000_Paths.Name = `Paths`

	// Path values setup
	__Path__000000_Tree_lines.Name = `Tree lines`
	__Path__000000_Tree_lines.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__Path__000000_Tree_lines.Color = `grey`
	__Path__000000_Tree_lines.FillOpacity = 0.500000
	__Path__000000_Tree_lines.Stroke = `black`
	__Path__000000_Tree_lines.StrokeWidth = 1.000000
	__Path__000000_Tree_lines.StrokeDashArray = ``
	__Path__000000_Tree_lines.StrokeDashArrayWhenSelected = ``
	__Path__000000_Tree_lines.Transform = `translate(0 960)`

	// Rect values setup
	__Rect__000000_Center_100_100.Name = `Center 100 100`
	__Rect__000000_Center_100_100.X = 146.000000
	__Rect__000000_Center_100_100.Y = 138.000000
	__Rect__000000_Center_100_100.Width = 254.000000
	__Rect__000000_Center_100_100.Height = 254.000000
	__Rect__000000_Center_100_100.RX = 4.000000
	__Rect__000000_Center_100_100.Color = ``
	__Rect__000000_Center_100_100.FillOpacity = 0.000000
	__Rect__000000_Center_100_100.Stroke = `blue`
	__Rect__000000_Center_100_100.StrokeWidth = 2.000000
	__Rect__000000_Center_100_100.StrokeDashArray = ``
	__Rect__000000_Center_100_100.StrokeDashArrayWhenSelected = ``
	__Rect__000000_Center_100_100.Transform = ``
	__Rect__000000_Center_100_100.IsSelectable = true
	__Rect__000000_Center_100_100.IsSelected = false
	__Rect__000000_Center_100_100.CanHaveLeftHandle = true
	__Rect__000000_Center_100_100.HasLeftHandle = false
	__Rect__000000_Center_100_100.CanHaveRightHandle = true
	__Rect__000000_Center_100_100.HasRightHandle = false
	__Rect__000000_Center_100_100.CanHaveTopHandle = true
	__Rect__000000_Center_100_100.HasTopHandle = false
	__Rect__000000_Center_100_100.IsScalingProportionally = true
	__Rect__000000_Center_100_100.CanHaveBottomHandle = true
	__Rect__000000_Center_100_100.HasBottomHandle = false
	__Rect__000000_Center_100_100.CanMoveHorizontaly = true
	__Rect__000000_Center_100_100.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Center_300_300.Name = `Center 300 300`
	__Rect__000001_Center_300_300.X = 341.000000
	__Rect__000001_Center_300_300.Y = 492.000000
	__Rect__000001_Center_300_300.Width = 203.000000
	__Rect__000001_Center_300_300.Height = 203.000000
	__Rect__000001_Center_300_300.RX = 5.000000
	__Rect__000001_Center_300_300.Color = ``
	__Rect__000001_Center_300_300.FillOpacity = 0.000000
	__Rect__000001_Center_300_300.Stroke = `blue`
	__Rect__000001_Center_300_300.StrokeWidth = 1.000000
	__Rect__000001_Center_300_300.StrokeDashArray = ``
	__Rect__000001_Center_300_300.StrokeDashArrayWhenSelected = ``
	__Rect__000001_Center_300_300.Transform = ``
	__Rect__000001_Center_300_300.IsSelectable = true
	__Rect__000001_Center_300_300.IsSelected = false
	__Rect__000001_Center_300_300.CanHaveLeftHandle = false
	__Rect__000001_Center_300_300.HasLeftHandle = false
	__Rect__000001_Center_300_300.CanHaveRightHandle = true
	__Rect__000001_Center_300_300.HasRightHandle = false
	__Rect__000001_Center_300_300.CanHaveTopHandle = false
	__Rect__000001_Center_300_300.HasTopHandle = false
	__Rect__000001_Center_300_300.IsScalingProportionally = true
	__Rect__000001_Center_300_300.CanHaveBottomHandle = false
	__Rect__000001_Center_300_300.HasBottomHandle = false
	__Rect__000001_Center_300_300.CanMoveHorizontaly = true
	__Rect__000001_Center_300_300.CanMoveVerticaly = true

	// RectAnchoredPath values setup
	__RectAnchoredPath__000000_Ref_anchored_path.Name = `Ref anchored path`
	__RectAnchoredPath__000000_Ref_anchored_path.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000000_Ref_anchored_path.X_Offset = 20.000000
	__RectAnchoredPath__000000_Ref_anchored_path.Y_Offset = 0.000000
	__RectAnchoredPath__000000_Ref_anchored_path.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__000000_Ref_anchored_path.ScalePropotionnally = true
	__RectAnchoredPath__000000_Ref_anchored_path.AppliedScaling = 0.264584
	__RectAnchoredPath__000000_Ref_anchored_path.Color = ``
	__RectAnchoredPath__000000_Ref_anchored_path.FillOpacity = 0.000000
	__RectAnchoredPath__000000_Ref_anchored_path.Stroke = `lightblue`
	__RectAnchoredPath__000000_Ref_anchored_path.StrokeWidth = 5.000000
	__RectAnchoredPath__000000_Ref_anchored_path.StrokeDashArray = ``
	__RectAnchoredPath__000000_Ref_anchored_path.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000000_Ref_anchored_path.Transform = `translate(50 50)`

	// RectAnchoredPath values setup
	__RectAnchoredPath__000001_new_rect_anchored_path.Name = `new rect anchored path`
	__RectAnchoredPath__000001_new_rect_anchored_path.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000001_new_rect_anchored_path.X_Offset = 0.000000
	__RectAnchoredPath__000001_new_rect_anchored_path.Y_Offset = 0.000000
	__RectAnchoredPath__000001_new_rect_anchored_path.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__000001_new_rect_anchored_path.ScalePropotionnally = true
	__RectAnchoredPath__000001_new_rect_anchored_path.AppliedScaling = 0.177403
	__RectAnchoredPath__000001_new_rect_anchored_path.Color = ``
	__RectAnchoredPath__000001_new_rect_anchored_path.FillOpacity = 0.000000
	__RectAnchoredPath__000001_new_rect_anchored_path.Stroke = `red`
	__RectAnchoredPath__000001_new_rect_anchored_path.StrokeWidth = 8.000000
	__RectAnchoredPath__000001_new_rect_anchored_path.StrokeDashArray = ``
	__RectAnchoredPath__000001_new_rect_anchored_path.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000001_new_rect_anchored_path.Transform = `scale(0.9 1)`

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.IsEditable = true

	// Setup of pointers
	__Layer__000000_Paths.Rects = append(__Layer__000000_Paths.Rects, __Rect__000000_Center_100_100)
	__Layer__000000_Paths.Rects = append(__Layer__000000_Paths.Rects, __Rect__000001_Center_300_300)
	__Layer__000000_Paths.Paths = append(__Layer__000000_Paths.Paths, __Path__000000_Tree_lines)
	__Rect__000000_Center_100_100.RectAnchoredPaths = append(__Rect__000000_Center_100_100.RectAnchoredPaths, __RectAnchoredPath__000000_Ref_anchored_path)
	__Rect__000001_Center_300_300.RectAnchoredPaths = append(__Rect__000001_Center_300_300.RectAnchoredPaths, __RectAnchoredPath__000001_new_rect_anchored_path)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Paths)
}



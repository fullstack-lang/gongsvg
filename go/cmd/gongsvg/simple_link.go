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

	__Layer__000000_Links := (&models.Layer{Name: `Links`}).Stage(stage)
	__Layer__000001_Rects := (&models.Layer{Name: `Rects`}).Stage(stage)

	__Link__000000_L1 := (&models.Link{Name: `L1`}).Stage(stage)

	__Rect__000000_End := (&models.Rect{Name: `End`}).Stage(stage)
	__Rect__000001_Start := (&models.Rect{Name: `Start`}).Stage(stage)

	__SVG__000000_test := (&models.SVG{Name: `test`}).Stage(stage)

	// Setup of values

	__Layer__000000_Links.Display = false
	__Layer__000000_Links.Name = `Links`

	__Layer__000001_Rects.Display = false
	__Layer__000001_Rects.Name = `Rects`

	__Link__000000_L1.Name = `L1`
	__Link__000000_L1.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_L1.IsBezierCurve = true
	__Link__000000_L1.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_L1.StartRatio = 0.650833
	__Link__000000_L1.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_L1.EndRatio = 0.617764
	__Link__000000_L1.CornerOffsetRatio = -0.763314
	__Link__000000_L1.CornerRadius = 0.000000
	__Link__000000_L1.HasEndArrow = false
	__Link__000000_L1.EndArrowSize = 0.000000
	__Link__000000_L1.HasStartArrow = true
	__Link__000000_L1.StartArrowSize = 9.000000
	__Link__000000_L1.Color = ``
	__Link__000000_L1.FillOpacity = 0.000000
	__Link__000000_L1.Stroke = `blue`
	__Link__000000_L1.StrokeOpacity = 1.000000
	__Link__000000_L1.StrokeWidth = 5.000000
	__Link__000000_L1.StrokeDashArray = ``
	__Link__000000_L1.StrokeDashArrayWhenSelected = ``
	__Link__000000_L1.Transform = ``

	__Rect__000000_End.Name = `End`
	__Rect__000000_End.X = 79.000000
	__Rect__000000_End.Y = 513.999969
	__Rect__000000_End.Width = 248.000000
	__Rect__000000_End.Height = 167.000000
	__Rect__000000_End.RX = 0.000000
	__Rect__000000_End.Color = `lightcyan`
	__Rect__000000_End.FillOpacity = 50.000000
	__Rect__000000_End.Stroke = ``
	__Rect__000000_End.StrokeOpacity = 0.000000
	__Rect__000000_End.StrokeWidth = 0.000000
	__Rect__000000_End.StrokeDashArray = ``
	__Rect__000000_End.StrokeDashArrayWhenSelected = ``
	__Rect__000000_End.Transform = ``
	__Rect__000000_End.IsSelectable = true
	__Rect__000000_End.IsSelected = false
	__Rect__000000_End.CanHaveLeftHandle = true
	__Rect__000000_End.HasLeftHandle = false
	__Rect__000000_End.CanHaveRightHandle = true
	__Rect__000000_End.HasRightHandle = false
	__Rect__000000_End.CanHaveTopHandle = true
	__Rect__000000_End.HasTopHandle = false
	__Rect__000000_End.IsScalingProportionally = false
	__Rect__000000_End.CanHaveBottomHandle = true
	__Rect__000000_End.HasBottomHandle = false
	__Rect__000000_End.CanMoveHorizontaly = true
	__Rect__000000_End.CanMoveVerticaly = true

	__Rect__000001_Start.Name = `Start`
	__Rect__000001_Start.X = 565.000000
	__Rect__000001_Start.Y = 189.000000
	__Rect__000001_Start.Width = 169.000000
	__Rect__000001_Start.Height = 200.000000
	__Rect__000001_Start.RX = 0.000000
	__Rect__000001_Start.Color = `bisque`
	__Rect__000001_Start.FillOpacity = 50.000000
	__Rect__000001_Start.Stroke = ``
	__Rect__000001_Start.StrokeOpacity = 0.000000
	__Rect__000001_Start.StrokeWidth = 0.000000
	__Rect__000001_Start.StrokeDashArray = ``
	__Rect__000001_Start.StrokeDashArrayWhenSelected = ``
	__Rect__000001_Start.Transform = ``
	__Rect__000001_Start.IsSelectable = true
	__Rect__000001_Start.IsSelected = false
	__Rect__000001_Start.CanHaveLeftHandle = true
	__Rect__000001_Start.HasLeftHandle = false
	__Rect__000001_Start.CanHaveRightHandle = true
	__Rect__000001_Start.HasRightHandle = false
	__Rect__000001_Start.CanHaveTopHandle = true
	__Rect__000001_Start.HasTopHandle = false
	__Rect__000001_Start.IsScalingProportionally = false
	__Rect__000001_Start.CanHaveBottomHandle = true
	__Rect__000001_Start.HasBottomHandle = false
	__Rect__000001_Start.CanMoveHorizontaly = true
	__Rect__000001_Start.CanMoveVerticaly = true

	__SVG__000000_test.Name = `test`
	__SVG__000000_test.DrawingState = models.NOT_DRAWING_LINK
	__SVG__000000_test.IsEditable = true

	// Setup of pointers
	__Layer__000000_Links.Links = append(__Layer__000000_Links.Links, __Link__000000_L1)
	__Layer__000001_Rects.Rects = append(__Layer__000001_Rects.Rects, __Rect__000001_Start)
	__Layer__000001_Rects.Rects = append(__Layer__000001_Rects.Rects, __Rect__000000_End)
	__Link__000000_L1.Start = __Rect__000001_Start
	__Link__000000_L1.End = __Rect__000000_End
	__SVG__000000_test.Layers = append(__SVG__000000_test.Layers, __Layer__000000_Links)
	__SVG__000000_test.Layers = append(__SVG__000000_test.Layers, __Layer__000001_Rects)
	__SVG__000000_test.StartRect = __Rect__000001_Start
	__SVG__000000_test.EndRect = __Rect__000000_End
}



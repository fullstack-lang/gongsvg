package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_issue_triangle models.StageStruct
var ___dummy__Time_issue_triangle time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_issue_triangle map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["issue_triangle"] = issue_triangleInjection
// }

// issue_triangleInjection will stage objects of database "issue_triangle"
func issue_triangleInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_ := (&models.Layer{Name: ``}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline
	__Polyline__000000_ := (&models.Polyline{Name: ``}).Stage(stage)

	// Declarations of staged instances of Rect

	// Declarations of staged instances of RectAnchoredPath

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_ := (&models.SVG{Name: ``}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_.Display = true
	__Layer__000000_.Name = ``

	// Polyline values setup
	__Polyline__000000_.Name = ``
	__Polyline__000000_.Points = `100 100 400 100`
	__Polyline__000000_.Color = ``
	__Polyline__000000_.FillOpacity = 0.000000
	__Polyline__000000_.Stroke = `red`
	__Polyline__000000_.StrokeOpacity = 0.500000
	__Polyline__000000_.StrokeWidth = 20.000000
	__Polyline__000000_.StrokeDashArray = ``
	__Polyline__000000_.StrokeDashArrayWhenSelected = ``
	__Polyline__000000_.Transform = ``

	// SVG values setup
	__SVG__000000_.Name = ``
	__SVG__000000_.IsEditable = false

	// Setup of pointers
	__Layer__000000_.Polylines = append(__Layer__000000_.Polylines, __Polyline__000000_)
	__SVG__000000_.Layers = append(__SVG__000000_.Layers, __Layer__000000_)
}



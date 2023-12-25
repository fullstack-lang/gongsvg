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

	// Declarations of staged instances of RectAnchoredPath

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

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.IsEditable = true

	// Setup of pointers
	__Layer__000000_Paths.Paths = append(__Layer__000000_Paths.Paths, __Path__000000_Tree_lines)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Paths)
}



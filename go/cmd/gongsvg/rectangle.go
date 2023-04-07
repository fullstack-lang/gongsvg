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

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Line

	// Declarations of staged instances of Path

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Sample := (&models.Rect{Name: `Sample`}).Stage(stage)

	// Declarations of staged instances of SVG
	__SVG__000000_Interactive_rectangle := (&models.SVG{Name: `Interactive rectangle`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Rect values setup
	__Rect__000000_Sample.Name = `Sample`
	__Rect__000000_Sample.X = 200.000000
	__Rect__000000_Sample.Y = 200.000000
	__Rect__000000_Sample.Width = 500.000000
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
	__SVG__000000_Interactive_rectangle.Display = true
	__SVG__000000_Interactive_rectangle.Name = `Interactive rectangle`

	// Setup of pointers
	__SVG__000000_Interactive_rectangle.Rects = append(__SVG__000000_Interactive_rectangle.Rects, __Rect__000000_Sample)
}



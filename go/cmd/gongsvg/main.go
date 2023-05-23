package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	gongsvg_go "github.com/fullstack-lang/gongsvg/go"
	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_static "github.com/fullstack-lang/gongsvg/go/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup  = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	unmarshall         = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	benchmark = flag.Int("benchmark", -1, "benchmark")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongsvg_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gongsvg/go/models", "main")
}

func main() {

	log.SetPrefix("gongsvg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongsvg_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *gongsvg_models.StageStruct
	if *marshallOnCommit != "" {
		// persistence in a SQLite file on disk in memory
		stage = gongsvg_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsvg/go/models")
	} else {
		// persistence in a SQLite file on disk
		stage = gongsvg_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsvg/go/models", "./test.db")
	}

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Checkout()
		stage.Marshall(file, "github.com/fullstack-lang/gongsvg/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}
		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	if *unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := gongsvg_models.ParseAstFile(stage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		stage.OnInitCommitFromFrontCallback = hook
	}

	if *benchmark != -1 {
		heightBetweenRect := 50.0

		stage.Reset()
		stage.Commit()

		svg := new(gongsvg_models.SVG).Stage(stage)
		svg.Name = "Bench"
		svg.IsEditable = true

		for i := 0; i < *benchmark; i++ {
			layer := new(gongsvg_models.Layer).Stage(stage)
			svg.Layers = append(svg.Layers, layer)
			layer.Name = fmt.Sprintf("%d", i)

			startRect := new(gongsvg_models.Rect).Stage(stage)
			layer.Rects = append(layer.Rects, startRect)
			startRect.X = 50
			startRect.Y = float64(i+1) * heightBetweenRect

			startRect.Width = 50
			startRect.Height = 50 / 1.618

			startRect.Stroke = gongsvg_models.Lightsalmon.ToString()
			startRect.StrokeWidth = 1
			startRect.StrokeDashArrayWhenSelected = "5 5"

			startRect.FillOpacity = 100
			startRect.Color = gongsvg_models.Lightsalmon.ToString()

			// moveability
			startRect.CanMoveHorizontaly = true
			startRect.CanMoveVerticaly = true

			// resizing
			startRect.IsSelectable = true
			startRect.CanHaveBottomHandle = true
			startRect.CanHaveLeftHandle = true
			startRect.CanHaveTopHandle = true
			startRect.CanHaveRightHandle = true

			title := new(gongsvg_models.RectAnchoredText).Stage(stage)
			title.Name = fmt.Sprintf("%d", i)
			title.Content = title.Name
			title.X_Offset = 0
			title.Y_Offset = 20
			title.RectAnchorType = gongsvg_models.RECT_ANCHOR_TOP
			title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
			title.FontWeight = "bold"
			title.Color = gongsvg_models.Black.ToString()
			title.FillOpacity = 1.0
			startRect.RectAnchoredTexts = append(startRect.RectAnchoredTexts, title)

			endRect := new(gongsvg_models.Rect).Stage(stage)
			layer.Rects = append(layer.Rects, endRect)
			endRect.X = 50 + 400
			endRect.Y = float64(i+1) * heightBetweenRect

			endRect.Width = 50
			endRect.Height = 50 / 1.618

			endRect.Stroke = gongsvg_models.Lightsalmon.ToString()
			endRect.StrokeWidth = 1
			endRect.StrokeDashArrayWhenSelected = "5 5"

			endRect.FillOpacity = 100
			endRect.Color = gongsvg_models.Lightsalmon.ToString()

			// moveability
			endRect.CanMoveHorizontaly = true
			endRect.CanMoveVerticaly = true

			// resizing
			endRect.IsSelectable = true
			endRect.CanHaveBottomHandle = true
			endRect.CanHaveLeftHandle = true
			endRect.CanHaveTopHandle = true
			endRect.CanHaveRightHandle = true

			link := new(gongsvg_models.Link).Stage(stage)
			link.Name = startRect.Name + " - to - " + endRect.Name

			linkLayer := new(gongsvg_models.Layer).Stage(stage)

			linkLayer.Links = append(linkLayer.Links, link)
			svg.Layers = append(svg.Layers, linkLayer)

			// configuration
			link.Stroke = gongsvg_models.Slategray.ToString()
			link.StrokeWidth = 3
			link.HasEndArrow = true
			link.EndArrowSize = 8

			link.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL
			link.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			link.StartRatio = 0.5
			link.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			link.EndRatio = 0.5
			link.CornerOffsetRatio = 1.4
			link.CornerRadius = 3

			link.Start = startRect
			link.End = endRect

			// add text to the arrow
			textAtArrowEnd := new(gongsvg_models.AnchoredText).Stage(stage)
			link.TextAtArrowEnd = append(link.TextAtArrowEnd, textAtArrowEnd)
			link.HasEndArrow = true

			textAtArrowEnd.Name = "A"
			textAtArrowEnd.Content = textAtArrowEnd.Name
			textAtArrowEnd.X_Offset = -50
			textAtArrowEnd.Y_Offset = 16
			textAtArrowEnd.Stroke = gongsvg_models.Black.ToString()
			textAtArrowEnd.StrokeWidth = 1
			textAtArrowEnd.Color = gongsvg_models.Black.ToString()
			textAtArrowEnd.FillOpacity = 100
			textAtArrowEnd.FontWeight = "normal"
		}
		stage.Commit()
	}

	gongdoc_load.Load(
		"gongsvg",
		"github.com/fullstack-lang/gongsvg/go/models",
		gongsvg_go.GoModelsDir,
		gongsvg_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

	// connects the rects to their behaviors
	gongsvg_models.DoNotKnowWhatThisFunctionIsFor(stage)

	stage.OnAfterRectUpdateCallback = new(gongsvg_models.RectOrchestrator)
	stage.OnAfterLineUpdateCallback = new(gongsvg_models.LineOrchestrator)
	stage.OnAfterSVGUpdateCallback = new(gongsvg_models.SVGOrchestrator)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

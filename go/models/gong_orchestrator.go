package models

// experimnetal, gong will generate an entry for a gongstruct only if
// it has one "OnAfterUpdate" function

// RectOrchestrator
type RectOrchestrator struct {
}

func (orchestrator *RectOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)

}

type SVGOrchestrator struct {
}

func (orchestrator *SVGOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedSVG, backRepoSVG *SVG) {

	stagedSVG.OnAfterUpdate(gongsvgStage, stagedSVG, backRepoSVG)

}

// LineOrchestrator
type LineOrchestrator struct {
}

func (orchestrator *LineOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLine, backRepoLine *Line) {

	stagedLine.OnAfterUpdate(gongsvgStage, stagedLine, backRepoLine)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	case Rect:
		stage.OnAfterRectUpdateCallback = new(RectOrchestrator)
	case Line:
		stage.OnAfterLineUpdateCallback = new(LineOrchestrator)
	case SVG:
		stage.OnAfterSVGUpdateCallback = new(SVGOrchestrator)
	}

}

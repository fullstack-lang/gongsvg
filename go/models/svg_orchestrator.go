package models

// SVGOrchestrator
type SVGOrchestrator struct {
}

func (orchestrator *SVGOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedSVG, backRepoSVG *SVG) {

	stagedSVG.OnAfterUpdate(gongsvgStage, stagedSVG, backRepoSVG)

}

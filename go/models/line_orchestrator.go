package models

// LineOrchestrator
type LineOrchestrator struct {
}

func (orchestrator *LineOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLine, backRepoLine *Line) {

	stagedLine.OnAfterUpdate(gongsvgStage, stagedLine, backRepoLine)
}

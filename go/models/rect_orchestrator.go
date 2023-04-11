package models

// RectOrchestrator
type RectOrchestrator struct {
}

func (orchestrator *RectOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)

}

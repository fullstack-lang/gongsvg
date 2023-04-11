package behavior

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// Orchestrator
type Orchestrator struct {
}

func (orchestrator *Orchestrator) OnAfterUpdate(
	gongsvgStage *gongsvg_models.StageStruct,
	stagedRect, backRepoRect *gongsvg_models.Rect) {

	stagedRect.Impl.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)

}

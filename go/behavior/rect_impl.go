package behavior

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type RectImpl struct {
}

func (rectImpl *RectImpl) OnAfterUpdate(stage *gongsvg_models.StageStruct, stagedRect, frontRect *gongsvg_models.Rect) {

	log.Println("Rect, OnAfterUpdate", stagedRect.Name)
}

package behavior

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// Connect every staged shape to a behavior models
func Connect(stage *gongsvg_models.StageStruct) {

	for rect := range stage.Rects {

		rectImpl := new(RectImpl)
		rect.Impl = rectImpl
	}
}

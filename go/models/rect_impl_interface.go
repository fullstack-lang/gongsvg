package models

type RectImplInterface interface {

	// OnAfterUpdate function is called each time a Rect is modified
	OnAfterUpdate(stage *StageStruct, stagedRect, frontRect *Rect)
}

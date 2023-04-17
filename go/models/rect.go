package models

import "log"

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool // alllow selected
	IsSelected   bool

	CanHaveLeftHandle bool
	HasLeftHandle     bool

	CanHaveRightHandle bool // allows HasHorizontalHandles
	HasRightHandle     bool // is true when selected

	CanMoveHorizontaly bool
	CanMoveVerticaly   bool

	Impl RectImplInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *StageStruct, _, frontRect *Rect) {

	log.Println("Rect, OnAfterUpdate", rect.Name)

	// behavior logic
	if frontRect.IsSelected != rect.IsSelected {
		rect.IsSelected = frontRect.IsSelected
		if frontRect.IsSelected && frontRect.CanHaveLeftHandle {
			rect.HasLeftHandle = true
		} else {
			rect.HasLeftHandle = false
		}
		if frontRect.IsSelected && frontRect.CanHaveRightHandle {
			rect.HasRightHandle = true
		} else {
			rect.HasRightHandle = false
		}
		rect.Commit(stage)
	}

	if rect.X != frontRect.X ||
		rect.Y != frontRect.Y ||
		rect.Width != frontRect.Width ||
		rect.Height != frontRect.Height {

		rect.X = frontRect.X
		rect.Y = frontRect.Y
		rect.Width = frontRect.Width
		rect.Height = frontRect.Height

		rect.Commit(stage)

		if rect.Impl != nil {
			rect.Impl.RectUpdated(rect)
		}
	}
}

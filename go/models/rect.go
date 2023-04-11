package models

import "log"

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool
	IsSelected   bool
}

func (rect *Rect) OnAfterUpdate(stage *StageStruct, stagedRect, frontRect *Rect) {

	log.Println("Rect, OnAfterUpdate", stagedRect.Name)
}

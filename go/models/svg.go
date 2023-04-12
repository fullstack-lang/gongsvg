package models

import "log"

type SVG struct {
	Name   string
	Layers []*Layer

	// For drawing a line between two rect
	DrawingState DrawingState
	StartRect    *Rect
	EndRect      *Rect
}

// OnAfterUpdate, notice that rect == stagedRect
func (svg *SVG) OnAfterUpdate(stage *StageStruct, _, frontSVG *SVG) {

	log.Println("Rect, OnAfterUpdate", svg.Name)

}

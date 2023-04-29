package models

import "log"

type Link struct {
	Name string

	Type LinkType

	Start           *Rect
	StartAnchorType AnchorType

	End           *Rect
	EndAnchorType AnchorType

	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	StartDirection DirectionType
	StartRatio     float64
	EndDirection   DirectionType
	EndRatio       float64

	// in case StartDirection is the same as EndDirection,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64

	ControlPoints []*Point

	Presentation
}

func (link *Link) OnAfterUpdate(stage *StageStruct, _, frontLink *Link) {

	log.Println("Link, OnAfterUpdate", link.Name)
}

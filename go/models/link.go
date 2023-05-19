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
	StartOrientation OrientationType

	EndOrientation OrientationType

	// legacy implementation

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64
	StartRatio        float64
	EndRatio          float64

	// end of legacy implementation

	// new implementation

	// horizontal & vertical cutoff for the link at start
	StartHC, StartVC float64

	// at middle
	MiddleHC, MiddleVC float64

	// at end
	EndHC, EndVC float64
	// end of new implementation

	// corner radius
	CornerRadius float64

	// Arrows
	HasEndArrow  bool
	EndArrowSize float64

	// to be displayed at the end
	TextAtArrowEnd []*AnchoredText

	// to be displayed at the start
	TextAtArrowStart []*AnchoredText

	// for non floating orthogonal anchors
	ControlPoints []*Point

	Presentation

	Impl LinkImplInterface
}

func (link *Link) OnAfterUpdate(stage *StageStruct, _, frontLink *Link) {

	log.Println("Link, OnAfterUpdate", link.Name)

	if link.Impl != nil {
		link.Impl.LinkUpdated(frontLink)
	}
}

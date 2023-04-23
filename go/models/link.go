package models

import "log"

type Link struct {
	Name string
	Presentation

	Start           *Rect
	StartAnchorType AnchorType

	End           *Rect
	EndAnchorType AnchorType
}

func (link *Link) OnAfterUpdate(stage *StageStruct, _, frontLink *Link) {

	log.Println("Link, OnAfterUpdate", link.Name)
}

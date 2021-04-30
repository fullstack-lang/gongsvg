package models

type SVG struct {
	Display bool

	Name string

	Rects []*Rect
	Texts []*Text
}

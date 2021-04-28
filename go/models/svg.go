package models

// swagger:model SVG
type SVG struct {
	Display bool

	Name string

	Rects []*Rect
}

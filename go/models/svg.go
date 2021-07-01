package models

type SVG struct {
	Display bool

	Name string

	Rects     []*Rect
	Texts     []*Text
	Lines     []*Line
	Ellipses  []*Ellipse
	Polylines []*Polyline
	Polygones []*Polygone
	Paths     []*Path
}

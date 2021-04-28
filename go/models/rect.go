package models

// swagger:model Rect
type Rect struct {
	GraphicalElement
	Name                string
	X, Y, Width, Height float64
	Color               string
}

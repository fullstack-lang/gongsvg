package models

type Rect struct {
	Name string

	// Impl is the pointer to the implementation in the models of interest
	Impl RectImplInterface

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool
	IsSelected   bool
}

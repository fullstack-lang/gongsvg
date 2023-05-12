package models

type RectAnchoredRect struct {
	Name string

	X, Y, Width, Height, RX float64

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType

	Presentation
}

package models

type RectAnchoredText struct {
	Name    string
	Content string

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType

	Presentation
	Animates []*Animate
}

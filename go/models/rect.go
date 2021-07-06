package models

type Rect struct {
	Name                    string
	X, Y, Width, Height, RX float64
	Presentation

	// Animates specifies svg <animate > elements
	Animates []*Animate
}

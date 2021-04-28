package models

// swagger:model GraphicalElement
type GraphicalElement struct {
	Fill string // The fill property paints the interior of the given graphical element.

	// // fill-opacity specifies the opacity of the painting operation used to paint the fill the current object.
	// <number>
	// The opacity of the fill. Any values outside the range 0 (fully transparent) to 1 (fully opaque) must be clamped to this range.
	// <percentage>
	// The opacity of the fill expressed as a percentage of the range 0 to 1.
	FillOpacity float64

	Stroke string // The stroke property paints along the outline of the given graphical element.

	// This property specifies the width of the stroke on the current object.
	// A zero value causes no stroke to be painted. A negative value is invalid.
	StrokeWidth   float64
	StrokeOpacity float64

	// stroke-linecap specifies the shape to be used at the end of open subpaths
	// when they are stroked, and the shape to be drawn for zero length subpaths whether they are open or closed.
	// The possible values are:
	StrokeLinecap StrokeLinecapStyle
}

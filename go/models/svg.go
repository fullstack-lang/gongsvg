package models

type SVG struct {
	Name string

	XML string // the XML representation of the SVG

	// not yet implement in gong
	// GraphicElements []GraphicElementInterface

	Rects []*Rect
}

func (svg *SVG) Output(output *string) {
	*output = `<?xml version="1.0" standalone="no"?>
	<svg width="5cm" height="4cm" version="1.1"
		 xmlns="http://www.w3.org/2000/svg">`

	for _, rect := range svg.Rects {
		rect.Output(output)
	}

	*output += `</svg>`

	return
}

package models

type Rect struct {
	Name                string
	X, Y, Width, Height float64
}

func (rect *Rect) Output(output *string) {
	*output += `
	  <rect x=".01cm" y=".01cm" width="4.98cm" height="3.98cm"
			fill="none" stroke="blue" stroke-width=".02cm" />
`

	return
}

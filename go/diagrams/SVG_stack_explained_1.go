package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongsvg/go/models"
)

var SVG_stack_explained_1 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Circle{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.Circle{}.CX,
				},
				{
					Field: models.Circle{}.CY,
				},
				{
					Field: models.Circle{}.Radius,
				},
			},
		},
		{
			Struct: &(models.Ellipse{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 160.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Ellipse{}.CX,
				},
				{
					Field: models.Ellipse{}.CY,
				},
				{
					Field: models.Ellipse{}.RX,
				},
				{
					Field: models.Ellipse{}.RY,
				},
			},
		},
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Line{}.X1,
				},
				{
					Field: models.Line{}.X2,
				},
				{
					Field: models.Line{}.Y1,
				},
				{
					Field: models.Line{}.Y2,
				},
			},
		},
		{
			Struct: &(models.Path{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 420.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Path{}.Definition,
				},
			},
		},
		{
			Struct: &(models.Polygone{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 520.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Polygone{}.Points,
				},
			},
		},
		{
			Struct: &(models.Polyline{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 630.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Polyline{}.Points,
				},
			},
		},
		{
			Struct: &(models.Presentation{}),
			Position: &uml.Position{
				X: 50.000000,
				Y: 840.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Fields: []*uml.Field{
				{
					Field: models.Presentation{}.Color,
				},
				{
					Field: models.Presentation{}.FillOpacity,
				},
				{
					Field: models.Presentation{}.Stroke,
				},
				{
					Field: models.Presentation{}.StrokeDashArray,
				},
				{
					Field: models.Presentation{}.StrokeWidth,
				},
				{
					Field: models.Presentation{}.Transform,
				},
			},
		},
		{
			Struct: &(models.Rect{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 730.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Fields: []*uml.Field{
				{
					Field: models.Rect{}.Height,
				},
				{
					Field: models.Rect{}.Width,
				},
				{
					Field: models.Rect{}.X,
				},
				{
					Field: models.Rect{}.Y,
				},
			},
		},
		{
			Struct: &(models.SVG{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.SVG{}.Circles,
					Middlevertice: &uml.Vertice{
						X: 410.000000,
						Y: 91.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Ellipses,
					Middlevertice: &uml.Vertice{
						X: 460.000000,
						Y: 216.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Lines,
					Middlevertice: &uml.Vertice{
						X: 470.000000,
						Y: 351.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Paths,
					Middlevertice: &uml.Vertice{
						X: 470.000000,
						Y: 446.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Polygones,
					Middlevertice: &uml.Vertice{
						X: 460.000000,
						Y: 556.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Polylines,
					Middlevertice: &uml.Vertice{
						X: 460.000000,
						Y: 671.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Rects,
					Middlevertice: &uml.Vertice{
						X: 410.000000,
						Y: 791.500000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.SVG{}.Texts,
					Middlevertice: &uml.Vertice{
						X: 360.000000,
						Y: 921.500000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.SVG{}.Display,
				},
			},
		},
		{
			Struct: &(models.Text{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 870.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Text{}.Content,
				},
				{
					Field: models.Text{}.X,
				},
				{
					Field: models.Text{}.Y,
				},
			},
		},
	},
}

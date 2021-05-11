// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class EllipseDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	CX?: number
	CY?: number
	RX?: number
	RY?: number
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string
	Transform?: string

	// insertion point for other declarations
	SVG_EllipsesDBID?: NullInt64
	SVG_EllipsesDBID_Index?: NullInt64 // store the index of the ellipse instance in SVG.Ellipses
	SVG_Ellipses_reverse?: SVGDB

}

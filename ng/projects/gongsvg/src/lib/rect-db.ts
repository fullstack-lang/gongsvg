// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class RectDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	X?: number
	Y?: number
	Width?: number
	Height?: number
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string
	Transform?: string

	// insertion point for other declarations
	SVG_RectsDBID?: NullInt64
	SVG_RectsDBID_Index?: NullInt64 // store the index of the rect instance in SVG.Rects
	SVG_Rects_reverse?: SVGDB

}

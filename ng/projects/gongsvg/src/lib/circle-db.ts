// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class CircleDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	CX?: number
	CY?: number
	Radius?: number
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string

	// insertion point for other declarations
	SVG_CirclesDBID?: NullInt64
	SVG_Circles_reverse?: SVGDB

}

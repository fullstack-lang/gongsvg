// insertion point for imports
import { AnimateDB } from './animate-db'
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
	Transform?: string

	// insertion point for other declarations
	CircleAnimations?: Array<AnimateDB>
	SVG_CirclesDBID?: NullInt64
	SVG_CirclesDBID_Index?: NullInt64 // store the index of the circle instance in SVG.Circles
	SVG_Circles_reverse?: SVGDB

}

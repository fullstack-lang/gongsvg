// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class PolygoneDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Points?: string
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string
	Transform?: string

	// insertion point for other declarations
	SVG_PolygonesDBID?: NullInt64
	SVG_PolygonesDBID_Index?: NullInt64 // store the index of the polygone instance in SVG.Polygones
	SVG_Polygones_reverse?: SVGDB

}

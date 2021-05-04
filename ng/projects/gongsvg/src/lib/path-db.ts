// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class PathDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Definition?: string
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string

	// insertion point for other declarations
	SVG_PathsDBID?: NullInt64
	SVG_Paths_reverse?: SVGDB

}

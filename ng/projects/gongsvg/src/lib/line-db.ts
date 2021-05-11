// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LineDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	X1?: number
	Y1?: number
	X2?: number
	Y2?: number
	Color?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeDashArray?: string
	Transform?: string

	// insertion point for other declarations
	SVG_LinesDBID?: NullInt64
	SVG_LinesDBID_Index?: NullInt64 // store the index of the line instance in SVG.Lines
	SVG_Lines_reverse?: SVGDB

}

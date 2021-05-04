// insertion point for imports
import { RectDB } from './rect-db'
import { TextDB } from './text-db'
import { CircleDB } from './circle-db'
import { LineDB } from './line-db'
import { EllipseDB } from './ellipse-db'
import { PolylineDB } from './polyline-db'
import { PolygoneDB } from './polygone-db'
import { PathDB } from './path-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class SVGDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Display?: string
	Name?: string

	// insertion point for other declarations
	Rects?: Array<RectDB>
	Texts?: Array<TextDB>
	Circles?: Array<CircleDB>
	Lines?: Array<LineDB>
	Ellipses?: Array<EllipseDB>
	Polylines?: Array<PolylineDB>
	Polygones?: Array<PolygoneDB>
	Paths?: Array<PathDB>
}

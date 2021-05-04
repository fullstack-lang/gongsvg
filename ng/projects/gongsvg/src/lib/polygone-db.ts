// insertion point for imports

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

	// insertion point for other declarations
}

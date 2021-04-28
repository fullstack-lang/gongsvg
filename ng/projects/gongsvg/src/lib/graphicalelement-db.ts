// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GraphicalElementDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Fill?: string
	FillOpacity?: number
	Stroke?: string
	StrokeWidth?: number
	StrokeOpacity?: number
	StrokeLinecap?: string

	// insertion point for other declarations
}

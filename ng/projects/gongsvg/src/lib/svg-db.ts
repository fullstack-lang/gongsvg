// insertion point for imports
import { RectDB } from './rect-db'

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
}

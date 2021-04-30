// insertion point for imports
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class TextDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	X?: number
	Y?: number
	Content?: string
	Color?: string

	// insertion point for other declarations
	SVG_TextsDBID?: NullInt64
	SVG_Texts_reverse?: SVGDB

}

// insertion point for imports
import { RectDB } from './rect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StartAnchorType: string = ""
	EndAnchorType: string = ""

	// insertion point for other declarations
	Start?: RectDB
	StartID: NullInt64 = new NullInt64 // if pointer is null, Start.ID = 0

	End?: RectDB
	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

}

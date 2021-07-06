// insertion point for imports
import { RectDB } from './rect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AnimateDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	AttributeName?: string
	Values?: string
	Dur?: string
	RepeatCount?: string

	// insertion point for other declarations
	Rect_AnimatesDBID?: NullInt64
	Rect_AnimatesDBID_Index?: NullInt64 // store the index of the animate instance in Rect.Animates
	Rect_Animates_reverse?: RectDB

}

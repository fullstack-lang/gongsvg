// insertion point for imports
import { CircleDB } from './circle-db'
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
	Circle_CircleAnimationsDBID?: NullInt64
	Circle_CircleAnimationsDBID_Index?: NullInt64 // store the index of the animate instance in Circle.CircleAnimations
	Circle_CircleAnimations_reverse?: CircleDB

	Rect_RectAnimationsDBID?: NullInt64
	Rect_RectAnimationsDBID_Index?: NullInt64 // store the index of the animate instance in Rect.RectAnimations
	Rect_RectAnimations_reverse?: RectDB

}

// insertion point for imports
import { AnimateDB } from './animate-db'
import { LinkDB } from './link-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnchoredTextDB {

	static GONGSTRUCT_NAME = "AnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	Link_TextAtArrowEndDBID: NullInt64 = new NullInt64
	Link_TextAtArrowEndDBID_Index: NullInt64  = new NullInt64 // store the index of the anchoredtext instance in Link.TextAtArrowEnd
	Link_TextAtArrowEnd_reverse?: LinkDB 

}
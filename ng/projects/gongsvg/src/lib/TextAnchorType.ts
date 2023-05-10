// generated from ng_file_enum.ts.go
export enum TextAnchorType {
	// insertion point	
	TEXT_ANCHOR_LEFT = "TEXT_ANCHOR_LEFT",
	TEXT_ANCHOR_RIGHT = "TEXT_ANCHOR_RIGHT",
	TEXT_ANCHOR_CENTER = "TEXT_ANCHOR_MIDDLE",
}

export interface TextAnchorTypeSelect {
	value: string;
	viewValue: string;
}

export const TextAnchorTypeList: TextAnchorTypeSelect[] = [ // insertion point	
	{ value: TextAnchorType.TEXT_ANCHOR_LEFT, viewValue: "TEXT_ANCHOR_LEFT" },
	{ value: TextAnchorType.TEXT_ANCHOR_RIGHT, viewValue: "TEXT_ANCHOR_RIGHT" },
	{ value: TextAnchorType.TEXT_ANCHOR_CENTER, viewValue: "TEXT_ANCHOR_MIDDLE" },
];

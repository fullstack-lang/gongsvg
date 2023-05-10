package models

type TextAnchorType string

// values for https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor
const (
	TEXT_ANCHOR_LEFT   TextAnchorType = "TEXT_ANCHOR_LEFT"
	TEXT_ANCHOR_RIGHT  TextAnchorType = "TEXT_ANCHOR_RIGHT"
	TEXT_ANCHOR_CENTER TextAnchorType = "TEXT_ANCHOR_MIDDLE"
)

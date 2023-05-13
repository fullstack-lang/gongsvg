package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Link models.StageStruct
var ___dummy__Time_Link time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Link ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Link map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.ANCHOR_BOTTOM": ref_models.ANCHOR_BOTTOM,

	"ref_models.ANCHOR_CENTER": ref_models.ANCHOR_CENTER,

	"ref_models.ANCHOR_LEFT": ref_models.ANCHOR_LEFT,

	"ref_models.ANCHOR_RIGHT": ref_models.ANCHOR_RIGHT,

	"ref_models.ANCHOR_TOP": ref_models.ANCHOR_TOP,

	"ref_models.Aliceblue": ref_models.Aliceblue,

	"ref_models.AnchorType": ref_models.AnchorType(""),

	"ref_models.AnchoredText": &(ref_models.AnchoredText{}),

	"ref_models.AnchoredText.Animates": (ref_models.AnchoredText{}).Animates,

	"ref_models.AnchoredText.Color": (ref_models.AnchoredText{}).Color,

	"ref_models.AnchoredText.Content": (ref_models.AnchoredText{}).Content,

	"ref_models.AnchoredText.FillOpacity": (ref_models.AnchoredText{}).FillOpacity,

	"ref_models.AnchoredText.Name": (ref_models.AnchoredText{}).Name,

	"ref_models.AnchoredText.Stroke": (ref_models.AnchoredText{}).Stroke,

	"ref_models.AnchoredText.StrokeDashArray": (ref_models.AnchoredText{}).StrokeDashArray,

	"ref_models.AnchoredText.StrokeDashArrayWhenSelected": (ref_models.AnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.AnchoredText.StrokeWidth": (ref_models.AnchoredText{}).StrokeWidth,

	"ref_models.AnchoredText.Transform": (ref_models.AnchoredText{}).Transform,

	"ref_models.AnchoredText.X_Offset": (ref_models.AnchoredText{}).X_Offset,

	"ref_models.AnchoredText.Y_Offset": (ref_models.AnchoredText{}).Y_Offset,

	"ref_models.Animate": &(ref_models.Animate{}),

	"ref_models.Animate.AttributeName": (ref_models.Animate{}).AttributeName,

	"ref_models.Animate.Dur": (ref_models.Animate{}).Dur,

	"ref_models.Animate.Name": (ref_models.Animate{}).Name,

	"ref_models.Animate.RepeatCount": (ref_models.Animate{}).RepeatCount,

	"ref_models.Animate.Values": (ref_models.Animate{}).Values,

	"ref_models.Antiquewhite": ref_models.Antiquewhite,

	"ref_models.Aqua": ref_models.Aqua,

	"ref_models.Aquamarine": ref_models.Aquamarine,

	"ref_models.Azure": ref_models.Azure,

	"ref_models.Beige": ref_models.Beige,

	"ref_models.Bisque": ref_models.Bisque,

	"ref_models.Black": ref_models.Black,

	"ref_models.Blanchedalmond": ref_models.Blanchedalmond,

	"ref_models.Blue": ref_models.Blue,

	"ref_models.Blueviolet": ref_models.Blueviolet,

	"ref_models.Brown": ref_models.Brown,

	"ref_models.Burlywood": ref_models.Burlywood,

	"ref_models.Cadetblue": ref_models.Cadetblue,

	"ref_models.Chartreuse": ref_models.Chartreuse,

	"ref_models.Chocolate": ref_models.Chocolate,

	"ref_models.Circle": &(ref_models.Circle{}),

	"ref_models.Circle.Animations": (ref_models.Circle{}).Animations,

	"ref_models.Circle.CX": (ref_models.Circle{}).CX,

	"ref_models.Circle.CY": (ref_models.Circle{}).CY,

	"ref_models.Circle.Color": (ref_models.Circle{}).Color,

	"ref_models.Circle.FillOpacity": (ref_models.Circle{}).FillOpacity,

	"ref_models.Circle.Name": (ref_models.Circle{}).Name,

	"ref_models.Circle.Radius": (ref_models.Circle{}).Radius,

	"ref_models.Circle.Stroke": (ref_models.Circle{}).Stroke,

	"ref_models.Circle.StrokeDashArray": (ref_models.Circle{}).StrokeDashArray,

	"ref_models.Circle.StrokeDashArrayWhenSelected": (ref_models.Circle{}).StrokeDashArrayWhenSelected,

	"ref_models.Circle.StrokeWidth": (ref_models.Circle{}).StrokeWidth,

	"ref_models.Circle.Transform": (ref_models.Circle{}).Transform,

	"ref_models.ColorType": ref_models.ColorType(""),

	"ref_models.Coral": ref_models.Coral,

	"ref_models.Cornflowerblue": ref_models.Cornflowerblue,

	"ref_models.Cornsilk": ref_models.Cornsilk,

	"ref_models.Crimson": ref_models.Crimson,

	"ref_models.Cyan": ref_models.Cyan,

	"ref_models.DRAWING_LINE": ref_models.DRAWING_LINE,

	"ref_models.Darkblue": ref_models.Darkblue,

	"ref_models.Darkcyan": ref_models.Darkcyan,

	"ref_models.Darkgoldenrod": ref_models.Darkgoldenrod,

	"ref_models.Darkgray": ref_models.Darkgray,

	"ref_models.Darkgreen": ref_models.Darkgreen,

	"ref_models.Darkgrey": ref_models.Darkgrey,

	"ref_models.Darkkhaki": ref_models.Darkkhaki,

	"ref_models.Darkmagenta": ref_models.Darkmagenta,

	"ref_models.Darkolivegreen": ref_models.Darkolivegreen,

	"ref_models.Darkorange": ref_models.Darkorange,

	"ref_models.Darkorchid": ref_models.Darkorchid,

	"ref_models.Darkred": ref_models.Darkred,

	"ref_models.Darksalmon": ref_models.Darksalmon,

	"ref_models.Darkseagreen": ref_models.Darkseagreen,

	"ref_models.Darkslateblue": ref_models.Darkslateblue,

	"ref_models.Darkslategray": ref_models.Darkslategray,

	"ref_models.Darkslategrey": ref_models.Darkslategrey,

	"ref_models.Darkturquoise": ref_models.Darkturquoise,

	"ref_models.Darkviolet": ref_models.Darkviolet,

	"ref_models.Deeppink": ref_models.Deeppink,

	"ref_models.Deepskyblue": ref_models.Deepskyblue,

	"ref_models.Dimgray": ref_models.Dimgray,

	"ref_models.Dimgrey": ref_models.Dimgrey,

	"ref_models.Dodgerblue": ref_models.Dodgerblue,

	"ref_models.DrawingState": ref_models.DrawingState(""),

	"ref_models.Ellipse": &(ref_models.Ellipse{}),

	"ref_models.Ellipse.Animates": (ref_models.Ellipse{}).Animates,

	"ref_models.Ellipse.CX": (ref_models.Ellipse{}).CX,

	"ref_models.Ellipse.CY": (ref_models.Ellipse{}).CY,

	"ref_models.Ellipse.Color": (ref_models.Ellipse{}).Color,

	"ref_models.Ellipse.FillOpacity": (ref_models.Ellipse{}).FillOpacity,

	"ref_models.Ellipse.Name": (ref_models.Ellipse{}).Name,

	"ref_models.Ellipse.RX": (ref_models.Ellipse{}).RX,

	"ref_models.Ellipse.RY": (ref_models.Ellipse{}).RY,

	"ref_models.Ellipse.Stroke": (ref_models.Ellipse{}).Stroke,

	"ref_models.Ellipse.StrokeDashArray": (ref_models.Ellipse{}).StrokeDashArray,

	"ref_models.Ellipse.StrokeDashArrayWhenSelected": (ref_models.Ellipse{}).StrokeDashArrayWhenSelected,

	"ref_models.Ellipse.StrokeWidth": (ref_models.Ellipse{}).StrokeWidth,

	"ref_models.Ellipse.Transform": (ref_models.Ellipse{}).Transform,

	"ref_models.Firebrick": ref_models.Firebrick,

	"ref_models.Floralwhite": ref_models.Floralwhite,

	"ref_models.Forestgreen": ref_models.Forestgreen,

	"ref_models.Fuchsia": ref_models.Fuchsia,

	"ref_models.Gainsboro": ref_models.Gainsboro,

	"ref_models.Ghostwhite": ref_models.Ghostwhite,

	"ref_models.Gold": ref_models.Gold,

	"ref_models.Goldenrod": ref_models.Goldenrod,

	"ref_models.Gray": ref_models.Gray,

	"ref_models.Green": ref_models.Green,

	"ref_models.Greenyellow": ref_models.Greenyellow,

	"ref_models.Grey": ref_models.Grey,

	"ref_models.Honeydew": ref_models.Honeydew,

	"ref_models.Hotpink": ref_models.Hotpink,

	"ref_models.Indianred": ref_models.Indianred,

	"ref_models.Indigo": ref_models.Indigo,

	"ref_models.Ivory": ref_models.Ivory,

	"ref_models.Khaki": ref_models.Khaki,

	"ref_models.LINK_TYPE_FLOATING_ORTHOGONAL": ref_models.LINK_TYPE_FLOATING_ORTHOGONAL,

	"ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS": ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS,

	"ref_models.Lavender": ref_models.Lavender,

	"ref_models.Lavenderblush": ref_models.Lavenderblush,

	"ref_models.Lawngreen": ref_models.Lawngreen,

	"ref_models.Layer": &(ref_models.Layer{}),

	"ref_models.Layer.Circles": (ref_models.Layer{}).Circles,

	"ref_models.Layer.Display": (ref_models.Layer{}).Display,

	"ref_models.Layer.Ellipses": (ref_models.Layer{}).Ellipses,

	"ref_models.Layer.Lines": (ref_models.Layer{}).Lines,

	"ref_models.Layer.Links": (ref_models.Layer{}).Links,

	"ref_models.Layer.Name": (ref_models.Layer{}).Name,

	"ref_models.Layer.Paths": (ref_models.Layer{}).Paths,

	"ref_models.Layer.Polygones": (ref_models.Layer{}).Polygones,

	"ref_models.Layer.Polylines": (ref_models.Layer{}).Polylines,

	"ref_models.Layer.Rects": (ref_models.Layer{}).Rects,

	"ref_models.Layer.Texts": (ref_models.Layer{}).Texts,

	"ref_models.Lemonchiffon": ref_models.Lemonchiffon,

	"ref_models.Lightblue": ref_models.Lightblue,

	"ref_models.Lightcoral": ref_models.Lightcoral,

	"ref_models.Lightcyan": ref_models.Lightcyan,

	"ref_models.Lightgoldenrodyellow": ref_models.Lightgoldenrodyellow,

	"ref_models.Lightgray": ref_models.Lightgray,

	"ref_models.Lightgreen": ref_models.Lightgreen,

	"ref_models.Lightgrey": ref_models.Lightgrey,

	"ref_models.Lightpink": ref_models.Lightpink,

	"ref_models.Lightsalmon": ref_models.Lightsalmon,

	"ref_models.Lightseagreen": ref_models.Lightseagreen,

	"ref_models.Lightskyblue": ref_models.Lightskyblue,

	"ref_models.Lightslategray": ref_models.Lightslategray,

	"ref_models.Lightslategrey": ref_models.Lightslategrey,

	"ref_models.Lightsteelblue": ref_models.Lightsteelblue,

	"ref_models.Lightyellow": ref_models.Lightyellow,

	"ref_models.Lime": ref_models.Lime,

	"ref_models.Limegreen": ref_models.Limegreen,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Animates": (ref_models.Line{}).Animates,

	"ref_models.Line.Color": (ref_models.Line{}).Color,

	"ref_models.Line.FillOpacity": (ref_models.Line{}).FillOpacity,

	"ref_models.Line.MouseClickX": (ref_models.Line{}).MouseClickX,

	"ref_models.Line.MouseClickY": (ref_models.Line{}).MouseClickY,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Stroke": (ref_models.Line{}).Stroke,

	"ref_models.Line.StrokeDashArray": (ref_models.Line{}).StrokeDashArray,

	"ref_models.Line.StrokeDashArrayWhenSelected": (ref_models.Line{}).StrokeDashArrayWhenSelected,

	"ref_models.Line.StrokeWidth": (ref_models.Line{}).StrokeWidth,

	"ref_models.Line.Transform": (ref_models.Line{}).Transform,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Linen": ref_models.Linen,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.Color": (ref_models.Link{}).Color,

	"ref_models.Link.ControlPoints": (ref_models.Link{}).ControlPoints,

	"ref_models.Link.CornerOffsetRatio": (ref_models.Link{}).CornerOffsetRatio,

	"ref_models.Link.CornerRadius": (ref_models.Link{}).CornerRadius,

	"ref_models.Link.End": (ref_models.Link{}).End,

	"ref_models.Link.EndAnchorType": (ref_models.Link{}).EndAnchorType,

	"ref_models.Link.EndArrowSize": (ref_models.Link{}).EndArrowSize,

	"ref_models.Link.EndOrientation": (ref_models.Link{}).EndOrientation,

	"ref_models.Link.EndRatio": (ref_models.Link{}).EndRatio,

	"ref_models.Link.FillOpacity": (ref_models.Link{}).FillOpacity,

	"ref_models.Link.HasEndArrow": (ref_models.Link{}).HasEndArrow,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.Start": (ref_models.Link{}).Start,

	"ref_models.Link.StartAnchorType": (ref_models.Link{}).StartAnchorType,

	"ref_models.Link.StartOrientation": (ref_models.Link{}).StartOrientation,

	"ref_models.Link.StartRatio": (ref_models.Link{}).StartRatio,

	"ref_models.Link.Stroke": (ref_models.Link{}).Stroke,

	"ref_models.Link.StrokeDashArray": (ref_models.Link{}).StrokeDashArray,

	"ref_models.Link.StrokeDashArrayWhenSelected": (ref_models.Link{}).StrokeDashArrayWhenSelected,

	"ref_models.Link.StrokeWidth": (ref_models.Link{}).StrokeWidth,

	"ref_models.Link.TextAtArrowEnd": (ref_models.Link{}).TextAtArrowEnd,

	"ref_models.Link.TextAtArrowStart": (ref_models.Link{}).TextAtArrowStart,

	"ref_models.Link.Transform": (ref_models.Link{}).Transform,

	"ref_models.Link.Type": (ref_models.Link{}).Type,

	"ref_models.LinkType": ref_models.LinkType(""),

	"ref_models.Magenta": ref_models.Magenta,

	"ref_models.Maroon": ref_models.Maroon,

	"ref_models.Mediumaquamarine": ref_models.Mediumaquamarine,

	"ref_models.Mediumblue": ref_models.Mediumblue,

	"ref_models.Mediumorchid": ref_models.Mediumorchid,

	"ref_models.Mediumpurple": ref_models.Mediumpurple,

	"ref_models.Mediumseagreen": ref_models.Mediumseagreen,

	"ref_models.Mediumslateblue": ref_models.Mediumslateblue,

	"ref_models.Mediumspringgreen": ref_models.Mediumspringgreen,

	"ref_models.Mediumturquoise": ref_models.Mediumturquoise,

	"ref_models.Mediumvioletred": ref_models.Mediumvioletred,

	"ref_models.Midnightblue": ref_models.Midnightblue,

	"ref_models.Mintcream": ref_models.Mintcream,

	"ref_models.Mistyrose": ref_models.Mistyrose,

	"ref_models.Moccasin": ref_models.Moccasin,

	"ref_models.NOT_DRAWING_LINE": ref_models.NOT_DRAWING_LINE,

	"ref_models.Navajowhite": ref_models.Navajowhite,

	"ref_models.Navy": ref_models.Navy,

	"ref_models.ORIENTATION_HORIZONTAL": ref_models.ORIENTATION_HORIZONTAL,

	"ref_models.ORIENTATION_VERTICAL": ref_models.ORIENTATION_VERTICAL,

	"ref_models.Oldlace": ref_models.Oldlace,

	"ref_models.Olive": ref_models.Olive,

	"ref_models.Olivedrab": ref_models.Olivedrab,

	"ref_models.Orange": ref_models.Orange,

	"ref_models.Orangered": ref_models.Orangered,

	"ref_models.Orchid": ref_models.Orchid,

	"ref_models.OrientationType": ref_models.OrientationType(""),

	"ref_models.POSITION_ON_ARROW_END": ref_models.POSITION_ON_ARROW_END,

	"ref_models.POSITION_ON_ARROW_START": ref_models.POSITION_ON_ARROW_START,

	"ref_models.Palegoldenrod": ref_models.Palegoldenrod,

	"ref_models.Palegreen": ref_models.Palegreen,

	"ref_models.Paleturquoise": ref_models.Paleturquoise,

	"ref_models.Palevioletred": ref_models.Palevioletred,

	"ref_models.Papayawhip": ref_models.Papayawhip,

	"ref_models.Path": &(ref_models.Path{}),

	"ref_models.Path.Animates": (ref_models.Path{}).Animates,

	"ref_models.Path.Color": (ref_models.Path{}).Color,

	"ref_models.Path.Definition": (ref_models.Path{}).Definition,

	"ref_models.Path.FillOpacity": (ref_models.Path{}).FillOpacity,

	"ref_models.Path.Name": (ref_models.Path{}).Name,

	"ref_models.Path.Stroke": (ref_models.Path{}).Stroke,

	"ref_models.Path.StrokeDashArray": (ref_models.Path{}).StrokeDashArray,

	"ref_models.Path.StrokeDashArrayWhenSelected": (ref_models.Path{}).StrokeDashArrayWhenSelected,

	"ref_models.Path.StrokeWidth": (ref_models.Path{}).StrokeWidth,

	"ref_models.Path.Transform": (ref_models.Path{}).Transform,

	"ref_models.Peachpuff": ref_models.Peachpuff,

	"ref_models.Peru": ref_models.Peru,

	"ref_models.Pink": ref_models.Pink,

	"ref_models.Plum": ref_models.Plum,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Polygone": &(ref_models.Polygone{}),

	"ref_models.Polygone.Animates": (ref_models.Polygone{}).Animates,

	"ref_models.Polygone.Color": (ref_models.Polygone{}).Color,

	"ref_models.Polygone.FillOpacity": (ref_models.Polygone{}).FillOpacity,

	"ref_models.Polygone.Name": (ref_models.Polygone{}).Name,

	"ref_models.Polygone.Points": (ref_models.Polygone{}).Points,

	"ref_models.Polygone.Stroke": (ref_models.Polygone{}).Stroke,

	"ref_models.Polygone.StrokeDashArray": (ref_models.Polygone{}).StrokeDashArray,

	"ref_models.Polygone.StrokeDashArrayWhenSelected": (ref_models.Polygone{}).StrokeDashArrayWhenSelected,

	"ref_models.Polygone.StrokeWidth": (ref_models.Polygone{}).StrokeWidth,

	"ref_models.Polygone.Transform": (ref_models.Polygone{}).Transform,

	"ref_models.Polyline": &(ref_models.Polyline{}),

	"ref_models.Polyline.Animates": (ref_models.Polyline{}).Animates,

	"ref_models.Polyline.Color": (ref_models.Polyline{}).Color,

	"ref_models.Polyline.FillOpacity": (ref_models.Polyline{}).FillOpacity,

	"ref_models.Polyline.Name": (ref_models.Polyline{}).Name,

	"ref_models.Polyline.Points": (ref_models.Polyline{}).Points,

	"ref_models.Polyline.Stroke": (ref_models.Polyline{}).Stroke,

	"ref_models.Polyline.StrokeDashArray": (ref_models.Polyline{}).StrokeDashArray,

	"ref_models.Polyline.StrokeDashArrayWhenSelected": (ref_models.Polyline{}).StrokeDashArrayWhenSelected,

	"ref_models.Polyline.StrokeWidth": (ref_models.Polyline{}).StrokeWidth,

	"ref_models.Polyline.Transform": (ref_models.Polyline{}).Transform,

	"ref_models.PositionOnArrowType": ref_models.PositionOnArrowType(""),

	"ref_models.Powderblue": ref_models.Powderblue,

	"ref_models.Purple": ref_models.Purple,

	"ref_models.RECT_ANCHOR_BOTTOM": ref_models.RECT_ANCHOR_BOTTOM,

	"ref_models.RECT_ANCHOR_BOTTOM_LEFT": ref_models.RECT_ANCHOR_BOTTOM_LEFT,

	"ref_models.RECT_ANCHOR_BOTTOM_RIGHT": ref_models.RECT_ANCHOR_BOTTOM_RIGHT,

	"ref_models.RECT_ANCHOR_CENTER": ref_models.RECT_ANCHOR_CENTER,

	"ref_models.RECT_ANCHOR_LEFT": ref_models.RECT_ANCHOR_LEFT,

	"ref_models.RECT_ANCHOR_RIGHT": ref_models.RECT_ANCHOR_RIGHT,

	"ref_models.RECT_ANCHOR_TOP": ref_models.RECT_ANCHOR_TOP,

	"ref_models.RECT_ANCHOR_TOP_LEFT": ref_models.RECT_ANCHOR_TOP_LEFT,

	"ref_models.RECT_ANCHOR_TOP_RIGHT": ref_models.RECT_ANCHOR_TOP_RIGHT,

	"ref_models.Rect": &(ref_models.Rect{}),

	"ref_models.Rect.Animations": (ref_models.Rect{}).Animations,

	"ref_models.Rect.CanHaveBottomHandle": (ref_models.Rect{}).CanHaveBottomHandle,

	"ref_models.Rect.CanHaveLeftHandle": (ref_models.Rect{}).CanHaveLeftHandle,

	"ref_models.Rect.CanHaveLeftHandle": (ref_models.Rect{}).CanHaveLeftHandle,

	"ref_models.Rect.CanHaveRightHandle": (ref_models.Rect{}).CanHaveRightHandle,

	"ref_models.Rect.CanHaveTopHandle": (ref_models.Rect{}).CanHaveTopHandle,

	"ref_models.Rect.CanMoveHorizontaly": (ref_models.Rect{}).CanMoveHorizontaly,

	"ref_models.Rect.CanMoveVerticaly": (ref_models.Rect{}).CanMoveVerticaly,

	"ref_models.Rect.Color": (ref_models.Rect{}).Color,

	"ref_models.Rect.FillOpacity": (ref_models.Rect{}).FillOpacity,

	"ref_models.Rect.HasBottomHandle": (ref_models.Rect{}).HasBottomHandle,

	"ref_models.Rect.HasLeftHandle": (ref_models.Rect{}).HasLeftHandle,

	"ref_models.Rect.HasLeftHandle": (ref_models.Rect{}).HasLeftHandle,

	"ref_models.Rect.HasRightHandle": (ref_models.Rect{}).HasRightHandle,

	"ref_models.Rect.HasTopHandle": (ref_models.Rect{}).HasTopHandle,

	"ref_models.Rect.Height": (ref_models.Rect{}).Height,

	"ref_models.Rect.IsSelectable": (ref_models.Rect{}).IsSelectable,

	"ref_models.Rect.IsSelected": (ref_models.Rect{}).IsSelected,

	"ref_models.Rect.Name": (ref_models.Rect{}).Name,

	"ref_models.Rect.RX": (ref_models.Rect{}).RX,

	"ref_models.Rect.RectAnchoredRects": (ref_models.Rect{}).RectAnchoredRects,

	"ref_models.Rect.RectAnchoredTexts": (ref_models.Rect{}).RectAnchoredTexts,

	"ref_models.Rect.Stroke": (ref_models.Rect{}).Stroke,

	"ref_models.Rect.StrokeDashArray": (ref_models.Rect{}).StrokeDashArray,

	"ref_models.Rect.StrokeDashArrayWhenSelected": (ref_models.Rect{}).StrokeDashArrayWhenSelected,

	"ref_models.Rect.StrokeWidth": (ref_models.Rect{}).StrokeWidth,

	"ref_models.Rect.Transform": (ref_models.Rect{}).Transform,

	"ref_models.Rect.Width": (ref_models.Rect{}).Width,

	"ref_models.Rect.X": (ref_models.Rect{}).X,

	"ref_models.Rect.Y": (ref_models.Rect{}).Y,

	"ref_models.RectAnchorType": ref_models.RectAnchorType(""),

	"ref_models.RectAnchoredRect": &(ref_models.RectAnchoredRect{}),

	"ref_models.RectAnchoredRect.Color": (ref_models.RectAnchoredRect{}).Color,

	"ref_models.RectAnchoredRect.FillOpacity": (ref_models.RectAnchoredRect{}).FillOpacity,

	"ref_models.RectAnchoredRect.Height": (ref_models.RectAnchoredRect{}).Height,

	"ref_models.RectAnchoredRect.HeightFollowRect": (ref_models.RectAnchoredRect{}).HeightFollowRect,

	"ref_models.RectAnchoredRect.Name": (ref_models.RectAnchoredRect{}).Name,

	"ref_models.RectAnchoredRect.RX": (ref_models.RectAnchoredRect{}).RX,

	"ref_models.RectAnchoredRect.RectAnchorType": (ref_models.RectAnchoredRect{}).RectAnchorType,

	"ref_models.RectAnchoredRect.Stroke": (ref_models.RectAnchoredRect{}).Stroke,

	"ref_models.RectAnchoredRect.StrokeDashArray": (ref_models.RectAnchoredRect{}).StrokeDashArray,

	"ref_models.RectAnchoredRect.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredRect{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredRect.StrokeWidth": (ref_models.RectAnchoredRect{}).StrokeWidth,

	"ref_models.RectAnchoredRect.Transform": (ref_models.RectAnchoredRect{}).Transform,

	"ref_models.RectAnchoredRect.Width": (ref_models.RectAnchoredRect{}).Width,

	"ref_models.RectAnchoredRect.WidthFollowRect": (ref_models.RectAnchoredRect{}).WidthFollowRect,

	"ref_models.RectAnchoredRect.X": (ref_models.RectAnchoredRect{}).X,

	"ref_models.RectAnchoredRect.X_Offset": (ref_models.RectAnchoredRect{}).X_Offset,

	"ref_models.RectAnchoredRect.Y": (ref_models.RectAnchoredRect{}).Y,

	"ref_models.RectAnchoredRect.Y_Offset": (ref_models.RectAnchoredRect{}).Y_Offset,

	"ref_models.RectAnchoredText": &(ref_models.RectAnchoredText{}),

	"ref_models.RectAnchoredText.Animates": (ref_models.RectAnchoredText{}).Animates,

	"ref_models.RectAnchoredText.Color": (ref_models.RectAnchoredText{}).Color,

	"ref_models.RectAnchoredText.Content": (ref_models.RectAnchoredText{}).Content,

	"ref_models.RectAnchoredText.FillOpacity": (ref_models.RectAnchoredText{}).FillOpacity,

	"ref_models.RectAnchoredText.FontSize": (ref_models.RectAnchoredText{}).FontSize,

	"ref_models.RectAnchoredText.FontWeight": (ref_models.RectAnchoredText{}).FontWeight,

	"ref_models.RectAnchoredText.Name": (ref_models.RectAnchoredText{}).Name,

	"ref_models.RectAnchoredText.RectAnchorType": (ref_models.RectAnchoredText{}).RectAnchorType,

	"ref_models.RectAnchoredText.Stroke": (ref_models.RectAnchoredText{}).Stroke,

	"ref_models.RectAnchoredText.StrokeDashArray": (ref_models.RectAnchoredText{}).StrokeDashArray,

	"ref_models.RectAnchoredText.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredText.StrokeWidth": (ref_models.RectAnchoredText{}).StrokeWidth,

	"ref_models.RectAnchoredText.TextAnchorType": (ref_models.RectAnchoredText{}).TextAnchorType,

	"ref_models.RectAnchoredText.Transform": (ref_models.RectAnchoredText{}).Transform,

	"ref_models.RectAnchoredText.X_Offset": (ref_models.RectAnchoredText{}).X_Offset,

	"ref_models.RectAnchoredText.Y_Offset": (ref_models.RectAnchoredText{}).Y_Offset,

	"ref_models.RectLinkLink": &(ref_models.RectLinkLink{}),

	"ref_models.RectLinkLink.End": (ref_models.RectLinkLink{}).End,

	"ref_models.RectLinkLink.Name": (ref_models.RectLinkLink{}).Name,

	"ref_models.RectLinkLink.Start": (ref_models.RectLinkLink{}).Start,

	"ref_models.RectLinkLink.TargetAnchorPosition": (ref_models.RectLinkLink{}).TargetAnchorPosition,

	"ref_models.Red": ref_models.Red,

	"ref_models.Rosybrown": ref_models.Rosybrown,

	"ref_models.Royalblue": ref_models.Royalblue,

	"ref_models.SIDE_BOTTOM": ref_models.SIDE_BOTTOM,

	"ref_models.SIDE_LEFT": ref_models.SIDE_LEFT,

	"ref_models.SIDE_RIGHT": ref_models.SIDE_RIGHT,

	"ref_models.SIDE_TOP": ref_models.SIDE_TOP,

	"ref_models.SVG": &(ref_models.SVG{}),

	"ref_models.SVG.DrawingState": (ref_models.SVG{}).DrawingState,

	"ref_models.SVG.EndRect": (ref_models.SVG{}).EndRect,

	"ref_models.SVG.Layers": (ref_models.SVG{}).Layers,

	"ref_models.SVG.Name": (ref_models.SVG{}).Name,

	"ref_models.SVG.StartRect": (ref_models.SVG{}).StartRect,

	"ref_models.Saddlebrown": ref_models.Saddlebrown,

	"ref_models.Salmon": ref_models.Salmon,

	"ref_models.Sandybrown": ref_models.Sandybrown,

	"ref_models.Seagreen": ref_models.Seagreen,

	"ref_models.Seashell": ref_models.Seashell,

	"ref_models.SideType": ref_models.SideType(""),

	"ref_models.Sienna": ref_models.Sienna,

	"ref_models.Silver": ref_models.Silver,

	"ref_models.Skyblue": ref_models.Skyblue,

	"ref_models.Slateblue": ref_models.Slateblue,

	"ref_models.Slategray": ref_models.Slategray,

	"ref_models.Slategrey": ref_models.Slategrey,

	"ref_models.Snow": ref_models.Snow,

	"ref_models.Springgreen": ref_models.Springgreen,

	"ref_models.Steelblue": ref_models.Steelblue,

	"ref_models.TEXT_ANCHOR_CENTER": ref_models.TEXT_ANCHOR_CENTER,

	"ref_models.TEXT_ANCHOR_LEFT": ref_models.TEXT_ANCHOR_LEFT,

	"ref_models.TEXT_ANCHOR_RIGHT": ref_models.TEXT_ANCHOR_RIGHT,

	"ref_models.Tan": ref_models.Tan,

	"ref_models.Teal": ref_models.Teal,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Animates": (ref_models.Text{}).Animates,

	"ref_models.Text.Color": (ref_models.Text{}).Color,

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.FillOpacity": (ref_models.Text{}).FillOpacity,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Stroke": (ref_models.Text{}).Stroke,

	"ref_models.Text.StrokeDashArray": (ref_models.Text{}).StrokeDashArray,

	"ref_models.Text.StrokeDashArrayWhenSelected": (ref_models.Text{}).StrokeDashArrayWhenSelected,

	"ref_models.Text.StrokeWidth": (ref_models.Text{}).StrokeWidth,

	"ref_models.Text.Transform": (ref_models.Text{}).Transform,

	"ref_models.Text.X": (ref_models.Text{}).X,

	"ref_models.Text.Y": (ref_models.Text{}).Y,

	"ref_models.TextAnchorType": ref_models.TextAnchorType(""),

	"ref_models.Thistle": ref_models.Thistle,

	"ref_models.Tomato": ref_models.Tomato,

	"ref_models.Turquoise": ref_models.Turquoise,

	"ref_models.Violet": ref_models.Violet,

	"ref_models.Wheat": ref_models.Wheat,

	"ref_models.White": ref_models.White,

	"ref_models.Whitesmoke": ref_models.Whitesmoke,

	"ref_models.Yellow": ref_models.Yellow,

	"ref_models.Yellowgreen": ref_models.Yellowgreen,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Link"] = LinkInjection
// }

// LinkInjection will stage objects of database "Link"
func LinkInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Links := (&models.Classdiagram{Name: `Links`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_1_Link := (&models.GongStructShape{Name: `NewDiagram_1-Link`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_1_Rect := (&models.GongStructShape{Name: `NewDiagram_1-Rect`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_1_RectLinkLink := (&models.GongStructShape{Name: `NewDiagram_1-RectLinkLink`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000001_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000002_Start := (&models.Link{Name: `Start`}).Stage(stage)
	__Link__000003_Start := (&models.Link{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Link := (&models.Position{Name: `Pos-NewDiagram_1-Link`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_1_Rect := (&models.Position{Name: `Pos-NewDiagram_1-Rect`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_1_RectLinkLink := (&models.Position{Name: `Pos-NewDiagram_1-RectLinkLink`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Link and NewDiagram_1-Rect`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Link and NewDiagram_1-Rect`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Link := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-RectLinkLink and NewDiagram_1-Link`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-RectLinkLink and NewDiagram_1-Rect`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Links.Name = `Links`
	__Classdiagram__000000_Links.IsInDrawMode = false

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_1_Link.Name = `NewDiagram_1-Link`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__GongStructShape__000000_NewDiagram_1_Link.Identifier = `ref_models.Link`
	__GongStructShape__000000_NewDiagram_1_Link.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_1_Link.NbInstances = 0
	__GongStructShape__000000_NewDiagram_1_Link.Width = 240.000000
	__GongStructShape__000000_NewDiagram_1_Link.Heigth = 63.000000
	__GongStructShape__000000_NewDiagram_1_Link.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_1_Rect.Name = `NewDiagram_1-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000001_NewDiagram_1_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000001_NewDiagram_1_Rect.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_1_Rect.NbInstances = 0
	__GongStructShape__000001_NewDiagram_1_Rect.Width = 240.000000
	__GongStructShape__000001_NewDiagram_1_Rect.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_1_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Name = `NewDiagram_1-RectLinkLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Identifier = `ref_models.RectLinkLink`
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.NbInstances = 0
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Width = 240.000000
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Heigth = 63.000000
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.End]
	__Link__000000_End.Identifier = `ref_models.Link.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000000_End.Fieldtypename = `ref_models.Rect`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink.End]
	__Link__000001_End.Identifier = `ref_models.RectLinkLink.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Link__000001_End.Fieldtypename = `ref_models.Link`
	__Link__000001_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000002_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Start]
	__Link__000002_Start.Identifier = `ref_models.Link.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000002_Start.Fieldtypename = `ref_models.Rect`
	__Link__000002_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_Start.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000003_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink.Start]
	__Link__000003_Start.Identifier = `ref_models.RectLinkLink.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000003_Start.Fieldtypename = `ref_models.Rect`
	__Link__000003_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_Start.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Link.X = 60.000000
	__Position__000000_Pos_NewDiagram_1_Link.Y = 90.000000
	__Position__000000_Pos_NewDiagram_1_Link.Name = `Pos-NewDiagram_1-Link`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_Rect.X = 540.000000
	__Position__000001_Pos_NewDiagram_1_Rect.Y = 80.000000
	__Position__000001_Pos_NewDiagram_1_Rect.Name = `Pos-NewDiagram_1-Rect`

	// Position values setup
	__Position__000002_Pos_NewDiagram_1_RectLinkLink.X = 280.000000
	__Position__000002_Pos_NewDiagram_1_RectLinkLink.Y = 390.000000
	__Position__000002_Pos_NewDiagram_1_RectLinkLink.Name = `Pos-NewDiagram_1-RectLinkLink`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.X = 333.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.Y = 28.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Link and NewDiagram_1-Rect`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.X = 403.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.Y = 218.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Link and NewDiagram_1-Rect`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Link.X = 162.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Link.Y = 283.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Link.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-RectLinkLink and NewDiagram_1-Link`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Rect.X = 640.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Rect.Y = 205.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Rect.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-RectLinkLink and NewDiagram_1-Rect`

	// Setup of pointers
	__Classdiagram__000000_Links.GongStructShapes = append(__Classdiagram__000000_Links.GongStructShapes, __GongStructShape__000001_NewDiagram_1_Rect)
	__Classdiagram__000000_Links.GongStructShapes = append(__Classdiagram__000000_Links.GongStructShapes, __GongStructShape__000000_NewDiagram_1_Link)
	__Classdiagram__000000_Links.GongStructShapes = append(__Classdiagram__000000_Links.GongStructShapes, __GongStructShape__000002_NewDiagram_1_RectLinkLink)
	__GongStructShape__000000_NewDiagram_1_Link.Position = __Position__000000_Pos_NewDiagram_1_Link
	__GongStructShape__000000_NewDiagram_1_Link.Links = append(__GongStructShape__000000_NewDiagram_1_Link.Links, __Link__000000_End)
	__GongStructShape__000000_NewDiagram_1_Link.Links = append(__GongStructShape__000000_NewDiagram_1_Link.Links, __Link__000002_Start)
	__GongStructShape__000001_NewDiagram_1_Rect.Position = __Position__000001_Pos_NewDiagram_1_Rect
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Position = __Position__000002_Pos_NewDiagram_1_RectLinkLink
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Links = append(__GongStructShape__000002_NewDiagram_1_RectLinkLink.Links, __Link__000003_Start)
	__GongStructShape__000002_NewDiagram_1_RectLinkLink.Links = append(__GongStructShape__000002_NewDiagram_1_RectLinkLink.Links, __Link__000001_End)
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect
	__Link__000001_End.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Link
	__Link__000002_Start.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Link_and_NewDiagram_1_Rect
	__Link__000003_Start.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_RectLinkLink_and_NewDiagram_1_Rect
}

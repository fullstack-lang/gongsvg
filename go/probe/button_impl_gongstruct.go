// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Animate":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAnimateFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		animate := new(models.Animate)
		FillUpForm(animate, formGroup, buttonImpl.playground)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCircleFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		FillUpForm(circle, formGroup, buttonImpl.playground)
	case "Ellipse":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewEllipseFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, formGroup, buttonImpl.playground)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLayerFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		layer := new(models.Layer)
		FillUpForm(layer, formGroup, buttonImpl.playground)
	case "Line":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLineFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		line := new(models.Line)
		FillUpForm(line, formGroup, buttonImpl.playground)
	case "Link":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, formGroup, buttonImpl.playground)
	case "LinkAnchoredText":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkAnchoredTextFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, formGroup, buttonImpl.playground)
	case "Path":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPathFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		path := new(models.Path)
		FillUpForm(path, formGroup, buttonImpl.playground)
	case "Point":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPointFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		point := new(models.Point)
		FillUpForm(point, formGroup, buttonImpl.playground)
	case "Polygone":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPolygoneFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		polygone := new(models.Polygone)
		FillUpForm(polygone, formGroup, buttonImpl.playground)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPolylineFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		polyline := new(models.Polyline)
		FillUpForm(polyline, formGroup, buttonImpl.playground)
	case "Rect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		rect := new(models.Rect)
		FillUpForm(rect, formGroup, buttonImpl.playground)
	case "RectAnchoredRect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredRectFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, formGroup, buttonImpl.playground)
	case "RectAnchoredText":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredTextFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, formGroup, buttonImpl.playground)
	case "RectLinkLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectLinkLinkFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, formGroup, buttonImpl.playground)
	case "SVG":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewSVGFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		svg := new(models.SVG)
		FillUpForm(svg, formGroup, buttonImpl.playground)
	case "Text":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewTextFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		text := new(models.Text)
		FillUpForm(text, formGroup, buttonImpl.playground)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Animate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("AttributeName", instanceWithInferedType.AttributeName, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Values", instanceWithInferedType.Values, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Dur", instanceWithInferedType.Dur, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RepeatCount", instanceWithInferedType.RepeatCount, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, playground)

	case *models.Ellipse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RY", instanceWithInferedType.RY, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Rects", instanceWithInferedType, &instanceWithInferedType.Rects, formGroup, playground)
		AssociationSliceToForm("Texts", instanceWithInferedType, &instanceWithInferedType.Texts, formGroup, playground)
		AssociationSliceToForm("Circles", instanceWithInferedType, &instanceWithInferedType.Circles, formGroup, playground)
		AssociationSliceToForm("Lines", instanceWithInferedType, &instanceWithInferedType.Lines, formGroup, playground)
		AssociationSliceToForm("Ellipses", instanceWithInferedType, &instanceWithInferedType.Ellipses, formGroup, playground)
		AssociationSliceToForm("Polylines", instanceWithInferedType, &instanceWithInferedType.Polylines, formGroup, playground)
		AssociationSliceToForm("Polygones", instanceWithInferedType, &instanceWithInferedType.Polygones, formGroup, playground)
		AssociationSliceToForm("Paths", instanceWithInferedType, &instanceWithInferedType.Paths, formGroup, playground)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)
		AssociationSliceToForm("RectLinkLinks", instanceWithInferedType, &instanceWithInferedType.RectLinkLinks, formGroup, playground)

	case *models.Line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X1", instanceWithInferedType.X1, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y1", instanceWithInferedType.Y1, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X2", instanceWithInferedType.X2, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y2", instanceWithInferedType.Y2, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		BasicFieldtoForm("MouseClickX", instanceWithInferedType.MouseClickX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MouseClickY", instanceWithInferedType.MouseClickY, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, playground)
		EnumTypeStringToForm("StartAnchorType", instanceWithInferedType.StartAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, playground)
		EnumTypeStringToForm("EndAnchorType", instanceWithInferedType.EndAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CornerRadius", instanceWithInferedType.CornerRadius, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasEndArrow", instanceWithInferedType.HasEndArrow, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndArrowSize", instanceWithInferedType.EndArrowSize, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("TextAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.TextAtArrowEnd, formGroup, playground)
		AssociationSliceToForm("TextAtArrowStart", instanceWithInferedType, &instanceWithInferedType.TextAtArrowStart, formGroup, playground)
		AssociationSliceToForm("ControlPoints", instanceWithInferedType, &instanceWithInferedType.ControlPoints, formGroup, playground)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)

	case *models.LinkAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.Path:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.Point:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Polygone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.Polyline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.Rect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, playground)
		BasicFieldtoForm("IsSelectable", instanceWithInferedType.IsSelectable, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanHaveLeftHandle", instanceWithInferedType.CanHaveLeftHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasLeftHandle", instanceWithInferedType.HasLeftHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanHaveRightHandle", instanceWithInferedType.CanHaveRightHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasRightHandle", instanceWithInferedType.HasRightHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanHaveTopHandle", instanceWithInferedType.CanHaveTopHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasTopHandle", instanceWithInferedType.HasTopHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanHaveBottomHandle", instanceWithInferedType.CanHaveBottomHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasBottomHandle", instanceWithInferedType.HasBottomHandle, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanMoveHorizontaly", instanceWithInferedType.CanMoveHorizontaly, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CanMoveVerticaly", instanceWithInferedType.CanMoveVerticaly, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("RectAnchoredTexts", instanceWithInferedType, &instanceWithInferedType.RectAnchoredTexts, formGroup, playground)
		AssociationSliceToForm("RectAnchoredRects", instanceWithInferedType, &instanceWithInferedType.RectAnchoredRects, formGroup, playground)

	case *models.RectAnchoredRect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("WidthFollowRect", instanceWithInferedType.WidthFollowRect, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HeightFollowRect", instanceWithInferedType.HeightFollowRect, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)

	case *models.RectAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("TextAnchorType", instanceWithInferedType.TextAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	case *models.RectLinkLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, playground)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, playground)
		BasicFieldtoForm("TargetAnchorPosition", instanceWithInferedType.TargetAnchorPosition, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)

	case *models.SVG:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Layers", instanceWithInferedType, &instanceWithInferedType.Layers, formGroup, playground)
		EnumTypeStringToForm("DrawingState", instanceWithInferedType.DrawingState, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("StartRect", instanceWithInferedType.StartRect, formGroup, playground)
		AssociationFieldToForm("EndRect", instanceWithInferedType.EndRect, formGroup, playground)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}


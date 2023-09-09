// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

type ButtonImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	Icon               gongtree_buttons.ButtonType
	tableStage         *form.StageStruct
	formStage          *form.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	tableStage *form.StageStruct,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.tableStage = tableStage
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r
	buttonImplGongstruct.backRepoOfInterest = backRepoOfInterest

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Animate":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAnimateFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		animate := new(models.Animate)
		FillUpForm(animate, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewCircleFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		FillUpForm(circle, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Ellipse":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewEllipseFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLayerFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		layer := new(models.Layer)
		FillUpForm(layer, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Line":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLineFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		line := new(models.Line)
		FillUpForm(line, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Link":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "LinkAnchoredText":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkAnchoredTextFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Path":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPathFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		path := new(models.Path)
		FillUpForm(path, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Point":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPointFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		point := new(models.Point)
		FillUpForm(point, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Polygone":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPolygoneFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		polygone := new(models.Polygone)
		FillUpForm(polygone, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPolylineFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		polyline := new(models.Polyline)
		FillUpForm(polyline, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Rect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		rect := new(models.Rect)
		FillUpForm(rect, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "RectAnchoredRect":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredRectFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "RectAnchoredText":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredTextFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "RectLinkLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewRectLinkLinkFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "SVG":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewSVGFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		svg := new(models.SVG)
		FillUpForm(svg, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Text":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewTextFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		text := new(models.Text)
		FillUpForm(text, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	stageOfInterest *models.StageStruct,
	formStage *form.StageStruct,
	formGroup *form.FormGroup,
	r *gin.Engine,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Animate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("AttributeName", instanceWithInferedType.AttributeName, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Values", instanceWithInferedType.Values, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Dur", instanceWithInferedType.Dur, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("RepeatCount", instanceWithInferedType.RepeatCount, instanceWithInferedType, formStage, formGroup)

	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, stageOfInterest, formStage, formGroup, r)

	case *models.Ellipse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("RY", instanceWithInferedType.RY, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Rects", instanceWithInferedType, &instanceWithInferedType.Rects, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Texts", instanceWithInferedType, &instanceWithInferedType.Texts, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Circles", instanceWithInferedType, &instanceWithInferedType.Circles, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Lines", instanceWithInferedType, &instanceWithInferedType.Lines, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Ellipses", instanceWithInferedType, &instanceWithInferedType.Ellipses, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Polylines", instanceWithInferedType, &instanceWithInferedType.Polylines, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Polygones", instanceWithInferedType, &instanceWithInferedType.Polygones, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Paths", instanceWithInferedType, &instanceWithInferedType.Paths, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("RectLinkLinks", instanceWithInferedType, &instanceWithInferedType.RectLinkLinks, stageOfInterest, formStage, formGroup, r)

	case *models.Line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X1", instanceWithInferedType.X1, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y1", instanceWithInferedType.Y1, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X2", instanceWithInferedType.X2, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y2", instanceWithInferedType.Y2, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("MouseClickX", instanceWithInferedType.MouseClickX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("MouseClickY", instanceWithInferedType.MouseClickY, instanceWithInferedType, formStage, formGroup)

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, stageOfInterest, formStage, formGroup)
		EnumTypeStringToForm("StartAnchorType", instanceWithInferedType.StartAnchorType, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("End", instanceWithInferedType.End, stageOfInterest, formStage, formGroup)
		EnumTypeStringToForm("EndAnchorType", instanceWithInferedType.EndAnchorType, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CornerRadius", instanceWithInferedType.CornerRadius, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasEndArrow", instanceWithInferedType.HasEndArrow, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EndArrowSize", instanceWithInferedType.EndArrowSize, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("TextAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.TextAtArrowEnd, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("TextAtArrowStart", instanceWithInferedType, &instanceWithInferedType.TextAtArrowStart, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("ControlPoints", instanceWithInferedType, &instanceWithInferedType.ControlPoints, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)

	case *models.LinkAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.Path:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.Point:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)

	case *models.Polygone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.Polyline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.Rect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("IsSelectable", instanceWithInferedType.IsSelectable, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanHaveLeftHandle", instanceWithInferedType.CanHaveLeftHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasLeftHandle", instanceWithInferedType.HasLeftHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanHaveRightHandle", instanceWithInferedType.CanHaveRightHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasRightHandle", instanceWithInferedType.HasRightHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanHaveTopHandle", instanceWithInferedType.CanHaveTopHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasTopHandle", instanceWithInferedType.HasTopHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanHaveBottomHandle", instanceWithInferedType.CanHaveBottomHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HasBottomHandle", instanceWithInferedType.HasBottomHandle, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanMoveHorizontaly", instanceWithInferedType.CanMoveHorizontaly, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CanMoveVerticaly", instanceWithInferedType.CanMoveVerticaly, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("RectAnchoredTexts", instanceWithInferedType, &instanceWithInferedType.RectAnchoredTexts, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("RectAnchoredRects", instanceWithInferedType, &instanceWithInferedType.RectAnchoredRects, stageOfInterest, formStage, formGroup, r)

	case *models.RectAnchoredRect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("WidthFollowRect", instanceWithInferedType.WidthFollowRect, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("HeightFollowRect", instanceWithInferedType.HeightFollowRect, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)

	case *models.RectAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("TextAnchorType", instanceWithInferedType.TextAnchorType, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	case *models.RectLinkLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("End", instanceWithInferedType.End, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("TargetAnchorPosition", instanceWithInferedType.TargetAnchorPosition, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)

	case *models.SVG:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Layers", instanceWithInferedType, &instanceWithInferedType.Layers, stageOfInterest, formStage, formGroup, r)
		EnumTypeStringToForm("DrawingState", instanceWithInferedType.DrawingState, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("StartRect", instanceWithInferedType.StartRect, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("EndRect", instanceWithInferedType.EndRect, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, formStage, formGroup)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, stageOfInterest, formStage, formGroup, r)

	default:
		_ = instanceWithInferedType
	}
}


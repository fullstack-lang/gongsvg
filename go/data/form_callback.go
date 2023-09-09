// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewAnimateFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	animate *models.Animate,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (animateFormCallback *AnimateFormCallback) {
	animateFormCallback = new(AnimateFormCallback)
	animateFormCallback.stageOfInterest = stageOfInterest
	animateFormCallback.tableStage = tableStage
	animateFormCallback.formStage = formStage
	animateFormCallback.animate = animate
	animateFormCallback.r = r
	animateFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type AnimateFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	animate            *models.Animate
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (animateFormCallback *AnimateFormCallback) OnSave() {

	log.Println("AnimateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	animateFormCallback.formStage.Checkout()

	if animateFormCallback.animate == nil {
		animateFormCallback.animate = new(models.Animate).Stage(animateFormCallback.stageOfInterest)
	}
	animate_ := animateFormCallback.animate
	_ = animate_

	// get the formGroup
	formGroup := animateFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(animate_.Name), formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&(animate_.AttributeName), formDiv)
		case "Values":
			FormDivBasicFieldToField(&(animate_.Values), formDiv)
		case "Dur":
			FormDivBasicFieldToField(&(animate_.Dur), formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&(animate_.RepeatCount), formDiv)
		}
	}

	animateFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Animate](
		animateFormCallback.stageOfInterest,
		animateFormCallback.tableStage,
		animateFormCallback.formStage,
		animateFormCallback.r,
		animateFormCallback.backRepoOfInterest,
	)
	animateFormCallback.tableStage.Commit()
}
func NewCircleFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	circle *models.Circle,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.stageOfInterest = stageOfInterest
	circleFormCallback.tableStage = tableStage
	circleFormCallback.formStage = formStage
	circleFormCallback.circle = circle
	circleFormCallback.r = r
	circleFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type CircleFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	circle            *models.Circle
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	// get the formGroup
	formGroup := circleFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(circle_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(circle_.CY), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(circle_.Radius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(circle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(circle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(circle_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(circle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(circle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(circle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(circle_.Transform), formDiv)
		}
	}

	circleFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.stageOfInterest,
		circleFormCallback.tableStage,
		circleFormCallback.formStage,
		circleFormCallback.r,
		circleFormCallback.backRepoOfInterest,
	)
	circleFormCallback.tableStage.Commit()
}
func NewEllipseFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	ellipse *models.Ellipse,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (ellipseFormCallback *EllipseFormCallback) {
	ellipseFormCallback = new(EllipseFormCallback)
	ellipseFormCallback.stageOfInterest = stageOfInterest
	ellipseFormCallback.tableStage = tableStage
	ellipseFormCallback.formStage = formStage
	ellipseFormCallback.ellipse = ellipse
	ellipseFormCallback.r = r
	ellipseFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type EllipseFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	ellipse            *models.Ellipse
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (ellipseFormCallback *EllipseFormCallback) OnSave() {

	log.Println("EllipseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ellipseFormCallback.formStage.Checkout()

	if ellipseFormCallback.ellipse == nil {
		ellipseFormCallback.ellipse = new(models.Ellipse).Stage(ellipseFormCallback.stageOfInterest)
	}
	ellipse_ := ellipseFormCallback.ellipse
	_ = ellipse_

	// get the formGroup
	formGroup := ellipseFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ellipse_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(ellipse_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(ellipse_.CY), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(ellipse_.RX), formDiv)
		case "RY":
			FormDivBasicFieldToField(&(ellipse_.RY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(ellipse_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(ellipse_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(ellipse_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(ellipse_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(ellipse_.Transform), formDiv)
		}
	}

	ellipseFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Ellipse](
		ellipseFormCallback.stageOfInterest,
		ellipseFormCallback.tableStage,
		ellipseFormCallback.formStage,
		ellipseFormCallback.r,
		ellipseFormCallback.backRepoOfInterest,
	)
	ellipseFormCallback.tableStage.Commit()
}
func NewLayerFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	layer *models.Layer,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (layerFormCallback *LayerFormCallback) {
	layerFormCallback = new(LayerFormCallback)
	layerFormCallback.stageOfInterest = stageOfInterest
	layerFormCallback.tableStage = tableStage
	layerFormCallback.formStage = formStage
	layerFormCallback.layer = layer
	layerFormCallback.r = r
	layerFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type LayerFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	layer            *models.Layer
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (layerFormCallback *LayerFormCallback) OnSave() {

	log.Println("LayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layerFormCallback.formStage.Checkout()

	if layerFormCallback.layer == nil {
		layerFormCallback.layer = new(models.Layer).Stage(layerFormCallback.stageOfInterest)
	}
	layer_ := layerFormCallback.layer
	_ = layer_

	// get the formGroup
	formGroup := layerFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Display":
			FormDivBasicFieldToField(&(layer_.Display), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(layer_.Name), formDiv)
		}
	}

	layerFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Layer](
		layerFormCallback.stageOfInterest,
		layerFormCallback.tableStage,
		layerFormCallback.formStage,
		layerFormCallback.r,
		layerFormCallback.backRepoOfInterest,
	)
	layerFormCallback.tableStage.Commit()
}
func NewLineFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	line *models.Line,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.stageOfInterest = stageOfInterest
	lineFormCallback.tableStage = tableStage
	lineFormCallback.formStage = formStage
	lineFormCallback.line = line
	lineFormCallback.r = r
	lineFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type LineFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	line            *models.Line
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (lineFormCallback *LineFormCallback) OnSave() {

	log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	// get the formGroup
	formGroup := lineFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(line_.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(line_.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(line_.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(line_.Y2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(line_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(line_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(line_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(line_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(line_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(line_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(line_.Transform), formDiv)
		case "MouseClickX":
			FormDivBasicFieldToField(&(line_.MouseClickX), formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&(line_.MouseClickY), formDiv)
		}
	}

	lineFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Line](
		lineFormCallback.stageOfInterest,
		lineFormCallback.tableStage,
		lineFormCallback.formStage,
		lineFormCallback.r,
		lineFormCallback.backRepoOfInterest,
	)
	lineFormCallback.tableStage.Commit()
}
func NewLinkFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	link *models.Link,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.stageOfInterest = stageOfInterest
	linkFormCallback.tableStage = tableStage
	linkFormCallback.formStage = formStage
	linkFormCallback.link = link
	linkFormCallback.r = r
	linkFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type LinkFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	link            *models.Link
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(link_.Type), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(link_.Start), linkFormCallback.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(link_.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(link_.End), linkFormCallback.stageOfInterest, formDiv)
		case "EndAnchorType":
			FormDivEnumStringFieldToField(&(link_.EndAnchorType), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&(link_.CornerRadius), formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&(link_.HasEndArrow), formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&(link_.EndArrowSize), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(link_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(link_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(link_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(link_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(link_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(link_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(link_.Transform), formDiv)
		}
	}

	linkFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.stageOfInterest,
		linkFormCallback.tableStage,
		linkFormCallback.formStage,
		linkFormCallback.r,
		linkFormCallback.backRepoOfInterest,
	)
	linkFormCallback.tableStage.Commit()
}
func NewLinkAnchoredTextFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	linkanchoredtext *models.LinkAnchoredText,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) {
	linkanchoredtextFormCallback = new(LinkAnchoredTextFormCallback)
	linkanchoredtextFormCallback.stageOfInterest = stageOfInterest
	linkanchoredtextFormCallback.tableStage = tableStage
	linkanchoredtextFormCallback.formStage = formStage
	linkanchoredtextFormCallback.linkanchoredtext = linkanchoredtext
	linkanchoredtextFormCallback.r = r
	linkanchoredtextFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type LinkAnchoredTextFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	linkanchoredtext            *models.LinkAnchoredText
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) OnSave() {

	log.Println("LinkAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredtextFormCallback.formStage.Checkout()

	if linkanchoredtextFormCallback.linkanchoredtext == nil {
		linkanchoredtextFormCallback.linkanchoredtext = new(models.LinkAnchoredText).Stage(linkanchoredtextFormCallback.stageOfInterest)
	}
	linkanchoredtext_ := linkanchoredtextFormCallback.linkanchoredtext
	_ = linkanchoredtext_

	// get the formGroup
	formGroup := linkanchoredtextFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(linkanchoredtext_.Content), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.Y_Offset), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontWeight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(linkanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(linkanchoredtext_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(linkanchoredtext_.Transform), formDiv)
		}
	}

	linkanchoredtextFormCallback.stageOfInterest.Commit()
	fillUpTable[models.LinkAnchoredText](
		linkanchoredtextFormCallback.stageOfInterest,
		linkanchoredtextFormCallback.tableStage,
		linkanchoredtextFormCallback.formStage,
		linkanchoredtextFormCallback.r,
		linkanchoredtextFormCallback.backRepoOfInterest,
	)
	linkanchoredtextFormCallback.tableStage.Commit()
}
func NewPathFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	path *models.Path,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (pathFormCallback *PathFormCallback) {
	pathFormCallback = new(PathFormCallback)
	pathFormCallback.stageOfInterest = stageOfInterest
	pathFormCallback.tableStage = tableStage
	pathFormCallback.formStage = formStage
	pathFormCallback.path = path
	pathFormCallback.r = r
	pathFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PathFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	path            *models.Path
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (pathFormCallback *PathFormCallback) OnSave() {

	log.Println("PathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pathFormCallback.formStage.Checkout()

	if pathFormCallback.path == nil {
		pathFormCallback.path = new(models.Path).Stage(pathFormCallback.stageOfInterest)
	}
	path_ := pathFormCallback.path
	_ = path_

	// get the formGroup
	formGroup := pathFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(path_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(path_.Definition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(path_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(path_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(path_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(path_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(path_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(path_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(path_.Transform), formDiv)
		}
	}

	pathFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Path](
		pathFormCallback.stageOfInterest,
		pathFormCallback.tableStage,
		pathFormCallback.formStage,
		pathFormCallback.r,
		pathFormCallback.backRepoOfInterest,
	)
	pathFormCallback.tableStage.Commit()
}
func NewPointFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	point *models.Point,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (pointFormCallback *PointFormCallback) {
	pointFormCallback = new(PointFormCallback)
	pointFormCallback.stageOfInterest = stageOfInterest
	pointFormCallback.tableStage = tableStage
	pointFormCallback.formStage = formStage
	pointFormCallback.point = point
	pointFormCallback.r = r
	pointFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PointFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	point            *models.Point
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (pointFormCallback *PointFormCallback) OnSave() {

	log.Println("PointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointFormCallback.formStage.Checkout()

	if pointFormCallback.point == nil {
		pointFormCallback.point = new(models.Point).Stage(pointFormCallback.stageOfInterest)
	}
	point_ := pointFormCallback.point
	_ = point_

	// get the formGroup
	formGroup := pointFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(point_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(point_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(point_.Y), formDiv)
		}
	}

	pointFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Point](
		pointFormCallback.stageOfInterest,
		pointFormCallback.tableStage,
		pointFormCallback.formStage,
		pointFormCallback.r,
		pointFormCallback.backRepoOfInterest,
	)
	pointFormCallback.tableStage.Commit()
}
func NewPolygoneFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	polygone *models.Polygone,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (polygoneFormCallback *PolygoneFormCallback) {
	polygoneFormCallback = new(PolygoneFormCallback)
	polygoneFormCallback.stageOfInterest = stageOfInterest
	polygoneFormCallback.tableStage = tableStage
	polygoneFormCallback.formStage = formStage
	polygoneFormCallback.polygone = polygone
	polygoneFormCallback.r = r
	polygoneFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PolygoneFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	polygone            *models.Polygone
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (polygoneFormCallback *PolygoneFormCallback) OnSave() {

	log.Println("PolygoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polygoneFormCallback.formStage.Checkout()

	if polygoneFormCallback.polygone == nil {
		polygoneFormCallback.polygone = new(models.Polygone).Stage(polygoneFormCallback.stageOfInterest)
	}
	polygone_ := polygoneFormCallback.polygone
	_ = polygone_

	// get the formGroup
	formGroup := polygoneFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polygone_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polygone_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polygone_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polygone_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polygone_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polygone_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polygone_.Transform), formDiv)
		}
	}

	polygoneFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Polygone](
		polygoneFormCallback.stageOfInterest,
		polygoneFormCallback.tableStage,
		polygoneFormCallback.formStage,
		polygoneFormCallback.r,
		polygoneFormCallback.backRepoOfInterest,
	)
	polygoneFormCallback.tableStage.Commit()
}
func NewPolylineFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	polyline *models.Polyline,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (polylineFormCallback *PolylineFormCallback) {
	polylineFormCallback = new(PolylineFormCallback)
	polylineFormCallback.stageOfInterest = stageOfInterest
	polylineFormCallback.tableStage = tableStage
	polylineFormCallback.formStage = formStage
	polylineFormCallback.polyline = polyline
	polylineFormCallback.r = r
	polylineFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PolylineFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	polyline            *models.Polyline
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (polylineFormCallback *PolylineFormCallback) OnSave() {

	log.Println("PolylineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polylineFormCallback.formStage.Checkout()

	if polylineFormCallback.polyline == nil {
		polylineFormCallback.polyline = new(models.Polyline).Stage(polylineFormCallback.stageOfInterest)
	}
	polyline_ := polylineFormCallback.polyline
	_ = polyline_

	// get the formGroup
	formGroup := polylineFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polyline_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polyline_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polyline_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polyline_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polyline_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polyline_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polyline_.Transform), formDiv)
		}
	}

	polylineFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Polyline](
		polylineFormCallback.stageOfInterest,
		polylineFormCallback.tableStage,
		polylineFormCallback.formStage,
		polylineFormCallback.r,
		polylineFormCallback.backRepoOfInterest,
	)
	polylineFormCallback.tableStage.Commit()
}
func NewRectFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	rect *models.Rect,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rectFormCallback *RectFormCallback) {
	rectFormCallback = new(RectFormCallback)
	rectFormCallback.stageOfInterest = stageOfInterest
	rectFormCallback.tableStage = tableStage
	rectFormCallback.formStage = formStage
	rectFormCallback.rect = rect
	rectFormCallback.r = r
	rectFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type RectFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	rect            *models.Rect
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rectFormCallback *RectFormCallback) OnSave() {

	log.Println("RectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectFormCallback.formStage.Checkout()

	if rectFormCallback.rect == nil {
		rectFormCallback.rect = new(models.Rect).Stage(rectFormCallback.stageOfInterest)
	}
	rect_ := rectFormCallback.rect
	_ = rect_

	// get the formGroup
	formGroup := rectFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rect_.RX), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rect_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rect_.Transform), formDiv)
		case "IsSelectable":
			FormDivBasicFieldToField(&(rect_.IsSelectable), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(rect_.IsSelected), formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveLeftHandle), formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&(rect_.HasLeftHandle), formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveRightHandle), formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&(rect_.HasRightHandle), formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveTopHandle), formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&(rect_.HasTopHandle), formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveBottomHandle), formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&(rect_.HasBottomHandle), formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&(rect_.CanMoveHorizontaly), formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&(rect_.CanMoveVerticaly), formDiv)
		}
	}

	rectFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Rect](
		rectFormCallback.stageOfInterest,
		rectFormCallback.tableStage,
		rectFormCallback.formStage,
		rectFormCallback.r,
		rectFormCallback.backRepoOfInterest,
	)
	rectFormCallback.tableStage.Commit()
}
func NewRectAnchoredRectFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	rectanchoredrect *models.RectAnchoredRect,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) {
	rectanchoredrectFormCallback = new(RectAnchoredRectFormCallback)
	rectanchoredrectFormCallback.stageOfInterest = stageOfInterest
	rectanchoredrectFormCallback.tableStage = tableStage
	rectanchoredrectFormCallback.formStage = formStage
	rectanchoredrectFormCallback.rectanchoredrect = rectanchoredrect
	rectanchoredrectFormCallback.r = r
	rectanchoredrectFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type RectAnchoredRectFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	rectanchoredrect            *models.RectAnchoredRect
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) OnSave() {

	log.Println("RectAnchoredRectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredrectFormCallback.formStage.Checkout()

	if rectanchoredrectFormCallback.rectanchoredrect == nil {
		rectanchoredrectFormCallback.rectanchoredrect = new(models.RectAnchoredRect).Stage(rectanchoredrectFormCallback.stageOfInterest)
	}
	rectanchoredrect_ := rectanchoredrectFormCallback.rectanchoredrect
	_ = rectanchoredrect_

	// get the formGroup
	formGroup := rectanchoredrectFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredrect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rectanchoredrect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rectanchoredrect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rectanchoredrect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rectanchoredrect_.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredrect_.RectAnchorType), formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.WidthFollowRect), formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.HeightFollowRect), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredrect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredrect_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredrect_.Transform), formDiv)
		}
	}

	rectanchoredrectFormCallback.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredRect](
		rectanchoredrectFormCallback.stageOfInterest,
		rectanchoredrectFormCallback.tableStage,
		rectanchoredrectFormCallback.formStage,
		rectanchoredrectFormCallback.r,
		rectanchoredrectFormCallback.backRepoOfInterest,
	)
	rectanchoredrectFormCallback.tableStage.Commit()
}
func NewRectAnchoredTextFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	rectanchoredtext *models.RectAnchoredText,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) {
	rectanchoredtextFormCallback = new(RectAnchoredTextFormCallback)
	rectanchoredtextFormCallback.stageOfInterest = stageOfInterest
	rectanchoredtextFormCallback.tableStage = tableStage
	rectanchoredtextFormCallback.formStage = formStage
	rectanchoredtextFormCallback.rectanchoredtext = rectanchoredtext
	rectanchoredtextFormCallback.r = r
	rectanchoredtextFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type RectAnchoredTextFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	rectanchoredtext            *models.RectAnchoredText
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) OnSave() {

	log.Println("RectAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredtextFormCallback.formStage.Checkout()

	if rectanchoredtextFormCallback.rectanchoredtext == nil {
		rectanchoredtextFormCallback.rectanchoredtext = new(models.RectAnchoredText).Stage(rectanchoredtextFormCallback.stageOfInterest)
	}
	rectanchoredtext_ := rectanchoredtextFormCallback.rectanchoredtext
	_ = rectanchoredtext_

	// get the formGroup
	formGroup := rectanchoredtextFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rectanchoredtext_.Content), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontSize), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.RectAnchorType), formDiv)
		case "TextAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.TextAnchorType), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredtext_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredtext_.Transform), formDiv)
		}
	}

	rectanchoredtextFormCallback.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredText](
		rectanchoredtextFormCallback.stageOfInterest,
		rectanchoredtextFormCallback.tableStage,
		rectanchoredtextFormCallback.formStage,
		rectanchoredtextFormCallback.r,
		rectanchoredtextFormCallback.backRepoOfInterest,
	)
	rectanchoredtextFormCallback.tableStage.Commit()
}
func NewRectLinkLinkFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	rectlinklink *models.RectLinkLink,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rectlinklinkFormCallback *RectLinkLinkFormCallback) {
	rectlinklinkFormCallback = new(RectLinkLinkFormCallback)
	rectlinklinkFormCallback.stageOfInterest = stageOfInterest
	rectlinklinkFormCallback.tableStage = tableStage
	rectlinklinkFormCallback.formStage = formStage
	rectlinklinkFormCallback.rectlinklink = rectlinklink
	rectlinklinkFormCallback.r = r
	rectlinklinkFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type RectLinkLinkFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	rectlinklink            *models.RectLinkLink
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rectlinklinkFormCallback *RectLinkLinkFormCallback) OnSave() {

	log.Println("RectLinkLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectlinklinkFormCallback.formStage.Checkout()

	if rectlinklinkFormCallback.rectlinklink == nil {
		rectlinklinkFormCallback.rectlinklink = new(models.RectLinkLink).Stage(rectlinklinkFormCallback.stageOfInterest)
	}
	rectlinklink_ := rectlinklinkFormCallback.rectlinklink
	_ = rectlinklink_

	// get the formGroup
	formGroup := rectlinklinkFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectlinklink_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(rectlinklink_.Start), rectlinklinkFormCallback.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(rectlinklink_.End), rectlinklinkFormCallback.stageOfInterest, formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&(rectlinklink_.TargetAnchorPosition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectlinklink_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectlinklink_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectlinklink_.Transform), formDiv)
		}
	}

	rectlinklinkFormCallback.stageOfInterest.Commit()
	fillUpTable[models.RectLinkLink](
		rectlinklinkFormCallback.stageOfInterest,
		rectlinklinkFormCallback.tableStage,
		rectlinklinkFormCallback.formStage,
		rectlinklinkFormCallback.r,
		rectlinklinkFormCallback.backRepoOfInterest,
	)
	rectlinklinkFormCallback.tableStage.Commit()
}
func NewSVGFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	svg *models.SVG,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (svgFormCallback *SVGFormCallback) {
	svgFormCallback = new(SVGFormCallback)
	svgFormCallback.stageOfInterest = stageOfInterest
	svgFormCallback.tableStage = tableStage
	svgFormCallback.formStage = formStage
	svgFormCallback.svg = svg
	svgFormCallback.r = r
	svgFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type SVGFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	svg            *models.SVG
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (svgFormCallback *SVGFormCallback) OnSave() {

	log.Println("SVGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.SVG).Stage(svgFormCallback.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	// get the formGroup
	formGroup := svgFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "DrawingState":
			FormDivEnumStringFieldToField(&(svg_.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(svg_.StartRect), svgFormCallback.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(svg_.EndRect), svgFormCallback.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(svg_.IsEditable), formDiv)
		}
	}

	svgFormCallback.stageOfInterest.Commit()
	fillUpTable[models.SVG](
		svgFormCallback.stageOfInterest,
		svgFormCallback.tableStage,
		svgFormCallback.formStage,
		svgFormCallback.r,
		svgFormCallback.backRepoOfInterest,
	)
	svgFormCallback.tableStage.Commit()
}
func NewTextFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	text *models.Text,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.stageOfInterest = stageOfInterest
	textFormCallback.tableStage = tableStage
	textFormCallback.formStage = formStage
	textFormCallback.text = text
	textFormCallback.r = r
	textFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type TextFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	text            *models.Text
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (textFormCallback *TextFormCallback) OnSave() {

	log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	// get the formGroup
	formGroup := textFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(text_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(text_.Y), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(text_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(text_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(text_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(text_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(text_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(text_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(text_.Transform), formDiv)
		}
	}

	textFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Text](
		textFormCallback.stageOfInterest,
		textFormCallback.tableStage,
		textFormCallback.formStage,
		textFormCallback.r,
		textFormCallback.backRepoOfInterest,
	)
	textFormCallback.tableStage.Commit()
}

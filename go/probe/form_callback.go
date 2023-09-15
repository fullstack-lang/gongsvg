// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewAnimateFormCallback(
	animate *models.Animate,
	playground *Playground,
) (animateFormCallback *AnimateFormCallback) {
	animateFormCallback = new(AnimateFormCallback)
	animateFormCallback.playground = playground
	animateFormCallback.animate = animate

	animateFormCallback.CreationMode = (animate == nil)

	return
}

type AnimateFormCallback struct {
	animate *models.Animate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (animateFormCallback *AnimateFormCallback) OnSave() {

	log.Println("AnimateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	animateFormCallback.playground.formStage.Checkout()

	if animateFormCallback.animate == nil {
		animateFormCallback.animate = new(models.Animate).Stage(animateFormCallback.playground.stageOfInterest)
	}
	animate_ := animateFormCallback.animate
	_ = animate_

	// get the formGroup
	formGroup := animateFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	animateFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Animate](
		animateFormCallback.playground,
	)
	animateFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if animateFormCallback.CreationMode {
		animateFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewAnimateFormCallback(
				nil,
				animateFormCallback.playground,
			),
		}).Stage(animateFormCallback.playground.formStage)
		animate := new(models.Animate)
		FillUpForm(animate, newFormGroup, animateFormCallback.playground)
		animateFormCallback.playground.formStage.Commit()
	}

}
func NewCircleFormCallback(
	circle *models.Circle,
	playground *Playground,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.playground = playground
	circleFormCallback.circle = circle

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.playground.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.playground.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	// get the formGroup
	formGroup := circleFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	circleFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.playground,
	)
	circleFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode {
		circleFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewCircleFormCallback(
				nil,
				circleFormCallback.playground,
			),
		}).Stage(circleFormCallback.playground.formStage)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.playground)
		circleFormCallback.playground.formStage.Commit()
	}

}
func NewEllipseFormCallback(
	ellipse *models.Ellipse,
	playground *Playground,
) (ellipseFormCallback *EllipseFormCallback) {
	ellipseFormCallback = new(EllipseFormCallback)
	ellipseFormCallback.playground = playground
	ellipseFormCallback.ellipse = ellipse

	ellipseFormCallback.CreationMode = (ellipse == nil)

	return
}

type EllipseFormCallback struct {
	ellipse *models.Ellipse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (ellipseFormCallback *EllipseFormCallback) OnSave() {

	log.Println("EllipseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ellipseFormCallback.playground.formStage.Checkout()

	if ellipseFormCallback.ellipse == nil {
		ellipseFormCallback.ellipse = new(models.Ellipse).Stage(ellipseFormCallback.playground.stageOfInterest)
	}
	ellipse_ := ellipseFormCallback.ellipse
	_ = ellipse_

	// get the formGroup
	formGroup := ellipseFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	ellipseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Ellipse](
		ellipseFormCallback.playground,
	)
	ellipseFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if ellipseFormCallback.CreationMode {
		ellipseFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewEllipseFormCallback(
				nil,
				ellipseFormCallback.playground,
			),
		}).Stage(ellipseFormCallback.playground.formStage)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, newFormGroup, ellipseFormCallback.playground)
		ellipseFormCallback.playground.formStage.Commit()
	}

}
func NewLayerFormCallback(
	layer *models.Layer,
	playground *Playground,
) (layerFormCallback *LayerFormCallback) {
	layerFormCallback = new(LayerFormCallback)
	layerFormCallback.playground = playground
	layerFormCallback.layer = layer

	layerFormCallback.CreationMode = (layer == nil)

	return
}

type LayerFormCallback struct {
	layer *models.Layer

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (layerFormCallback *LayerFormCallback) OnSave() {

	log.Println("LayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layerFormCallback.playground.formStage.Checkout()

	if layerFormCallback.layer == nil {
		layerFormCallback.layer = new(models.Layer).Stage(layerFormCallback.playground.stageOfInterest)
	}
	layer_ := layerFormCallback.layer
	_ = layer_

	// get the formGroup
	formGroup := layerFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Display":
			FormDivBasicFieldToField(&(layer_.Display), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(layer_.Name), formDiv)
		}
	}

	layerFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Layer](
		layerFormCallback.playground,
	)
	layerFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if layerFormCallback.CreationMode {
		layerFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewLayerFormCallback(
				nil,
				layerFormCallback.playground,
			),
		}).Stage(layerFormCallback.playground.formStage)
		layer := new(models.Layer)
		FillUpForm(layer, newFormGroup, layerFormCallback.playground)
		layerFormCallback.playground.formStage.Commit()
	}

}
func NewLineFormCallback(
	line *models.Line,
	playground *Playground,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.playground = playground
	lineFormCallback.line = line

	lineFormCallback.CreationMode = (line == nil)

	return
}

type LineFormCallback struct {
	line *models.Line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (lineFormCallback *LineFormCallback) OnSave() {

	log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.playground.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.playground.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	// get the formGroup
	formGroup := lineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	lineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Line](
		lineFormCallback.playground,
	)
	lineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if lineFormCallback.CreationMode {
		lineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewLineFormCallback(
				nil,
				lineFormCallback.playground,
			),
		}).Stage(lineFormCallback.playground.formStage)
		line := new(models.Line)
		FillUpForm(line, newFormGroup, lineFormCallback.playground)
		lineFormCallback.playground.formStage.Commit()
	}

}
func NewLinkFormCallback(
	link *models.Link,
	playground *Playground,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.playground = playground
	linkFormCallback.link = link

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.playground.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.playground.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(link_.Type), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(link_.Start), linkFormCallback.playground.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(link_.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(link_.End), linkFormCallback.playground.stageOfInterest, formDiv)
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

	linkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.playground,
	)
	linkFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode {
		linkFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				nil,
				linkFormCallback.playground,
			),
		}).Stage(linkFormCallback.playground.formStage)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.playground)
		linkFormCallback.playground.formStage.Commit()
	}

}
func NewLinkAnchoredTextFormCallback(
	linkanchoredtext *models.LinkAnchoredText,
	playground *Playground,
) (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) {
	linkanchoredtextFormCallback = new(LinkAnchoredTextFormCallback)
	linkanchoredtextFormCallback.playground = playground
	linkanchoredtextFormCallback.linkanchoredtext = linkanchoredtext

	linkanchoredtextFormCallback.CreationMode = (linkanchoredtext == nil)

	return
}

type LinkAnchoredTextFormCallback struct {
	linkanchoredtext *models.LinkAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) OnSave() {

	log.Println("LinkAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredtextFormCallback.playground.formStage.Checkout()

	if linkanchoredtextFormCallback.linkanchoredtext == nil {
		linkanchoredtextFormCallback.linkanchoredtext = new(models.LinkAnchoredText).Stage(linkanchoredtextFormCallback.playground.stageOfInterest)
	}
	linkanchoredtext_ := linkanchoredtextFormCallback.linkanchoredtext
	_ = linkanchoredtext_

	// get the formGroup
	formGroup := linkanchoredtextFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	linkanchoredtextFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LinkAnchoredText](
		linkanchoredtextFormCallback.playground,
	)
	linkanchoredtextFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if linkanchoredtextFormCallback.CreationMode {
		linkanchoredtextFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewLinkAnchoredTextFormCallback(
				nil,
				linkanchoredtextFormCallback.playground,
			),
		}).Stage(linkanchoredtextFormCallback.playground.formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, newFormGroup, linkanchoredtextFormCallback.playground)
		linkanchoredtextFormCallback.playground.formStage.Commit()
	}

}
func NewPathFormCallback(
	path *models.Path,
	playground *Playground,
) (pathFormCallback *PathFormCallback) {
	pathFormCallback = new(PathFormCallback)
	pathFormCallback.playground = playground
	pathFormCallback.path = path

	pathFormCallback.CreationMode = (path == nil)

	return
}

type PathFormCallback struct {
	path *models.Path

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pathFormCallback *PathFormCallback) OnSave() {

	log.Println("PathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pathFormCallback.playground.formStage.Checkout()

	if pathFormCallback.path == nil {
		pathFormCallback.path = new(models.Path).Stage(pathFormCallback.playground.stageOfInterest)
	}
	path_ := pathFormCallback.path
	_ = path_

	// get the formGroup
	formGroup := pathFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	pathFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Path](
		pathFormCallback.playground,
	)
	pathFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pathFormCallback.CreationMode {
		pathFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewPathFormCallback(
				nil,
				pathFormCallback.playground,
			),
		}).Stage(pathFormCallback.playground.formStage)
		path := new(models.Path)
		FillUpForm(path, newFormGroup, pathFormCallback.playground)
		pathFormCallback.playground.formStage.Commit()
	}

}
func NewPointFormCallback(
	point *models.Point,
	playground *Playground,
) (pointFormCallback *PointFormCallback) {
	pointFormCallback = new(PointFormCallback)
	pointFormCallback.playground = playground
	pointFormCallback.point = point

	pointFormCallback.CreationMode = (point == nil)

	return
}

type PointFormCallback struct {
	point *models.Point

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pointFormCallback *PointFormCallback) OnSave() {

	log.Println("PointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointFormCallback.playground.formStage.Checkout()

	if pointFormCallback.point == nil {
		pointFormCallback.point = new(models.Point).Stage(pointFormCallback.playground.stageOfInterest)
	}
	point_ := pointFormCallback.point
	_ = point_

	// get the formGroup
	formGroup := pointFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	pointFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Point](
		pointFormCallback.playground,
	)
	pointFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pointFormCallback.CreationMode {
		pointFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewPointFormCallback(
				nil,
				pointFormCallback.playground,
			),
		}).Stage(pointFormCallback.playground.formStage)
		point := new(models.Point)
		FillUpForm(point, newFormGroup, pointFormCallback.playground)
		pointFormCallback.playground.formStage.Commit()
	}

}
func NewPolygoneFormCallback(
	polygone *models.Polygone,
	playground *Playground,
) (polygoneFormCallback *PolygoneFormCallback) {
	polygoneFormCallback = new(PolygoneFormCallback)
	polygoneFormCallback.playground = playground
	polygoneFormCallback.polygone = polygone

	polygoneFormCallback.CreationMode = (polygone == nil)

	return
}

type PolygoneFormCallback struct {
	polygone *models.Polygone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (polygoneFormCallback *PolygoneFormCallback) OnSave() {

	log.Println("PolygoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polygoneFormCallback.playground.formStage.Checkout()

	if polygoneFormCallback.polygone == nil {
		polygoneFormCallback.polygone = new(models.Polygone).Stage(polygoneFormCallback.playground.stageOfInterest)
	}
	polygone_ := polygoneFormCallback.polygone
	_ = polygone_

	// get the formGroup
	formGroup := polygoneFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	polygoneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Polygone](
		polygoneFormCallback.playground,
	)
	polygoneFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if polygoneFormCallback.CreationMode {
		polygoneFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewPolygoneFormCallback(
				nil,
				polygoneFormCallback.playground,
			),
		}).Stage(polygoneFormCallback.playground.formStage)
		polygone := new(models.Polygone)
		FillUpForm(polygone, newFormGroup, polygoneFormCallback.playground)
		polygoneFormCallback.playground.formStage.Commit()
	}

}
func NewPolylineFormCallback(
	polyline *models.Polyline,
	playground *Playground,
) (polylineFormCallback *PolylineFormCallback) {
	polylineFormCallback = new(PolylineFormCallback)
	polylineFormCallback.playground = playground
	polylineFormCallback.polyline = polyline

	polylineFormCallback.CreationMode = (polyline == nil)

	return
}

type PolylineFormCallback struct {
	polyline *models.Polyline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (polylineFormCallback *PolylineFormCallback) OnSave() {

	log.Println("PolylineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polylineFormCallback.playground.formStage.Checkout()

	if polylineFormCallback.polyline == nil {
		polylineFormCallback.polyline = new(models.Polyline).Stage(polylineFormCallback.playground.stageOfInterest)
	}
	polyline_ := polylineFormCallback.polyline
	_ = polyline_

	// get the formGroup
	formGroup := polylineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	polylineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Polyline](
		polylineFormCallback.playground,
	)
	polylineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if polylineFormCallback.CreationMode {
		polylineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewPolylineFormCallback(
				nil,
				polylineFormCallback.playground,
			),
		}).Stage(polylineFormCallback.playground.formStage)
		polyline := new(models.Polyline)
		FillUpForm(polyline, newFormGroup, polylineFormCallback.playground)
		polylineFormCallback.playground.formStage.Commit()
	}

}
func NewRectFormCallback(
	rect *models.Rect,
	playground *Playground,
) (rectFormCallback *RectFormCallback) {
	rectFormCallback = new(RectFormCallback)
	rectFormCallback.playground = playground
	rectFormCallback.rect = rect

	rectFormCallback.CreationMode = (rect == nil)

	return
}

type RectFormCallback struct {
	rect *models.Rect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectFormCallback *RectFormCallback) OnSave() {

	log.Println("RectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectFormCallback.playground.formStage.Checkout()

	if rectFormCallback.rect == nil {
		rectFormCallback.rect = new(models.Rect).Stage(rectFormCallback.playground.stageOfInterest)
	}
	rect_ := rectFormCallback.rect
	_ = rect_

	// get the formGroup
	formGroup := rectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	rectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Rect](
		rectFormCallback.playground,
	)
	rectFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectFormCallback.CreationMode {
		rectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewRectFormCallback(
				nil,
				rectFormCallback.playground,
			),
		}).Stage(rectFormCallback.playground.formStage)
		rect := new(models.Rect)
		FillUpForm(rect, newFormGroup, rectFormCallback.playground)
		rectFormCallback.playground.formStage.Commit()
	}

}
func NewRectAnchoredRectFormCallback(
	rectanchoredrect *models.RectAnchoredRect,
	playground *Playground,
) (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) {
	rectanchoredrectFormCallback = new(RectAnchoredRectFormCallback)
	rectanchoredrectFormCallback.playground = playground
	rectanchoredrectFormCallback.rectanchoredrect = rectanchoredrect

	rectanchoredrectFormCallback.CreationMode = (rectanchoredrect == nil)

	return
}

type RectAnchoredRectFormCallback struct {
	rectanchoredrect *models.RectAnchoredRect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) OnSave() {

	log.Println("RectAnchoredRectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredrectFormCallback.playground.formStage.Checkout()

	if rectanchoredrectFormCallback.rectanchoredrect == nil {
		rectanchoredrectFormCallback.rectanchoredrect = new(models.RectAnchoredRect).Stage(rectanchoredrectFormCallback.playground.stageOfInterest)
	}
	rectanchoredrect_ := rectanchoredrectFormCallback.rectanchoredrect
	_ = rectanchoredrect_

	// get the formGroup
	formGroup := rectanchoredrectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	rectanchoredrectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredRect](
		rectanchoredrectFormCallback.playground,
	)
	rectanchoredrectFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredrectFormCallback.CreationMode {
		rectanchoredrectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredRectFormCallback(
				nil,
				rectanchoredrectFormCallback.playground,
			),
		}).Stage(rectanchoredrectFormCallback.playground.formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, newFormGroup, rectanchoredrectFormCallback.playground)
		rectanchoredrectFormCallback.playground.formStage.Commit()
	}

}
func NewRectAnchoredTextFormCallback(
	rectanchoredtext *models.RectAnchoredText,
	playground *Playground,
) (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) {
	rectanchoredtextFormCallback = new(RectAnchoredTextFormCallback)
	rectanchoredtextFormCallback.playground = playground
	rectanchoredtextFormCallback.rectanchoredtext = rectanchoredtext

	rectanchoredtextFormCallback.CreationMode = (rectanchoredtext == nil)

	return
}

type RectAnchoredTextFormCallback struct {
	rectanchoredtext *models.RectAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) OnSave() {

	log.Println("RectAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredtextFormCallback.playground.formStage.Checkout()

	if rectanchoredtextFormCallback.rectanchoredtext == nil {
		rectanchoredtextFormCallback.rectanchoredtext = new(models.RectAnchoredText).Stage(rectanchoredtextFormCallback.playground.stageOfInterest)
	}
	rectanchoredtext_ := rectanchoredtextFormCallback.rectanchoredtext
	_ = rectanchoredtext_

	// get the formGroup
	formGroup := rectanchoredtextFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	rectanchoredtextFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredText](
		rectanchoredtextFormCallback.playground,
	)
	rectanchoredtextFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredtextFormCallback.CreationMode {
		rectanchoredtextFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredTextFormCallback(
				nil,
				rectanchoredtextFormCallback.playground,
			),
		}).Stage(rectanchoredtextFormCallback.playground.formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, newFormGroup, rectanchoredtextFormCallback.playground)
		rectanchoredtextFormCallback.playground.formStage.Commit()
	}

}
func NewRectLinkLinkFormCallback(
	rectlinklink *models.RectLinkLink,
	playground *Playground,
) (rectlinklinkFormCallback *RectLinkLinkFormCallback) {
	rectlinklinkFormCallback = new(RectLinkLinkFormCallback)
	rectlinklinkFormCallback.playground = playground
	rectlinklinkFormCallback.rectlinklink = rectlinklink

	rectlinklinkFormCallback.CreationMode = (rectlinklink == nil)

	return
}

type RectLinkLinkFormCallback struct {
	rectlinklink *models.RectLinkLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectlinklinkFormCallback *RectLinkLinkFormCallback) OnSave() {

	log.Println("RectLinkLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectlinklinkFormCallback.playground.formStage.Checkout()

	if rectlinklinkFormCallback.rectlinklink == nil {
		rectlinklinkFormCallback.rectlinklink = new(models.RectLinkLink).Stage(rectlinklinkFormCallback.playground.stageOfInterest)
	}
	rectlinklink_ := rectlinklinkFormCallback.rectlinklink
	_ = rectlinklink_

	// get the formGroup
	formGroup := rectlinklinkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectlinklink_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(rectlinklink_.Start), rectlinklinkFormCallback.playground.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(rectlinklink_.End), rectlinklinkFormCallback.playground.stageOfInterest, formDiv)
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

	rectlinklinkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectLinkLink](
		rectlinklinkFormCallback.playground,
	)
	rectlinklinkFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectlinklinkFormCallback.CreationMode {
		rectlinklinkFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewRectLinkLinkFormCallback(
				nil,
				rectlinklinkFormCallback.playground,
			),
		}).Stage(rectlinklinkFormCallback.playground.formStage)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, newFormGroup, rectlinklinkFormCallback.playground)
		rectlinklinkFormCallback.playground.formStage.Commit()
	}

}
func NewSVGFormCallback(
	svg *models.SVG,
	playground *Playground,
) (svgFormCallback *SVGFormCallback) {
	svgFormCallback = new(SVGFormCallback)
	svgFormCallback.playground = playground
	svgFormCallback.svg = svg

	svgFormCallback.CreationMode = (svg == nil)

	return
}

type SVGFormCallback struct {
	svg *models.SVG

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (svgFormCallback *SVGFormCallback) OnSave() {

	log.Println("SVGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.playground.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.SVG).Stage(svgFormCallback.playground.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	// get the formGroup
	formGroup := svgFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "DrawingState":
			FormDivEnumStringFieldToField(&(svg_.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(svg_.StartRect), svgFormCallback.playground.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(svg_.EndRect), svgFormCallback.playground.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(svg_.IsEditable), formDiv)
		}
	}

	svgFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.SVG](
		svgFormCallback.playground,
	)
	svgFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if svgFormCallback.CreationMode {
		svgFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewSVGFormCallback(
				nil,
				svgFormCallback.playground,
			),
		}).Stage(svgFormCallback.playground.formStage)
		svg := new(models.SVG)
		FillUpForm(svg, newFormGroup, svgFormCallback.playground)
		svgFormCallback.playground.formStage.Commit()
	}

}
func NewTextFormCallback(
	text *models.Text,
	playground *Playground,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.playground = playground
	textFormCallback.text = text

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (textFormCallback *TextFormCallback) OnSave() {

	log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.playground.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.playground.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	// get the formGroup
	formGroup := textFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	textFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Text](
		textFormCallback.playground,
	)
	textFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if textFormCallback.CreationMode {
		textFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewTextFormCallback(
				nil,
				textFormCallback.playground,
			),
		}).Stage(textFormCallback.playground.formStage)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.playground)
		textFormCallback.playground.formStage.Commit()
	}

}

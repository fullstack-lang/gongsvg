// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Animate" {
		fillUpTable[models.Animate](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		fillUpTable[models.Circle](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ellipse" {
		fillUpTable[models.Ellipse](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Layer" {
		fillUpTable[models.Layer](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line" {
		fillUpTable[models.Line](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LinkAnchoredText" {
		fillUpTable[models.LinkAnchoredText](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Path" {
		fillUpTable[models.Path](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Point" {
		fillUpTable[models.Point](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polygone" {
		fillUpTable[models.Polygone](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polyline" {
		fillUpTable[models.Polyline](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rect" {
		fillUpTable[models.Rect](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredRect" {
		fillUpTable[models.RectAnchoredRect](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredText" {
		fillUpTable[models.RectAnchoredText](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectLinkLink" {
		fillUpTable[models.RectLinkLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SVG" {
		fillUpTable[models.SVG](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text" {
		fillUpTable[models.Text](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Animate:
		fillUpTable[models.Animate](playground)
	case *models.Circle:
		fillUpTable[models.Circle](playground)
	case *models.Ellipse:
		fillUpTable[models.Ellipse](playground)
	case *models.Layer:
		fillUpTable[models.Layer](playground)
	case *models.Line:
		fillUpTable[models.Line](playground)
	case *models.Link:
		fillUpTable[models.Link](playground)
	case *models.LinkAnchoredText:
		fillUpTable[models.LinkAnchoredText](playground)
	case *models.Path:
		fillUpTable[models.Path](playground)
	case *models.Point:
		fillUpTable[models.Point](playground)
	case *models.Polygone:
		fillUpTable[models.Polygone](playground)
	case *models.Polyline:
		fillUpTable[models.Polyline](playground)
	case *models.Rect:
		fillUpTable[models.Rect](playground)
	case *models.RectAnchoredRect:
		fillUpTable[models.RectAnchoredRect](playground)
	case *models.RectAnchoredText:
		fillUpTable[models.RectAnchoredText](playground)
	case *models.RectLinkLink:
		fillUpTable[models.RectLinkLink](playground)
	case *models.SVG:
		fillUpTable[models.SVG](playground)
	case *models.Text:
		fillUpTable[models.Text](playground)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	playground *Playground,
) {

	playground.tableStage.Reset()
	playground.tableStage.Commit()

	table := new(gongtable.Table).Stage(playground.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	playground.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](playground.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			playground.stageOfInterest,
			playground.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(playground.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, playground)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				structInstance,
			),
		}).Stage(playground.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(playground.tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(playground.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(playground.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.playground = playground
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Animate:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Animate Form",
			OnSave: NewAnimateFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Circle Form",
			OnSave: NewCircleFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Ellipse:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Ellipse Form",
			OnSave: NewEllipseFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Layer:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Layer Form",
			OnSave: NewLayerFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Line:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Line Form",
			OnSave: NewLineFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Link Form",
			OnSave: NewLinkFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.LinkAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update LinkAnchoredText Form",
			OnSave: NewLinkAnchoredTextFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Path:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Path Form",
			OnSave: NewPathFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Point:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Point Form",
			OnSave: NewPointFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Polygone:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Polygone Form",
			OnSave: NewPolygoneFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Polyline:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Polyline Form",
			OnSave: NewPolylineFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Rect:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Rect Form",
			OnSave: NewRectFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.RectAnchoredRect:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update RectAnchoredRect Form",
			OnSave: NewRectAnchoredRectFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.RectAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update RectAnchoredText Form",
			OnSave: NewRectAnchoredTextFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.RectLinkLink:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update RectLinkLink Form",
			OnSave: NewRectLinkLinkFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.SVG:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update SVG Form",
			OnSave: NewSVGFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Text:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Text Form",
			OnSave: NewTextFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	}
	formStage.Commit()

}

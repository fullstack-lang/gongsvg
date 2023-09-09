// generated code - do not edit
package data

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	r                  *gin.Engine
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct,
	r *gin.Engine,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.tableStage = tableStage
	nodeImplGongstruct.formStage = formStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	nodeImplGongstruct.backRepoOfInterest = backRepoOfInterest
	nodeImplGongstruct.r = r
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

	tableStage := nodeImplGongstruct.tableStage

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Animate" {
		fillUpTable[models.Animate](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		fillUpTable[models.Circle](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ellipse" {
		fillUpTable[models.Ellipse](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Layer" {
		fillUpTable[models.Layer](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line" {
		fillUpTable[models.Line](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LinkAnchoredText" {
		fillUpTable[models.LinkAnchoredText](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Path" {
		fillUpTable[models.Path](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Point" {
		fillUpTable[models.Point](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polygone" {
		fillUpTable[models.Polygone](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polyline" {
		fillUpTable[models.Polyline](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rect" {
		fillUpTable[models.Rect](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredRect" {
		fillUpTable[models.RectAnchoredRect](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredText" {
		fillUpTable[models.RectAnchoredText](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectLinkLink" {
		fillUpTable[models.RectLinkLink](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SVG" {
		fillUpTable[models.SVG](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text" {
		fillUpTable[models.Text](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	stageOfInterest *models.StageStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) {

	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			stageOfInterest,
			backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance,
			tableStage,
			formStage,
			stageOfInterest,
			r,
			backRepoOfInterest)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				structInstance,
			),
		}).Stage(tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.tableStage = tableStage
	rowUpdate.formStage = formStage
	rowUpdate.stageOfInterest = stageOfInterest
	rowUpdate.r = r
	rowUpdate.backRepoOfInterest = backRepoOfInterest
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance           *T
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Animate:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewAnimateFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewCircleFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Ellipse:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewEllipseFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Layer:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLayerFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Line:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLineFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.LinkAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLinkAnchoredTextFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Path:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPathFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Point:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPointFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Polygone:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPolygoneFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Polyline:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPolylineFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Rect:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewRectFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.RectAnchoredRect:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredRectFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.RectAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewRectAnchoredTextFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.RectLinkLink:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewRectLinkLinkFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.SVG:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewSVGFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Text:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewTextFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	}
	formStage.Commit()

}

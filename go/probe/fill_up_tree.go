package probe

import (
	"fmt"
	"sort"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	playground.treeStage.Reset()

	// create tree
	sidebar := (&gongtree_models.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&gongtree_models.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		switch gongStruct.Name {
		// insertion point
		case "Animate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Animate](playground.stageOfInterest)
			for animate := range set {
				nodeInstance := (&gongtree_models.Node{Name: animate.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(animate, "Animate", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Circle](playground.stageOfInterest)
			for circle := range set {
				nodeInstance := (&gongtree_models.Node{Name: circle.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(circle, "Circle", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Ellipse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Ellipse](playground.stageOfInterest)
			for ellipse := range set {
				nodeInstance := (&gongtree_models.Node{Name: ellipse.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(ellipse, "Ellipse", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Layer":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Layer](playground.stageOfInterest)
			for layer := range set {
				nodeInstance := (&gongtree_models.Node{Name: layer.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(layer, "Layer", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Line](playground.stageOfInterest)
			for line := range set {
				nodeInstance := (&gongtree_models.Node{Name: line.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(line, "Line", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](playground.stageOfInterest)
			for link := range set {
				nodeInstance := (&gongtree_models.Node{Name: link.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(link, "Link", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LinkAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LinkAnchoredText](playground.stageOfInterest)
			for linkanchoredtext := range set {
				nodeInstance := (&gongtree_models.Node{Name: linkanchoredtext.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(linkanchoredtext, "LinkAnchoredText", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Path":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Path](playground.stageOfInterest)
			for path := range set {
				nodeInstance := (&gongtree_models.Node{Name: path.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(path, "Path", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Point":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Point](playground.stageOfInterest)
			for point := range set {
				nodeInstance := (&gongtree_models.Node{Name: point.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(point, "Point", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polygone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Polygone](playground.stageOfInterest)
			for polygone := range set {
				nodeInstance := (&gongtree_models.Node{Name: polygone.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(polygone, "Polygone", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polyline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Polyline](playground.stageOfInterest)
			for polyline := range set {
				nodeInstance := (&gongtree_models.Node{Name: polyline.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(polyline, "Polyline", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Rect](playground.stageOfInterest)
			for rect := range set {
				nodeInstance := (&gongtree_models.Node{Name: rect.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(rect, "Rect", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredRect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectAnchoredRect](playground.stageOfInterest)
			for rectanchoredrect := range set {
				nodeInstance := (&gongtree_models.Node{Name: rectanchoredrect.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(rectanchoredrect, "RectAnchoredRect", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectAnchoredText](playground.stageOfInterest)
			for rectanchoredtext := range set {
				nodeInstance := (&gongtree_models.Node{Name: rectanchoredtext.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(rectanchoredtext, "RectAnchoredText", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectLinkLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectLinkLink](playground.stageOfInterest)
			for rectlinklink := range set {
				nodeInstance := (&gongtree_models.Node{Name: rectlinklink.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(rectlinklink, "RectLinkLink", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SVG":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SVG](playground.stageOfInterest)
			for svg := range set {
				nodeInstance := (&gongtree_models.Node{Name: svg.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(svg, "SVG", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Text](playground.stageOfInterest)
			for text := range set {
				nodeInstance := (&gongtree_models.Node{Name: text.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(text, "Text", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}

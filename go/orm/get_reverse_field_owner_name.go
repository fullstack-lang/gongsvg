// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongsvg/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Animate:
		tmp := GetInstanceDBFromInstance[models.Animate, AnimateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Animations":
			if tmp.Circle_AnimationsDBID.Int64 != 0 {
				id := uint(tmp.Circle_AnimationsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoCircle.Map_CircleDBID_CirclePtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Ellipse_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Ellipse_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoEllipse.Map_EllipseDBID_EllipsePtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Line_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Line_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLine.Map_LineDBID_LinePtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.LinkAnchoredText_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.LinkAnchoredText_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Path_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Path_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoPath.Map_PathDBID_PathPtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Polygone_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Polygone_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Polyline_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Polyline_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[id]
				res = reservePointerTarget.Name
			}
		case "Animations":
			if tmp.Rect_AnimationsDBID.Int64 != 0 {
				id := uint(tmp.Rect_AnimationsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.RectAnchoredText_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.RectAnchoredText_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[id]
				res = reservePointerTarget.Name
			}
		case "Animates":
			if tmp.Text_AnimatesDBID.Int64 != 0 {
				id := uint(tmp.Text_AnimatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoText.Map_TextDBID_TextPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Circles":
			if tmp.Layer_CirclesDBID.Int64 != 0 {
				id := uint(tmp.Layer_CirclesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Ellipse:
		tmp := GetInstanceDBFromInstance[models.Ellipse, EllipseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Ellipses":
			if tmp.Layer_EllipsesDBID.Int64 != 0 {
				id := uint(tmp.Layer_EllipsesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Layer:
		tmp := GetInstanceDBFromInstance[models.Layer, LayerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Layers":
			if tmp.SVG_LayersDBID.Int64 != 0 {
				id := uint(tmp.SVG_LayersDBID.Int64)
				reservePointerTarget := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Lines":
			if tmp.Layer_LinesDBID.Int64 != 0 {
				id := uint(tmp.Layer_LinesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Links":
			if tmp.Layer_LinksDBID.Int64 != 0 {
				id := uint(tmp.Layer_LinksDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.LinkAnchoredText:
		tmp := GetInstanceDBFromInstance[models.LinkAnchoredText, LinkAnchoredTextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "TextAtArrowEnd":
			if tmp.Link_TextAtArrowEndDBID.Int64 != 0 {
				id := uint(tmp.Link_TextAtArrowEndDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
				res = reservePointerTarget.Name
			}
		case "TextAtArrowStart":
			if tmp.Link_TextAtArrowStartDBID.Int64 != 0 {
				id := uint(tmp.Link_TextAtArrowStartDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Path:
		tmp := GetInstanceDBFromInstance[models.Path, PathDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Paths":
			if tmp.Layer_PathsDBID.Int64 != 0 {
				id := uint(tmp.Layer_PathsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Point:
		tmp := GetInstanceDBFromInstance[models.Point, PointDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "ControlPoints":
			if tmp.Link_ControlPointsDBID.Int64 != 0 {
				id := uint(tmp.Link_ControlPointsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Polygone:
		tmp := GetInstanceDBFromInstance[models.Polygone, PolygoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Polygones":
			if tmp.Layer_PolygonesDBID.Int64 != 0 {
				id := uint(tmp.Layer_PolygonesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Polyline:
		tmp := GetInstanceDBFromInstance[models.Polyline, PolylineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Polylines":
			if tmp.Layer_PolylinesDBID.Int64 != 0 {
				id := uint(tmp.Layer_PolylinesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Rect:
		tmp := GetInstanceDBFromInstance[models.Rect, RectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Rects":
			if tmp.Layer_RectsDBID.Int64 != 0 {
				id := uint(tmp.Layer_RectsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.RectAnchoredRect:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredRect, RectAnchoredRectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "RectAnchoredRects":
			if tmp.Rect_RectAnchoredRectsDBID.Int64 != 0 {
				id := uint(tmp.Rect_RectAnchoredRectsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.RectAnchoredText:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredText, RectAnchoredTextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "RectAnchoredTexts":
			if tmp.Rect_RectAnchoredTextsDBID.Int64 != 0 {
				id := uint(tmp.Rect_RectAnchoredTextsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.RectLinkLink:
		tmp := GetInstanceDBFromInstance[models.RectLinkLink, RectLinkLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "RectLinkLinks":
			if tmp.Layer_RectLinkLinksDBID.Int64 != 0 {
				id := uint(tmp.Layer_RectLinkLinksDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.SVG:
		tmp := GetInstanceDBFromInstance[models.SVG, SVGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Texts":
			if tmp.Layer_TextsDBID.Int64 != 0 {
				id := uint(tmp.Layer_TextsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
				res = reservePointerTarget.Name
			}
		}

	default:
		_ = inst
	}
	return
}

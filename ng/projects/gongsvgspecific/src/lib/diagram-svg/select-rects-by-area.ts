
import { SelectAreaConfig, SweepDirection } from "../svg-event.service";
import { DiagramSvgComponent } from "./diagram-svg.component";

import * as gongsvg from 'gongsvg'
import { getFunctionName } from "./get.function.name";

export function selectRectsByArea(diagramSvgComponent: DiagramSvgComponent) {
    let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

    if (diagramSvgComponent.PointAtMouseUp.X > diagramSvgComponent.PointAtMouseDown.X) {
        selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
    } else {
        selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
    }

    selectAreaConfig.TopLeft = [
        Math.min(diagramSvgComponent.PointAtMouseDown.X, diagramSvgComponent.PointAtMouseUp.X),
        Math.min(diagramSvgComponent.PointAtMouseDown.Y, diagramSvgComponent.PointAtMouseUp.Y)]
    selectAreaConfig.BottomRigth = [
        Math.max(diagramSvgComponent.PointAtMouseDown.X, diagramSvgComponent.PointAtMouseUp.X),
        Math.max(diagramSvgComponent.PointAtMouseDown.Y, diagramSvgComponent.PointAtMouseUp.Y)]

    for (let layer of diagramSvgComponent.gongsvgFrontRepo!.Layers_array) {
        for (let rect of layer.Rects) {
            switch (selectAreaConfig.SweepDirection) {
                case SweepDirection.LEFT_TO_RIGHT:
                    if (rect.X > selectAreaConfig.TopLeft[0] &&
                        rect.X + rect.Width < selectAreaConfig.BottomRigth[0] &&
                        rect.Y > selectAreaConfig.TopLeft[1] &&
                        rect.Y + rect.Height < selectAreaConfig.BottomRigth[1]) {
                        if (!rect.IsSelected) {
                            diagramSvgComponent.selectRect(rect);
                        }
                    }
                    break;
                case SweepDirection.RIGHT_TO_LEFT:
                    // rectangle has to be partialy boxed in the rect
                    if (rect.X < selectAreaConfig.BottomRigth[0] &&
                        rect.X + rect.Width > selectAreaConfig.TopLeft[0] &&
                        rect.Y < selectAreaConfig.BottomRigth[1] &&
                        rect.Y + rect.Height > selectAreaConfig.TopLeft[1]) {
                        console.log(getFunctionName(), "selecting rect", rect.Name);
                        if (!rect.IsSelected) {
                            console.log(getFunctionName(), "selecting rect", rect.Name);
                            rect.IsSelected = true;
                            diagramSvgComponent.manageHandles(rect);
                            diagramSvgComponent.rectService.updateRect(rect, diagramSvgComponent.GONG__StackPath, diagramSvgComponent.gongsvgFrontRepoService.frontRepo).subscribe(
                                _ => {
                                }
                            );
                        }
                    }
                    break;
            }
        }
    }
}
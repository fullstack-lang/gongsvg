

import { DiagramSvgComponent } from "./diagram-svg.component";

import * as gongsvg from 'gongsvg'

// informBackEndOfEndOfLinkDrawing
//
// informs the back end with 2 updates
// on the first update, the svg is updated with the state DRAWING_LINK
// on the second, the svg is updated with the state NOT_DRAWING_LINK
//
// the back ends shall interpret those calls in order to create the link between

// start end end rects
export function informBackEndOfEndOfLinkDrawing(diagramSvgComponent: DiagramSvgComponent) {
    diagramSvgComponent.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINK
    diagramSvgComponent.svgService.updateSVG(diagramSvgComponent.svg, diagramSvgComponent.GONG__StackPath, diagramSvgComponent.gongsvgFrontRepoService.frontRepo).subscribe(
        () => {

            diagramSvgComponent.gongsvgFrontRepoService.pull(diagramSvgComponent.GONG__StackPath).subscribe(
                gongsvgsFrontRepo => {
                    diagramSvgComponent.gongsvgFrontRepo = gongsvgsFrontRepo;

                    if (diagramSvgComponent.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
                        diagramSvgComponent.svg = diagramSvgComponent.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0];

                        // back to normal state
                        diagramSvgComponent.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINK;
                        diagramSvgComponent.svgService.updateSVG(diagramSvgComponent.svg, diagramSvgComponent.GONG__StackPath, diagramSvgComponent.gongsvgFrontRepoService.frontRepo).subscribe();

                        // set the isEditable
                        diagramSvgComponent.isEditableService.setIsEditable(diagramSvgComponent.svg!.IsEditable);
                    } else {
                        return;
                    }
                }
            );
        }
    );
}
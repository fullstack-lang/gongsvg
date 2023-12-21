import { Component, Input, OnDestroy, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'

import { manageHandles } from '../manage.handles';
import { Segment, drawSegments } from '../draw.segments';
import { getOrientation } from '../get.orientation';
import { getArcPath } from '../get.arc.path';
import { getEndArrowPath } from '../get.end.arrow.path';
import { swapSegment } from '../swap.segment';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { getStartPosition } from '../get.start.position';
import { getEndPosition } from '../get.end.position';
import { SvgEventService } from '../svg-event.service';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { Observable, timer } from 'rxjs';

@Component({
  selector: 'lib-diagram-svg',
  templateUrl: './diagram-svg.component.html',
  styleUrls: ['./diagram-svg.component.css']
})
export class DiagramSvgComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""

  // temporary, will be computed dynamicaly
  svgWidth = 3000
  svgHeight = 4000

  // for use in the template
  RectAnchorType = gongsvg.RectAnchorType
  LinkType = gongsvg.LinkType

  // svg is the singloton that is displayed. A svg
  // is the root of the directed acyclic graph containing
  // all svg elements
  svg = new gongsvg.SVGDB

  //
  // LINKS MANAGEMENT
  //
  // the components draws all svg elements directly.
  //
  // however, links of type LINK_TYPE_FLOATING_ORTHOGONAL are a set of line
  // in this case, each link is associated with a set of segment
  //
  map_Link_Segment: Map<gongsvg.LinkDB, Segment[]> = new (Map<gongsvg.LinkDB, Segment[]>)
  getSegments(link: gongsvg.LinkDB): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    return getOrientation(segment)
  }

  // Add this method to ArcComponent
  getArcPath(link: gongsvg.LinkDB, segment: Segment, nextSegment: Segment): string {
    return getArcPath(link, segment, nextSegment)
  }

  getEndArrowPath(link: gongsvg.LinkDB, segment: Segment, arrowSize: number): string {
    return getEndArrowPath(link, segment, arrowSize)
  }

  getStartArrowPath(link: gongsvg.LinkDB, segment: Segment, arrowSize: number): string {
    let inverseSegment = swapSegment(segment)
    let path = this.getEndArrowPath(link, inverseSegment, arrowSize)
    return path
  }

  resetAllLinksPreviousStartEndRects() {
    for (let link of this.gongsvgFrontRepo!.Links_array) {
      this.map_Link_PreviousStart.set(link, structuredClone(link.Start!))
      this.map_Link_PreviousEnd.set(link, structuredClone(link.End!))
    }
  }

  // for change detection, we need to store start and end rect of all links
  map_Link_PreviousStart: Map<gongsvg.LinkDB, gongsvg.RectDB> = new (Map<gongsvg.LinkDB, gongsvg.RectDB>)
  map_Link_PreviousEnd: Map<gongsvg.LinkDB, gongsvg.RectDB> = new (Map<gongsvg.LinkDB, gongsvg.RectDB>)

  //
  // RECT ANCHOR MANAGEMENT
  //

  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  rectDragging: boolean = false
  draggedRect: gongsvg.RectDB | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'

  //
  // RECT LINK LINK MANAGEMENT
  //
  // Elements for managing links between a rect and a link
  public getStartPosition(rectLinkLink: gongsvg.RectLinkLinkDB): Coordinate {
    return getStartPosition(rectLinkLink, this.map_Link_Segment)
  }

  public getEndPosition(rectLinkLink: gongsvg.RectLinkLinkDB): Coordinate {
    return getEndPosition(rectLinkLink, this.map_Link_Segment)
  }

  //
  // MULTI SHAPE SELECTION
  //
  selectionRectDrawing: boolean = false
  rectX = 100
  rectY = 100
  width = 300
  height = 40

  //
  // LINK BETWEEN RECT DRAWING
  //
  linkDrawing: boolean = false
  startX = 0
  startY = 0
  endX = 0
  endY = 0

  //
  // BACKEND MANAGEMENT
  //
  public gongsvgFrontRepo?: gongsvg.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private isEditableService: IsEditableService,

    private rectService: gongsvg.RectService,
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,

    private refreshService: RefreshService,
  ) { }

  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.refresh()
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack
        }
      }
    )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        if (this.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
          this.svg = this.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0]

          // set the isEditable
          this.isEditableService.setIsEditable(this.svg!.IsEditable)
        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        // compute segments for links
        this.map_Link_Segment.clear()

        for (let layer of this.gongsvgFrontRepo.Layers_array) {
          for (let link of layer.Links) {
            let segmentsParams = {
              StartRect: link.Start!,
              EndRect: link.End!,
              StartDirection: link.StartOrientation! as gongsvg.OrientationType,
              EndDirection: link.EndOrientation! as gongsvg.OrientationType,
              StartRatio: link.StartRatio,
              EndRatio: link.EndRatio,
              CornerOffsetRatio: link.CornerOffsetRatio,
              CornerRadius: link.CornerRadius,
            }

            let segments = drawSegments(segmentsParams)
            this.map_Link_Segment.set(link, segments)
          }
        }

        this.resetAllLinksPreviousStartEndRects()
      }
    )
  }

  ngOnDestroy() {

  }

  //
  // UTILITIES
  //

  //
  // USER INTERACTION MNGT
  //
  onmousemove(event: MouseEvent): void { }

  onmouseup(event: MouseEvent): void { }
  mousedown(event: MouseEvent): void { }

  rectMouseDown(event: MouseEvent, rect: gongsvg.RectDB): void { }
  rectMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void { }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void { }
  anchorMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void { }

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.LinkDB): void { }
  linkMouseUp(event: MouseEvent, link: gongsvg.LinkDB, segmentNumber: number = 0): void { }

  textAnchoredMouseDown(
    link: gongsvg.LinkDB,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void { }

  textAnchoredMouseUp(link: gongsvg.LinkDB, event: MouseEvent): void { }


  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.RectDB) {
    manageHandles(rect)
  }

}

import { ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';

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
import { StateEnumType } from './state.enum';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { getFunctionName } from './get.function.name';

@Component({
  selector: 'lib-diagram-svg',
  templateUrl: './diagram-svg.component.html',
  styleUrls: ['./diagram-svg.component.css']
})
export class DiagramSvgComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""

  //
  // state of the component
  //
  StateEnumType = StateEnumType
  State: StateEnumType = StateEnumType.NOT_EDITABLE


  // temporary, will be computed dynamicaly
  svgWidth = 3000
  svgHeight = 4000

  // svg is the singloton that is displayed. A svg
  // is the root of the directed acyclic graph containing
  // all svg elements
  svg = new gongsvg.SVGDB

  //
  // USER INTERACTION MANAGEMENT
  //
  PointAtMouseDown = new gongsvg.PointDB
  PointAtMouseMove = new gongsvg.PointDB
  PointAtMouseUp = new gongsvg.PointDB

  //
  // RECT MANAGEMENT
  //
  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }

  //
  // LINK MANAGEMENT
  //
  // the components draws all svg elements directly.
  //
  // however, links of type LINK_TYPE_FLOATING_ORTHOGONAL are a set of line
  // in this case, each link is associated with a set of segment
  //
  LinkType = gongsvg.LinkType
  map_Link_Segment: Map<gongsvg.LinkDB, Segment[]> = new (Map<gongsvg.LinkDB, Segment[]>)
  getSegments(link: gongsvg.LinkDB): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    return getOrientation(segment)
  }

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
  RectAnchorType = gongsvg.RectAnchorType
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  rectDragging: boolean = false
  draggedRect: gongsvg.RectDB | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.RectDB) {
    manageHandles(rect)
  }

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

    private changeDetectorRef: ChangeDetectorRef,
  ) { }

  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // this component is a "push mode component"
    // because the template calls many functions whose state cannot be computed
    // by the change detector ("dirty" or "clean").
    // therefore, the extra complexity is needed
    this.changeDetectorRef.detach()

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

          console.log(getFunctionName(), "state switch, before", this.State)
          if (this.isEditableService.getIsEditable()) {
            this.State = StateEnumType.WAITING_FOR_USER_INPUT
          } else {
            this.State = StateEnumType.NOT_EDITABLE
          }
          console.log(getFunctionName(), "state switch, current", this.State)

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

        // Manually trigger change detection
        this.changeDetectorRef.detectChanges()
      }
    )
  }

  ngOnDestroy() {

  }

  //
  // USER INTERACTION MNGT
  //

  //
  // computeShapeStates align shapes state on component state
  //
  // Shapes have states. For instance, Rect can be selected or not.
  // The State of shape must be conformed with the component state.
  computeShapeStates() {
    for (let layer of this.gongsvgFrontRepo!.Layers_array) {
      for (let rect of layer.Rects) {
        switch (this.State) {
          case StateEnumType.NOT_EDITABLE:
            rect.IsSelected = false
            break;
          case StateEnumType.WAITING_FOR_USER_INPUT:
          case StateEnumType.MULTI_RECTS_SELECTION:
          case StateEnumType.LINK_DRAWING:
            rect.IsSelected = false
            break;
          case StateEnumType.RECTS_DRAGGING:
          case StateEnumType.ANCHORS_DRAGGING:
          case StateEnumType.TEXT_ANCHOR_DRAGGING:
            rect.IsSelected = false
            break;
          case StateEnumType.LINK_DRAGGING:
            rect.IsSelected = false
            break;
        }
      }
    }
  }

  //
  // processGenericMouseUp performs all mouse up stuff
  processMouseUp() {
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.MULTI_RECTS_SELECTION) {
      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

  Math = Math

  onmousemove(event: MouseEvent, source?: string): void {
    this.PointAtMouseMove = mouseCoordInComponentRef(event)

    // case when the user releases the shift key
    if (this.State == StateEnumType.MULTI_RECTS_SELECTION && !event.shiftKey) {
      console.log(getFunctionName(), "end user release shift key")
      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    console.log(getFunctionName(), source)

    this.changeDetectorRef.detectChanges()
  }

  backgroundOnMouseDown(event: MouseEvent): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && event.shiftKey) {

      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.MULTI_RECTS_SELECTION
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }
  backgroundOnMouseUp(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp()
  }

  rectMouseDown(event: MouseEvent, rect: gongsvg.RectDB): void { }
  rectMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp()
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {

  }
  anchorMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp()
  }

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.LinkDB): void { }
  linkMouseUp(event: MouseEvent, link: gongsvg.LinkDB, segmentNumber: number = 0): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp()
  }

  textAnchoredMouseDown(
    link: gongsvg.LinkDB,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void { }

  textAnchoredMouseUp(link: gongsvg.LinkDB, event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

}

import { ChangeDetectorRef, Component, ElementRef, Input, OnDestroy, OnInit, QueryList, ViewChildren } from '@angular/core';

import * as gongsvg from 'gongsvg'

import { manageHandles } from '../manage.handles';
import { Segment, drawSegments, drawSegmentsFromLink } from '../draw.segments';
import { getOrientation } from '../get.orientation';
import { getArcPath } from '../get.arc.path';
import { getEndArrowPath } from '../get.end.arrow.path';
import { swapSegment } from '../swap.segment';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { getStartPosition } from '../get.start.position';
import { getEndPosition } from '../get.end.position';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { Observable, timer } from 'rxjs';
import { StateEnumType } from './state.enum';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { getFunctionName } from './get.function.name';
import { informBackEndOfEndOfLinkDrawing } from './inform-backend-end-of-link-drawing';
import { selectRectsByArea } from './select-rects-by-area';
import { LinkConf, computeLinkFromMouseEvent } from '../compute.link.from.mouse.event';
import { updateLinkFromCursor } from '../update.link.from.cursor';

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

  dragThreshold = 5

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

  draggedLink: gongsvg.LinkDB | undefined
  draggedSegmentNumber = 0

  //
  // RECT ANCHOR MANAGEMENT
  //
  RectAnchorType = gongsvg.RectAnchorType
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  draggedRect: gongsvg.RectDB | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.RectDB) {
    manageHandles(rect)
  }

  //
  // LINK ANCHORED TEXT MANAGEMENT
  //

  // the link whose anchored text is dragged
  draggedLinkAnchoredTextLink: gongsvg.LinkDB | undefined
  draggedTextIndex = 0
  draggedText: gongsvg.LinkAnchoredTextDB | undefined

  // LinkAtMouseDown is the clone of the Link when mouse down
  AnchoredTextAtMouseDown: gongsvg.LinkAnchoredTextDB | undefined

  // is the link anchored text at the start or the end of the arrows
  draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType = gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START

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
    public gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    public svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    public isEditableService: IsEditableService,

    public rectService: gongsvg.RectService,
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,

    private refreshService: RefreshService,

    private changeDetectorRef: ChangeDetectorRef,
  ) { }

  @ViewChildren('#background2') backgroundElement: QueryList<ElementRef> | undefined;

  @ViewChildren('svgRef') svgElements: QueryList<ElementRef> | undefined;


  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // this component is a "push mode component"
    // because the template calls many functions whose state cannot be computed
    // by the change detector ("dirty" or "clean").
    // therefore, the extra complexity is needed otherwise the template is perpetually
    // computed
    this.changeDetectorRef.detach()

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
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

        console.log("svg", this.backgroundElement?.length)
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
        let unselectRect = false
        switch (this.State) {
          case StateEnumType.NOT_EDITABLE:
            unselectRect = true
            break;
          case StateEnumType.WAITING_FOR_USER_INPUT:
            break;
          case StateEnumType.MULTI_RECTS_SELECTION:
            break;
          case StateEnumType.LINK_DRAWING:
            unselectRect = true
            break;
          case StateEnumType.RECTS_DRAGGING:
            break;
          case StateEnumType.ANCHORS_DRAGGING:
            break;
          case StateEnumType.LINK_ANCHORED_TEXT_DRAGGING:
            unselectRect = true
            break;
          case StateEnumType.LINK_DRAGGING:
            unselectRect = true
            break;
        }
        if (unselectRect && rect.IsSelected) {
          this.unselectRect(rect)
        }
      }
    }
  }

  //
  // processGenericMouseUp performs all mouse up stuff
  processMouseUp(event: MouseEvent) {
    console.log(getFunctionName(), "state at entry", this.State)

    // when the mouse has not moved more than a threshold
    // all rects are unselected
    let distanceMoved = this.Math.sqrt(
      (this.PointAtMouseDown.X - this.PointAtMouseUp.X) *
      (this.PointAtMouseDown.X - this.PointAtMouseUp.X) +
      (this.PointAtMouseDown.Y - this.PointAtMouseUp.Y) *
      (this.PointAtMouseDown.Y - this.PointAtMouseUp.Y))

    // the user clicks in the void to unselect all rect
    if (distanceMoved < this.dragThreshold && this.State == StateEnumType.WAITING_FOR_USER_INPUT) {
      console.log(getFunctionName(), "distanceMoved below threshold in state", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT

      this.unselectAllRects();
    }

    // the use clicks on a rect for selecting it if it is not selected or
    // unselecting it if it is already selected
    // if shift is pressed, all selected rect stay selected, otherwise, they are unselected
    if (distanceMoved < this.dragThreshold && this.State == StateEnumType.RECTS_DRAGGING) {
      console.log(getFunctionName(), "distanceMoved below threshold in state", this.State)

      // if the shift key is not pressed, unselect all other rects
      if (!event.shiftKey) {
        this.unselectAllRects()
      }

      console.assert(this.draggedRect != undefined, "no dragged rect")
      if (!this.draggedRect?.IsSelected) {
        this.selectRect(this.draggedRect!)
      } else {
        this.unselectRect(this.draggedRect!)
      }
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

    }

    if (this.State == StateEnumType.MULTI_RECTS_SELECTION) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

      selectRectsByArea(this)
    }

    if (this.State == StateEnumType.LINK_DRAWING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

      informBackEndOfEndOfLinkDrawing(this)
    }

    if (this.State == StateEnumType.LINK_ANCHORED_TEXT_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)

      console.assert(this.draggedText != undefined, "no dragged text")
      this.anchoredTextService.updateLinkAnchoredText(
        this.draggedText!, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()
    }

    if (this.State == StateEnumType.LINK_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
      this.linkService.updateLink(this.draggedLink!,
        this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()
      document.body.style.cursor = ''
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }



  Math = Math

  public selectRect(rect: gongsvg.RectDB) {
    console.log(getFunctionName(), "selecting rect", rect.Name);
    rect.IsSelected = true;
    this.manageHandles(rect);
    this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectRect(rect: gongsvg.RectDB) {
    console.log(getFunctionName(), "unselecting rect", rect.Name);
    rect.IsSelected = false;
    this.manageHandles(rect);
    this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectAllRects() {
    for (let layer of this.gongsvgFrontRepo!.Layers_array) {
      for (let rect of layer.Rects) {
        if (rect.IsSelected) {
          this.unselectRect(rect)
        }
      }
    }
  }

  onmousemove(event: MouseEvent, source?: string): void {
    this.PointAtMouseMove = mouseCoordInComponentRef(event)
    // console.log(getFunctionName(), source, event.buttons)

    // case when the user releases the shift key
    if (this.State == StateEnumType.MULTI_RECTS_SELECTION && !event.shiftKey) {
      console.log(getFunctionName(), "end user release shift key")
      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    if (this.State == StateEnumType.LINK_ANCHORED_TEXT_DRAGGING) {
      let deltaX = this.PointAtMouseMove.X - this.PointAtMouseDown!.X
      let deltaY = this.PointAtMouseMove.Y - this.PointAtMouseDown!.Y

      // console.log("gongsvg Text dragging, deltaX", deltaX, "deltaY", deltaY)

      if (this.draggedText == undefined) {
        console.log("Problem : this.draggedText should not be undefined")
        return
      }
      this.draggedText.X_Offset = this.AnchoredTextAtMouseDown!.X_Offset + deltaX
      this.draggedText.Y_Offset = this.AnchoredTextAtMouseDown!.Y_Offset + deltaY
    }

    if (this.State == StateEnumType.LINK_DRAGGING) {
      document.body.style.cursor = ''

      updateLinkFromCursor(
        this.draggedLink!,
        this.draggedSegmentNumber,
        this.map_Link_Segment.get(this.draggedLink!)!,
        this.map_Link_PreviousStart.get(this.draggedLink!)!,
        this.map_Link_PreviousEnd.get(this.draggedLink!)!,
        this.PointAtMouseDown,
        this.PointAtMouseMove,
      )

      let segments = drawSegmentsFromLink(this.draggedLink!)
      this.map_Link_Segment.set(this.draggedLink!, segments)
    }

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

  backgroundDragOver(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnClick(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnDragEnd(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnMouseUp(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  rectMouseDown(event: MouseEvent, rect: gongsvg.RectDB): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey) {
      this.State = StateEnumType.RECTS_DRAGGING
      this.draggedRect = rect
    }
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && event.altKey) {
      this.State = StateEnumType.LINK_DRAWING
      this.svg.StartRect = rect
    }

    console.log(getFunctionName(), "state at exit", this.State)
  }

  rectMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.LINK_DRAWING) {
      this.svg.EndRect = rect
    }
    this.processMouseUp(event)
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {

  }
  anchorMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.LinkDB): void {
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      // this link shit to dragging state
      this.draggedLink = link
      this.draggedSegmentNumber = segmentNumber
    }
  }
  linkMouseUp(event: MouseEvent, link: gongsvg.LinkDB, segmentNumber: number = 0): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  textAnchoredMouseDown(
    link: gongsvg.LinkDB,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_ANCHORED_TEXT_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      this.draggedTextIndex = anchoredTextIndex
      this.draggedSegmentPositionOnArrow = draggedSegmentPositionOnArrow as gongsvg.PositionOnArrowType
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        this.draggedText = link.TextAtArrowStart![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowStart[this.draggedTextIndex])

      }
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        this.draggedText = link.TextAtArrowEnd![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowEnd[this.draggedTextIndex])
      }
    }
  }

  textAnchoredMouseUp(link: gongsvg.LinkDB, event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

}
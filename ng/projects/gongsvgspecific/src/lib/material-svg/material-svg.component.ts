import { AfterViewInit, Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { ShapeMouseEvent } from '../shape.mouse.event';
import { MouseEventService } from '../mouse-event.service';
import { AngularDragEndEventService } from '../angular-drag-end-event.service';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { SizeTrackerService } from '../size-tracker.service';

import { SegmentsParams, Segment, createPoint, drawSegments, Offset } from '../draw.segments';


@Component({
  selector: 'lib-material-svg',
  templateUrl: './material-svg.component.html',
  styleUrls: ['./material-svg.component.css']
})
export class MaterialSvgComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

  // for use in the template
  RectAnchorType = gongsvg.RectAnchorType
  LinkType = gongsvg.LinkType

  // the components draws all elements directly but the links when they are 
  // LINK_TYPE_FLOATING_ORTHOGONAL, in this case, each link is associated with a 
  // set of segments
  // map_Link_Segments = new (Map < gongsvg.LineDB, []segment>)
  map_Link_Segment: Map<gongsvg.LinkDB, Segment[]> = new (Map<gongsvg.LinkDB, Segment[]>)
  getSegments(link: gongsvg.LinkDB): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  // temporary, will be computed dynamicaly
  svgWidth = 3000
  svgHeight = 4000
  private onResize = (): void => {
    this.updateSize();
  }
  rectDragging: boolean = false
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'

  // RectAtMouseDown is the clone of the Rect when mouse down
  private RectAtMouseDown: gongsvg.RectDB | undefined

  // to compute wether it was a select / dragging event
  distanceMoved = 0
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private dragThreshold = 5;

  private updateSize() {
    this.svgWidth = window.innerWidth;
    this.svgHeight = window.innerHeight;
    console.log(`New size: Width = ${this.width}, Height = ${this.height}`);

    this.refresh()
  }


  public gongsvgFrontRepo?: gongsvg.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  svg = new gongsvg.SVGDB
  linkStartRectangleID: number = 0

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];


  // if true, the end user is shiftKey + mouse down from one rectangle
  // to another
  linkDrawing: boolean = false
  startX = 0;
  startY = 0;
  endX = 0;
  endY = 0;

  selectionRectDrawing: boolean = false
  rectX = 100;
  rectY = 100;
  width = 300;
  height = 40;

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private mouseEventService: MouseEventService,
    private isEditableService: IsEditableService,
    private refreshRequestService: RefreshService,
    private sizeTracker: SizeTrackerService,
    private rectService: gongsvg.RectService,
    private refreshService: RefreshService,
  ) {

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseDownEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {
        console.log('SvgComponent, Alt Mouse down event occurred on rectangle ', shapeMouseEvent.ShapeID)
        this.linkStartRectangleID = shapeMouseEvent.ShapeID

        // refactorable of Rect name
        let rect = this.gongsvgFrontRepo?.getMap<gongsvg.RectDB>(gongsvg.RectDB.GONGSTRUCT_NAME).get(shapeMouseEvent.ShapeID)

        if (rect == undefined) {
          return
        }

        this.linkDrawing = true
        this.startX = shapeMouseEvent.Point.X
        this.startY = shapeMouseEvent.Point.Y
      }))

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseDragEvent$.subscribe((shapeMouseEvent: ShapeMouseEvent) => {

      this.endX = shapeMouseEvent.Point.X
      this.endY = shapeMouseEvent.Point.Y
      // console.log('SvgComponent, Received Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
    }))

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseUpEvent$.subscribe((rectangleID: number) => {
      console.log('SvgComponent, Mouse up event occurred on rectangle ', rectangleID);
      this.linkDrawing = false

      this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
    }))

    this.subscriptions.push(svgEventService.multiShapeSelectDragEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        let actualX = shapeMouseEvent.Point.X
        let actualY = shapeMouseEvent.Point.Y

        this.rectX = Math.min(this.startX, actualX);
        this.rectY = Math.min(this.startY, actualY);
        this.width = Math.abs(actualX - this.startX);
        this.height = Math.abs(actualY - this.startY);
      }))

    this.subscriptions.push(svgEventService.mouseShiftKeyMouseUpEvent$.subscribe(
      (shapeMouseEvent) => {

        this.selectionRectDrawing = false
        this.endX = shapeMouseEvent.Point.X
        this.endY = shapeMouseEvent.Point.Y

        let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

        if (this.endX > this.startX) {
          selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
          selectAreaConfig.TopLeft = [this.startX, this.startY]
          selectAreaConfig.BottomRigth = [this.endX, this.endY]
        } else {
          selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
          selectAreaConfig.TopLeft = [this.endX, this.endY]
          selectAreaConfig.BottomRigth = [this.startX, this.startY]
        }

        this.svgEventService.emitMultiShapeSelectEnd(selectAreaConfig)
      }))

    this.subscriptions.push(refreshRequestService.refreshRequest$.subscribe(
      _ => {
        this.refresh()
      }))

    this.subscriptions.push(
      mouseEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
          if (rect == undefined) {
            return
          }

          if (this.anchorDragging || this.rectDragging || rect.IsSelected) {

            console.log("rect: mouseMouseDownEvent, ", rect!.Name)

            this.distanceMoved = 0
            this.RectAtMouseDown = structuredClone(rect)
            this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)
          }
        }
      )
    )

    this.subscriptions.push(
      mouseEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
          if (rect == undefined) {
            return
          }

          if (!this.isEditableService.getIsEditable()) {
            return
          }

          if (this.anchorDragging) {
            if (this.activeAnchor === 'left') {
              rect.X = this.RectAtMouseDown!.X + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
              rect.Width = this.RectAtMouseDown!.Width - (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'right') {
              rect.Width = this.RectAtMouseDown!.Width + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'top') {
              rect.Y = this.RectAtMouseDown!.Y + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
              rect.Height = this.RectAtMouseDown!.Height - (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            } else if (this.activeAnchor === 'bottom') {
              rect.Height = this.RectAtMouseDown!.Height + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            }
            return // we don't want the move move to be interpreted by the rect
          }

          if (this.PointAtMouseDown && (this.rectDragging || rect.IsSelected)) {
            const deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
            const deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
            this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

            if (rect.CanMoveHorizontaly) {
              rect.X = this.RectAtMouseDown!.X + deltaX
            }
            if (rect.CanMoveVerticaly) {
              rect.Y = this.RectAtMouseDown!.Y + deltaY
            }
          }
        }
      )
    )

    this.subscriptions.push(mouseEventService.mouseMouseUpEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {
        let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
        if (rect == undefined) {
          return
        }

        if (!this.isEditableService.getIsEditable()) {
          return
        }

        if (shapeMouseEvent.ShapeID != 0 && this.distanceMoved > this.dragThreshold) {
          if (this.anchorDragging || this.rectDragging || rect.IsSelected) {
            rect.IsSelected = false
            this.manageHandles(rect)
            this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
              _ => {
                this.refreshService.emitRefreshRequestEvent(0)
              }
            )
          }

        }
        if (this.distanceMoved <= this.dragThreshold) {
          if (rect.IsSelectable &&
            shapeMouseEvent.ShapeType == gongsvg.RectDB.GONGSTRUCT_NAME &&
            shapeMouseEvent.ShapeID == rect.ID) {
            console.log("rect, mouseEventService.mouseMouseUpEvent$.subscribe, from the shape: ", rect?.Name)
            rect.IsSelected = !rect.IsSelected
            this.manageHandles(rect)
            this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
              _ => {
                this.refreshService.emitRefreshRequestEvent(0)
              }
            )
          }

          // mouseup emited from the background let to unselect selected shapes
          if (rect.IsSelectable && rect.IsSelected && shapeMouseEvent.ShapeID == 0) {
            console.log("rect, mouseEventService.mouseMouseUpEvent$.subscribe: from the svg", rect.Name)
            rect.IsSelected = false
            this.manageHandles(rect)
            this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
              _ => {
                this.refreshService.emitRefreshRequestEvent(0)
              }
            )
          }
        }

        this.rectDragging = false
        this.anchorDragging = false
        // console.log('Rect ', this.Rect.Name, 'Mouse down event occurred on rectangle ', rectangleID);
      })
    )
  }

  ngOnInit(): void {

    this.updateSize();
    window.addEventListener('resize', this.onResize);

    // console.log("Svg component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
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
      }

    )
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
    window.removeEventListener('resize', this.onResize);
  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    let svgArray = this.gongsvgFrontRepo?.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)
    // update the svg
    if (svgArray?.length == 1) {
      this.svg = svgArray[0]
    } else {
      return
    }

    if (this.svg.Layers == undefined) {
      return
    }

    if (this.svg.DrawingState != gongsvg.DrawingState.NOT_DRAWING_LINE) {
      // console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.NOT_DRAWING_LINE)
    }

    this.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINE

    let rectMap = this.gongsvgFrontRepo!.getMap<gongsvg.RectDB>(gongsvg.RectDB.GONGSTRUCT_NAME)

    this.svg.StartRect = rectMap.get(startRectangleID)
    this.svg.EndRect = rectMap.get(endRectangleID)

    this.svgService.updateSVG(this.svg, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
      () => {

        this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
          gongsvgsFrontRepo => {
            this.gongsvgFrontRepo = gongsvgsFrontRepo

            if (this.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
              this.svg = this.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0]

              // back to normal state
              this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE
              this.svgService.updateSVG(this.svg, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()

              // set the isEditable
              this.isEditableService.setIsEditable(this.svg!.IsEditable)
            } else {
              return
            }
          }
        )
      }
    )
  }

  mousedown(event: MouseEvent): void {
    if (event.shiftKey) {

      this.selectionRectDrawing = true

      let point = mouseCoordInComponentRef(event)

      this.startX = point.X
      this.startY = point.Y
    }
  }

  mousemove(event: MouseEvent): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: 0,
      ShapeType: "",
      Point: mouseCoordInComponentRef(event),
    }

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      // console.log('SvgComponent, ALT Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(shapeMouseEvent)
      return
    }
    if (event.shiftKey) {

      this.svgEventService.emitMultiShapeSelectDrag(shapeMouseEvent)
      // console.log('SvgComponent, SHIFT Mouse drag event occurred', this.selectionRectDrawing, this.rectX, this.rectY, this.width, this.height);
    }

    if (!event.shiftKey && !event.altKey) {
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
      // console.log("svg background, mouse move", x, y)
    }
  }


  onmouseup(event: MouseEvent): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: 0,
      ShapeType: "",
      Point: mouseCoordInComponentRef(event),
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent(shapeMouseEvent)
    }

    if (!event.shiftKey && !event.altKey) {
      // in case of dragging something, when the mouse moves fast, it can reach the SVG background
      // in this case, one forward the mouse event on the event bus
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  exportSVG() {
    const serializer = new XMLSerializer();
    const svgString = serializer.serializeToString(this.drawingArea!.nativeElement);
    return svgString;
  }

  downloadSVG() {
    const svgString = this.exportSVG();
    const blob = new Blob([svgString], { type: 'image/svg+xml' });
    const url = window.URL.createObjectURL(blob);

    // Create a link element
    const downloadLink = document.createElement('a');
    downloadLink.href = url;
    downloadLink.download = 'image.svg';

    // Attach the link to the document
    document.body.appendChild(downloadLink);
    downloadLink.click();

    // Clean up after to avoid memory leaks
    document.body.removeChild(downloadLink);
  }


  rectMouseDown(event: MouseEvent, rect: gongsvg.RectDB): void {
    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    } else {
      console.log('RectComponent, Alt Mouse down event occurred on rectangle ', rect.Name, rect.ID, event.clientX, event.clientY);

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.rectangleEventService.emitRectAltKeyMouseDownEvent(shapeMouseEvent)

    }
  }

  rectMouseMove(event: MouseEvent, rect: gongsvg.RectDB): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: rect.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    // console.log("RectComponent DragRect : ", deltaX, deltaY, distanceMoved)

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      // console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', rect.Name, event.clientX, event.clientY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(shapeMouseEvent)
      return
    }

    if (event.shiftKey) {
      this.svgEventService.emitMultiShapeSelectDrag(shapeMouseEvent)
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  rectMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {

    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: rect.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    if (!event.altKey && !event.shiftKey) {
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }

    if (event.altKey) {
      console.log('RectComponent, Alt Mouse up event occurred on rectangle ', rect.Name, rect.ID, event.clientX, event.clientY);

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(rect.ID);
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent(shapeMouseEvent)
    }
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  anchorMouseMove(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  anchorMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = 'bottom'
      rect.IsSelected = false
      this.manageHandles(rect)
      this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
        _ => {
          this.refreshService.emitRefreshRequestEvent(0)
        }
      )
    }
  }

  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.RectDB) {

    if (rect.IsSelected && rect.CanHaveLeftHandle) {
      rect.HasLeftHandle = true
    } else {
      rect.HasLeftHandle = false
    }
    if (rect.IsSelected && rect.CanHaveRightHandle) {
      rect.HasRightHandle = true
    } else {
      rect.HasRightHandle = false
    }
    if (rect.IsSelected && rect.CanHaveTopHandle) {
      rect.HasTopHandle = true
    } else {
      rect.HasTopHandle = false
    }
    if (rect.IsSelected && rect.CanHaveBottomHandle) {
      rect.HasBottomHandle = true
    } else {
      rect.HasBottomHandle = false
    }
  }
}

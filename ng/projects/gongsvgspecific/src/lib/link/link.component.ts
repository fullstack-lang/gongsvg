import { AfterViewInit, Component, ElementRef, Input, OnInit } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { ConnectorParams, Segment, createPoint, drawSegments } from './draw.segments';
import { LinkEventService } from '../link-event.service';
import { Point } from 'leaflet';
import { Subscription } from 'rxjs';
import { ShapeMouseEvent } from '../shape.mouse.event';


@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent implements OnInit, AfterViewInit {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

  nbControlPoints = 0
  isFloatingOrthogonal = false

  // to compute wether it was a select / dragging event
  dragging = false

  // offset between the cursor at the start and the top left corner
  offsetX = 0;
  offsetY = 0;
  distanceMoved = 0

  private mousePosRelativeToShapeAtMouseDown: { x: number; y: number } = { x: 0, y: 0 }

  private dragThreshold = 5;
  isSelected = false
  // LinkAtMouseDown is the clone of the Link when mouse down
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private LinkAtMouseDown: gongsvg.LinkDB | undefined


  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private linkService: gongsvg.LinkService,
    private linkEventService: LinkEventService,
    private elementRef: ElementRef) {

    this.subscriptions.push(
      linkEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {
          this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)
          this.LinkAtMouseDown = structuredClone(this.Link!)
        })
    )

    this.subscriptions.push(
      linkEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (shapeMouseEvent.ShapeID === this.Link?.ID && this.dragging) {
            let segment = this.segments![shapeMouseEvent.SegmentNumber]

            if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
              let deltaY = shapeMouseEvent.Point.Y - segment.StartPoint.Y
              if (segment.Number == 0 && deltaY != 0) {
                let newRatio = (shapeMouseEvent.Point.Y - this.Link!.Start!.Y) / this.Link!.Start!.Height
                // console.log("shapeMouseEvent.Point.Y\t\t", shapeMouseEvent.Point.Y)
                // console.log("this.Link!.Start!.Y\t\t", this.Link!.Start!.Y)
                // console.log("this.Link!.Start!.Height\t", this.Link!.Start!.Height)
                // console.log("deltaY\t\t\t\t", deltaY)
                // console.log("Old ratio\t\t\t", this.Link!.StartRatio)
                // console.log("New ratio\t\t\t", newRatio)

                if (newRatio < 0) { newRatio = 0 }
                if (newRatio > 1) { newRatio = 1 }
                this.Link!.StartRatio = newRatio
              }

              // last segment
              if (segment.Number == this.segments!.length - 1 && deltaY != 0) {

                let newRatio = (shapeMouseEvent.Point.Y - this.Link!.End!.Y) / this.Link!.End!.Height

                if (newRatio < 0) { newRatio = 0 }
                if (newRatio > 1) { newRatio = 1 }
                this.Link!.EndRatio = newRatio
              }

              // middle segment
              if (this.segments!.length == 3 && segment.Number == 1 && deltaY != 0) {
                let newRatio = (shapeMouseEvent.Point.Y - this.Link!.Start!.Y) / this.Link!.Start!.Height
                this.Link!.CornerOffsetRatio = newRatio
              }
            }
            if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
              let deltaX = shapeMouseEvent.Point.X - segment.StartPoint.X
              if (segment.Number == 0 && deltaX != 0) {
                let newRatio = (shapeMouseEvent.Point.X - this.Link!.Start!.X) / this.Link!.Start!.Width

                if (newRatio < 0) { newRatio = 0 }
                if (newRatio > 1) { newRatio = 1 }
                this.Link!.StartRatio = newRatio
              }

              // last segment
              if (segment.Number == this.segments!.length - 1 && deltaX != 0) {

                let newRatio = (shapeMouseEvent.Point.X - this.Link!.End!.X) / this.Link!.End!.Width

                if (newRatio < 0) { newRatio = 0 }
                if (newRatio > 1) { newRatio = 1 }
                this.Link!.EndRatio = newRatio
              }

              // middle segment
              if (this.segments!.length == 3 && segment.Number == 1 && deltaX != 0) {
                let newRatio = (shapeMouseEvent.Point.X - this.Link!.Start!.X) / this.Link!.Start!.Width
                this.Link!.CornerOffsetRatio = newRatio
              }
            }
          }

        })
    )

    this.subscriptions.push(
      linkEventService.mouseMouseUpEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (shapeMouseEvent.ShapeID === this.Link?.ID && this.dragging) {

            let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
            let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
            this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

            if (this.distanceMoved < this.dragThreshold) {
              console.log("Link, link selected selected: ", this.Link?.Name)
            } else {
              this.linkService.updateLink(this.Link!, this.GONG__StackPath).subscribe(
                link => {
                  // console.log("Updated", link.ID)
                }
              )
            }

          }
          this.dragging = false
        })
    )
  }


  ngOnInit(): void {
    // console.log("LinkComponent init: ", this.Link?.Name)

    this.isFloatingOrthogonal = this.Link!.Type == gongsvg.LinkType.LINK_TYPE_FLOATING_ORTHOGONAL

    if (this.Link) {
      if (this.Link.ControlPoints) {
        this.nbControlPoints = this.Link.ControlPoints.length
      }
    }
  }

  public getPosition(rect: gongsvg.RectDB | undefined, position: string | undefined): Coordinate {

    let coordinate: Coordinate = [0, 0]

    if (rect == undefined || position == undefined) {
      return coordinate
    }

    switch (position) {
      case gongsvg.AnchorType.ANCHOR_BOTTOM:
        coordinate = [rect.X + rect.Width / 2, rect.Y + rect.Height]
        break;
      case gongsvg.AnchorType.ANCHOR_TOP:
        coordinate = [rect.X + rect.Width / 2, rect.Y]
        break;
      case gongsvg.AnchorType.ANCHOR_LEFT:
        coordinate = [rect.X, rect.Y + rect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_RIGHT:
        coordinate = [rect.X + rect.Width, rect.Y + rect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_CENTER:
        coordinate = [rect.X + rect.Width / 2, rect.Y + rect.Height / 2]
        break;
    }

    return coordinate
  }

  linkMouseDown(event: MouseEvent, segmentNumber: number): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      // this link shit to dragging state
      this.dragging = true

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
        SegmentNumber: segmentNumber
      }
      this.linkEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  linkMouseMove(event: MouseEvent, segmentNumber: number = 0): void {

    if (!event.altKey && !event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
        SegmentNumber: segmentNumber
      }
      this.linkEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  linkMouseUp(event: MouseEvent, segmentNumber: number = 0): void {

    // console.log("Link : linkMouseUp", this.Link?.Name)
    if (!event.altKey && !event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
        SegmentNumber: segmentNumber
      }
      this.linkEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  connectorParams: ConnectorParams | undefined
  segments: Segment[] | undefined

  drawConnector(): boolean {

    let link = this.Link!

    this.connectorParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartDirection! as gongsvg.OrientationType,
      EndDirection: link.EndDirection! as gongsvg.OrientationType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
    }

    this.segments = drawSegments(this.connectorParams)

    return true
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
      return 'horizontal';
    } else if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
      return 'vertical';
    } else {
      return null; // You can return null or another value if the line is not strictly horizontal or vertical
    }
  }

  pageX: number = 0
  pageY: number = 0
  getSvgTopLeftCoordinates() {
    const svgElement = this.elementRef.nativeElement.querySelector('svg');
    const svgRect = svgElement.getBoundingClientRect();
    this.pageX = svgRect.left + window.pageXOffset;
    this.pageY = svgRect.top + window.pageYOffset;

    // console.log('SVG Top-Left Corner:', this.pageX, this.pageY);
  }

  ngAfterViewInit() {
    this.getSvgTopLeftCoordinates();
  }
}

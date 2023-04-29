import { AfterViewInit, Component, ElementRef, Input, OnInit } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { ConnectorParams, createPoint, drawConnector } from './draw.connector';
import { LinkEventService, ShapeMouseEvent } from '../link-event.service';
import { Point } from 'leaflet';
import { Subscription } from 'rxjs';

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
  private startPosition: { x: number; y: number } = { x: 0, y: 0 }
  private dragThreshold = 5;
  isSelected = false

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private linkEventService: LinkEventService,
    private elementRef: ElementRef) {

    this.subscriptions.push(
      linkEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          this.offsetX = shapeMouseEvent.Point.X - this.Link!.Start!.X
          this.offsetY = shapeMouseEvent.Point.X - this.Link!.Start!.X

          this.startPosition = { x: this.offsetX, y: this.offsetY }
        })
    );

    this.subscriptions.push(
      linkEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (shapeMouseEvent.shapeID === this.Link?.ID && this.dragging) {

            console.log("LinkComponent processing mouse move service event: ",
              this.Link?.Name,
              shapeMouseEvent.Point.X,
              shapeMouseEvent.Point.Y)
          }

        })
    );
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

  linkMouseDown(event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      // this link shit to dragging state
      this.dragging = true

      let shapeMouseEvent: ShapeMouseEvent = {
        shapeID: this.Link!.ID,
        Point: createPoint(x, y)
      }
      this.linkEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  linkMouseMove(event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      const deltaX = x - this.startPosition.x
      const deltaY = y - this.startPosition.y
      this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

      let shapeMouseEvent: ShapeMouseEvent = {
        shapeID: this.Link!.ID,
        Point: createPoint(deltaX, deltaY)
      }
      this.linkEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  linkMouseLeave(event: MouseEvent): void {

    this.dragging = false

    if (!event.altKey && !event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      const deltaX = x - this.startPosition.x
      const deltaY = y - this.startPosition.y
      this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

      let shapeMouseEvent: ShapeMouseEvent = {
        shapeID: this.Link!.ID,
        Point: createPoint(x, y)
      }
      this.linkEventService.emitMouseLeaveEvent(shapeMouseEvent)
    }
  }

  linkMouseUp(event: MouseEvent): void {

    this.dragging = false

    console.log("Link : ", this.Link?.Name)

    if (!event.altKey && !event.shiftKey) {

      if (this.distanceMoved < this.dragThreshold) {
        console.log("Link, link selected selected: ", this.Link?.Name)
      }
    }
  }

  connectorParams: ConnectorParams | undefined
  segments: gongsvg.PointDB[][] | undefined

  drawConnector(): boolean {

    let link = this.Link!

    this.connectorParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartDirection! as gongsvg.DirectionType,
      EndDirection: link.EndDirection! as gongsvg.DirectionType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
    }

    this.segments = drawConnector(this.connectorParams)

    return true
  }

  getOrientation(segment: gongsvg.PointDB[]): 'horizontal' | 'vertical' | null {
    if (segment[0].Y === segment[1].Y) {
      return 'horizontal';
    } else if (segment[0].X === segment[1].X) {
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

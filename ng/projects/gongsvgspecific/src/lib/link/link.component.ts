import { AfterViewInit, Component, ElementRef, Input, OnInit } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { shortestFloatingOrthogonalConnector } from './compute.floating.connectors';
import { ConnectorParams, drawConnector } from './compute.floating.connectors.2';

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
  distanceMoved = 0
  private startPosition: { x: number; y: number } = { x: 0, y: 0 }
  private dragThreshold = 5;
  isSelected = false

  constructor(
    private elementRef: ElementRef) { }


  ngOnInit(): void {
    console.log("LinkComponent init: ", this.Link?.Name)

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
    }
  }

  linkMouseMove(event: MouseEvent): void {

    let x = event.clientX - this.pageX
    let y = event.clientY - this.pageY

    const deltaX = x - this.startPosition.x;
    const deltaY = y - this.startPosition.y;
    this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY);
  }


  linkMouseUp(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {

      if (this.distanceMoved < this.dragThreshold) {

        console.log("Link, link selected selected: ", this.Link?.Name)

      }
    }
  }

  connectorParams: ConnectorParams | undefined
  points: gongsvg.PointDB[][] | undefined

  drawConnector(): boolean {

    let link = this.Link!

    this.connectorParams = {
      startRect: link.Start!,
      endRect: link.End!,
      startDirection: gongsvg.DirectionType.DIRECTION_HORIZONTAL,
      endDirection: gongsvg.DirectionType.DIRECTION_VERTICAL,
      startRatio: link.StartRatio,
      endRatio: link.EndRatio,
    }

    this.points = drawConnector(this.connectorParams)

    return true
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

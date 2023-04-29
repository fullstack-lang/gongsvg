import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { ElementRef, Renderer2, AfterViewInit } from '@angular/core';


import { RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { Subscription } from 'rxjs';

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit, OnDestroy, AfterViewInit {

  @Input() Rect: gongsvg.RectDB = new gongsvg.RectDB()
  @Input() GONG__StackPath: string = ""

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' | null = null;

  rectDragging: boolean = false;

  // offset between the cursor at the start and the top left corner
  private RectAtMouseDown: gongsvg.RectDB | undefined

  // to compute wether it was a select / dragging event
  distanceMoved = 0
  private mousePosRelativeToSvgAtMouseDown: { x: number; y: number } = { x: 0, y: 0 }
  private dragThreshold = 5;

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private rectService: gongsvg.RectService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private elementRef: ElementRef) {

    this.subscriptions.push(
      rectangleEventService.mouseRectMouseDownEvent$.subscribe(
        ({ rectangleID: rectangleID, MousePosRelativeSVG: coordinate }) => {

          let x = coordinate[0]
          let y = coordinate[1]

          this.RectAtMouseDown = structuredClone(this.Rect)

          this.mousePosRelativeToSvgAtMouseDown = { x: x, y: y }
        })
    );

    this.subscriptions.push(
      rectangleEventService.mouseRectMouseUpEvent$.subscribe(
        (rectangleID: number) => {

          if (rectangleID != this.Rect.ID) {
            if (this.Rect.IsSelected) {
              this.Rect.IsSelected = false
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
            }
          }
          // console.log('Rect ', this.Rect.Name, 'Mouse down event occurred on rectangle ', rectangleID);
        })
    );

    this.subscriptions.push(
      svgEventService.multiShapeSelectEndEvent$.subscribe(
        (selectAreaConfig: SelectAreaConfig) => {

          if (this.Rect.IsSelected) {
            return
          }

          switch (selectAreaConfig.SweepDirection) {
            case SweepDirection.LEFT_TO_RIGHT:

              // rectangle has to be in boxed in the rect
              if (
                this.Rect.X > selectAreaConfig.TopLeft[0] &&
                this.Rect.X + this.Rect.Width < selectAreaConfig.BottomRigth[0] &&
                this.Rect.Y > selectAreaConfig.TopLeft[1] &&
                this.Rect.Y + this.Rect.Height < selectAreaConfig.BottomRigth[1]
              ) {
                this.Rect.IsSelected = true
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
              }
              break
            case SweepDirection.RIGHT_TO_LEFT:

              // rectangle has to be partialy boxed in the rect
              if (
                this.Rect.X < selectAreaConfig.BottomRigth[0] &&
                this.Rect.X + this.Rect.Width > selectAreaConfig.TopLeft[0] &&
                this.Rect.Y < selectAreaConfig.BottomRigth[1] &&
                this.Rect.Y + this.Rect.Height > selectAreaConfig.TopLeft[1]
              ) {
                this.Rect.IsSelected = true
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
              }
              break
          }
        })
    );

    this.subscriptions.push(
      rectangleEventService.mouseRectMouseDragEvent$.subscribe(
        ({ rectangleID: rectangleID, MousePosRelativeSVG: mousePosRelativeSVG }) => {

          if (this.rectDragging || this.Rect.IsSelected) {
            if (this.Rect?.CanMoveHorizontaly) {
              console.log("Rect : mouseRectMouseDragEvent",
                "this.Rect.X", this.Rect.X,
                "mousePosRelativeSVG[0]", mousePosRelativeSVG[0],
                "this.mousePosRelativeToSvgAtMouseDown.x", this.mousePosRelativeToSvgAtMouseDown.x
              )

              this.Rect.X = this.RectAtMouseDown!.X +
                (mousePosRelativeSVG[0] - this.mousePosRelativeToSvgAtMouseDown.x)
            }
            if (this.Rect?.CanMoveVerticaly) {
              this.Rect.Y = this.RectAtMouseDown!.Y +
                (mousePosRelativeSVG[1] - this.mousePosRelativeToSvgAtMouseDown.y)
            }
          }
        })
    );

  }

  ngOnInit(): void {
  }

  onSVGClick(event: MouseEvent) {
    console.log("rect, onSVGClick(): ", this.Rect?.Name)
  }

  onRectClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (this.rectDragging) {
      this.rectDragging = false;
      return
    }

    if (!event.altKey && !event.shiftKey) {


    }
  }

  rectMouseDown(event: MouseEvent): void {

    if (!event.altKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true;

      let mousePosRelativeSvg: { x: number; y: number } =
      {
        x: event.clientX - this.pageX,
        y: event.clientY - this.pageY
      }

      let mouseEvent = {
        rectangleID: this.Rect.ID,
        MousePosRelativeSVG: [mousePosRelativeSvg.x, mousePosRelativeSvg.y] as [number, number]
      }
      this.rectangleEventService.emitRectMouseDownEvent(mouseEvent)
    } else {
      console.log("startRectDrag + shiftKey : ", this.Rect?.Name)

      const actualX = event.clientX - this.pageX
      const actualY = event.clientY - this.pageY

      this.rectangleEventService.emitRectAltKeyMouseDownEvent(
        this.Rect.ID,
        [actualX, actualY]);
    }
  }

  rectMouseMove(event: MouseEvent): void {

    if (!this.rectDragging) {
      return
    }

    let mousePosRelativeSvg: { x: number; y: number } =
    {
      x: event.clientX - this.pageX,
      y: event.clientY - this.pageY
    }

    const deltaX = mousePosRelativeSvg.x - this.mousePosRelativeToSvgAtMouseDown.x;
    const deltaY = mousePosRelativeSvg.y - this.mousePosRelativeToSvgAtMouseDown.y;
    this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY);

    // console.log("RectComponent DragRect : ", deltaX, deltaY, distanceMoved)

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', this.Rect.Name, event.clientX, event.clientY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(
        [mousePosRelativeSvg.x, mousePosRelativeSvg.y])
      return
    }

    if (event.shiftKey) {
      this.svgEventService.emitMultiShapeSelectDrag([mousePosRelativeSvg.x, mousePosRelativeSvg.y])
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    let mouseEvent = {
      rectangleID: this.Rect.ID,
      MousePosRelativeSVG: [mousePosRelativeSvg.x, mousePosRelativeSvg.y] as [number, number]
    }

    this.rectangleEventService.emitRectMouseDragEvent(mouseEvent)
  }

  rectMouseUp(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {

      this.rectDragging = false

      if (this.distanceMoved > this.dragThreshold) {
        this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
      } else {
        if (this.Rect?.IsSelectable) {
          console.log("rect, onRectClick() toggle selected: ", this.Rect?.Name)
          this.Rect.IsSelected = !this.Rect.IsSelected
          this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()

          this.rectangleEventService.emitRectMouseUpEvent(this.Rect.ID)
        }
      }
    }

    if (event.altKey) {
      console.log("endRectDrag + shiftKey rect : ", this.Rect?.Name)

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(this.Rect.ID);
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent([event.clientX, event.clientY])
    }
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  // 
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


  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;
  }

  anchorMouseMove(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!this.anchorDragging) {
      return;
    }

    const newX = event.clientX - this.pageX
    const newY = event.clientY - this.pageY

    if (anchor === 'left') {
      const originalRightEdge = this.Rect.X + this.Rect.Width;
      this.Rect.X = newX;
      this.Rect.Width = originalRightEdge - newX;
    } else if (anchor === 'right') {
      this.Rect.Width = newX - this.Rect.X;
    } else if (anchor === 'top') {
      const originalBottomEdge = this.Rect.Y + this.Rect.Height;
      this.Rect.Y = newY;
      this.Rect.Height = originalBottomEdge - newY;
    } else if (anchor === 'bottom') {
      this.Rect.Height = newY - this.Rect.Y;
    }
  }

  anchorMouseUp(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = null;
      this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
    }
  }
}

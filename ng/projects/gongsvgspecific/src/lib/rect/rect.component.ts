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

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private rectService: gongsvg.RectService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private elementRef: ElementRef,
    private renderer: Renderer2) {

    this.subscriptions.push(
      rectangleEventService.mouseRectMouseDownEvent$.subscribe(
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

  }

  ngOnInit(): void {
  }

  onSVGClick(event: MouseEvent) {
    console.log("rect, onSVGClick(): ", this.Rect?.Name)
  }

  onRectClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      if (this.Rect?.IsSelectable) {
        console.log("rect, onRectClick() toggle selected: ", this.Rect?.Name)
        this.Rect.IsSelected = !this.Rect.IsSelected
        this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()

        this.rectangleEventService.emitRectMouseDownEvent(this.Rect.ID)
      }
    }
  }

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' | null = null;

  rectDragging: boolean = false;
  offsetX = 0;
  offsetY = 0;

  startAnchorDrag(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;
  }

  anchorDrag(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
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

  endAnchorDrag(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = null;
      this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
    }
  }

  startRectDrag(event: MouseEvent): void {

    if (!event.altKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true;

      this.offsetX = event.clientX - this.Rect.X
      this.offsetY = event.clientY - this.Rect.Y
    } else {
      console.log("startRectDrag + shiftKey : ", this.Rect?.Name)

      const actualX = event.clientX - this.pageX
      const actualY = event.clientY - this.pageY

      this.rectangleEventService.emitRectAltKeyMouseDownEvent(
        this.Rect.ID,
        [actualX, actualY]);
    }
  }

  dragRect(event: MouseEvent): void {

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', this.Rect.Name, event.clientX, event.clientY);

      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      this.rectangleEventService.emitRectAltKeyMouseDragEvent([x, y])
      return
    }

    if (event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      this.svgEventService.emitMultiShapeSelectDrag([x, y])
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!this.rectDragging) {
      return;
    }

    if (this.Rect?.CanMoveHorizontaly) {
      this.Rect.X = event.clientX - this.offsetX;
    }
    if (this.Rect?.CanMoveVerticaly) {
      this.Rect.Y = event.clientY - this.offsetY;
    }
  }

  endRectDrag(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      this.rectDragging = false;
      this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
    }

    if (event.altKey) {
      console.log("endRectDrag + shiftKey rect : ", this.Rect?.Name)

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(this.Rect.ID);
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
}

import { Component, OnInit, Input } from '@angular/core';
import { RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit {

  @Input() Rect: gongsvg.RectDB = new gongsvg.RectDB()
  @Input() GONG__StackPath: string = ""

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  constructor(
    private rectService: gongsvg.RectService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService) { }

  ngOnInit(): void {
  }

  onClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey) {
      if (this.Rect?.IsSelectable) {
        console.log("rect, onClick() toggle selected: ", this.Rect?.Name)
        this.Rect!.IsSelected = !this.Rect!.IsSelected
        this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
      }
    }
  }

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' | null = null;

  rectDragging: boolean = false;
  offsetX = 0;
  offsetY = 0;

  startAnchorDrag(event: MouseEvent, anchor: 'left' | 'right'| 'top' | 'bottom' ): void {
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
  
    const newX = event.clientX;
    const newY = event.clientY;
  
    if (anchor === 'left') {
      const originalRightEdge = this.Rect!.X + this.Rect!.Width;
      this.Rect!.X = newX;
      this.Rect!.Width = originalRightEdge - newX;
    } else if (anchor === 'right') {
      this.Rect!.Width = newX - this.Rect!.X;
    } else if (anchor === 'top') {
      const originalBottomEdge = this.Rect!.Y + this.Rect!.Height;
      this.Rect!.Y = newY;
      this.Rect!.Height = originalBottomEdge - newY;
    } else if (anchor === 'bottom') {
      this.Rect!.Height = newY - this.Rect!.Y;
    }
  }

  endAnchorDrag(event: MouseEvent): void {
    if (!event.altKey) {
      this.anchorDragging = false;
      this.activeAnchor = null;
      this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
    }
  }


  // management for line drawing
  private selectedRect: any
  private svg: gongsvg.SVGDB = new gongsvg.SVGDB()

  startRectDrag(event: MouseEvent): void {

    if (!event.altKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true;

      this.offsetX = event.clientX - this.Rect!.X
      this.offsetY = event.clientY - this.Rect!.Y
    } else {
      console.log("startRectDrag + shiftKey : ", this.Rect?.Name)

      this.rectangleEventService.emitRectAltKeyMouseDownEvent(this.Rect!.ID);
    }
  }

  dragRect(event: MouseEvent): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (event.altKey) {
      this.rectangleEventService.emitRectAltKeyMouseDragEvent([event.clientX, event.clientY])
    }

    if (!this.rectDragging) {
      return;
    }

    if (this.Rect?.CanMoveHorizontaly) {
      this.Rect!.X = event.clientX - this.offsetX;
    }
    if (this.Rect?.CanMoveVerticaly) {
      this.Rect!.Y = event.clientY - this.offsetY;
    }
  }

  endRectDrag(event: MouseEvent): void {
    if (!event.altKey) {
      this.rectDragging = false;
      this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
    } else {
      console.log("endRectDrag + shiftKey rect : ", this.Rect?.Name)

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(this.Rect!.ID);
    }
  }
}

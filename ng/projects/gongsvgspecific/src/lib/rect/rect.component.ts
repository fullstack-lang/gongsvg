import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit {

  @Input() Rect?: gongsvg.RectDB
  @Input() GONG__StackPath: string = ""

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  constructor(
    private rectService: gongsvg.RectService,
    private svgService: gongsvg.SVGService) { }

  ngOnInit(): void {
  }

  onClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    // console.log("rect, onClick() : ", this.Rect?.Name)
    if (this.Rect?.IsSelectable) {
      this.Rect!.IsSelected = !this.Rect!.IsSelected
      this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
    }
  }

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | null = null;

  rectDragging: boolean = false;
  offsetX = 0;
  offsetY = 0;

  startDrag(event: MouseEvent, anchor: 'left' | 'right'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;
  }

  drag(event: MouseEvent, anchor: 'left' | 'right'): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!this.anchorDragging) {
      return;
    }

    const newX = event.clientX;


    if (anchor === 'left') {
      const originalRightEdge = this.Rect!.X + this.Rect!.Width;
      this.Rect!.X = newX;
      this.Rect!.Width = originalRightEdge - newX;
    } else if (anchor === 'right') {
      this.Rect!.Width = newX - this.Rect!.X;
    }
  }

  endDrag(): void {
    this.anchorDragging = false;
    this.activeAnchor = null;
    this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()

  }


  // management for line drawing
  private selectedRect: any
  private svg: gongsvg.SVGDB = new gongsvg.SVGDB()

  startRectDrag(event: MouseEvent): void {

    if (!event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true;

      this.offsetX = event.clientX - this.Rect!.X
      this.offsetY = event.clientY - this.Rect!.Y
    } else {
      console.log("startRectDrag rect : ", this.Rect?.Name)

      // get the SVG and inform it that there it swithed in line drawing mode
      this.svgService.getSVGs(this.GONG__StackPath).subscribe(
        (svgArray: gongsvg.SVGDB[]) => {
          if (svgArray.length != 1) {
            console.log("problem with svg, length ", svgArray.length, " is not 1")
            return
          }
          this.svg = svgArray[0]

          if (this.svg.DrawingState != gongsvg.DrawingState.NOT_DRAWING_LINE) {
            console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.NOT_DRAWING_LINE)
          }

          this.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINE

          this.svg.StartRectID.Valid = true
          this.svg.StartRectID.Int64 = this.Rect!.ID

          this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe()
        }
      )
    }
  }

  dragRect(event: MouseEvent): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

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
    if (!event.shiftKey) {
      this.rectDragging = false;
      this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
    } else {
      console.log("endRectDrag rect : ", this.Rect?.Name)

      // get the SVG and inform it that there it swithed in line drawing mode
      this.svgService.getSVGs(this.GONG__StackPath).subscribe(
        (svgArray: gongsvg.SVGDB[]) => {
          if (svgArray.length != 1) {
            console.log("problem with svg, length ", svgArray.length, " is not 1")
            return
          }
          this.svg = svgArray[0]

          if (this.svg.DrawingState != gongsvg.DrawingState.DRAWING_LINE) {
            console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.DRAWING_LINE)
          }

          this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE

          this.svg.EndRectID.Valid = true
          this.svg.EndRectID.Int64 = this.Rect!.ID

          this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe()
        }
      )
    }
  }
}

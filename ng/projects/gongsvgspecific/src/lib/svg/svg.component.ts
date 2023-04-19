import { AfterViewInit, Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit, OnDestroy, AfterViewInit {

  @Input() GONG__StackPath: string = ""

  public gongsvgFrontRepo?: gongsvg.FrontRepo


  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommiNbFromBagetCommitNbFromBack = -1
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


  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private elementRef: ElementRef, 
  ) {

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseDownEvent$.subscribe(
        ({ rectangleID: rectangleID, Coordinate : coordinate} ) => {
        console.log('SvgComponent, Mouse down event occurred on rectangle ', rectangleID, " at ", coordinate)
        this.linkStartRectangleID = rectangleID

        let rect = this.gongsvgFrontRepo?.Rects.get(rectangleID)

        if (rect == undefined) {
          return
        }

        this.linkDrawing = true
        this.startX = coordinate[0]
        this.startY = coordinate[1]
      })
    );

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseDragEvent$.subscribe((coordinate: Coordinate) => {

        this.endX = coordinate[0]
        this.endY = coordinate[1]
        console.log('SvgComponent, Received Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
      })
    )

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseUpEvent$.subscribe((rectangleID: number) => {
        console.log('SvgComponent, Mouse up event occurred on rectangle ', rectangleID);
        this.linkDrawing = false

        this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
      })
    )

    // Add a global mouse up event listener
    document.addEventListener('mouseup', this.handleGlobalMouseUp)
  }



  ngOnInit(): void {

    // console.log("Svg component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommiNbFromBagetCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.refresh()
          this.lastCommiNbFromBagetCommitNbFromBack = commiNbFromBagetCommitNbFromBack
        }
      }
    )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        if (this.gongsvgFrontRepo.SVGs_array.length == 1) {
          this.svg = this.gongsvgFrontRepo.SVGs_array[0]
        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        this.svg.Layers.sort((t1, t2) => {
          let t1_revPointerID_Index = t1.SVG_LayersDBID_Index
          let t2_revPointerID_Index = t2.SVG_LayersDBID_Index

          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }

    )
  }

  fillColor = 'rgb(255, 0, 0)';
  Text = "Toto"
  width = 150
  height = 200
  changeColor() {
    const r = Math.floor(Math.random() * 256);
    const g = Math.floor(Math.random() * 256);
    const b = Math.floor(Math.random() * 256);
    this.fillColor = `rgb(${r}, ${g}, ${b})`;
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  handleGlobalMouseUp = (event: MouseEvent) => {
    // console.log("Layer Component, global mouse up")
    this.linkDrawing = false
  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    let svgArray = this.gongsvgFrontRepo?.SVGs_array
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
      console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.NOT_DRAWING_LINE)
    }

    this.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINE

    this.svg.StartRectID.Valid = true
    this.svg.StartRectID.Int64 = startRectangleID

    this.svg.EndRectID.Valid = true
    this.svg.EndRectID.Int64 = endRectangleID

    this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe(
      () => {
        // back to normal state
        this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE
        this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe()
      }
    )
  }

  onClick(event: MouseEvent) {
    console.log("SVG : on click()")
    // event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    // an event is emitted for all rects to go on a unselect mode
    this.rectangleEventService.emitRectMouseDownEvent(0)
  }

  dragLine(event: MouseEvent): void {


    const actualX = event.clientX - this.pageX
    const actualY = event.clientY - this.pageY
    
    // we want this event to bubble to the SVG element
    if (event.altKey) {
      console.log('SvgComponent, Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);

      this.rectangleEventService.emitRectAltKeyMouseDragEvent([actualX, actualY])
      return
    }
  }

  endDragLine(event: MouseEvent): void {

    if (event.altKey) {
      console.log("SvgComponent, End drag line")
    }
  }

  pageX: number = 0
  pageY: number = 0
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

  ngAfterViewInit() {
    const offset = this.getDivOffset(this.drawingArea!.nativeElement);
    this.pageX = offset.offsetX
    this.pageY = offset.offsetY
  }

  getDivOffset(div: HTMLDivElement): { offsetX: number; offsetY: number } {
    const rect = div.getBoundingClientRect();
    const offsetX = rect.left + window.pageXOffset;
    const offsetY = rect.top + window.pageYOffset;
    return { offsetX, offsetY };
  }
}

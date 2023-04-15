import { Component, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit, OnDestroy {

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
  // for link drawing
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
    private rectangleEventService: RectangleEventService
  ) {

    this.subscriptions.push(
      rectangleEventService.mouseRectShiftKeyDownEvent$.subscribe((rectangleID: number) => {
        console.log('Mouse down event occurred on rectangle ', rectangleID);
        this.linkStartRectangleID = rectangleID

        let rect = this.gongsvgFrontRepo?.Rects.get(rectangleID)

        if (rect == undefined) {
          return
        }

        this.linkDrawing = true
        this.startX = rect.X + rect.Width / 2;
        this.startY = rect.Y + rect.Height / 2;
      })
    );

    this.subscriptions.push(
      rectangleEventService.mouseRectShiftKeyMouseUpEvent$.subscribe((rectangleID: number) => {
        console.log('Mouse up event occurred on rectangle ', rectangleID);
        this.linkDrawing = false

        this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
      })
    )

    this.subscriptions.push(
      rectangleEventService.mouseRectShiftKeyDragEvent$.subscribe((coordinate: Coordinate) => {

        this.endX = coordinate[0]
        this.endY = coordinate[1]
      })
    )

    // Add a global mouse up event listener
    document.addEventListener('mouseup', this.handleGlobalMouseUp)
  }



  ngOnInit(): void {

    console.log("Svg component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

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
    console.log("Layer Component, global mouse up")
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

}

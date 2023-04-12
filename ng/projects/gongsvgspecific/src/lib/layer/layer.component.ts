import { Component, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-layer',
  templateUrl: './layer.component.html',
  styleUrls: ['./layer.component.css']
})
export class LayerComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""

  public gongsvgFrontRepo?: gongsvg.FrontRepo

  public Rects = new Array<gongsvg.RectDB>()
  public Texts = new Array<gongsvg.TextDB>()
  public Lines = new Array<gongsvg.LineDB>()
  public Circles = new Array<gongsvg.CircleDB>()
  public Ellipses = new Array<gongsvg.EllipseDB>()
  public Paths = new Array<gongsvg.PathDB>()
  public Polygones = new Array<gongsvg.PolygoneDB>()
  public Polylines = new Array<gongsvg.PolylineDB>()

  /**
 * the component is refreshed when modification are performed in the back repo 
 * 
 * the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
 * if the commit number has increased, it pulls the front repo and redraw the diagram
 */
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommiNbFromBagetCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  layer_?: gongsvg.LayerDB
  layerArray = new Array<gongsvg.LayerDB>()
  svg = new gongsvg.SVGDB
  linkStartRectangleID: number = 0

  private subscriptions: Subscription[] = [];

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService
  ) {

    this.subscriptions.push(
      rectangleEventService.mouseDownEvent$.subscribe((rectangleID: number) => {
        console.log('Mouse down event occurred on rectangle ', rectangleID);
        this.linkStartRectangleID = rectangleID
      })
    );

    this.subscriptions.push(
      rectangleEventService.mouseUpEvent$.subscribe((rectangleID: number) => {
        console.log('Mouse up event occurred on rectangle ', rectangleID);

        this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
      })
    );

  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    this.svgService.getSVGs(this.GONG__StackPath).subscribe(
      (svgArray: gongsvg.SVGDB[]) => {
        // update the svg
        if (svgArray.length == 1) {
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
          }
        )
      }
    )



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

        this.layerArray = this.svg.Layers

        this.layer_ = this.gongsvgFrontRepo.Layers_array[0]
        // console.log("svgSingloton " + this.layer_?.Name)

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

}

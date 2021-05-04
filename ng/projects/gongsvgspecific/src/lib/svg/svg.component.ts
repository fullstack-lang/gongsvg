import { Component, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit {

  SVG: string

  rect: number

  public gongsvgFrontRepo: gongsvg.FrontRepo

  public Rects = new Array<gongsvg.RectDB>()
  public Texts = new Array<gongsvg.TextDB>()
  public Lines = new Array<gongsvg.LineDB>()
  public Circles = new Array<gongsvg.CircleDB>()
  public Ellipses = new Array<gongsvg.EllipseDB>()

  /**
 * the component is refreshed when modification are performed in the back repo 
 * 
 * the checkCommitNbTimer polls the commit number of the back repo
 * if the commit number has increased, it pulls the front repo and redraw the diagram
 */
  checkCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  currTime: number

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgCommitNbService: gongsvg.CommitNbService,
  ) { }

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongsvgCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )
      }
    )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull().subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        this.Rects = new Array<gongsvg.RectDB>()
        this.Texts = new Array<gongsvg.TextDB>()
        this.Lines = new Array<gongsvg.LineDB>()
        this.Circles = new Array<gongsvg.CircleDB>()
        this.Ellipses = new Array<gongsvg.EllipseDB>()

        for (let svg of this.gongsvgFrontRepo.SVGs_array) {

          if (svg.Display) {

            svg.Rects?.forEach(
              rect => {
                this.Rects.push(rect)
              }
            )

            svg.Texts?.forEach(
              Text => {
                this.Texts.push(Text)
              }
            )

            svg.Lines?.forEach(
              Line => {
                this.Lines.push(Line)
              }
            )

            svg.Circles?.forEach(
              Circle => {
                this.Circles.push(Circle)
              }
            )

            svg.Ellipses?.forEach(
              Ellipse => {
                this.Ellipses.push(Ellipse)
              }
            )
          }
        }

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

}

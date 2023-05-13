import { Component, DoCheck, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { compareRectGeometries } from '../compare.rect.geometries';

@Component({
  selector: 'lib-rect-link-link',
  templateUrl: './rect-link-link.component.svg',
  styleUrls: ['./rect-link-link.component.css']
})
export class RectLinkLinkComponent implements OnInit, DoCheck {

  @Input() RectLinkLink: gongsvg.RectLinkLinkDB = new gongsvg.RectLinkLinkDB()
  @Input() GONG__StackPath: string = ""

  previousStart: gongsvg.RectDB | undefined
  previousEnd_StartRect: gongsvg.RectDB | undefined
  previousEnd_EndRect: gongsvg.RectDB | undefined

  ngOnInit(): void {
    this.resetPreviousState()
  }

  public getStartPosition(): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = this.previousStart!.X + this.previousStart!.Width / 2
    coordinate[1] = this.previousStart!.Y + this.previousStart!.Height / 2

    return coordinate
  }

  public getEndPosition(): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = this.previousEnd_StartRect!.X + this.previousEnd_StartRect!.Width / 2
    coordinate[1] = this.previousEnd_StartRect!.Y + this.previousEnd_StartRect!.Height / 2

    return coordinate
  }

  ngDoCheck(): void {

    let hasStartChanged = !compareRectGeometries(this.previousStart!, this.RectLinkLink!.Start!)
    let hasEnd_StartChanged = !compareRectGeometries(this.previousEnd_StartRect!, this.RectLinkLink!.End!.Start!)
    let hasEnd_EndChanged = !compareRectGeometries(this.previousEnd_EndRect!, this.RectLinkLink!.End!.End!)

    if (hasStartChanged || hasEnd_StartChanged || hasEnd_EndChanged) {
      // if (hasStartChanged || hasEndChanged) {
      // this.drawSegments()
      this.resetPreviousState()
    }
  }

  resetPreviousState() {
    this.previousStart = structuredClone(this.RectLinkLink!.Start!)
    this.previousEnd_StartRect = structuredClone(this.RectLinkLink!.End!.Start!)
    this.previousEnd_EndRect = structuredClone(this.RectLinkLink!.End!.End!)
  }
}

import { Component, DoCheck, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { compareRectGeometries } from '../compare.rect.geometries';

@Component({
  selector: 'lib-rect-link-link',
  templateUrl: './rect-link-link.component.svg',
  styleUrls: ['./rect-link-link.component.css']
})
export class RectLinkLinkComponent implements OnInit, DoCheck  {

  @Input() RectLinkLink: gongsvg.RectLinkLinkDB = new gongsvg.RectLinkLinkDB()
  @Input() GONG__StackPath: string = ""

  previousStart: gongsvg.RectDB | undefined
  previousEnd_StartRect: gongsvg.RectDB | undefined
  previousEnd_EndRect: gongsvg.RectDB | undefined

  ngOnInit(): void {
    this.resetPreviousState()
  }

  public getStartPosition(rect: gongsvg.RectDB): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = rect.X + rect.Width/2
    coordinate[1] = rect.Y + rect.Height/2

    return coordinate
  }

  public getEndPosition(link: gongsvg.LinkDB): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = link.Start!.X + link.Start!.Width/2
    coordinate[1] = link.Start!.Y + link.Start!.Height/2

    return coordinate
  }

  ngDoCheck(): void {

    let hasStartChanged = !compareRectGeometries(this.previousStart!, this.RectLinkLink!.Start!)

    if (hasStartChanged ) {
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

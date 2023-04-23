import { Component, Input, OnInit } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';

@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent implements OnInit {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

  ngOnInit(): void {
    console.log("LinkComponent init: ", this.Link?.Name)
  }

  public getPosition(rect: gongsvg.RectDB | undefined, position: string | undefined): Coordinate {

    let coordinate: Coordinate = [0, 0]

    if (rect == undefined || position == undefined) {
      return coordinate
    }

    switch (position) {
      case gongsvg.AnchorType.ANCHOR_BOTTOM:
        coordinate = [rect.X + rect.Width / 2, rect.Y + rect.Height]
        break;
      case gongsvg.AnchorType.ANCHOR_TOP:
        coordinate = [rect.X + rect.Width / 2, rect.Y]
        break;
      case gongsvg.AnchorType.ANCHOR_LEFT:
        coordinate = [rect.X, rect.Y + rect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_RIGHT:
        coordinate = [rect.X + rect.Width, rect.Y + rect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_CENTER:
        break;
    }

    return coordinate
  }

}

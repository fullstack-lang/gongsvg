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
  anchorRadius = 5; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color


  constructor(
    private rectService: gongsvg.RectService) { }

  ngOnInit(): void {
  }

  onClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element
    this.Rect!.Selected = true
    this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
  }

  onSvgClick(event: MouseEvent) {
    this.Rect!.Selected = false
    this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
  }
}

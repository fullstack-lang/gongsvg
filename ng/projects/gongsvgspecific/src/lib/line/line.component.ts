import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'lib-line',
  templateUrl: './line.component.svg',
  styleUrls: ['./line.component.css']
})
export class LineComponent implements OnInit {

  @Input() X1: number
  @Input() Y1: number
  @Input() X2: number
  @Input() Y2: number

  @Input() Color: string
  @Input() Stroke: string
  @Input() StrokeWidth: string
  @Input() FillOpacity: number
  @Input() StrokeDashArray: string
  @Input() Transform: string

  constructor() { }

  ngOnInit(): void {
    console.log("line X1 " + this.X1)
  }

}
